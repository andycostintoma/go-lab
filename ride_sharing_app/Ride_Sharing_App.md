## **Asynchronous Communication**
### **Introduction to Asynchronous Architecture**
![Screenshot 2025-11-16 at 20.14.05.png](Screenshot_2025-11-16_at_20.14.05.png)
![Screenshot 2025-11-16 at 20.24.31.png](Screenshot_2025-11-16_at_20.24.31.png)
![Screenshot 2025-11-16 at 20.23.34.png](Screenshot_2025-11-16_at_20.23.34.png)
![Screenshot 2025-11-16 at 20.23.23.png](Screenshot_2025-11-16_at_20.23.23.png)
## **1. Why We Need Asynchronous Communication**
We introduce an asynchronous communication layer to remove direct service dependencies.
### Key motivations:
- Avoid blocking or waiting for slow services
- Prevent failures when other services are down
- Increase resilience and reliability
- Enable message buffering and retries
- Support distributed system behavior
So far we used:
- HTTP → initial
- gRPC → efficient internal comms
- WebSockets → real-time UI updates
Now we add **message queues** (RabbitMQ).
---
## **2. The Problem With the Current Synchronous Flow**
When a user schedules a trip:
- TripService must notify DriverService directly
- If DriverService is down → the flow fails
- If DriverService is overloaded → slow or broken flow
- No buffering or retry logic
- TripService is tightly coupled to DriverService
This does not scale and is not fault tolerant.
---
## **3. How Asynchronous Messaging Fixes This**
We use RabbitMQ as a **middleman**:
### New behavior:
- TripService publishes an event: `trip.created`
- DriverService consumes it when ready
- DriverService finds a driver and publishes `trip.driver_assigned`
- TripService consumes it and updates the trip
- User is notified via WebSocket
### Why this works:
- TripService does not care if DriverService is online
- RabbitMQ stores messages until consumed
- Flow becomes reliable and decoupled
- Services scale independently
---
## **4. RabbitMQ Concepts (What You Need to Know First)**
### **Message Broker**
A component that accepts, stores, and routes messages.
### **Producer**
A program that **sends** messages.
Examples:
- TripService → sends `trip.created`
- DriverService → sends `trip.driver_assigned`
### **Consumer**
A program that **reads** messages.
Examples:
- DriverService → reads `trip.created`
- TripService → reads `trip.driver_assigned`
### **Message**
Has:
- Body → business data
- Headers → metadata (trace IDs, correlation IDs)
### **Queue**
A buffer where messages wait until a consumer is available.
### **Broker responsibility**
- Stores undelivered messages
- Routes messages correctly
- Balances load across consumers
- Guarantees delivery (when configured)
---
## **5. Understanding the Event-Driven Trip Flow**
### Trip Creation Flow:
1. User schedules a trip
2. API Gateway → TripService (gRPC)
3. TripService publishes `trip.created`
4. DriverService consumes event → determines best driver
5. DriverService publishes `trip.driver_assigned`
6. TripService consumes event → updates DB
7. TripService notifies user via WebSocket
This is **true asynchronous behavior**: no waiting for direct responses.
---
## **6. RabbitMQ vs Kafka**
Instructor’s reasoning:
- RabbitMQ is easier to teach
- Perfect for beginner/intermediate workloads
- Matches our project’s traffic scale
- Concepts later transfer to Kafka
Kafka is for very heavy throughput — unnecessary here.
---
## **7. Benefits of Moving to Asynchronous Architecture**
- No waiting for downstream services
- No flow interruption when a service is offline
- Messages persist until consumed
- Better reliability and fault tolerance
- Horizontal scaling of consumers
- Strong decoupling
- More natural distributed behavior
## **Connecting Services to RabbitMQ**
### **1. Establishing a Connection to the Broker**
Each microservice must establish a long-lived connection to RabbitMQ. The recommended Go client is:
```
github.com/rabbitmq/amqp091-go
```
This client provides the AMQP protocol implementation used throughout the system.
A connection string follows the format:
```
amqp://username:password@host:port/
```
Hardcoding this value is discouraged. Instead, each service retrieves the URI from an environment variable. This allows development, staging, and production to use different credentials without altering the application code.
```go
rabbitMqURI := env.GetString(
    "RABBITMQ_URI",
    "amqp://guest:guest@rabbitmq:5672/",
)
```
If the variable is absent, the fallback points to the RabbitMQ Kubernetes service.
---
## **2. Integrating the Environment Variable Through Kubernetes**
Each service receives the RabbitMQ URI from Kubernetes secrets:
```yaml
env:
  - name: RABBITMQ_URI
    valueFrom:
      secretKeyRef:
        name: rabbitmq-credentials
        key: uri
```
This ensures that the credentials stored in `rabbitmq-credentials` are injected directly into the service’s runtime environment.
---
## **3. Creating a Reusable Messaging Component**
To avoid duplicating connection logic across services, the messaging layer is abstracted inside a shared package. This package encapsulates:
- Establishing the AMQP connection
- Closing the connection
- (Later) managing channels, declaring queues, exchanges, and publishing/consuming messages
### **Data structure**
```go
type RabbitMQ struct {
    conn *amqp.Connection
}
```
### **Constructor**
```go
func NewRabbitMQ(uri string) (*RabbitMQ, error) {
    conn, err := amqp.Dial(uri)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
    }
    rmq := &RabbitMQ{conn: conn}
    return rmq, nil
}
```
### **Closing the connection**
```go
func (r *RabbitMQ) Close() {
    if r.conn != nil {
        r.conn.Close()
    }
}
```
This method is invoked at the end of each service’s lifecycle, typically using:
```go
defer rabbitmq.Close()
```
---
## **4. Adding the Connection to Each Service**
Each microservice now follows the same initialization pattern:
1. Read the RabbitMQ URI from the environment
2. Initialize the shared messaging component
3. Handle any connection errors
4. Defer the connection close
5. Log successful initialization
### **Example**
```go
rabbitMqURI := env.GetString("RABBITMQ_URI", "amqp://guest:guest@rabbitmq:5672/")
rabbitmq, err := messaging.NewRabbitMQ(rabbitMqURI)
if err != nil {
    log.Fatal(err)
}
defer rabbitmq.Close()
log.Println("Starting RabbitMQ connection")
```
This ensures consistency across all services and prepares the system for publishing and consuming messages.
---
## **5. Verifying Broker Connectivity**
Once both TripService and DriverService integrate this initialization:
- The RabbitMQ dashboard lists active connections for each service
- Tilt logs display successful connection messages
- Kubernetes confirms correct secret injection
This completes the foundational setup required for introducing asynchronous messaging for events such as:
- `trip.created`
- `trip.driver_assigned`
- `trip.accepted`
- and later, payment-related events
## **Publishing the First Message**
### **1. Adding Channel Support to the Messaging Layer**
Publishing and consuming messages in RabbitMQ requires a **channel**. A channel is a lightweight virtual connection inside the main AMQP connection and is the primary API used to declare queues, publish messages, and subscribe to them.
The shared messaging component (`shared/messaging/rabbitmq.go`) is extended to include:
```go
Channel *amqp.Channel
```
The constructor initializes both the AMQP connection and the channel:
```go
ch, err := conn.Channel()
if err != nil {
    conn.Close()
    return nil, fmt.Errorf("failed to create channel: %v", err)
}
```
If channel creation fails, the connection is closed to prevent resource leaks.
The `Close()` method now closes both the connection and channel:
```go
if r.Channel != nil {
    r.Channel.Close()
}
```
This ensures all RabbitMQ resources are properly released when the service shuts down.
---
## **2. Centralized Setup of Queues and Exchanges**
Queue and exchange declarations belong inside the messaging layer, not inside the services. A dedicated setup method is added:
```go
func (r *RabbitMQ) setupExchangesAndQueues() error
```
This method is responsible for creating all required RabbitMQ resources. At this stage, it defines a single queue named `"hello"`:
```go
_, err := r.Channel.QueueDeclare(
    "hello",
    false, // durable
    false, // auto-delete
    false, // exclusive
    false, // no-wait
    nil,   // args
)
```
The queue is declared with `durable=false` intentionally. This allows observing how message persistence works before switching to a durable configuration later in the module.
If any setup step fails:
- All RabbitMQ resources are closed
- The constructor returns an error to the service
This ensures incorrect configurations never leave the system in a half-initialized state.
---
## **3. Creating a Higher-Level Publish Method**
Publishing logic is encapsulated inside the messaging component so that services never interact directly with the AMQP library.
A method is introduced:
```go
func (r *RabbitMQ) PublishMessage(ctx context.Context, routingKey string, message string) error
```
This method wraps:
- `PublishWithContext`
- Routing key
- Message body encoding
- AMQP publishing options
The initial implementation uses:
- The default exchange (`""`)
- A temporary routing key `"hello"`
- Plain text messages
Later chapters extend this to JSON payloads, typed events, and exchange-based routing.
---
## **4. Introducing the Trip Event Publisher**
A dedicated publisher is created inside the Trip Service. It acts as a thin domain-specific wrapper around the shared messaging layer.
Location:
```
services/trip-service/internal/infrastructure/events/trip_publisher.go
```
### Structure
```go
type TripEventPublisher struct {
    rabbitmq *messaging.RabbitMQ
}
```
### Constructor
```go
func NewTripEventPublisher(rabbitmq *messaging.RabbitMQ) *TripEventPublisher
```
### Publishing a Trip Created Event
```go
func (p *TripEventPublisher) PublishTripCreated(ctx context.Context) error {
    return p.rabbitmq.PublishMessage(ctx, "hello", "hello world")
}
```
This keeps the trip-specific publishing logic separate from the messaging implementation.
---
## **5. Integrating the Publisher Into the Trip Service**
The Trip Service initializes the publisher after establishing the RabbitMQ connection:
```go
publisher := events.NewTripEventPublisher(rabbitmq)
```
The gRPC handler receives the publisher as a dependency:
```go
grpc.NewGRPCHandler(grpcServer, svc, publisher)
```
The handler stores and uses it:
```go
type gRPCHandler struct {
    service   domain.TripService
    publisher *events.TripEventPublisher
}
```
When a trip is created, the handler publishes an event:
```go
if err := h.publisher.PublishTripCreated(ctx); err != nil {
    return nil, status.Errorf(codes.Internal, "failed to publish the trip created event: %v", err)
}
```
This cleanly separates:
- Domain logic
- Transport logic
- Messaging infrastructure
---
## **6. Triggering the First Published Message**
Once everything is wired:
1. A user schedules a trip through the UI
2. The Trip Service successfully executes the gRPC `CreateTrip` method
3. The Trip Event Publisher sends `"hello world"` to the `"hello"` queue
4. RabbitMQ confirms receipt
In the RabbitMQ management console:
- The `"hello"` queue appears
- Messages accumulate in the queue
- They remain unacknowledged because no consumer exists yet
Inspecting the queue shows:
- Number of messages
- Message body (`hello world`)
- Message headers
- Delivery status
This verifies message publishing is functioning correctly.
---
## **7. Preparing for Consumption**
At this point:
- The queue exists
- Messages are successfully published
- No consumer is reading them yet
- Messages remain pending and ready
The next step is implementing the consumer inside the Driver Service, which will retrieve and acknowledge the `trip.created` events.
## **Durability of Queues and Messages**
### **1. The Need for Durability**
When messages are published without consumers, they accumulate in the queue, which is expected.
However, in the initial configuration, restarting RabbitMQ (or the entire Kubernetes environment) results in:
- All queues disappearing
- All messages being lost
This happens because both the **queue** and the **messages** were non-durable, and RabbitMQ was running inside a **non-persistent Kubernetes Deployment**.
Durability requires three things:
1. Durable queues
2. Persistent messages
3. Persistent storage on the RabbitMQ pod
---
## **2. Durable Queues**
A queue is durable if it survives broker restarts.
In the messaging setup, durability is enabled by switching:
```go
_, err := r.Channel.QueueDeclare(
    "hello",
    false, // durable (incorrect)
```
to:
```go
_, err := r.Channel.QueueDeclare(
    "hello",
    true, // durable
```
Durable queues must be declared **before** messages are published and cannot be modified afterward.
Because queues are immutable, replacing a queue's durability requires deleting and recreating it.
---
## **3. Persistent Messages**
Messages must also be configured as persistent. This is controlled through the AMQP `DeliveryMode`.
Initial code used transient messages:
```go
amqp.Publishing{
    ContentType: "text/plain",
    Body:        []byte(message),
}
```
This is updated to:
```go
amqp.Publishing{
    ContentType: "text/plain",
    Body:        []byte(message),
    DeliveryMode: amqp.Persistent,
}
```
Persistent messages ensure that RabbitMQ writes them to disk rather than keeping them only in memory.
Both queue and message durability are required together.
A durable queue with transient messages still loses messages on restart.
---
## **4. Kubernetes Deployment vs StatefulSet**
Even with durable queues and persistent messages, RabbitMQ was still losing all data.
The root cause was that RabbitMQ was running as a **Deployment**, which does not preserve pod data between restarts.
Deployments recreate pods with empty filesystems.
To persist any broker data, RabbitMQ must run as a **StatefulSet** with its own persistent volume.
The RabbitMQ manifest is updated accordingly:
### Kind change:
```yaml
kind: StatefulSet
```
### Required service name:
```yaml
serviceName: "rabbitmq"
```
### Persistent storage:
```yaml
volumeClaimTemplates:
  - metadata:
      name: rabbitmq-data
    spec:
      accessModes: ["ReadWriteOnce"]
      resources:
        requests:
          storage: 1Gi
```
### Volume mount inside the container:
```yaml
volumeMounts:
  - name: rabbitmq-data
    mountPath: /var/lib/rabbitmq
```
This ensures:
- Mnesia database files persist
- Queues persist
- Messages persist
Now RabbitMQ behaves like a database with its own persisted state.
---
## **5. Cleaning Up Old Deployments**
Switching from Deployment → StatefulSet requires removing the old pod and service:
```
kubectl delete deployment rabbitmq
kubectl delete service rabbitmq
tilt down
tilt up
```
This forces Kubernetes to recreate RabbitMQ as a StatefulSet with its new persistent volume.
---
## **6. Verifying Persistence**
The durability changes are tested:
1. Publish a trip event
2. Confirm the message exists in the `"hello"` queue
3. Trigger a full RabbitMQ restart
4. Refresh the management UI
If durability is working, the following remain intact:
- **The queue** is still present
- **The message** is still present
After switching to:
- durable queue
- persistent delivery mode
- StatefulSet with persistent volume
both the queue and its messages survive a restart.
---
## **7. Summary**
Durability requires a complete chain:
- **Durable queue** → survives restart
- **Persistent messages** → written to disk
- **StatefulSet + Persistent Volume** → preserves RabbitMQ internal storage
With all three applied, the system now persists events across:
- Pod restarts
- Node restarts
- Full Kubernetes restarts
- Machine reboots
This is critical in production environments where message loss is unacceptable.
## **Consuming Messages from RabbitMQ**
### **1. Overview**
With message publishing and durability established, the next step is to introduce a consumer that can retrieve and process messages from the queue. The Driver Service becomes the first consumer in the system, retrieving trip-related events and reacting to them asynchronously.
The design goal is to create a clean, reusable consumption flow that can support multiple consumers and multiple queues without duplicating logic across services.
---
## **2. Introducing the Consumer Structure**
Each consumer is modeled as its own component. In the Driver Service, a new file defines the logic for receiving trip events:
```
services/driver-service/trip_consumer.go
```
The consumer has access to the shared RabbitMQ abstraction:
```go
type tripConsumer struct {
    rabbitmq *messaging.RabbitMQ
}
```
A constructor initializes the component:
```go
func NewTripConsumer(rabbitmq *messaging.RabbitMQ) *tripConsumer {
    return &tripConsumer{rabbitmq: rabbitmq}
}
```
This structure allows each event type (such as trip creation, trip acceptance, payment events) to define its own consumer with its own business logic.
---
## **3. Consuming Messages Through a Shared Abstraction**
The shared messaging component in `shared/messaging/rabbitmq.go` exposes a high-level method for consuming messages:
```go
func (r *RabbitMQ) ConsumeMessages(queueName string, handler MessageHandler) error
```
### Message Handler Definition
A handler is a function with a unified signature:
```go
type MessageHandler func(context.Context, amqp.Delivery) error
```
This defines how individual messages should be processed.
Each service implements its own handler for its own business logic.
---
## **4. Implementing the Consume Loop**
The consumption workflow is built around the AMQP `Channel.Consume` method:
```go
msgs, err := r.Channel.Consume(
    queueName,
    "",
    true,  // auto-ack
    false, // exclusive
    false, // no-local
    false, // no-wait
    nil,
)
```
The method returns a Go channel of messages.
A goroutine continuously listens for incoming messages:
```go
go func() {
    for msg := range msgs {
        if err := handler(ctx, msg); err != nil {
            log.Fatalf("failed to handle the message: %v", err)
        }
    }
}()
```
### Key characteristics:
- The consumer runs in its own goroutine
- The process is non-blocking
- Each message is delegated to the service-specific handler
- Failures in message handling are logged immediately
- Auto-acknowledgment is enabled for now (later replaced with manual acks)
This design allows multiple consumers to run concurrently, each reacting independently to messages.
---
## **5. Defining the Trip Event Consumer**
Inside the Driver Service, the consumer declares a handler for the `"hello"` queue:
```go
func (c *tripConsumer) Listen() error {
    return c.rabbitmq.ConsumeMessages("hello", func(ctx context.Context, msg amqp091.Delivery) error {
        log.Printf("driver received message: %v", msg)
        return nil
    })
}
```
This handler currently logs the incoming message.
Later, this is extended to include:
- Trip matching logic
- Route computation
- Driver notification via WebSocket
- Integration with domain services
---
## **6. Starting the Consumer**
The consumer is initialized in `driver-service/main.go`:
```go
consumer := NewTripConsumer(rabbitmq)
go func() {
    if err := consumer.Listen(); err != nil {
        log.Fatalf("Failed to listen to the message: %v", err)
    }
}()
```
A goroutine is used so the service can continue serving gRPC requests while listening in the background.
---
## **7. End-to-End Behavior**
With both the publisher and consumer active:
1. A trip triggers the publication of a message to the `"hello"` queue
2. RabbitMQ stores the message durably
3. The Driver Service’s consumer starts up (even after a restart)
4. The message is immediately delivered and processed
5. Messages never get stuck in the queue unless the consumer is unavailable
This demonstrates the asynchronous nature of the system:
- The Trip Service and Driver Service are completely decoupled
- A service can go offline and still receive all pending messages once it restarts
## **Message Delivery, Load Distribution, and Acknowledgments**
### **1. Understanding Message Delivery Across Multiple Instances**
When multiple replicas of the same service consume from the same queue, RabbitMQ distributes messages in a **round-robin** fashion. This is called **competing consumers**, or a **work queue** pattern.
### Behavior:
- Each message is delivered to **exactly one** consumer
- Consumers take turns
- Load is spread evenly across replicas
- This works even if one replica is slower than the others
### Example:
With 2 replicas of Driver Service running:
1. First trip event → Replica A
2. Second trip event → Replica B
3. Third trip event → Replica A
This is the default behavior and is ideal for parallelizing workloads.
### Limitation:
Round-robin dispatch **does not consider service health or load**.
A slow or struggling instance will still receive messages.
RabbitMQ does support more advanced patterns (priority queues, consumer prefetch, dead-lettering, retry queues). This lesson lays the groundwork for those improvements.
---
## **2. The Importance of Acknowledgments**
Up to this point, messages were being consumed with:
```go
auto-ack = true
```
Meaning:
- As soon as RabbitMQ *delivers* a message to the consumer, it is automatically marked as processed.
- If the service crashes while processing the message, the message is **lost permanently**.
This creates a dangerous scenario:
- Work is not completed
- But RabbitMQ and the system *believe it was*
This breaks reliability.
---
## **3. Switching to Manual Acknowledgments**
To avoid losing work, we switch to:
```go
auto-ack = false
```
This ensures:
- A message stays *unacknowledged* until the service explicitly acknowledges it
- If the service crashes or throws an error before sending `Ack`, the message is returned to the queue
### Updated Consume Configuration:
```go
msgs, err := r.Channel.Consume(
    queueName,
    "",
    false, // auto-ack OFF
    false,
    false,
    false,
    nil,
)
```
Now the consumer is responsible for calling:
```go
msg.Ack(false)
```
or
```go
msg.Nack(false, false)
```
---
## **4. The Correct Handling Flow**
The design wraps all consumption and acknowledgment logic inside the messaging abstraction:
```go
for msg := range msgs {
    if err := handler(ctx, msg); err != nil {
        msg.Nack(false, false)
        continue
    }
    msg.Ack(false)
}
```
### Behaviors:
- If handler returns `nil` → message is acknowledged
- If handler returns an error → message is **negatively acknowledged**
- The message is not requeued (for now)
- The loop continues to the next message
- The consumer does not crash from temporary errors
This is a critical improvement:
- The system no longer loses messages
- A faulty consumer no longer acknowledges half-processed work
- Failed messages stay visible and can later be routed into a **dead-letter queue**
---
## **5. Handling Consumer Failures**
When the Driver Service intentionally fails inside its handler:
```go
return errors.New("simulated crash")
```
RabbitMQ behavior:
1. The message is delivered
2. Driver Service fails before acknowledging
3. RabbitMQ returns the message to **Ready** state
4. The consumer restarts
5. The message is delivered again
6. It can then be processed successfully
The system behaves reliably even when consumers crash unexpectedly.
This is essential for:
- Retry logic
- Graceful recovery
- Ensuring no message is lost
- Guaranteeing eventual consistency
---
## **6. Safer Nack Behavior**
Instead of re-queuing failed messages immediately (which can cause infinite redelivery loops), the code uses:
```go
msg.Nack(false, false)
```
Meaning:
- Do not acknowledge
- Do not requeue
- Let dead-letter policies decide what happens next
This is a safe default until a proper **retry strategy** or **dead-letter exchange** is introduced.
---
## **7. Final Behavior Summary**
### **Before**
- Auto-ack = true
- Messages marked as delivered before processing
- Crashes caused message loss
- No error handling
- No requeue or retry
### **After**
- Auto-ack = false
- Messages processed inside a controlled loop
- Explicit `Ack` on success
- Explicit `Nack` on failure
- No message loss
- Faulty messages remain unprocessed
- Consumers do not crash the entire service
## **Fair Dispatching and Smarter Message Distribution**
### **1. The Problem With Default Round-Robin Delivery**
RabbitMQ’s default dispatching strategy is **round-robin**.
When multiple consumers listen to the same queue, RabbitMQ alternates deliveries:
1. Message → Consumer A
2. Message → Consumer B
3. Message → Consumer A
4. Message → Consumer B
    
    …and so on.
    
This produces **balanced distribution**, but it does **not** consider:
- Whether a consumer is busy
- Whether a consumer is slow
- Whether a consumer is stuck or under heavy load
### Example Problem
If Consumer A receives a long-running message and takes 15 seconds, RabbitMQ still forwards new messages to Consumer A (because round-robin doesn’t check whether the consumer is available).
Meanwhile, Consumer B sits idle.
This leads to:
- Higher latency
- Unfair workloads
- Bottlenecks on busy consumers
- Wasted capacity on idle consumers
We want smarter behavior.
---
## **2. Fair Dispatching (Prefetch Count = 1)**
RabbitMQ provides a built-in solution called **fair dispatch**.
It is implemented by configuring **QoS (Quality of Service)** with a `prefetchCount`.
```go
r.Channel.Qos(
    1,     // The consumer can only have ONE unacknowledged message
    0,
    false,
)
```
### What This Solves
Fair dispatch ensures:
- A consumer receives a new message **only after it has acknowledged the previous one**
- Busy consumers do not receive new messages
- Idle consumers receive messages sooner
- Workload is distributed based on availability, not position in the round-robin cycle
This is a **huge improvement** in system throughput when consumers perform variable or heavy work.
---
## **3. How Fair Dispatch Works Internally**
### With round-robin:
- Unacknowledged messages can pile up for the same consumer
- Heavy consumers can get overwhelmed
- Idle consumers wait unnecessarily
### With fair dispatch:
- Each consumer is allowed only **one pending message**
- RabbitMQ examines which consumer has 0 unacknowledged messages
- That consumer receives the next event
- There is no starvation or queue buildup on slow consumers
This converts your system into **a true work queue**, where consumers compete based on availability.
---
## **4. Implementing Fair Dispatch in Code**
In the messaging abstraction:
```go
err := r.Channel.Qos(
    1,     // Prefetch 1 unacked message
    0,
    false,
)
```
If this fails, the consumer fails to initialize.
Fair dispatch is applied **per consumer**, not globally, because `global=false`.
This ensures each driver service instance independently enforces the limit.
---
## **5. Testing the Behavior**
To demonstrate the difference:
### Setup:
- Driver Service scaled to **two replicas**
- The consumer handler was modified to simulate heavy work:
```go
time.Sleep(15 * time.Second)
```
### Observed Result:
- First message → Consumer A
- Consumer A becomes busy immediately
- Second message → Consumer B
- Both consumers enter “busy” state
- Third message stays in queue (in “Ready” state)
- When Consumer A finishes processing, message is delivered to A
- When Consumer B finishes, any remaining messages are delivered to B
### RabbitMQ UI Shows:
- Increasing “Ready” count while consumers are busy
- Two “Unacked” messages during processing
- Once processing completes, “Ready” returns to zero
This confirms the correctness of fair dispatching.
---
## **6. End Result**
Thanks to prefetch=1:
- Consumers now process one message at a time
- No consumer is overloaded
- Idle consumers get tasks immediately
- Work is evenly and **intelligently** distributed
- Latency and throughput improve
- The system performs consistently under load
- The message queue never overwhelms a slow worker
This is exactly how resource-heavy microservices should behave.
## **Exchanges, Routing Keys, and the Publish/Subscribe Model**
### **1. The Limitations of Direct-to-Queue Messaging**
So far, messages have been published **directly to a queue**, and services consume directly from that queue.
This approach works for small systems—but it breaks down as the system grows.
### **Problems With Direct Queue Publishing**
#### 1. **Tight Coupling**
The producer must know:
- The exact queue name
- The intended consumer
- The routing structure
Example:
```go
PublishMessage(ctx, "hello", ...)
```
This hardcodes knowledge that “trip created” must go to queue `"hello"`—which makes the system harder to evolve.
### 2. **Limited Routing**
A queue can be consumed by many consumers, but a **message can only go into one queue**.
If you want to deliver the same message to two consumers, you must:
```
Publish to Queue A
Publish to Queue B
```
This duplicates messages, clutters code, and leads to brittle logic.
### 3. **Not Scalable**
With several services, each consuming different event types, direct queue publishing becomes unmanageable:
- Too many queue names
- Too many publish calls
- No clear abstraction around routing logic
This is why large, event-driven architectures **never publish directly to queues**.
---
## **2. Introducing Exchanges**
RabbitMQ solves these problems using an abstraction called the **exchange**.
### **Core Concept**
A **producer sends a message to an exchange**, not to a queue.
The exchange:
- Receives the message
- Looks at its routing rules
- Delivers the message to one or multiple queues
Producers no longer care about queue names.
### **Routing Flow**
```
Producer  →  Exchange  →  0..N Queues  →  Consumers
```
Producers only specify:
- The exchange name
- The routing key
That’s it.
---
## **3. Why Exchanges Improve the Architecture**
### **Benefit 1 — Loose Coupling**
Producers know only:
- Exchange name (e.g., `"trip.exchange"`)
- Routing key (e.g., `"trip.events.created"`)
They don’t know:
- Which queue receives the message
- Which service processes it
- How many consumers exist
This allows:
- Adding/removing consumers with no producer changes
- Adding new services without touching old ones
- Clear evolvability as the system grows
### **Benefit 2 — Multi-Consumer Delivery**
A single message can go to multiple queues.
Example:
When a **driver is assigned** to a trip:
- The **Trip Service** needs to update the trip state
- The **Payment Service** needs to create a payment session
Instead of publishing the same message twice, the producer publishes once:
```
Exchange: trip.exchange
Routing key: trip.events.driver_assigned
```
Multiple queues bind to this routing key and receive the same event.
### **Benefit 3 — Clean Separation of Responsibilities**
Producers describe what happened.
Exchanges decide where the message goes.
Queues buffer the message.
Consumers decide what to do with it.
This creates:
- Modularity
- Scalability
- Extensibility
---
## **4. Types of Exchanges**
RabbitMQ supports several routing strategies.
We focus only on the ones relevant to real-world systems.
### **A. Direct Exchange**
Routing key must **match exactly**.
```
Routing key: "order.paid"
Binding key: "order.paid"  → receives
Binding key: "order.*"     → ignored (because direct)
```
Useful for:
- One-to-one message routing
- Very specific bindings
### **B. Fanout Exchange**
Delivers messages to **all queues**, ignoring routing keys.
Use cases:
- Broadcast
- System-wide notifications
- Shutdown events
We do **not** need it for this project.
### **C. Topic Exchange (Recommended)**
The most flexible and common.
Supports **wildcards** and **patterns**.
Routing keys are dot-separated words:
```
trip.events.created
trip.events.driver_assigned
payment.commands.process
driver.events.status_changed
```
Consumers bind using patterns:
- `"trip.events.*"` → receives all trip events
- `"*.commands.*"` → receives all commands
- `"trip.#"` → receives all trip-related messages
- `"#"` → receives everything
This model lets different services subscribe to exactly the messages they care about.
### **Why Topic Exchanges Are Ideal**
- Extremely flexible
- Easy to scale
- Clean and expressive naming
- Perfect for microservices
- No need to duplicate messages
- Simple to add new event types
This is the strategy the project will adopt.
---
## **5. Designing Routing Keys**
Routing keys should describe **what happened**.
Recommended pattern:
```
<domain>.<type>.<action>
```
Examples:
```
trip.events.created
trip.events.driver_assigned
payment.commands.create_session
payment.events.session_completed
driver.events.location_updated
```
### Why this works:
- Clear separation of domains
- Easy to filter by topic
- Easy to expand
- Self-documenting
---
## **6. Queue Bindings**
Queues bind to exchanges with a **binding key**—a pattern that describes which messages they want.
Examples:
### Driver Service:
```
Bind to: trip.events.created
Bind to: trip.events.cancelled
```
### Trip Service:
```
Bind to: trip.events.driver_assigned
```
### Payment Service:
```
Bind to: trip.events.driver_assigned
Bind to: payment.commands.create_session
```
This is how one event can flow to multiple services seamlessly.
---
## **7. Real Example for This Project**
Here is exactly what the final architecture will look like:
### **Producer (Driver Service)**
Publishes to:
```
Exchange: trip.exchange
Routing key: trip.events.driver_assigned
```
### **Exchange**
Receives routing key → finds two matching queues:
```
Queue: update-trip-on-driver-assigned
Queue: begin-payment-on-driver-assigned
```
### **Two services receive the same message**
- Trip Service updates the trip with the driver info
- Payment Service starts a payment workflow
Producers publish only **once**, completely unaware of who consumes.
This is the power of topic exchanges.
## **Implementing Exchanges, Topics, and Declarative Queue Bindings**
### **1. Purpose of Moving to Exchanges**
In a growing microservices system, publishing messages directly to queues leads to coupling and limited routing flexibility.
A producer becomes aware of queue names, and sending the same event to multiple services requires publishing multiple times.
To solve this, messages are published to an **exchange**, and queues bind to the exchange using **routing keys**.
This allows:
- Decoupling producers from consumers
- Routing a single event to multiple queues
- Selective message delivery using patterns
- A unified, scalable asynchronous contract between services
---
## **2. Centralized Messaging Contracts**
A shared folder defines a set of **AMQP routing keys**.
These represent the official message “contract” between services.
Example contract constants:
- `TripEventCreated`
- `TripEventDriverNotInterested`
Contracts ensure all services use the same routing keys without string duplication.
---
## **3. Defining the Exchange**
A dedicated exchange is created for all trip-related asynchronous communication.
```go
const TripExchange = "trip"
```
The exchange is declared with type **topic**, enabling flexible wildcard routing.
```go
r.Channel.ExchangeDeclare(
    TripExchange,
    "topic",
    true,   // durable
    false,
    false,
    false,
    nil,
)
```
All trip events flow through this exchange.
---
## **4. Declarative Queue Creation and Binding**
A queue must declare which routing keys it is interested in.
To make this scalable, a helper method encapsulates queue creation and binding logic:
```go
func (r *RabbitMQ) declareAndBindQueue(
    queueName string,
    messageTypes []string,
    exchange string,
) error
```
This function:
1. Declares a durable queue
2. Iterates over a list of routing keys
3. Creates a binding for each key
This pattern allows each queue to listen to multiple event types.
---
## **5. Creating the First Queue**
A queue responsible for locating available drivers is defined:
```go
const FindAvailableDriversQueue = "find_available_drivers"
```
This queue is bound to the routing keys:
- `trip.events.created`
- `trip.events.driver_not_interested`
Meaning it receives:
- Newly created trip events
- Events indicating a driver rejected a trip
These bindings are set during RabbitMQ initialization:
```go
r.declareAndBindQueue(
    FindAvailableDriversQueue,
    []string{
        contracts.TripEventCreated,
        contracts.TripEventDriverNotInterested,
    },
    TripExchange,
)
```
This establishes a clear and descriptive routing relationship.
---
## **6. Rewriting the Trip Publisher**
Messages are now published to the **exchange**, not directly to queues.
```go
p.rabbitmq.PublishMessage(
    ctx,
    contracts.TripEventCreated,
    "Trip has been created",
)
```
Internally:
```go
Channel.PublishWithContext(
    ctx,
    TripExchange,
    routingKey,
    ...
)
```
The routing key determines which queues receive the message.
---
## **7. Updating the Consumer to Subscribe to a Queue**
The driver service listens on the queue related to finding drivers:
```go
rabbitmq.ConsumeMessages(
    messaging.FindAvailableDriversQueue,
    handler,
)
```
Any event whose routing key matches this queue’s bindings is delivered to this consumer.
---
## **Publishing and Consuming JSON Messages (Structured AMQP Payloads)**
### **1. Why Move Beyond String Messages**
Simple string payloads are not sufficient for a real microservices system.
Services must exchange **rich, structured data**, such as the full trip information (`TripModel`).
This requires:
- JSON encoding before publishing
- JSON decoding inside consumers
- A shared, predictable contract for message structure
- A typed payload model for each event
This lesson introduces a **unified AMQP message format** and JSON-based serialization for all future event flow.
---
## **2. Standard AMQP Message Contract**
A shared struct defines the envelope for every published event:
```go
type AmqpMessage struct {
    OwnerID string
    Data    []byte
}
```
This wrapper provides:
- `OwnerID` → identifies who triggered the event
- `Data` → raw JSON payload of the event-specific body
All services publish and consume **the same top-level structure**.
---
## **3. Standardizing the Payload Type**
Each event has its own typed payload structure.
For trip creation events:
```go
type TripEventData struct {
    Trip *pb.Trip `json:"trip"`
}
```
Using `pb.Trip` ensures:
- A globally shared type
- No dependency on domain-layer structs from other services
- A consistent schema across the system
The TripModel is converted to `*pb.Trip` using a helper:
```go
func (t *TripModel) ToProto() *pb.Trip { ... }
```
---
## **4. Publishing JSON Messages**
The trip publisher now builds a strongly typed payload:
```go
payload := messaging.TripEventData{
    Trip: trip.ToProto(),
}
```
The payload is JSON-encoded:
```go
tripEventJSON, _ := json.Marshal(payload)
```
This JSON is wrapped inside the shared envelope:
```go
contracts.AmqpMessage{
    OwnerID: trip.UserID,
    Data:    tripEventJSON,
}
```
Finally, the full message is published through RabbitMQ:
```go
p.rabbitmq.PublishMessage(ctx, contracts.TripEventCreated, amqpMessage)
```
The `PublishMessage` method itself serializes the envelope into JSON before sending.
---
## **5. JSON Serialization Inside RabbitMQ Publisher**
`PublishMessage` now receives a typed `AmqpMessage` and marshals it:
```go
jsonMsg, _ := json.Marshal(message)
```
This JSON becomes the AMQP body:
```go
amqp.Publishing{
    Body: jsonMsg,
}
```
All published messages are now structured, typed, and encoded using the shared AMQP contract.
---
## **6. Consuming JSON Messages**
The consumer performs two unmarshalling steps:
### **Step 1 — Decode the AMQP envelope**
```go
var envelope contracts.AmqpMessage
json.Unmarshal(msg.Body, &envelope)
```
Now the service has access to:
- `envelope.OwnerID`
- `envelope.Data` (raw JSON payload)
### **Step 2 — Decode the typed payload**
```go
var payload messaging.TripEventData
json.Unmarshal(envelope.Data, &payload)
```
### **Result**
The driver service now receives:
```go
payload.Trip
```
A rich `*pb.Trip` object containing:
- Trip ID
- User ID
- Fare info
- Price
- Route
- Driver (nil until assigned)
---
## **7. Benefits of the Structured JSON Messaging System**
### **Consistency**
Every message uses the same top-level envelope.
### **Interoperability**
Any service can consume events simply by:
1. Decoding the envelope
2. Decoding the typed payload
### **Type Safety**
Payloads are defined explicitly (e.g., `TripEventData`).
### **Scalability**
New event types only require:
- A new payload struct
- A new routing key
- A binding to the exchange
### **RabbitMQ-Friendly**
JSON keeps messages human-readable in the RabbitMQ UI.
---
## **8. Full Flow Summary**
1. **Trip is created** → TripModel generated
2. TripModel → **converted to protobuf**
3. Protobuf trip → **marshaled to JSON payload**
4. JSON payload → **wrapped in AmqpMessage**
5. AmqpMessage → **marshaled and published**
6. Consumer receives Message → **decodes envelope**
7. Envelope.Data → **decoded into event-specific struct**
8. Consumer logic executes with a rich `pb.Trip` object
---
If you want, I can continue with the next lesson and produce the notes in the same structured, book-style format.
## **Finding & Notifying the Suitable Driver (Full Event-Driven Flow)**
### **1. Goal of This Lesson**
This lesson completes the *first major async workflow* in the system:
### **Trip Created → Driver Service → Find Driver → Publish Events**
We already:
- Published **TripEventCreated**
- Consumed it inside the **Driver Service**
- Decoded JSON into strong types
Now we:
1. **Select a suitable driver** based on the incoming trip request.
2. **Publish two possible follow-up events**:
    - `TripEventNoDriversFound`
    - `DriverCmdTripRequest`
3. Prepare the system for real-time WebSocket notifications (next module).
---
## **2. Adding Driver Lookup Logic**
The Driver Service already keeps an **in-memory registry of connected drivers**:
```go
drivers map[string]*DriverModel
```
We add a lightweight matching function:
```go
func (s *Service) FindAvailableDrivers(packageType string) []string {
    var matchingDrivers []string
    for _, driver := range s.drivers {
        if driver.Driver.PackageSlug == packageType {
            matchingDrivers = append(matchingDrivers, driver.Driver.Id)
        }
    }
    if len(matchingDrivers) == 0 {
        return []string{}
    }
    return matchingDrivers
}
```
### What it does
- Loops through registered drivers
- Picks those offering the **same package** as the trip
- Returns **a list of matching driver IDs**
- For now, we use the first driver only
    
    (later you can improve to closest driver, ranking, etc.)
    
---
## **3. Understanding the Consumer Logic**
The driver-side consumer listens to:
- `TripEventCreated`
- `TripEventDriverNotInterested`
Both should trigger the **same handler**, because both mean:
> “We need to find a driver for this trip.”
> 
This is handled via a `switch` on the routing key:
```go
switch msg.RoutingKey {
case contracts.TripEventCreated, contracts.TripEventDriverNotInterested:
    return c.handleFindAndNotifyDrivers(ctx, payload)
}
```
Anything else → logged as unknown event.
---
## **4. Handling the Trip Event**
### **4.1 Find Suitable Drivers**
```go
suitableIDs := c.service.FindAvailableDrivers(payload.Trip.SelectedFare.PackageSlug)
```
### **4.2 Case 1: No drivers found**
Publish:
```go
TripEventNoDriversFound
```
This event is consumed later (WebSocket → notify rider).
```go
c.rabbitmq.PublishMessage(ctx, contracts.TripEventNoDriversFound, contracts.AmqpMessage{
    OwnerID: payload.Trip.UserID,
})
```
### **4.3 Case 2: Driver found**
Pick the first driver:
```go
suitableDriverID := suitableIDs[0]
```
Prepare payload:
```go
marshalledEvent, _ := json.Marshal(payload)
```
Publish:
```go
c.rabbitmq.PublishMessage(ctx, contracts.DriverCmdTripRequest, contracts.AmqpMessage{
    OwnerID: suitableDriverID,
    Data:    marshalledEvent,
})
```
This event means:
> “Driver, someone is requesting a trip — check the route & accept/decline.”
> 
In the *next module*, this will appear instantly on the driver’s mobile UI.
---
## **5. Why We Use `OwnerID`**
Events always include `OwnerID`:
- If notifying the **rider**, OwnerID = rider’s ID
- If notifying the **driver**, OwnerID = driver’s ID
The **API Gateway WebSocket layer** routes messages to the right connected user:
```
AMQP message → API Gateway → WebSocket connection → Specific user ID
```
This enables:
- Real-time driver popup (“New trip request”)
- Real-time rider popup (“No drivers available”)
---
## **6. Full Flow Overview**
### **1️⃣ Trip created**
Trip Service publishes:
```
trip.events.created
```
### **2️⃣ Driver Service consumes the event**
Consumes `FindAvailableDriversQueue`.
### **3️⃣ Driver Service identifies drivers**
Compares:
```
driver.PackageSlug == trip.PackageSlug
```
### **4️⃣ Publish next event**
#### **Case A — No drivers available**
Publish:
```
trip.events.no_drivers_found
OwnerID = trip.UserID
```
### **Case B — Driver found**
Publish:
```
driver.cmd.trip.request
OwnerID = foundDriverID
Data = full trip payload (JSON)
```
## **Real-time Notifications**
### **Notifications Module – WebSocket & RabbitMQ Integration**
#### **1. Purpose of the Module**
The goal of this module is to complete the **real-time communication loop** between:
- **Backend microservices** (Trip Service, Driver Service)
- **API Gateway**
- **Browser clients** (Driver UI, Rider UI)
This module enables:
- Drivers to receive *trip requests* instantly
- Riders to receive *status updates* instantly
- Services to react to driver actions (accept/decline)
Everything is now connected through **RabbitMQ → API Gateway → WebSockets**.
---
### **2. Current State Before Notifications**
Up until now, the event flow looks like this:
1. **Trip Service**
    
    Publishes → `trip.events.created`
    
2. **Driver Service**
    
    Consumes → Finds a suitable driver → Publishes:
    
    - `driver.cmd.trip.request`
    - or `trip.events.no_drivers_found`
So far everything is internal.
**But none of this reaches the actual users yet.**
---
### **3. The Problem We Must Solve**
Services can talk to each other (via RabbitMQ).
But users (driver & rider) only exist on the **frontend**, connected via WebSockets to the **API Gateway**.
So the missing piece:
> How do we deliver RabbitMQ events to the correct user’s WebSocket connection?
> 
We need a bridge:
```
RabbitMQ → API Gateway → Specific WebSocket connection
```
---
### **4. WebSocket Delivery Requirements**
To send a message from API Gateway → browser, we need:
1. The **user’s ID**
    
    (Already included in each event as `OwnerID`)
    
2. The **user’s active WebSocket connection**
    
    (Not stored yet — we must fix this)
    
3. The **ability to route messages to the correct connection**
    
    (This will be implemented in this module)
    
---
### **5. Storing WebSocket Connections**
When a user (driver or rider) connects to the UI, the WebSocket handler receives:
```go
type Client struct {
    ID string         // user ID
    Conn *websocket.Conn
}
```
We must store these connections somewhere.
### **Why?**
Because when an event targeting that user arrives from RabbitMQ, we must do:
```go
connections[userID].WriteJSON(eventPayload)
```
---
### **6. Where to Store Connections**
We face a trade-off.
### **Simplest approach (used in course):**
Store connections **in memory inside the API Gateway**:
```go
map[string]*websocket.Conn
```
Good:
- Easy
- No extra infrastructure
Bad:
- ❌ Gateway cannot be scaled horizontally
- ❌ Losing the pod = losing all live connections
### **Real-world solution (recommended for production):**
Use **Redis** to store active user → gateway instance mappings.
But the course keeps it simple:
**We store WebSocket connections in memory.**
---
### **7. Event Routing Logic**
Every RabbitMQ event contains:
```go
OwnerID string
Data    []byte
```
So the API Gateway listens to RabbitMQ:
```
driver.cmd.trip.request → OwnerID = driverID
trip.events.no_drivers_found → OwnerID = riderID
```
The Gateway then:
1. Looks up the WebSocket connection by OwnerID
2. Sends the JSON message through that WebSocket
This enables:
- Drivers receiving **“New Trip Request”** popups
- Riders receiving **“No drivers found”** messages
- Later: Riders receiving **“Driver accepted”**
- Later: Drivers receiving **“Trip payment link”**
---
### **8. Reverse Direction: Browser → Service**
The flow must also work the other way:
When a driver clicks:
```
Accept trip
```
The frontend emits a WebSocket event:
```
driver.cmd.trip.accepted
```
The API Gateway:
- Receives this WebSocket message
- Converts it into a RabbitMQ event
- Publishes to the exchange
- Trip Service consumes it
Thus completing the two-way communication path.
---
### **9. Full Architecture After This Module**
```
             ┌─────────────────────┐
             │     Trip Service     │
             └───────▲──────────────┘
                     │
               AMQP Events
                     │
┌─────────────┐   ┌──┴─────────┐   ┌───────────────┐
│  Driver UI   │⇄ │ API Gateway │ ⇄│  RabbitMQ       │
└─────────────┘   └──┬─────────┘   └───────────────┘
                     │
               WebSocket Events
                     │
             ┌───────▼────────────┐
             │    Rider UI         │
             └─────────────────────┘
```
Both users and services communicate through **the same central routing system**, but using different protocols:
- Services ↔ RabbitMQ
- Users ↔ WebSockets
The **API Gateway sits in the middle**.
## **WebSocket Connection Manager**
### **1. Purpose of the Connection Manager**
This lesson introduces a core building block required for real-time communication:
### **The Connection Manager**
It is responsible for:
- Tracking all active WebSocket connections
- Associating each connection with a **user ID**
- Sending typed messages back to the correct user
- Handling concurrent access with proper locking
- Providing a safe abstraction around WebSocket operations
This enables the API Gateway to function as a **message router** between:
- RabbitMQ events → specific users
- WebSocket events → RabbitMQ
---
### **2. The Connection Manager Structure**
A new file was added under `shared/messaging`:
```go
connection_manager.go
```
It contains:
### **A thread-safe structure:**
```go
type ConnectionManager struct {
    connections map[string]*connWrapper
    mutex       sync.RWMutex
}
```
### **A wrapper around the actual WebSocket connection:**
```go
type connWrapper struct {
    conn  *websocket.Conn
    mutex sync.Mutex
}
```
This is required because **WebSocket writes are not thread-safe**.
---
### **3. What the Connection Manager Provides**
#### **Upgrade()**
Replaces repeated gorilla upgrader code:
```go
connManager.Upgrade(w, r)
```
### **Add()**
Store new user → WebSocket connection:
```go
connManager.Add(userID, conn)
```
### **Remove()**
Clean up when client disconnects:
```go
connManager.Remove(userID)
```
### **Get()**
Optional lookup:
```go
connManager.Get(userID)
```
### **SendMessage()**
Unified JSON send API:
```go
connManager.SendMessage(userID, contracts.WSMessage{ ... })
```
This allows any component (including the upcoming **RabbitMQ consumer**) to route messages to the correct online user.
---
### **4. Integrating Connection Manager Into API Gateway**
The gateway now holds a **global instance**:
```go
var connManager = messaging.NewConnectionManager()
```
This allows all WebSocket handlers to reuse it.
---
### **5. Updating the Riders WebSocket Handler**
#### **Old flow:**
- Manually upgrade connection
- Store nothing
- Could not route messages to specific user
### **New flow:**
**Upgrade:**
```go
conn, err := connManager.Upgrade(w, r)
```
**Store the connection:**
```go
connManager.Add(userID, conn)
defer connManager.Remove(userID)
```
This enables targeted notifications:
```
SendMessage(userID, ...)
```
---
### **6. Updating the Drivers WebSocket Handler**
Same improvements as riders, plus:
### **Better cleanup:**
On disconnect:
- Removes WebSocket connection
- Unregisters driver from driver service
- Closes gRPC client
### **Driver registration now uses SendMessage()**
Instead of manually doing:
```go
conn.WriteJSON(...)
```
We now do:
```go
connManager.SendMessage(userID, contracts.WSMessage{
    Type: contracts.DriverCmdRegister,
    Data: driverData.Driver,
})
```
This unifies all gateway → client messaging.
---
### **7. Benefits of This Abstraction**
#### **1. Reusable**
All services can now send messages to users through a single API.
### **2. Safer**
- Mutex protected map
- Per-connection mutex to prevent concurrent writes
- No repeated upgrader code
### **3. Scalable design**
While this implementation stores connections in memory, the abstraction makes it easy to:
- Replace map storage with Redis
- Allow horizontal gateway scaling
- Persist connections across gateway instances
(The course keeps in-memory for simplicity.)
## Queue Consumer
### Overview
To complete the real-time driver–rider interaction, the API Gateway must begin consuming RabbitMQ queues and forward those messages directly to the appropriate WebSocket client. This enables two essential flows:
- **Driver receives trip requests** in real time.
- **Rider receives updates** when no driver is available or when a driver is assigned.
This step connects RabbitMQ → API Gateway → WebSocket clients.
### Key Problem
Internal services (Trip Service and Driver Service) publish events to RabbitMQ, but **users** (drivers and riders) live outside the microservice ecosystem, connected only through WebSockets in the API Gateway.
Therefore:
- Services publish events → RabbitMQ.
- Gateway consumes those events → identifies the target user (OwnerID).
- Gateway pushes the message → correct WebSocket connection.
To implement this, the gateway needs:
1. A **Connection Manager** to store all active WebSocket connections.
2. A **Queue Consumer** to receive RabbitMQ events and forward them to users.
3. A way to **initialize queue consumers inside WebSocket handlers**.
### Connection Manager
The Connection Manager is responsible for managing user WebSocket sessions.
It provides:
- **Upgrade**: wraps WebSocket upgrade logic.
- **Add**: store a WebSocket connection with user ID as key.
- **Remove**: delete a connection when socket closes.
- **Get**: fetch a connection by user ID.
- **SendMessage**: send a JSON message to a specific user safely.
Connections are stored in a map and guarded by mutexes for thread safety.
### Queue Consumer
The Queue Consumer listens to a specific RabbitMQ queue and forwards messages to users.
Flow:
1. Consume messages from a queue.
2. Unmarshal them into `AmqpMessage`.
3. Extract:
    - `OwnerID` → the user who should receive the notification
    - `Data` → the event payload
4. Build a WebSocket message:
    - `Type = msg.RoutingKey`
    - `Data = payload`
5. Send the message using the Connection Manager.
This allows *any service* to notify *any user* through RabbitMQ → Gateway → WebSocket.
### Integrating Queue Consumers Into the API Gateway
Inside each WebSocket handler (drivers and riders), the gateway now:
1. Upgrades the WebSocket connection.
2. Extracts the authenticated user ID.
3. Registers the connection in the Connection Manager.
4. Starts consuming one or more queues relevant to that user type.
5. For each message in those queues, forwards the event to the correct user.
Example for drivers:
- The gateway listens to `driver_cmd_trip_request` queue.
- When a driver is selected for a trip:
    - Driver Service publishes `driver.cmd.trip_request`.
    - API Gateway receives it.
    - Sends the message over the driver’s WebSocket.
### Binding the New Queue in RabbitMQ
The gateway needs a queue that receives trip-request commands meant for drivers.
A new queue is added:
- **driver_cmd_trip_request**
It is bound to the routing key:
- **driver.cmd.trip_request**
This ensures that any service publishing this event causes RabbitMQ to forward the event to the driver-specific queue, which the gateway listens to.
### End-to-End Result
Once the gateway starts consuming RabbitMQ events and forwarding them through WebSocket connections:
- When a rider creates a trip, the Driver Service selects the most suitable driver and publishes a trip request event.
- The API Gateway receives this event through the `driver_cmd_trip_request` queue.
- The gateway forwards the trip preview and route details to the driver’s WebSocket.
- The driver sees a real-time modal prompting them to accept or decline.
- When the driver responds, the gateway will publish their acceptance/decline back to RabbitMQ.
- The Trip Service will then continue the booking process.
This completes the cross-service real-time notification loop required for matching riders and drivers.
## Handling Incoming Driver Messages in the API Gateway
### Overview
At this stage of the system, drivers are already receiving trip requests via WebSocket (after the Driver Service publishes a trip-event command through RabbitMQ). Now the gateway must handle **driver replies** — specifically **trip accept** and **trip decline** actions — and forward them back into RabbitMQ so that backend services can continue the flow.
This connects the last missing piece of the request/response loop:
Driver UI → WebSocket → API Gateway → RabbitMQ → Trip Service.
Additionally, the gateway must notify riders when no drivers are available, which now becomes part of the WebSocket consumption logic on the rider side.
### Parsing Driver WebSocket Messages
Driver messages arrive as raw bytes over WebSocket.
The gateway now:
1. Defines a small internal struct:
    - `type` → identifying the action (e.g., `"driver.cmd.trip_accept"`)
    - `data` → raw message payload to forward into RabbitMQ
2. Unmarshals the incoming bytes into this struct.
3. Switches behavior based on the message type:
    - `driver.cmd.location` → currently ignored (reserved for later real-time location tracking)
    - `driver.cmd.trip_accept` and `driver.cmd.trip_decline` → forwarded to RabbitMQ
Only accepted message types are forwarded; all unknown types are logged for debugging.
This isolates the gateway from caring about internal logic.
The gateway simply forwards structured driver intents.
### Publishing Driver Responses Back to RabbitMQ
Once a driver accepts or declines a trip, the gateway wraps the information into an `AmqpMessage`:
- `OwnerID` → the driver’s user ID (useful for backend correlation)
- `Data` → the JSON payload sent from the frontend
Then the gateway publishes the message using the driver’s message type as the routing key.
This means that the Trip Service can simply listen for:
- `driver.cmd.trip_accept`
- `driver.cmd.trip_decline`
…without caring about WebSocket details.
### Adding New Queues for Backend Consumers
Two new queues are introduced and bound:
1. **driver_trip_response**
    - Handles:
        - `driver.cmd.trip_accept`
        - `driver.cmd.trip_decline`
    - Consumed by the Trip Service in the next stage.
2. **notify_driver_no_drivers_found**
    - Handles:
        - `trip.event.no_drivers_found`
    - Consumed by the rider’s WebSocket handler to notify the customer that no drivers were available.
Both queues are declared and bound in `setupExchangesAndQueues` so RabbitMQ routes messages correctly.
### Consuming "No Drivers Found" for Riders
On the rider WebSocket handler, the gateway now starts a queue consumer for:
- `notify_driver_no_drivers_found`
When the Driver Service determines that no driver matches the trip request, it publishes:
- Routing key: `trip.event.no_drivers_found`
- OwnerID: the rider’s user ID
The queue consumer:
1. Reads this event from RabbitMQ
2. Creates a `WSMessage`
3. Sends it to the rider’s WebSocket connection
This produces the correct UI behavior where the rider sees an immediate notification that no drivers were available.
### Resulting Flow
#### When a driver accepts or declines:
1. Driver clicks an action in the UI.
2. WebSocket sends `{type, data}` to the API Gateway.
3. Gateway parses the message and publishes to RabbitMQ.
4. Trip Service consumes the message and continues the business flow (handled next).
### When no drivers are found:
1. Driver Service publishes `trip.event.no_drivers_found`.
2. RabbitMQ routes the event to `notify_driver_no_drivers_found`.
3. API Gateway’s queue consumer reads it.
4. Gateway pushes the message to the rider’s WebSocket.
5. Rider sees UI notification.
This completes real-time, cross-service communication for both positive and negative outcomes.
## Listening for Trip Accept event
### What This Module Achieves
The goal of this module is to complete the **driver decision flow** by having the **Trip Service** consume the *accept* and *decline* events that originate from:
Driver UI → WebSocket → API Gateway → RabbitMQ → Trip Service
Once the Trip Service consumes these events, it must:
- Validate the trip exists
- Update the trip state (status + driver assignment)
- Notify the rider that a driver has been assigned
- (Later) trigger payment workflow
This closes the loop for real-time trip confirmation.
### The New Driver Response Consumer
A new consumer is added inside the Trip Service.
It listens to the `driver_trip_response` queue, which is bound to:
- `driver.cmd.trip_accept`
- `driver.cmd.trip_decline`
This consumer:
1. Unmarshals the `AmqpMessage` envelope
2. Extracts the `DriverTripResponseData` payload
3. Switches based on routing key (accept vs decline)
4. Executes the appropriate business logic
It parallels the structure of the existing trip consumer, but tailored for driver decisions.
### DriverTripResponseData Structure
A new globally shared structure is introduced to decode the driver’s decision:
- `Driver` → protobuf driver model
- `TripID` → the trip being accepted/declined
- `RiderID` → rider who requested the trip (used later for notifications)
This matches the shape of the message sent from the frontend via the gateway.
### Handling Trip Acceptance
The acceptance flow is the important path:
1. **Fetch the trip**
    
    Uses `GetTripByID`.
    
    If not found → return error.
    
2. **Update trip**
    
    Calls `UpdateTrip(tripID, "accepted", driver)`
    
    This updates:
    
    - trip status
    - trip driver details (ID, name, car plate, picture)
3. **Re-fetch the trip**
    
    Ensures the updated state is correct before broadcasting.
    
4. **Publish driver-assigned event**
    
    Sends `trip.event.driver_assigned` to notify the rider.
    
    This event is delivered to the Rider’s WebSocket via the gateway.
    
This is the event that triggers the UI panel:
**“Your driver is on the way.”**
### Handling Trip Decline
The decline branch is intentionally simple:
- Log it
- Return nil
In a production system, additional logic might include:
- reassigning another driver
- retrying algorithms
- notifying the rider
But this is left intentionally incomplete for clarity.
### TripService Additions
The Trip Service interface and implementation gain two new methods:
- `GetTripByID`
- `UpdateTrip`
These are implemented in both:
- the service layer
- the in-memory repository
`UpdateTrip` also converts the protobuf driver into the local trip driver struct.
### Binding Driver Assigned Events
A new queue is introduced:
- `notify_driver_assign`
It is bound to:
- `trip.event.driver_assigned`
The API gateway rider WebSocket handler subscribes to this queue so it can push the event live to the rider UI.
### New Runtime Flow
#### On trip acceptance:
1. Driver accepts → WebSocket → gateway
2. Gateway publishes message to RabbitMQ
3. Trip Service pulls from `driver_trip_response`
4. Trip Service:
    - validates trip
    - updates status
    - sets driver
    - publishes `trip.event.driver_assigned`
5. RabbitMQ delivers event to `notify_driver_assign`
6. Gateway pushes update to rider’s WebSocket
7. Rider sees **driver assigned** UI update
This completes the trip confirmation lifecycle.
If you want, I can now generate notes for the next module (payments + Stripe).
## Trip Decline Flow
### Purpose of This Flow
The decline flow enables the system to react when a driver rejects a trip request. Instead of failing the request entirely, the system automatically attempts to find another suitable driver using the exact same logic as the original trip-created event.
### How Declines Are Processed
1. The driver clicks **Decline** in the UI.
2. The UI sends a WebSocket message to the API Gateway.
3. The API Gateway publishes the message to RabbitMQ using the routing key
    
    **driver.cmd.trip_decline**.
    
4. The **Trip Service** listens to the `driver_trip_response` queue.
5. When it receives a decline:
    - It loads the trip from storage.
    - It republishes a *new* event to RabbitMQ triggering the driver-selection logic again.
### Why Republishing Works
The Trip Service republishes the message under the routing key
**trip.event.driver_not_interested**.
This routing key is already bound to the same queue used by the *original* Trip Created event.
That queue is processed by the **Driver Service → Trip Consumer**, which performs driver matching.
Therefore, publishing this event seamlessly re-triggers:
- `FindAvailableDrivers`
- Randomized driver choice
- Notification to whichever driver is selected next
Exactly the same logic used for the first driver.
### Randomizing Driver Selection
Prior to this commit, the system always picked the first matching driver.
This caused a bug:
- If Driver A and Driver B were both suitable
- And Driver A always handled messages first
- Declining would always return the trip to Driver A
    
    → Driver B would never receive the request
    
The fix introduces:
```
rand.Intn(len(suitableIDs))
```
This selects a **random driver** from the suitable list.
Now requeueing a declined trip fairly redistributes it across all available drivers.
### Decline Handler Logic (Trip Service)
The decline handler performs three core steps:
1. **Fetch Trip**
    - Ensures the trip exists before requeueing.
2. **Build the payload**
    - Wraps the trip inside `TripEventData`
    - Marshals to JSON
3. **Republish message**
    - Sends `trip.event.driver_not_interested`
    - Owner is set to the **rider ID**
    - Data contains the trip JSON payload
This cleanly restarts the driver-search workflow without duplicating logic.
### End-to-End Decline Scenario
1. Rider requests SUV trip
2. Driver A receives it
3. Driver A clicks **Decline**
4. Trip Service republishes driver_not_interested
5. Driver Service re-runs matching
6. Driver B now receives the trip
7. Either Driver B…
    - Accepts → Trip assigned
    - Declines → The cycle repeats again
## Payment Service Setup
### Purpose of This Module
This module introduces the final component of the project: the **Payment Service**.
Its role is to handle payment session creation, integrate with Stripe, react to trip events, and complete the end-to-end flow of scheduling, accepting, and paying for a trip.
Because all previous modules already taught you how to scaffold microservices, this module provides the Payment Service **pre-built**, and your job is simply to connect it into the system.
### Provided Resources
The module includes the entire starter code for:
- `Tiltfile` integration
- Docker build configuration
- Kubernetes Deployment and Service
- Environment variables
- Domain, service, and type definitions
- Main service bootstrap with graceful shutdown
- RabbitMQ integration hooks
These files should be copy-pasted directly into your project from the resources.
### Service Structure Overview
#### **Domain Layer**
Defines the abstractions used throughout the Payment Service:
- **Service interface**
    - `CreatePaymentSession(ctx, tripID, userID, driverID, amount, currency)`
- **PaymentProcessor interface**
    - An adapter used to plug Stripe (or PayPal, etc.)
    - Useful for mocking during unit testing
This separation allows Stripe logic to live in an isolated implementation file.
### **Service Implementation**
The internal service (`paymentService`) orchestrates:
- Generating metadata (trip ID, user ID, driver ID)
- Calling the injected `PaymentProcessor`
- Creating a `PaymentIntent` with necessary fields:
    - TripID
    - UserID
    - DriverID
    - Amount
    - Currency
    - StripeSessionID
    - Timestamps
This `PaymentIntent` will later be forwarded through RabbitMQ to notify other services.
### **Types Package**
Defines:
- Payment statuses:
    - pending
    - success
    - failed
    - cancelled
- Payment model (used for DB storage)
- PaymentIntent model (used for initiating sessions)
- PaymentConfig (Stripe keys + URLs)
These ensure type-safety across the payment flow.
### Configuration
#### **Environment Variables**
`app-config.yaml` provides:
- `STRIPE_SUCCESS_URL`
- `STRIPE_CANCEL_URL`
Deployment expects:
- `STRIPE_SECRET_KEY`
- `RABBITMQ_URI`
If `STRIPE_SECRET_KEY` is missing, the service exits immediately.
This ensures your Stripe configuration must be present before the system boots.
### Infrastructure Integration
#### **Kubernetes Deployment**
The Payment Service is fully configured with:
- Pod definition
- Resource requests/limits
- Environment variable wiring
- ClusterIP service (port 9004)
### **Tilt Development Workflow**
The `Tiltfile` is updated to:
- Build the service
- Enable live update
- Sync shared code and compiled binaries
- Deploy to Kubernetes
- Wait for RabbitMQ as a dependency
This matches the pattern used for all existing services.
## Stripe Integration and API Keys
### Overview
To enable the Payment Service, we must integrate Stripe. Stripe requires two sets of API keys:
1. **Secret Key (backend only)**
2. **Publishable Key (frontend only)**
Stripe provides **test keys** that work even without an account. These are safe to use for development.
### Stripe Test Keys
Stripe exposes public documentation containing:
- `sk_test_...` → secret key used by the backend
- `pk_test_...` → publishable key used by the frontend
These keys allow you to:
- Create checkout sessions
- Simulate payments
- Test redirects
- Inspect webhooks later
### Adding the Secret Key to Kubernetes
#### Step 1: Create the Kubernetes Secret
Create a new secret under `infra/development/k8s/secrets.yaml`:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: stripe-secrets
type: Opaque
data:
  stripe-secret-key: <base64_encoded_sk_test_key>
```
Notes:
- Replace `<base64_encoded_sk_test_key>` with your Base64-encoded Stripe secret.
- Use the exact key name `stripe-secret-key` because the deployment expects this.
### Step 2: Activate the Stripe Secret in the Deployment
In `payment-service-deployment.yaml`, uncomment the Stripe key section:
```yaml
- name: STRIPE_SECRET_KEY
  valueFrom:
    secretKeyRef:
      name: stripe-secrets
      key: stripe-secret-key
```
This injects the secret into the container.
### Step 3: Verify the Payment Service Boots
The Payment Service’s `main.go` validates:
```go
if stripeCfg.StripeSecretKey == "" {
    log.Fatalf("STRIPE_SECRET_KEY is not set")
}
```
If the secret is missing, the service exits.
After adding the secret, running:
```
tilt up
```
should show:
```
Starting RabbitMQ connection
Payment Service running...
```
### Adding the Publishable Key to the Frontend
#### Step 1: Create `.env.local` inside the web folder
```
web/.env.local
```
### Step 2: Add the publishable key
```
NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=pk_test_...
```
Note:
- Must start with `NEXT_PUBLIC_` or Next.js will **not** expose it to the browser.
## Stripe Payment Processor Integration
### Overview
This lesson introduces the **Stripe payment processor implementation** inside the Payment Service. The goal is to implement the backend logic that creates Stripe Checkout Sessions, using the Stripe SDK via a clean **PaymentProcessor interface**.
The Payment Service already contains:
- A domain layer defining the `PaymentProcessor` and `Service` interfaces
- A `paymentService` implementing the main business logic
- A `PaymentConfig` struct for Stripe configuration
- A starting `main.go` that loads config, initializes Stripe, and sets up RabbitMQ
This lesson focuses on implementing the **Stripe-backed PaymentProcessor**.
### Stripe Client Structure
A new folder is created:
```
services/payment-service/internal/infrastructure/stripe/
```
Inside, `stripe.go` implements the `PaymentProcessor` interface.
### Stripe Client Initialization
Before a Stripe client can create sessions, the Stripe SDK requires:
- Setting the global API key
- Selecting the correct version (`v81` in this project)
In the constructor:
```go
stripe.Key = config.StripeSecretKey
```
This allows the Stripe SDK to issue authenticated requests.
### Stripe Client Definition
```go
type stripeClient struct {
    config *types.PaymentConfig
}
```
This struct holds the Stripe configuration (secret key, success URL, cancel URL).
The constructor returns a `domain.PaymentProcessor`:
```go
func NewStripeClient(config *types.PaymentConfig) domain.PaymentProcessor {
    stripe.Key = config.StripeSecretKey
    return &stripeClient{
        config: config,
    }
}
```
### CreatePaymentSession Logic
The required method from the `PaymentProcessor` interface is:
```go
CreatePaymentSession(ctx context.Context, amount int64, currency string, metadata map[string]string) (string, error)
```
This method:
- Builds `stripe.CheckoutSessionParams`
- Includes success URL, cancel URL, metadata, amount, currency
- Creates a Checkout Session with:
```go
result, err := session.New(params)
```
- Returns the created session ID (`result.ID`)
This ID will later be sent to the frontend as a payment link.
### Removal of GetSessionStatus
The earlier interface included `GetSessionStatus`.
The lesson removes it because the course uses **webhooks** for payment confirmation instead of polling Stripe.
This simplifies the PaymentProcessor to a single method.
### Payment Service Integration
`main.go` now initializes Stripe:
```go
paymentProcessor := stripe.NewStripeClient(stripeCfg)
svc := service.NewPaymentService(paymentProcessor)
```
This means:
- The business layer (Service) only depends on the interface (clean architecture)
- Stripe integration stays fully isolated inside the infrastructure layer
## Listening to the Payment Event
### Event Flow Overview
When a driver accepts a trip, the system triggers a multi-service workflow:
- The **Trip Service** updates the trip status to *accepted*.
- The Trip Service publishes a **PaymentCmdCreateSession** command.
- The **Payment Service** consumes that event and creates a Stripe Checkout Session.
- The Payment Service publishes a **PaymentEventSessionCreated** event.
- The **API Gateway** listens for this event and forwards it to the rider via WebSocket.
- The rider UI displays a *Pay Now* button linked to Stripe Checkout.
### Trip Service – Publishing the Payment Command
When the driver accepts a trip:
- The trip is fetched and updated.
- A `PaymentTripResponseData` payload is created containing:
    - tripID
    - userID
    - driverID
    - amount (in cents)
    - currency
- The service publishes the command:
    - **Routing key:** `payment.cmd.create_session`
    - **Queue:** `payment_trip_response`
- This tells the Payment Service to create a checkout session.
### Messaging Layer – New Events and Queues
The messaging module defines new queue names:
- **PaymentTripResponseQueue** – receives `payment.cmd.create_session`
- **NotifyPaymentSessionCreatedQueue** – emits `payment.event.session_created`
New event types:
- **PaymentTripResponseData** – data passed from Trip → Payment
- **PaymentEventSessionCreatedData** – stripe session returned to frontend
RabbitMQ binds:
- `payment.cmd.create_session` → `payment_trip_response`
- `payment.event.session_created` → `notify_payment_session_created`
### Payment Service – Trip Consumer
A new consumer listens to **PaymentTripResponseQueue**:
- Unmarshals the event
- Calls the Payment Service’s `CreatePaymentSession`
- Receives a Stripe session ID
- Publishes a new event with:
    - tripID
    - sessionID
    - amount (converted from cents → dollars)
    - currency
    - owner = riderID
This triggers the WebSocket notification to the rider.
### Payment Service – Stripe Checkout Session
The Payment Processor implementation:
- Uses `stripe-go` V81
- Loads secret keys from Kubernetes secrets
- Builds a Stripe Checkout Session with:
    - success URL
    - cancel URL
    - metadata (tripID, userID, driverID)
    - line items
    - amount
    - currency
Returns the generated `sessionID`.
### Payment Service – Main Setup
The main file:
- Loads config from env
- Creates the Stripe client
- Creates the payment service
- Starts RabbitMQ
- Starts the **TripConsumer** in a goroutine
- Waits for graceful shutdown
### API Gateway – Forwarding Stripe Session to UI
The rider WebSocket handler loads these queues:
- `notify_driver_no_drivers_found`
- `notify_driver_assign`
- `notify_payment_session_created`
When a message arrives:
- The gateway pushes it to the corresponding WebSocket client
- The frontend receives the `payment.event.session_created` message
- The rider UI displays the payment link
### Frontend – Stripe Public Key
A `.env.local` file is added:
- `NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=<test-key>`
This key is used by Next.js to initialize Stripe Checkout on the client.
### Stripe Test Keys and Behavior
Stripe test mode allows:
- Fake payments
- Fake card numbers
- No real money movement
Common test card:
- `4242 4242 4242 4242` – successful payment
When the user pays:
- Stripe redirects to `?payment=success`
- The UI marks the payment successful
## Stripe Webhook Flow and Final Payment Integration
### Overview
The webhook flow connects Stripe’s external payment confirmation to your internal system.
Once a payment succeeds, Stripe sends a signed webhook request to your API Gateway.
The API Gateway verifies the signature, extracts metadata, publishes a **PaymentEventSuccess**, and the Trip Service marks the trip as *payed*.
### Stripe Test Account and Keys
- You log into Stripe Dashboard → enable *Test Mode*.
- Two keys must be replaced with your Stripe account keys:
    - Secret Key (backend)
    - Publishable Key (frontend)
- These replace the earlier public “_test” keys you used without an account.
### Stripe CLI Setup
To receive webhooks in local development:
- Install Stripe CLI
- Run `stripe login`
- Then run a command like `stripe listen --forward-to localhost:<gateway-port>/webhook/stripe`
- CLI prints a *Webhook Secret*
- This secret must be stored as a Kubernetes Secret and passed to the API Gateway as `STRIPE_WEBHOOK_KEY`
### Kubernetes Updates for Webhook Key
A new secret is added under `stripe-secrets`:
- Key: `stripe-webhook-key`
- Value: your Stripe webhook secret from CLI
API Gateway Deployment mounts it into env:
- `STRIPE_WEBHOOK_KEY`
- Referenced in the webhook handler for signature validation
### API Gateway – Webhook Handler
A new HTTP handler is exposed at `/webhook/stripe`.
Responsibilities:
- Read raw request body
- Load `STRIPE_WEBHOOK_KEY`
- Verify signature using `webhook.ConstructEventWithOptions`
- Parse the Stripe event
- If event type is `checkout.session.completed`:
    - Extract metadata (`trip_id`, `user_id`, `driver_id`)
    - Create `PaymentStatusUpdateData`
    - Publish **PaymentEventSuccess** to RabbitMQ
This is the official source of truth confirming a real Stripe payment.
### Published Event – PaymentEventSuccess
When the payment succeeds:
- API Gateway publishes:
    - Routing key: `payment.event.success`
    - Queue: `payment_success`
- Payload includes:
    - tripID
    - userID
    - driverID
### Messaging Layer Updates
New queue:
- **NotifyPaymentSuccessQueue** = `"payment_success"`
RabbitMQ now binds:
- Event → Queue
    - `payment.event.success` → `payment_success`
### Trip Service – Payment Consumer
A new consumer is added:
- Listens on `payment_success`
- Unmarshals `PaymentStatusUpdateData`
- Calls `UpdateTrip()` with status `"payed"`
- Prints: *"Trip has been completed and payed."*
This finalizes the trip lifecycle.
### Full End-to-End Payment Sequence
1. Rider requests trip
2. Driver accepts trip
3. Trip Service publishes `payment.cmd.create_session`
4. Payment Service creates Stripe checkout session
5. Payment Service publishes `payment.event.session_created`
6. API Gateway pushes event to rider WebSocket
7. Rider pays via Stripe Checkout
8. Stripe sends webhook → `/webhook/stripe`
9. API Gateway verifies signature and publishes `payment.event.success`
10. Trip Service marks trip as **payed**
### Stripe Dashboard—What You Can See
Stripe records each step:
- PaymentIntent created
- PaymentIntent succeeded
- Charge succeeded
- Checkout Session completed
- Metadata included (tripID, userID, driverID)
All of these confirm the payment lifecycle and help debug issues.
## Observability
### **Intro to Distributed Tracing**
#### Why Observability Matters in a Microservice System
As the ride-sharing project grows, debugging becomes harder because:
- Multiple services interact asynchronously
- Failures can occur at many layers (API Gateway, Trip Service, external OSRM API, Payment Service, WebSockets, Stripe, etc.)
- Logs alone are not enough to understand end-to-end request behavior
Observability provides a clear view of how a request travels through the system, making it easier to:
- Identify where failures happen
- Measure performance bottlenecks
- Debug multi-service flows
- Ensure reliability as the system scales
### What Distributed Tracing Solves
Distributed tracing allows you to track a request as it moves through:
- The API layer
- Internal microservices
- External service calls
- Message queues
- Database operations
- Payment providers
With tracing:
- Each request has a **Trace ID** shared across services
- Each operation is a **Span**
- Spans form a hierarchical tree showing exactly where time was spent
- You get a waterfall visualization showing every step in the request lifecycle
This is crucial for debugging flows like trip preview:
```
Rider → API Gateway → Trip Service → OSRM API → Trip Service → Gateway → Rider
```
If preview fails, tracing instantly tells you whether:
- The UI call failed
- Gateway routing failed
- Trip Service errored
- OSRM API timed out
- JSON parsing failed
- Network latency spiked
### What a Trace and Span Actually Represent
A **trace** represents a full end-to-end request.
A **span** represents a single unit of work inside a trace.
Examples of spans:
- HTTP request to API Gateway
- gRPC call from API Gateway → Trip Service
- MongoDB query in Trip Service
- Call to OSRM API
- Publish message to RabbitMQ
- Payment session creation
- Webhook processing
Spans include:
- Start time
- End time
- Duration
- Attributes (method, URL, IDs, metadata)
- Status (OK / ERR)
A trace may contain dozens of spans across multiple services.
### How Services Share Trace Context
To connect spans from multiple services, the system propagates:
- TraceID
- SpanID
- Metadata
Propagation can occur through:
- HTTP headers
- gRPC metadata
- RabbitMQ message headers
If propagation is implemented correctly:
- The entire distributed system behaves like one connected timeline
- Jaeger will show the full cross-service timeline of any request
### Instruments for Observability: OpenTelemetry
OpenTelemetry (OTel) is a unified standard used to instrument distributed systems.
It is **not** a backend; it's a framework for:
- Generating traces
- Generating logs
- Generating metrics
- Exporting telemetry to any backend (Jaeger, Prometheus, Grafana, Honeycomb, Datadog)
It includes:
- APIs
- SDKs
- Propagators
- Exporters
OpenTelemetry exists for almost all languages.
In this course, you use Go.
### Key Components of OpenTelemetry in This Project
1. **Automatic instrumentation**
    
    Adds tracing automatically for:
    
    - gRPC servers & clients
    - HTTP servers & clients
    - RabbitMQ publish/consume
    - Database queries (MongoDB)
2. **Manual instrumentation**
    
    You manually create spans around:
    
    - Business logic
    - External API calls
    - Payment flows
    - Complex operations
This gives fine-grained detail that automatic tooling cannot guess.
### The Telemetry Pipeline
OpenTelemetry processes data through several stages:
1. **Instrumentation**
    
    Code generates telemetry (spans, logs, metrics)
    
2. **Processor**
    
    Adds batching, filtering, sampling
    
3. **Exporter**
    
    Sends telemetry to a backend
    
4. **Backend**
    
    Stores and indexes traces
    
    In this project: **Jaeger**
    
5. **Frontend**
    
    Used for visualization
    
    In this project: **Jaeger UI**
    
### Jaeger: The Observability Frontend
Jaeger provides:
- Timeline views
- Visual trace waterfall
- Span details
- Error visualization
- Service-to-service dependency graphs
- Latency breakdowns
You will be able to see exactly:
- How long each service took
- Where bottlenecks are
- Whether external APIs are slow
- How many retries occurred
- How RabbitMQ delivery performed
### Benefits for This Ride-Sharing System
With observability added, you can debug:
- Slow trip previews
- Failed trip creation
- Long OSRM routing times
- Payment session failures
- Stripe webhook delays
- RabbitMQ bottlenecks
- gRPC timeout chains
This transforms the system from a black box into a fully traceable workflow.
### Types of Instrumentation You Will Implement
1. **Global tracer setup**
    
    Every service initializes the same OTel provider
    
2. **Automatic instrumentation**
    - gRPC interceptor middleware
    - HTTP middleware
    - RabbitMQ wrappers
3. **Manual spans in critical areas**
    - External API calls
    - Business logic steps
    - Payment service communication
    - Stripe checkout creation
    - Trip assignment logic
### Trade-Offs and Code Verbosity
Instrumentation adds:
- More lines of code
- More context passing
- More imports
But it improves:
- Reliability
- Debuggability
- Production readiness
A balanced approach is used:
- Automatic instrumentation for frameworks
- Manual spans for core business flows
## Tracing Setup
### Overview
To introduce distributed tracing across all microservices, the system requires a **centralized tracing module**. This module lives in `shared/tracing` and provides reusable tracing initialization for all services. Each service calls this module once during startup, supplying its own configuration.
### Tracing Configuration
Each service initializes tracing by calling `InitTracer` with a configuration object containing:
- **ServiceName** – identifies the service in the tracing backend
- **Environment** – separates traces per environment (e.g., development, production)
- **JaegerEndpoint** – the endpoint of the tracing exporter (Jaeger will be added later)
### Designing the Shared Tracing Module
The tracing module is structured around three core components:
1. **Exporter**
    
    Responsible for sending spans to an external tracing backend (Jaeger, console exporter, etc.).
    
    The exporter will be implemented later.
    
2. **TracerProvider**
    
    The core object that manages span creation, batching, and exporting.
    
3. **Propagator**
    
    Responsible for propagating trace context across boundaries (HTTP, gRPC, RabbitMQ).
    
    This enables multi-service trace continuity.
    
The tracing setup function should return:
- A **shutdown function** to cleanly flush spans on service exit
- An **error** if the initialization fails
### InitTracer Function
Each service calls `InitTracer`, which:
1. Creates an exporter (currently `nil`, implemented later)
2. Builds the TracerProvider with:
    - the exporter
    - a Resource object describing the service
    - batching settings
3. Registers the provider globally through `otel.SetTracerProvider`
4. Sets the global propagator
5. Returns the shutdown function
### Creating the Resource Object
A resource describes the identity of the process emitting telemetry.
For each service, the resource includes:
- `service.name`
- `deployment.environment`
These fields allow Jaeger to group, filter, and analyze traces across services and environments.
### Creating the TracerProvider
The TracerProvider combines:
- The exporter (to be added later)
- The resource descriptor
- The batching processor
All spans created anywhere in the service will be handled by this TracerProvider.
### Propagator Setup
Distributed tracing requires that trace context is forwarded across service boundaries.
The tracing module enables two propagation mechanisms:
- **W3C TraceContext** (trace IDs and span IDs)
- **Baggage** (key-value metadata propagated with traces)
These are combined into a composite propagator and registered as the global default.
## Jaeger Integration and Exporter Setup
### Adding Jaeger to the Development Environment
A dedicated Jaeger instance is added to the Kubernetes development environment through a new deployment file (`infra/development/k8s/jaeger.yaml`). The Jaeger setup uses the `jaegertracing/all-in-one` image, which provides both:
- The Jaeger UI (port 16686)
- The collector endpoint for receiving spans (port 14268)
A Kubernetes Service exposes these two ports inside the cluster.
To integrate Jaeger into Tilt, the Tiltfile now includes:
- The Jaeger YAML resource
- A `jaeger` Tilt resource with port forwards for:
    - `16686:16686` (UI)
    - `14268:14268` (collector)
This allows local access to the Jaeger UI at [http://localhost:16686](http://localhost:16686/).
### Extending App Configuration
A new configuration value is added to `app-config.yaml`:
- `JAEGER_ENDPOINT: "http://jaeger:14268/api/traces"`
Each service consumes this endpoint as an environment variable to know where spans should be exported.
The following services now receive `JAEGER_ENDPOINT`:
- API Gateway
- Driver Service
- Payment Service
- Trip Service
### Initializing Tracing in Each Service
Every service now initializes OpenTelemetry tracing during its startup sequence. The steps are the same across all services:
1. Build a `tracing.Config` with:
    - `ServiceName`
    - `Environment`
    - `JaegerEndpoint`
2. Call `tracing.InitTracer`
3. Handle any initialization error (services fail fast if tracing cannot start)
4. Create a root context with cancellation
5. Defer both `cancel()` and the tracer’s shutdown function
This guarantees proper initialization and graceful flushing of spans on service termination.
### Implementing the Jaeger Exporter
The tracing module now includes the exporter initialization.
A new private function `newExporter` is introduced:
```go
func newExporter(endpoint string) (sdktrace.SpanExporter, error) {
    return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
}
```
This exporter sends spans to the Jaeger collector using the HTTP endpoint defined in configuration.
The exporter is created inside `InitTracer`, replacing the previous `nil` placeholder.
### Updated Tracing Initialization Flow
The `InitTracer` function now performs three core tasks:
1. **Create the Jaeger exporter**
2. **Build the TraceProvider** using:
    - the exporter
    - batch processing
    - the service resource
3. **Set up global propagation**
Finally, it returns the Provider’s shutdown function so that each service can flush spans on exit.
### Verifying Jaeger Startup
After running `tilt up`, you should now see:
- A running Jaeger pod and service
- A new Tilt resource labelled as tooling
- The Jaeger UI accessible locally via port-forwarding
Initially, no services appear in Jaeger because no spans have been emitted yet; spans will appear once instrumentation is added in upcoming steps.
## HTTP Instrumentation
### Overview
This section introduces **HTTP-level tracing instrumentation** in the API Gateway using OpenTelemetry.
The goal is to automatically and manually create **spans** for every incoming HTTP request, so that Jaeger can display full request waterfalls (parent/child spans) and enrich traces with HTTP metadata such as method, status code, and user-agent.
### Global Tracer Access
A global tracer is now retrieved via:
- `GetTracer(name string)` returning a `trace.Tracer`
- This avoids passing tracer instances through functions
- All spans created via this tracer are attached to the **global tracer provider** configured during service startup
### Manual Span Instrumentation in Handlers
Each HTTP handler now begins with:
- `ctx, span := tracer.Start(r.Context(), "<operation>")`
- `defer span.End()`
- Replacing all `r.Context()` usages with the traced `ctx`
Example operations instrumented:
- `handleTripStart`
- `handleTripPreview`
- `handleStripeWebhook`
These now emit spans with:
- The handler’s name (`handleTripStart`, etc.)
- The full call chain (gRPC calls, RabbitMQ publish, etc.)
### Automatic Instrumentation via Middleware Wrapper
A new HTTP middleware wrapper was added in `shared/tracing/http.go`:
- `WrapHandlerFunc(handler http.HandlerFunc, operation string)`
- Returns `otelhttp.NewHandler(handler, operation)`
The wrapper adds:
- Automatic creation of a root span for the handler
- Standardized HTTP span attributes (status, method, bytes sent/received, client IP, protocol, user-agent)
- Metrics enrichment
This wrapper is now used across all gateway routes.
### Updated Route Registration
The API Gateway’s handlers now use `mux.Handle` with wrapped handlers:
- POST `/trip/preview`
- POST `/trip/start`
- WS `/ws/drivers`
- WS `/ws/riders`
- POST `/webhook/stripe`
Example:
```go
mux.Handle("POST /trip/preview",
    tracing.WrapHandlerFunc(enableCORS(handleTripPreview), "/trip/preview"),
)
```
### Handler Behavior Changes
Services now:
- Use traced context when calling gRPC to other services
- Use traced context when publishing events to RabbitMQ
- Generate full trace chains visible in Jaeger
### Results in Jaeger
After this commit:
- Every HTTP request produces a full waterfall in Jaeger
- Automatic spans show request metadata
- Manual spans show deeper internal operations (gRPC calls, webhook flow)
- WebSocket handlers are also traced
- Stripe webhook callbacks appear as traces from external services
## gRPC Tracing Instrumentation
### Overview
This lesson adds **OpenTelemetry tracing** to all gRPC communication across the microservices.
After completing HTTP tracing, the next step is ensuring visibility inside all gRPC calls—both **client-side** and **server-side**.
The goal:
Every gRPC request (preview trip, start trip, driver updates, payment processing, etc.) should generate spans that connect seamlessly with the API Gateway spans, producing full cross-service traces inside Jaeger.
### Tracing Architecture for gRPC
There are two instrumentation layers required:
### 1. **Server-side tracing**
Every service that exposes a gRPC server must attach an OpenTelemetry stats handler.
Services affected:
- **driver-service**
- **trip-service**
### 2. **Client-side tracing**
Every gRPC client must attach a tracing stats handler:
- API Gateway → Driver Service
- API Gateway → Trip Service
- Payment Service → internal gRPC calls (if any)
### Tracing Utilities (shared package)
#### `WithTracingInterceptors`
Creates gRPC server options that automatically apply OpenTelemetry handlers:
```go
func WithTracingInterceptors() []grpc.ServerOption {
    return []grpc.ServerOption{
        grpc.StatsHandler(newServerHandler()),
    }
}
```
### `DialOptionsWithTracing`
Creates gRPC dial options for clients:
```go
func DialOptionsWithTracing() []grpc.DialOption {
    return []grpc.DialOption{
        grpc.WithStatsHandler(newClientHandler()),
    }
}
```
### Client/Server handlers
Both rely on OpenTelemetry’s official `otelgrpc` instrumentation:
```go
func newClientHandler() stats.Handler {
    return otelgrpc.NewClientHandler(
        otelgrpc.WithTracerProvider(otel.GetTracerProvider()),
    )
}
func newServerHandler() stats.Handler {
    return otelgrpc.NewServerHandler(
        otelgrpc.WithTracerProvider(otel.GetTracerProvider()),
    )
}
```
These handlers automatically:
- create spans for all gRPC requests
- include metadata like RPC method name, status, deadlines, payload size
- propagate context upstream/downstream
### Applying Tracing in gRPC Servers
#### Driver Service
```go
grpcServer := grpcserver.NewServer(tracing.WithTracingInterceptors()...)
```
### Trip Service
```go
grpcServer := grpcserver.NewServer(tracing.WithTracingInterceptors()...)
```
Both services now emit spans whenever:
- a rider requests a trip preview
- a rider starts a trip
- a driver updates availability/status
- payment consumer updates the trip
### Applying Tracing in gRPC Clients
#### Driver Client
```go
dialOptions := append(
    tracing.DialOptionsWithTracing(),
    grpc.WithTransportCredentials(insecure.NewCredentials()),
)
conn, err := grpc.NewClient(driverServiceURL, dialOptions...)
```
### Trip Client
Same pattern applied:
```go
dialOptions := append(
    tracing.DialOptionsWithTracing(),
    grpc.WithTransportCredentials(insecure.NewCredentials()),
)
conn, err := grpc.NewClient(tripServiceURL, dialOptions...)
```
### What This Enables in Jaeger
After enabling these interceptors:
- API Gateway span → gRPC client span → gRPC server span
    
    Creates a full chain of visibility.
    
- The **System Architecture** tab in Jaeger now displays:
    - API Gateway → Trip Service
    - API Gateway → Driver Service
    - Call counts between services
- Each request waterfall shows:
    - HTTP entry span
    - gRPC client span
    - gRPC server span
    - Any additional spans inside the service
This completes observability for all **synchronous** communication in the system.
The next step in the course:
**Add tracing to RabbitMQ (asynchronous messaging)** to complete the full picture.
## Asynchronous RabbitMQ Tracing
### Overview
This lesson completes the observability stack by adding **OpenTelemetry tracing** to all **asynchronous RabbitMQ messaging**.
Unlike HTTP or gRPC (which pass context in-process), asynchronous messaging requires manually **injecting** and **extracting** trace context into message headers, because:
- messages travel across the network,
- producers and consumers run in different services,
- spans must still form a single trace in Jaeger.
The goal is full visibility of the async workflow:
- API Gateway → Trip Service (publish)
- Trip Service → Driver Service (consume/publish)
- Payment Service → Trip Service (consume)
- Stripe Webhook → Payment → Trip
### The Problem Being Solved
gRPC tracing only covers synchronous calls.
But the core ride-sharing workflow heavily depends on async events:
- driver assignment
- trip created
- trip accepted
- payment events
- webhook events
- background updates
Without RabbitMQ tracing, the Jaeger waterfall would show gaps between services.
### Context Propagation Strategy
OpenTelemetry requires:
1. **Injecting trace context into message headers** when publishing.
2. **Extracting headers** on the consumer side to continue the trace.
RabbitMQ supports arbitrary key–value headers (`amqp.Table`), so these are used as the carrier for trace metadata.
### AMQP Carrier Implementation
A custom `amqpHeadersCarrier` is created to satisfy the `TextMapCarrier` interface.
This allows the Otel propagator to:
- write trace data into the AMQP headers (`Inject`)
- read trace data when consuming (`Extract`)
The carrier maps string keys to interface values inside `amqp.Table`.
### Traced Publisher
When publishing a message:
```go
ctx, span := tracer.Start(ctx, "rabbitmq.publish")
```
**What the publisher does:**
- creates a span for the publish operation
- attaches metadata (exchange, routing key, owner_id)
- injects the trace context into `msg.Headers`
- calls the underlying publish function (`r.publish`)
- marks the span as error if publish fails
This allows the asynchronous message to **carry the trace** across process boundaries.
### Traced Consumer
When a message is received:
```go
ctx := propagator.Extract(context.Background(), carrier)
ctx, span := tracer.Start(ctx, "rabbitmq.consume")
```
**What the consumer does:**
- extracts the context from the incoming message headers
- starts a new span that becomes a child of the publisher’s span
- attaches attributes like exchange, routing key, owner_id
- executes the application’s handler logic
- sets error status on failure
This produces the asynchronous link in Jaeger’s waterfall.
### Updating the RabbitMQ Abstraction
#### PublishMessage
The original publishing code is wrapped:
```go
return tracing.TracedPublisher(ctx, TripExchange, routingKey, msg, r.publish)
```
`r.publish` is now a small internal method that performs the actual AMQP publish.
### ConsumeMessages
The consumer loop is updated to wrap each message:
```go
if err := tracing.TracedConsumer(msg, func(ctx context.Context, d amqp.Delivery) error {
    ...
}); err != nil {
    log.Printf("Error processing message: %v", err)
}
```
This ensures **every** event creates proper spans.
### Result in Jaeger
After updating the system, Jaeger displays:
- spans for publishing events
- spans for consuming events
- correct parent/child relationships across services
- full cross-service waterfall covering the entire ride flow
- visible gaps representing real network delays
- metadata attached: routing keys, queues, owners, payload details
For example, a `/trip/start` request now produces:
1. HTTP handler span (API Gateway)
2. gRPC span (API → Trip Service)
3. RabbitMQ publish span (Trip Service)
4. RabbitMQ consume span (Driver Service)
5. Driver Service publish span
6. Trip Service consume span
7. Payment Service publish/consume
8. Stripe webhook span
All automatically connected in a single end-to-end trace.
### Architecture View
Jaeger’s “System Architecture” tab now shows a fully connected graph:
- API Gateway
- Trip Service
- Driver Service
- Payment Service
With arrows indicating message flow and counts.
## Reliability
### **Understanding DLQ and Retries**
#### Introduction
As we prepare for the production-ready deployment of the system, this module focuses on improving **reliability**, **message safety**, and **observability**.
The goal is to take the architecture from “good” to “robust,” by handling failure cases that inevitably occur in distributed systems.
### Current Strengths of the System
Before introducing new improvements, we review what the system already does well:
### Message Acknowledgment
Each consumer manually **acknowledges** messages only after successful processing.
If processing fails, messages are **nacked** and requeued.
This gives us full control over:
- when a message is considered successfully processed
- when a failure is logged
- when a retry should occur
### Durable, Persistent Messages
All outgoing messages are published with:
- `DeliveryModePersistent`
- Queues declared durable
This ensures:
- messages survive pod restarts
- messages survive RabbitMQ restarts
- no data loss in the event of crashes
### Monitoring & Observability
We already implemented:
- **HTTP tracing**
- **gRPC tracing**
- **RabbitMQ async tracing**
- End-to-end Jaeger visibility
This means errors, latency, and cross-service behavior are visible and traceable.
### Recovery & Resilience
The system currently:
- properly handles message failures
- logs and surfaces errors that affect message flow
- maintains good modular isolation between services
Overall, we are in a *very healthy* architectural state.
---
## What We Can Improve Next
To harden the system for real production scenarios, we introduce three critical reliability patterns:
### 1. Dead Letter Queues (DLQ)
#### 2. Retry Mechanisms (exponential backoff)
### 3. Circuit Breakers & Idempotency (future improvements)
This module focuses on **Dead Letter Queues**.
---
## What Is a Dead Letter Queue?
A **Dead Letter Queue (DLQ)** is a special queue used to store messages that:
- cannot be processed,
- consistently produce errors,
- exceed their max number of retries,
- contain invalid or corrupted data ("poison messages").
Think of a DLQ as a **parking lot for failed messages**.
Messages placed there are *not lost*, but preserved for:
- debugging
- auditing
- automated recovery
- developer investigation
### Why DLQs Matter
In distributed systems, failure is inevitable.
### DLQs prevent message loss
Without a DLQ:
- failed messages disappear forever
- debugging becomes nearly impossible
- inconsistent system states may appear
With a DLQ:
- all problematic messages are preserved
- no data is lost
### DLQs isolate poison messages
A poison message is one that:
- is malformed
- is missing required fields
- cannot be unmarshaled
- always causes a processing error
If processed repeatedly, it can:
- break consumers
- cause infinite retry loops
- clog the queue
A DLQ isolates these messages safely.
### DLQs improve debugging and visibility
Stored failed messages give us insight into:
- data corruption
- unexpected business cases
- integration issues
- downstream service outages
For modern systems, DLQs are not optional—they are foundational.
---
## Typical DLQ Workflow
A standard scenario that demonstrates DLQ usefulness:
### 1. Trip service publishes `trip.driver_assigned`
```
Trip → RabbitMQ → Payment Service
```
### 2. Payment service processes the message
But Stripe API is temporarily **down**.
### 3. The system retries the message
With exponential backoff (implemented in the next lesson).
### 4. After N failed attempts
The message is rejected.
### 5. RabbitMQ moves the message to the Dead Letter Queue
Along with:
- original routing key
- error reason
- number of attempts
- timestamp
### 6. Developers (or automated jobs) investigate the DLQ
This ensures:
- visibility
- safety
- no silent failures
- no lost payments or broken flows
---
## Why This Matters for Your Project
The ride-sharing backend is:
- asynchronous
- event-driven
- distributed
- real-time
This architecture **depends heavily on messaging**, so failures in message processing must be handled with extreme care.
DLQs give us:
- stronger correctness guarantees
- better resilience
- transparency in failure modes
- protection against corrupted data
- easier production debugging
---
If you want, I can continue with the next lesson: **Retry Mechanisms (Exponential Backoff)**.
## Retry Mechanism Added to RabbitMQ Consumers
### Overview
This update integrates a full retry system with exponential backoff into the RabbitMQ consumers. Instead of immediately failing or losing messages when a handler errors, the consumer now retries processing several times before finally rejecting the message and marking it for the Dead Letter Queue. This creates a much more resilient and production-ready async pipeline.
### Why This Matters
Previously, failed messages were either lost or stuck waiting until a service came back up. Now the system handles transient failures automatically, and only after repeated failures does it discard the message, enriched with metadata for later investigation.
### Retry Logic
The retry package provides configurable exponential backoff. The consumer wraps the handler with:
```go
retry.WithBackoff(ctx, cfg, func() error { return handler(ctx, d) })
```
This retries the handler based on the default retry configuration until success or until `MaxRetries` is reached.
### Failure Handling and Metadata
If all retries fail, the message is enriched with diagnostic metadata:
- `x-death-reason`: the error returned by the handler
- `x-origin-exchange`: which exchange produced the message
- `x-original-routing-key`: the routing key of the message
- `x-retry-count`: number of retries attempted
These headers follow the message into the Dead Letter Queue on the next module, allowing developers or automated systems to inspect why the message failed.
### Reject and DLQ Forwarding
After adding metadata, the message is explicitly rejected with:
```go
d.Reject(false)
```
With `requeue=false`, RabbitMQ will forward the message to a DLQ when it is configured. Until then, the message will simply disappear, validating the retry flow logic.
### Successful Processing
Messages are acknowledged only when the handler ultimately succeeds:
```go
msg.Ack(false)
```
This ensures no message is acknowledged prematurely.
### Testing Scenario
To simulate failure, the payment service was temporarily disabled.
After scheduling a trip:
1. The UI gets stuck waiting for payment.
2. RabbitMQ shows the payment event stuck in the payment queue.
3. With retries enabled and a forced error in the handler, logs show retries at 1s, 2s, 4s.
4. After max retries, the message is rejected and leaves the queue.
This confirms that the retry logic, metadata tagging, and rejection flow are working correctly.
## Dead Letter Queue (DLQ) and Dead Letter Exchange (DLX) Setup
### Overview
This update introduces a full Dead Letter Exchange + Dead Letter Queue flow into the messaging layer. Any message that fails after all retry attempts is no longer lost — it is now routed to a centralized DLQ where failures can be inspected, debugged, and audited. This brings the system to true production-grade reliability.
### What Was Added
The system now includes:
### Dead Letter Exchange (DLX)
A new exchange named `dlx` is declared.
Its purpose: receive any rejected/failed messages from all queues.
### Dead Letter Queue (DLQ)
A single queue named `dead_letter_queue` is declared and bound to the DLX using the wildcard `#`.
This means **any** message sent to the DLX will land in this queue.
### Queue Reconfiguration
Every existing functional queue in the system is now created with:
```go
args := amqp.Table{"x-dead-letter-exchange": DeadLetterExchange}
```
This instructs RabbitMQ that **if the consumer rejects a message**, the broker should route it to the DLX.
### Integration With Retry Logic
Because the previous module added retry logic (exponential backoff), the new flow now works as:
1. A consumer receives a message
2. Retry is attempted N times
3. After final failure:
    - The message is enriched with metadata headers:
        - `x-death-reason`
        - `x-origin-exchange`
        - `x-original-routing-key`
        - `x-retry-count`
    - The message is **rejected without requeue**
4. RabbitMQ forwards the message to the DLX → DLQ
5. The message waits in the DLQ for debugging or automated recovery
### RabbitMQ Behavior After Deployment
When a queue is re-declared with new arguments, RabbitMQ refuses it (queues are immutable).
Because of this, the existing queues had to be manually deleted so that RabbitMQ could recreate them with the new DLX settings.
Once deleted, the system recreated all queues correctly, and each queue now shows in the UI:
```
x-dead-letter-exchange: dlx
```
### Testing the DLQ
To validate the implementation:
1. The payment service was modified to return an intentional error.
2. A new trip was scheduled.
3. The retry flow triggered (1s → 2s → 4s).
4. The message failed permanently.
5. The message disappeared from the payment queue.
6. It appeared in `dead_letter_queue`.
7. The DLQ message showed the correct metadata headers.
### Result
Your messaging architecture is now fully production-ready:
- Retries
- Durable queues
- Persistent messages
- Dead Letter Queue for all failures
- Context-rich debugging metadata
- Automatic routing via DLX
- Compatible with the existing tracing instrumentation
The system can now safely handle real-world failures, transient outages, misformatted messages, or poison messages — without losing data and without blocking the pipelines.
## MongoDB
### MongoDB for Our Ride-Sharing Project
#### What MongoDB Is
- MongoDB is a **NoSQL document database**
    - NoSQL = not based on traditional relational tables/joins, but key–value / document style
- It stores data as **JSON-like documents**
    - Documents = objects (similar to Go structs / JS objects)
    - Collections = groups of documents (similar to a table, but flexible)
- It is **document-oriented**
    - Data is stored in rich objects instead of rows
    - Very natural fit if you already think in JSON / structs
- It uses an internal **`_id` field** of type `ObjectID`
    - This is the unique primary key MongoDB uses
    - In Go we’ve already prepared for this with `primitive.ObjectID` in the trip domain models
### Schema and Flexibility
- MongoDB is **schema-less / schema-flexible**
    - You don’t need to define a rigid schema up front (like in Postgres/MySQL)
    - You can **add or remove fields over time** without migrations
    - Perfect for an evolving project where:
        - Trip data may gain new fields
        - You don’t want to constantly run SQL migrations
- Compared to relational DBs:
    - Traditional DBs: fixed tables, strict columns, migrations
    - MongoDB: documents can differ slightly between records, and that’s fine
### Why MongoDB Is a Good Fit Here
- **No SQL required to start**
    - You interact with it via the **MongoDB Go driver API** (e.g., `collection.InsertOne`, `Find`, etc.)
    - No need to write SQL queries to get value from a DB quickly
- **JSON-like documents match our domain**
    - Trips, drivers, payments all map nicely to nested JSON / struct documents
- **Great for early-stage microservices**
    - The data model is still moving
    - You avoid schema migration overhead while you iterate
- **Geo features available**
    - Mongo supports **geospatial queries**
    - Useful if later you want location-based searching (drivers near rider, etc.)
### Microservices + One Database per Service
- Common microservice pattern: **“one database per service”**
    - Each service owns its data and its storage
    - Avoids tight coupling and cross-service DB sharing
- How this maps to our project:
    - Right now, only **Trip Service** persists data (by design, for simplicity)
    - In a “real” system you might have:
        - Trip Service → `trip_db` (Mongo or relational)
        - Driver Service → `driver_db`
        - Payment Service → `payment_db`
    - Each could even use **different technologies**:
        - Trips → MongoDB
        - Payments → Postgres
        - Analytics → ClickHouse, etc.
- Why Mongo helps here:
    - Spinning up a Mongo instance per service is **lightweight and fast**
    - Compared to starting a heavy SQL DB instance per service (Postgres, MySQL)
### Managed MongoDB (MongoDB Atlas)
- **Managed solution** = MongoDB Atlas
    - Cloud platform from the MongoDB company
    - You get:
        - Automatic provisioning
        - Backups, monitoring, metrics
        - Built-in security & authentication
        - Easy scaling (bigger cluster, replica sets, etc.)
- **Big benefits for us**:
    - **No maintenance**
        - No patching, no security upgrades, no OS management
    - **Free tier**
        - Ideal for learning, pet projects, and this course
    - Good defaults for:
        - Backups
        - Monitoring
        - Alerting
- **Trade-off**:
    - In general, managed = more expensive than self-hosting
    - But for us:
        - Free tier removes this concern
        - Simplicity > theoretical cost optimization
### Self-Hosted MongoDB on Kubernetes
- **Self-hosted** = run MongoDB as a pod inside our Kubernetes cluster
    - Similar to how we run RabbitMQ
    - We control:
        - MongoDB version
        - Storage class / persistent volumes
        - Backup strategy
        - Security hardening
- **Advantages**:
    - **Full control** over configuration and deployment
    - Potentially **cheaper** at scale if you already pay for the cluster resources
    - Can integrate deeply with your existing infra (storage, monitoring, backup tools)
- **Disadvantages**:
    - You must handle:
        - Upgrades and patches
        - Replica set setup and failover
        - Backup and restore procedures
        - Security (TLS, auth, network policies)
    - More moving parts = more ways to break production
### Strategy for This Course
- We’ll introduce **MongoDB first via a managed approach** (MongoDB Atlas):
    - You get a **real persistent DB in the cloud** with minimal friction
    - Perfect for learning and for a portfolio project
- Then we’ll also **show how to self-host MongoDB on Kubernetes**:
    - So you understand both models
    - And can choose later based on:
        - Cost
        - Control
        - Operational maturity
- Key idea:
    - For learning and early projects → **Managed Atlas** is usually the best move
    - For heavy, cost-sensitive, and infra-mature setups → **Self-hosted** might be worth it
## MongoDB Atlas Setup + Backend Integration
### Creating a Free MongoDB Atlas Cluster
#### Step 1: Create a MongoDB Atlas Account
Sign up on MongoDB Atlas and create or select a project.
### Step 2: Create a Free Cluster
Choose the **Free Tier**, select any cloud provider, name the cluster (for example `ride-sharing`), and create it.
### Step 3: Wait for Atlas to Initialize Everything
Atlas prepares:
- Network access
- Default environment config
- Cluster provisioning
### Step 4: Create a Database User
This must be done manually.
Go to **Connect → Drivers** and create:
- A username
- A password
### Step 5: Get the Connection String
Select “Drivers → Go”, copy the connection URI, and replace:
- `<username>` with your database username
- `<password>` with your database password
    
    Remove the angle brackets.
    
Example:
```
mongodb+srv://myuser:mypassword@cluster0.abc123.mongodb.net/?retryWrites=true&w=majority
```
Store this string temporarily. It will go inside your Kubernetes secret.
### Creating the Shared MongoDB Package
#### db/mongodb.go
The shared DB package introduces a reusable MongoDB setup.
### Key Components
**MongoConfig**
Holds:
- `URI` from environment variable
- `Database` name (in this case `"ride-sharing"`)
**NewMongoDefaultConfig**
Loads:
```
URI: os.Getenv("MONGODB_URI")
Database: "ride-sharing"
```
**NewMongoClient**
- Validates config
- Connects to MongoDB via the official Go driver
- Pings the server to confirm connectivity
- Logs a successful connection
**GetDatabase**
Returns the database instance for use inside repositories.
### Integrating MongoDB in Trip Service
#### Step 1: Add Environment Variable to Deployment
Inside `trip-service-deployment.yaml`:
```
- name: MONGODB_URI
  valueFrom:
    secretKeyRef:
      name: mongodb
      key: uri
```
### Step 2: Add Kubernetes Secret
Create a secret named `mongodb` with:
```
uri: "<your-complete-atlas-connection-uri>"
```
### Step 3: Initialize MongoDB in trip-service
Inside `cmd/main.go`:
```
mongoClient, err := db.NewMongoClient(ctx, db.NewMongoDefaultConfig())
defer mongoClient.Disconnect(ctx)
mongoDb := db.GetDatabase(mongoClient, db.NewMongoDefaultConfig())
log.Printf(mongoDb.Name())
```
This ensures:
- A live MongoDB connection
- Graceful shutdown with `Disconnect`
- A reference to the DB for repositories on the next lesson
## MongoDB Repository Integration
### Overview
The final database step is replacing the in-memory trip repository with a fully persistent MongoDB-backed implementation. This is done by introducing a `mongoRepository` that satisfies the same interface as the in-memory repo, allowing the service layer to remain unchanged. This demonstrates the benefit of the repository pattern: you can replace an entire storage engine without touching business logic.
### MongoDB Repository File
A new file was added:
`services/trip-service/internal/infrastructure/repository/mongodb.go`
This file contains the full implementation for:
- Creating trips
- Fetching trips by ID
- Updating trips (status and assigned driver)
- Saving ride fares
- Fetching ride fares by ID
Each method interacts with the appropriate Mongo collection:
- `db.TripsCollection` (`"trips"`)
- `db.RideFaresCollection` (`"ride_fares"`)
### Key Behaviors
#### CreateTrip
- Inserts the trip document into MongoDB.
- MongoDB automatically generates `_id` (ObjectID).
- That ID is assigned back to the TripModel struct.
### GetTripByID
- Converts string ID → `ObjectID`
- Retrieves the document from Mongo
- Decodes BSON into a Go struct (`TripModel`)
### UpdateTrip
- Builds a `$set` update document
- Conditionally adds the driver to the update
- Runs `UpdateOne` and ensures that at least one document was modified
- If not modified, returns an error (`trip not found`)
### SaveRideFare / GetRideFareByID
Works exactly like the Trip version, but on the `ride_fares` collection.
### Replacing the Repository in main.go
Before:
```
inmemRepo := repository.NewInmemRepository()
svc := service.NewService(inmemRepo)
```
After:
```
mongoDBRepo := repository.NewMongoRepository(mongoDb)
svc := service.NewService(mongoDBRepo)
```
This is the only change needed in the application logic.
### Why This Works Smoothly
Because the service layer depends only on an **interface**, not a concrete storage engine.
Both repositories implement the same methods:
- `CreateTrip`
- `GetTripByID`
- `UpdateTrip`
- `SaveRideFare`
- `GetRideFareByID`
### UI Verification
Inside MongoDB Atlas → **Browse Collections**, you will now see:
**Database:** `ride-sharing`
Collections:
- `trips`
- `ride_fares`
A created trip will appear with fields like:
```
{
  _id: ObjectId(...),
  status: "accepted",
  driver: {...},
  origin: {...},
  destination: {...}
}
```
After paying:
- Refresh the collection
- The trip’s `status` updates to `"paid"`
Ride fares are also stored every time a trip is previewed.
### Optional Practice: Cleaning Old Ride Fares with TTL
MongoDB supports TTL (time-to-live) indexes.
A useful extension is:
- Automatically delete `ride_fares` that are older than X minutes.
- This prevents the collection from endlessly growing.
You would set an index like:
```
db.ride_fares.createIndex(
    { createdAt: 1 },
    { expireAfterSeconds: 3600 }
)
```
This is optional and left as an exercise.
If you want, I can produce the full TTL implementation (model update, repository changes, index creation).
## Fixing MongoDB Document Casing with BSON Tags
### Why This Change Was Needed
After switching to MongoDB, the stored documents had inconsistent field names:
- Mixed casing (`UserID` vs `userID`)
- Duplicate identifier fields (`_id` and `ID`)
- Uncontrolled serialization because we weren’t using any tags
MongoDB uses BSON (Binary JSON) for storage.
Just like JSON tags, BSON tags let you control:
- Field names
- Optional fields
- Mapping between Go structs and Mongo documents
Adding proper `bson` tags ensures consistent, predictable database documents.
### Updating the Domain Models
Two domain models were updated:
- `TripModel`
- `RideFareModel`
All fields are now explicitly mapped with BSON tags.
### ID Field Fix
The `_id` field problem is solved by:
```go
ID primitive.ObjectID `bson:"_id,omitempty"`
```
This ensures:
- MongoDB writes and reads only one `_id` field
- The Go field `ID` is correctly populated
- The field is omitted on inserts until Mongo assigns it
### Example: Updated RideFareModel
```go
type RideFareModel struct {
    ID                primitive.ObjectID     `bson:"_id,omitempty"`
    UserID            string                 `bson:"userID"`
    PackageSlug       string                 `bson:"packageSlug"`
    TotalPriceInCents float64                `bson:"totalPriceInCents"`
    Route             *types.OsrmApiResponse `bson:"route"`
}
```
### Example: Updated TripModel
```go
type TripModel struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    UserID   string             `bson:"userID"`
    Status   string             `bson:"status"`
    RideFare *RideFareModel     `bson:"rideFare"`
    Driver   *pb.TripDriver     `bson:"driver"`
}
```
### What This Fix Achieves
- Clean, consistent casing across the whole database
- No duplicate ID fields
- Proper marshaling/unmarshaling via BSON
- A stable schema for future queries, indexing, or aggregations
### Verifying the Fix
After dropping the old collections and creating a new trip:
- The documents appear cleanly structured
- Fields follow your desired casing style
- `_id` is correctly mapped
- No redundant fields appear
Example on MongoDB Atlas:
```
{
  _id: ObjectId("..."),
  userID: "123",
  status: "accepted",
  rideFare: {...},
  driver: {...}
}
```
Everything now aligns with idiomatic MongoDB document design and your preferred casing.
## Production Deployment
### Overview
The goal of this module is to take the full microservices system you built in development and deploy it to Google Cloud using Kubernetes (GKE), Google Artifact Registry for container images, and MongoDB Atlas for managed database storage.
You deploy everything manually so you understand each moving part before introducing CI/CD later if you want.
### Required Changes Before Deployment
The production setup mirrors the development one, but with a few extra services and differences:
- Production has its own Kubernetes manifests under `infra/production/k8s/`
- You must build and push Docker images for all Go services to Google Artifact Registry
- RabbitMQ and Jaeger are deployed inside the cluster
- MongoDB is not self-hosted; the trip service connects to MongoDB Atlas
- Secrets must exist in production as Kubernetes secrets
- All manifests must use your actual `PROJECT_ID` and region
### Preparing the Production Environment
You start by fixing up the production folder:
- Add production Dockerfiles for driver and payment service
- Add production K8s deployment files for:
    - api-gateway
    - trip-service
    - driver-service
    - payment-service
    - rabbitmq
    - jaeger
- Add `app-config.yaml` with production environment variables
- Add a `secrets.yaml` (copied from development) that contains:
    - MongoDB connection URI
    - RabbitMQ credentials
    - Stripe credentials
    - External API URLs (OSRM)
Every file that references your image name uses the pattern:
`{REGION}-docker.pkg.dev/{PROJECT_ID}/ride-sharing/<service-name>`
### Building Docker Images
Each service has its own production Dockerfile.
You must build all four images using:
- `-platform linux/amd64` (needed on Apple Silicon)
- The correct production Dockerfile path
- The correct destination tag in Artifact Registry
Example:
`docker build -t europe-west1-docker.pkg.dev/<project>/ride-sharing/api-gateway:latest --platform linux/amd64 -f infra/production/docker/api-gateway.Dockerfile .`
Repeat this for driver, trip, and payment services.
### Pushing to Artifact Registry
Before pushing, authenticate Docker to Artifact Registry:
`gcloud auth configure-docker europe-west1-docker.pkg.dev`
Then push all images:
`docker push europe-west1-docker.pkg.dev/<project>/ride-sharing/<service-name>:latest`
You verify the images appear in Artifact Registry.
### Creating the GKE Cluster
You create a Kubernetes cluster in the Google Cloud console:
- Name: `ride-sharing`
- Region: the same region used for your Artifact Registry containers
After creation, connect your local machine:
`gcloud container clusters get-credentials ride-sharing --region <region> --project <project>`
Now all `kubectl` commands target the cloud cluster.
### Applying the Kubernetes Manifests (Correct Order)
Apply everything **in sequence** to avoid dependency failures.
### 1. Config and Secrets
`kubectl apply -f app-config.yaml`
`kubectl apply -f secrets.yaml`
Secrets include MongoDB, RabbitMQ, Stripe, and external APIs.
### 2. Observability and Messaging
Apply Jaeger and RabbitMQ:
- Jaeger exposes:
    - 16686 for UI
    - 14268 for tracing collector
- RabbitMQ runs as a StatefulSet with:
    - Persistent volume
    - Readiness and liveness probes
    - `rabbitmq:3-management` image
Wait until both pods are fully running.
### 3. API Gateway
The API gateway starts first because the frontend communicates with it.
It pulls environment variables from `app-config` and secrets, and connects to RabbitMQ and Jaeger.
Once its pod shows READY, you retrieve the external IP:
`kubectl get service api-gateway`
This IP is used for frontend testing.
### 4. Business Services
Apply in this order:
- Driver service
- Trip service
- Payment service
Trip service additionally needs:
- `MONGODB_URI` from the production secret
- `OSRM_API` from secrets
- Jaeger endpoint
- RabbitMQ URI
When Trip service starts, MongoDB connection fails at first because Atlas blocks network access.
You must go to Mongo Atlas → Network Access → Allow access from anywhere (0.0.0.0/0).
Once active, the trip service reconnects successfully.
### Verifying the Deployment
To test without deploying a frontend, run the Next.js app locally:
1. Set `NEXT_PUBLIC_API_URL` to the external IP of the API gateway
2. Run `npm install` then `npm run dev`
3. Open [http://localhost:3000](http://localhost:3000/)
4. Request a trip, assign driver, pay through Stripe
5. Watch backend logs via:
    
    `kubectl logs deploy/<service> -f`
    
Everything should flow exactly like in local Tilt.
### Accessing Jaeger or RabbitMQ via Port Forwarding
Because they are ClusterIP-only, you use:
Jaeger UI:
`kubectl port-forward deploy/jaeger 16686:16686`
→ open [http://localhost:16686](http://localhost:16686/)
RabbitMQ:
`kubectl port-forward statefulset/rabbitmq 15672:15672`
→ open [http://localhost:15672](http://localhost:15672/)
You can debug traces and queues on your local machine from the production cluster.
### Optional Frontend Deployment
You can deploy the Next.js frontend for free using Vercel:
- Import the GitHub repo
- Set `NEXT_PUBLIC_API_URL` to your gateway external IP
- Set Stripe publishable key
- Deploy
Vercel auto-detects that it’s a Next.js project.
### Cleanup
To avoid costs, delete the GKE cluster when finished:
`gcloud container clusters delete ride-sharing --region <region>`
This removes all nodes and stops billing.
