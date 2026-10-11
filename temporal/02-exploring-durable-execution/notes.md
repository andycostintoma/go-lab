# Temporal 102: Exploring Durable Execution

## Table of Contents

- [Welcome](#welcome)
  - [Prerequisites](#prerequisites)
  - [Course Goals](#course-goals)
  - [Example Code and Learning Approach](#example-code-and-learning-approach)
  - [Examples Used in the Course](#examples-used-in-the-course)
- [Course Outcomes](#course-outcomes)
  - [Evaluate Durable Execution](#evaluate-durable-execution)
  - [Apply Development Best Practices](#apply-development-best-practices)
  - [Debug Workflow Execution](#debug-workflow-execution)
  - [Prepare for Production](#prepare-for-production)
- [Durable Execution System](#durable-execution-system)
  - [What Is a Durable Execution System?](#what-is-a-durable-execution-system)
  - [Developer Productivity](#developer-productivity)
- [Temporal Application Structure](#temporal-application-structure)
  - [Workflows and Activities](#workflows-and-activities)
  - [Workers and Application Code](#workers-and-application-code)
  - [The Temporal Cluster](#the-temporal-cluster)
  - [Temporal Cloud](#temporal-cloud)
  - [Separation of Application and Cluster](#separation-of-application-and-cluster)
- [How Errors Affect Workflow Execution](#how-errors-affect-workflow-execution)
  - [Activity Errors](#activity-errors)
  - [Workflow Errors](#workflow-errors)
  - [Returning Errors from Application Code](#returning-errors-from-application-code)
  - [Cross-Language Error Handling](#cross-language-error-handling)
- [Backwards-Compatible Evolution of Input Parameters and Return Values](#backwards-compatible-evolution-of-input-parameters-and-return-values)
  - [Example: Struct-Based Activity Input and Output](#example-struct-based-activity-input-and-output)
  - [Compatibility Considerations](#compatibility-considerations)
  - [Optional Self-Study: Using Structs](#optional-self-study-using-structs)
- [Choosing Names for Task Queues](#choosing-names-for-task-queues)
  - [Task Queue Names](#task-queue-names)
  - [Sharing a Task Queue Constant](#sharing-a-task-queue-constant)
  - [Running Multiple Worker Processes](#running-multiple-worker-processes)
- [Choosing Workflow IDs](#choosing-workflow-ids)
  - [Business-Meaningful Identifiers](#business-meaningful-identifiers)
  - [Uniqueness Within a Namespace](#uniqueness-within-a-namespace)
  - [How Temporal Handles a Workflow ID Conflict](#how-temporal-handles-a-workflow-id-conflict)
- [Workflow ID Reuse Policy and Retention Period](#workflow-id-reuse-policy-and-retention-period)
  - [Workflow ID Reuse Policies](#workflow-id-reuse-policies)
  - [Specifying a Workflow ID Reuse Policy](#specifying-a-workflow-id-reuse-policy)
  - [Retention Period](#retention-period)
  - [Archival and Export](#archival-and-export)
  - [Retention Periods Do Not Affect Open Workflows](#retention-periods-do-not-affect-open-workflows)
- [Logging in Workflows and Activities](#logging-in-workflows-and-activities)
  - [Replay-Aware Workflow Logging](#replay-aware-workflow-logging)
  - [Using the Logger Interface](#using-the-logger-interface)
  - [Adding Shared Log Fields](#adding-shared-log-fields)
  - [Activity Logging](#activity-logging)
  - [Limitations of the Default Logger](#limitations-of-the-default-logger)
  - [Integrating Another Logging Implementation](#integrating-another-logging-implementation)
- [Accessing Results](#accessing-results)
  - [Asynchronous Execution Requests](#asynchronous-execution-requests)
  - [Waiting on Execution Results](#waiting-on-execution-results)
  - [Deferring Access to Execution Results](#deferring-access-to-execution-results)
- [Using Appropriate Timeouts](#using-appropriate-timeouts)
  - [Match Timeouts to the Actual Work](#match-timeouts-to-the-actual-work)
  - [Choosing a Start-to-Close Timeout](#choosing-a-start-to-close-timeout)
- [What Is a Timer?](#what-is-a-timer)
  - [Durable Delays](#durable-delays)
  - [Duration and Precision](#duration-and-precision)
- [Use Cases for Timers](#use-cases-for-timers)
- [Pausing Workflow Execution for a Specified Duration](#pausing-workflow-execution-for-a-specified-duration)
- [Running Code at a Specific Point in the Future](#running-code-at-a-specific-point-in-the-future)
- [What Happens to a Timer if the Worker Crashes?](#what-happens-to-a-timer-if-the-worker-crashes)
- [Hands-On Exercise #1: Observing Durable Execution](#hands-on-exercise-1-observing-durable-execution)
  - [Files and Objectives](#files-and-objectives)
  - [Parts A–C: Logging and Timer](#parts-ac-logging-and-timer)
  - [Part D: Observe Worker Recovery](#part-d-observe-worker-recovery)
  - [Inspect the Results](#inspect-the-results)
  - [Optional: Integrate Zap](#optional-integrate-zap)

## Welcome

### Prerequisites

This course is intended for experienced software developers with basic Go proficiency and foundational Temporal knowledge. **Temporal 101** is the recommended prerequisite.

You should already be able to:

- Develop and execute Workflows and Activities in Go.
- Configure and run a Worker.
- Navigate the Temporal Web UI.
- Explain the high-level interactions between a Temporal application and the Temporal Cluster during Workflow Execution.

### Course Goals

Temporal 102 moves beyond basic application development toward skills needed for production deployment. It focuses on:

- Understanding common problems encountered by Temporal developers and why they occur.
- Identifying, solving, and preventing those problems.
- Building a deeper understanding of Temporal through key concepts and best practices.

### Example Code and Learning Approach

The examples use more Go features than Temporal 101, including **structs, pointers, and interfaces**, as well as third-party packages such as **`testify`** for unit tests. Expert Go knowledge is not required.

The code is designed to teach specific Temporal concepts, rather than serve as a production-ready template. Some examples contain **intentional flaws** so that failures can be observed and understood in a learning environment.

### Examples Used in the Course

| Example | Learning purpose |
| :--- | :--- |
| Translation Workflow | Replaces Temporal 101's Spanish greeting/farewell microservice with support for other words and additional languages. The first exercise uses structs for business inputs and results to demonstrate best practices. |
| Pizza-order Workflow | Shows how Workflow code produces **Commands** sent to the Temporal Cluster, which lead to **History Events**. These concepts underpin history replay and recovery after a Worker crash. |
| Loan-processing Workflow | Demonstrates a long-running Workflow, why Workflow code must be **deterministic**, and what can happen when it is not. |

## Course Outcomes

The course covers Temporal across the **full development lifecycle**: testing, debugging, deploying, and updating applications. Its objectives fall into four areas.

### Evaluate Durable Execution

- Explain how **History Replay** recreates the state of a Workflow Execution.
- Use **Timers** to introduce delays and determine how a Worker crash affects them.
- Summarize the relationship between **Temporal SDK API calls, Worker Commands, and History Events**.
- Explain determinism in the context of Workflow Execution and its implications for application code.

### Apply Development Best Practices

- Model business input parameters and return values as **data structures** to support backwards-compatible evolution.
- Use logging to report information from Workflows and Activities during execution.
- Identify, correct, and avoid common sources of Workflow non-determinism.
- Select Workflow IDs appropriate to the use case.
- Develop and execute unit tests for Workflows and Activities.
- Use **mocks** to isolate Workflow tests from Activity implementations.

### Debug Workflow Execution

- Analyze **Event History** to diagnose execution problems.
- Use the Web UI to trace both open and closed Workflow Executions.
- View the stack trace of a current Workflow Execution.
- Distinguish the different states of Workflows and Activities.
- Explain the role of **Sticky Task Queues** in Workflow Execution.
- Describe how returning errors from Workflows and Activities affects execution.
- Observe the consequences of running non-deterministic Workflow code.
- **Reset** a Workflow Execution to recover from a bad deployment.
- Distinguish **cancellation** from **termination**.
- Terminate Workflows using the CLI, SDK, and Web UI.

### Prepare for Production

- Illustrate the logical and physical views of a production Temporal application deployment.
- Configure a Temporal Client to connect to **Temporal Cloud** or a secure self-hosted cluster.
- Identify which changes can be safely deployed **without versioning**.

## Durable Execution System

### What Is a Durable Execution System?

Temporal is a **durable execution system**: it preserves application execution state so code can recover automatically from failures and continue making progress.

Failures can range from a **network timeout** to a **kernel panic on an application server**. Rather than requiring the application to reconstruct its progress from scratch, Temporal retains the information needed to recover Workflow state.

Durable execution supports reliable orchestration; application code still needs correct business logic and retry-safe external side effects.

### Developer Productivity

Developers normally spend substantial effort handling failures and timeouts. Temporal provides **higher-level abstractions**, recovery mechanisms, and built-in scalability so more of that effort can go toward business logic.

The **Web UI** also makes execution observable. It lets you inspect past and current Workflow Executions, including their:

- Event History.
- Input parameters.
- Return values, when available.

In this course, the Web UI is used to **debug Workflow Execution and verify that a fix resolves the problem**.

## Temporal Application Structure

### Workflows and Activities

A **Workflow** defines the sequence of steps that makes up an application's main business logic. Temporal provides language-specific **SDKs** with APIs and libraries for implementing applications in general-purpose languages. In this course, Workflow and Activity Definitions are Go functions using the Go SDK.

Workflow code must be **deterministic**. “The same input produces the same output” is an introductory explanation, but the more precise requirement is **replay compatibility: Workflow code must produce the sequence of Commands expected by the recorded Event History**. The course explores why this matters, what happens when it is violated, and how to avoid non-determinism.

**Activities** encapsulate business logic that is non-deterministic or prone to failure, such as external I/O. Workflows schedule Activities, and Temporal automatically retries retryable failures according to their Retry Policy and applicable timeouts. This handles transient failures without a custom retry loop in Workflow code.

### Workers and Application Code

The **Temporal Cluster orchestrates execution; it does not execute application code**. Workers poll Task Queues and execute Workflow and Activity code.

The SDK provides the Worker implementation. The developer writes code to:

1. Configure a Worker with a Task Queue name.
2. Register the Workflow and Activity Definitions it can execute.
3. Start the Worker.

The course uses **Temporal Application Code** to collectively mean:

- Workflow Definitions.
- Activity Definitions.
- Worker configuration and startup code.

A complete Temporal application includes both this developer-written code and the supporting functionality supplied by the SDK.

### The Temporal Cluster

A self-hosted **Temporal Cluster** consists of the Temporal Server—its Frontend Service and backend services—plus the database used for persistence. Its support enables durable, scalable, and reliable Workflow Execution.

Optional supporting components can include:

- **Elasticsearch** for advanced visibility, depending on the deployment's visibility configuration; supported SQL visibility is another option.
- **Grafana** for operational dashboards showing cluster and application health, backed by a metrics data source.

### Temporal Cloud

**Temporal Cloud** provides a managed service fulfilling the same orchestration and persistence roles as a self-hosted cluster. It removes the need for your team to deploy, secure, operate, and maintain the Temporal Server infrastructure.

The operational responsibilities differ, but the core Workflow programming and execution model is the same. In this course, **“Temporal Cluster” generally also applies to Temporal Cloud**, unless a distinction is explicitly made.

Moving an application from a self-hosted cluster to Temporal Cloud typically requires minimal code changes, primarily to the Temporal Client's connection settings, such as endpoint, namespace, and authentication. The course covers these settings later.

### Separation of Application and Cluster

**Application code executes outside the Temporal Cluster**, whether the cluster is self-hosted or provided by Temporal Cloud.

```text
Temporal application                         Temporal Cluster / Cloud
┌──────────────────────────────┐             ┌──────────────────────────┐
│ Developer-written code       │             │ Frontend Service         │
│ • Workflow Definitions       │  polls and  │ Backend services         │
│ • Activity Definitions       │◄───────────►│ Task Queues              │
│ • Worker configuration       │ reports     │ Persistent Event History │
│                              │ results     │ and execution state      │
│ SDK-provided Worker + Client │             │                          │
└──────────────────────────────┘             └──────────────────────────┘
```

In production, Workers and the cluster typically run on separate machines or containers and can even run in different data centers. Workers must be able to reach the cluster's Frontend Service. Using Temporal Cloud manages the cluster infrastructure; your team still deploys and operates the application Workers.

## How Errors Affect Workflow Execution

Returning an error signals failure, but **Activity attempt failure and Workflow Execution failure have different consequences**.

| Where the error is returned | Default behavior |
| :--- | :--- |
| Activity Definition | The attempt fails; Temporal retries retryable errors according to the Activity's Retry Policy and applicable timeouts. |
| Workflow Definition | The Workflow Execution fails; there is no Workflow Retry Policy by default. |

### Activity Errors

Activities perform error-prone operations, such as querying a database. Failures might involve a temporary outage, rejected credentials, or invalid SQL.

The default Activity Retry Policy uses **exponential backoff**:

```text
1s → 2s → 4s → 8s → 16s → 32s → 64s → 100s → 100s → …
```

These are delays between attempts, not the duration of the attempts themselves. By default, there is no attempt-count limit. Retries can stop when the Activity succeeds, is canceled, encounters a non-retryable error, reaches an applicable overall timeout, or its Workflow closes. A Start-to-Close Timeout limits one attempt and can trigger another retry; it does not limit the entire retry sequence.

Temporary outages may resolve without intervention. Persistent problems require a fix: for example, correct credentials or SQL, then deploy and restart the Activity Workers if code changed. A subsequent attempt can use the corrected implementation, provided retries are still allowed and the updated Worker receives the task.

Customize retries through **`workflow.ActivityOptions.RetryPolicy`**. Different calls to `workflow.ExecuteActivity` can use different options and policies to match their business requirements.

**One `ExecuteActivity` call can result in multiple executions of the Activity code.** Make external side effects **idempotent**: repeating the same logical operation should have the same effect as performing it once. For operations such as payments, this commonly requires an idempotency key or deduplication in the external system.

While retries continue, the Activity Future remains pending. If the Activity ultimately fails, its Future's `Get` returns an error to Workflow code. The Workflow can handle that error with a fallback or compensation; Activity failure does not automatically mean the Workflow must fail.

### Workflow Errors

Workflows have **no Retry Policy by default**, and adding one is unusual. If the Workflow Definition returns an error, the Workflow Execution is marked as failed and is not automatically retried under the default configuration.

Activities are retried because external operations often encounter transient failures. Workflow code instead performs deterministic orchestration and should not directly perform those external operations. It decides how to handle an Activity's eventual failure and whether to return an error itself.

Prefer correcting a recoverable underlying problem while execution can still make progress, rather than failing the Workflow unnecessarily. Return an error when failure is the intended business outcome or recovery cannot be completed.

**Workflow Execution failure is distinct from Workflow Task failure.** A Worker crash or a Workflow-code problem such as a non-determinism error does not, by itself, mean the Workflow Definition returned an error and closed the execution as failed. Such problems can prevent progress while the execution remains open.

### Returning Errors from Application Code

Use ordinary Go errors; a Temporal-specific API is not required. For example, an Activity making an HTTP request can return the original error:

```go
resp, err := http.Get(url)
if err != nil {
	return "", err
}
defer resp.Body.Close()
```

This fragment assumes a function returning `(string, error)` and continues with response handling after the shown code. Returning the original error preserves its diagnostic information, whereas replacing it with `errors.New("request failed")` discards the cause. In an Activity, prefer an HTTP request using the Activity's context so cancellation propagates to the request.

The SDK's **`temporal` package** also provides error constructors for explicit application error types and **non-retryable errors**. Use these when the failure needs semantics beyond a generic Go error, such as indicating that another Activity attempt cannot resolve invalid input.

### Cross-Language Error Handling

The Worker automatically converts returned errors into Temporal's **language-neutral failure representation**. The cluster can persist failure details consistently, and SDKs can decode them into their own language's error model.

This supports applications spanning multiple SDKs—for example, a TypeScript client starting a Go Workflow that schedules Python and Java Activities. Each Activity still needs a suitable Worker and Task Queue routing; language-neutral failures allow errors to cross these boundaries without depending on a shared native error type.

## Backwards-Compatible Evolution of Input Parameters and Return Values

**Use a single struct for business input and a struct for the result** of Workflow and Activity Definitions. This lets their data evolve without changing the function signature.

- A Workflow must take **`workflow.Context`** as its first parameter.
- An Activity should take **`context.Context`** as its first parameter, although it is not required. It enables cancellation notification and other Activity SDK features.
- These SDK-supplied contexts are separate from the serialized business input.

Both definitions can accept multiple business parameters, but changing their **number, order, or types** can break compatibility with existing executions and callers. Adding fields to an input struct avoids changing the parameter list. A result struct similarly allows additional output fields without changing the return type.

### Example: Struct-Based Activity Input and Output

Temporal 101's greeting Activity accepted a name and returned a Spanish greeting:

```go
func GreetInSpanish(ctx context.Context, name string) (string, error)
```

Supporting additional languages by adding a separate `languageCode string` parameter would change that signature. Starting with structs gives the data room to evolve:

```go
type GreetingInput struct {
	Name         string
	LanguageCode string
}

type GreetingOutput struct {
	Greeting string
}
```

The Activity now has a stable struct-based signature. With `context` and `fmt` imported, this illustrative implementation supports French and returns an error for other language codes:

```go
func GetTranslatedGreeting(ctx context.Context, input GreetingInput) (GreetingOutput, error) {
	if input.LanguageCode == "fr" {
		return GreetingOutput{
			Greeting: fmt.Sprintf("Bonjour, %s", input.Name),
		}, nil
	}

	return GreetingOutput{}, fmt.Errorf("unsupported language code: %q", input.LanguageCode)
}
```

Additional languages can be implemented without adding function parameters. Additional result fields can be added to `GreetingOutput` as requirements evolve. In a real Activity, classify unsupported input appropriately if retrying it cannot help; the ordinary error above is retryable by default.

### Compatibility Considerations

**Structs support compatible evolution; they do not make every change compatible.** With the default JSON encoding, adding a field is generally compatible at the serialization level, but the implementation must still handle older payloads:

- A field absent from stored input decodes to its Go **zero value**. Define a suitable default or otherwise preserve the behavior expected by older executions.
- Renaming fields, changing their types, or making a new field mandatory can break compatibility.
- Result changes must also remain compatible with callers and Workflows consuming older recorded results.
- Keep serialized fields exported so the default JSON converter can encode them.

The initial migration from `string` inputs/results to structs **is not backwards-compatible** with the old payload shape. Adopt the struct-based design early rather than assuming an existing string-based execution can decode into the new type.

Payload compatibility is also separate from Workflow determinism: a stable function signature does not make a change to the Workflow's Command sequence replay-compatible.

### Optional Self-Study: Using Structs

The **`01-using-structs/`** example demonstrates struct-based inputs and results for both Workflow and Activity Definitions. It contains `microservice/`, `start/`, and `worker/` directories for exploring the complete application flow.

## Choosing Names for Task Queues

### Task Queue Names

The Temporal Cluster maintains **Task Queues** that Workers poll for work. The client specifies a Task Queue when starting a Workflow, and a suitable Worker must poll that same queue in the same namespace.

These excerpts show the matching values in starter and Worker code:

```go
options := client.StartWorkflowOptions{
	ID:        "my-workflow",
	TaskQueue: "my-task-queue-name",
}

run, err := c.ExecuteWorkflow(ctx, options, ProcessOrderWorkflow, input)
```

```go
w := worker.New(c, "my-task-queue-name", worker.Options{})
```

**Task Queues are created dynamically when first used.** A mismatch does not immediately produce a “queue not found” error. Instead, the client and Worker use different queues: the Worker can still communicate with the cluster, but it does not receive tasks from the queue selected by the client. Without a suitable Worker polling that queue, the Workflow cannot progress.

### Sharing a Task Queue Constant

When starter and Worker code can share a package, define the queue name **once as a constant**:

```go
package app

const TaskQueueName = "my-task-queue-name"
```

Reference it in the starter:

```go
options := client.StartWorkflowOptions{
	ID:        "my-workflow",
	TaskQueue: app.TaskQueueName,
}

run, err := c.ExecuteWorkflow(ctx, options, ProcessOrderWorkflow, input)
```

And in the Worker:

```go
w := worker.New(c, app.TaskQueueName, worker.Options{})
```

This prevents the two values from drifting apart. Our **`01-using-structs/`** example follows this pattern: its starter and Worker both reference `translation.TaskQueueName`.

A shared Go constant is not always practical, particularly when the starter is maintained separately or written in another language. In that case, keep the exact queue name consistent across their configuration.

### Running Multiple Worker Processes

For production, run **at least two Worker processes for each Task Queue**, with compatible definitions registered. If one crashes, another can recover Workflow state and continue processing tasks, including tasks for new executions.

- **Availability:** A single Worker process is no longer the sole execution resource for that queue. Place replicas on separate machines or failure domains when protection from host failure is needed.
- **Scalability:** Additional Workers increase processing capacity, subject to cluster and downstream-system limits.
- **Recovery:** Workflow state can be reconstructed from Event History. An in-flight Activity on a crashed Worker may need to time out and be retried; recovery is not necessarily immediate.

Completed Activities whose results are recorded in history are not re-executed merely because Workflow state is replayed. Activities that may be retried still need retry-safe side effects.

## Choosing Workflow IDs

### Business-Meaningful Identifiers

A **Workflow ID** should identify the business process represented by the execution. Including a business identifier makes the execution easier to find and lets repeated start requests refer to the same logical operation.

Examples:

- Order processing: `process-order-number-90743812`.
- Loan management: `manage-loan-account-28430614`.

This starter excerpt assumes `input.OrderNumber` is a string:

```go
options := client.StartWorkflowOptions{
	ID:        "process-order-number-" + input.OrderNumber,
	TaskQueue: app.TaskQueueName,
}

run, err := c.ExecuteWorkflow(ctx, options, ProcessOrderWorkflow, input)
```

Use a stable ID for the same business operation rather than generating a new ID for every repeated request. This prevents concurrent duplicate Workflow Executions with that ID; it does not replace idempotency for external Activity side effects.

### Uniqueness Within a Namespace

Temporal allows **at most one running Workflow Execution with a given Workflow ID in a namespace**.

- The constraint applies **across all Workflow Types** in that namespace.
- An order Workflow and a loan Workflow cannot run simultaneously with the same ID, even though their types differ.
- The same ID can be used independently in another namespace.

Prefixes such as `process-order-number-` and `manage-loan-account-` help distinguish business processes and avoid unintended collisions.

**Workflow ID and Run ID are different:** the Workflow ID identifies the business process, while a Run ID uniquely identifies a particular run. Whether an ID can be reused after an execution closes is controlled separately by the **Workflow ID reuse policy**.

### How Temporal Handles a Workflow ID Conflict

With the Go SDK's default `ExecuteWorkflow` behavior, attempting to start a Workflow whose ID is already running **does not start a second execution**. Instead, the call returns a handle to the existing execution. The new request's input does not update the running Workflow.

To receive an error instead, set this option:

```go
options := client.StartWorkflowOptions{
	ID:                                       "process-order-number-" + input.OrderNumber,
	TaskQueue:                                app.TaskQueueName,
	WorkflowExecutionErrorWhenAlreadyStarted: true,
}
```

The SDK also supports an explicit **`WorkflowIDConflictPolicy`** for choosing how to handle an already-running ID, including using the existing execution or terminating it and starting another. Choose the behavior deliberately for the business use case. Conflict policy concerns a running execution; reuse policy concerns an ID whose previous execution has closed.

## Workflow ID Reuse Policy and Retention Period

### Workflow ID Reuse Policies

Only one Workflow Execution with a given Workflow ID can run within a namespace at a time. A **Workflow ID Reuse Policy** further controls whether a new execution can use the ID of a **closed execution** whose data Temporal still retains.

| Policy | Behavior |
| :--- | :--- |
| `WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE` | **Default:** allow reuse after the previous execution closes, regardless of its outcome. |
| `WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY` | Allow reuse only if the previous execution ended unsuccessfully: failed, timed out, canceled, or terminated. Reject reuse after successful completion. |
| `WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE` | Reject reuse regardless of the previous execution's outcome, while its data is retained. |
| `WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING` | **Deprecated:** terminate an execution currently running with the ID and start a new one; otherwise allow reuse after closure. |

The course presents `TERMINATE_IF_RUNNING` as a fourth reuse option, but the [official Temporal API definition](https://github.com/temporalio/api/blob/master/temporal/api/enums/v1/workflow.proto) marks it as **deprecated**. Its recommended replacement separates closed-ID reuse from running-ID conflicts:

```go
WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
WorkflowIDConflictPolicy: enums.WORKFLOW_ID_CONFLICT_POLICY_TERMINATE_EXISTING,
```

These fields belong in `client.StartWorkflowOptions`. **`WorkflowIDConflictPolicy` requires Temporal Server v1.24.0 or later.** When using the older `TERMINATE_IF_RUNNING` option, leave `WorkflowIDConflictPolicy` unspecified.

### Specifying a Workflow ID Reuse Policy

Import the API enums alongside the SDK client package:

```go
import (
	enums "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
)
```

Set **`WorkflowIDReusePolicy`** in the start options. This excerpt allows reuse only after an unsuccessful execution:

```go
options := client.StartWorkflowOptions{
	ID:                    "example-workflow-id",
	TaskQueue:             shared.TaskQueueName,
	WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY,
}

we, err := c.ExecuteWorkflow(ctx, options, example.MyWorkflow)
```

The Temporal CLI also supports specifying a reuse policy when starting a Workflow; consult `temporal workflow start --help` for the values accepted by the installed CLI version.

### Retention Period

The **Retention Period** is the namespace-level duration after a Workflow Execution closes before its stored data becomes eligible for automatic deletion. Retention limits storage growth and the operational cost of keeping closed execution data.

The course describes **1–30 days** as a typical range. Supported ranges and longer-retention options depend on the deployment and version; choose retention according to operational and access requirements.

**Reuse policies are not permanent business uniqueness constraints.** Temporal can enforce them only while it retains the previous execution's data. After that data is deleted, even `REJECT_DUPLICATE` cannot prevent a new execution from using the same ID based on the old execution. If an operation must never be repeated, preserve that business rule in a durable application record or external deduplication mechanism.

### Archival and Export

Long-term audit requirements may call for execution records months or years after closure. Temporal offers ways to preserve history outside normal cluster retention:

- **Self-hosted Archival:** When enabled and configured, copies execution history and/or visibility data to supported long-term storage, such as a filesystem or Amazon S3, before deletion from active persistence.
- **Temporal Cloud Export:** The course describes a configured export of namespace history data to S3 on an hourly basis. Availability and configuration depend on the Cloud offering.

Preserving history externally supports long-term inspection; **archived or exported records do not extend Workflow ID reuse enforcement** in active persistence.

### Retention Periods Do Not Affect Open Workflows

**The retention countdown begins when a Workflow Execution closes**, not when it starts. A five-day retention period does not delete the state of a Workflow that runs for ten years.

Open executions retain the state and history needed to continue running. This is separate from Event History size limits: a long-running Workflow may still need **Continue-As-New** to keep each run's history bounded.

## Logging in Workflows and Activities

### Replay-Aware Workflow Logging

Logging helps trace the execution path at runtime. The Go SDK provides a **`log.Logger` interface**, a basic default implementation, and accessors for Workflow and Activity loggers.

During **History Replay**, a Worker reconstructs Workflow state by re-executing Workflow code against recorded Events. **`workflow.GetLogger(ctx)` suppresses Workflow log messages during replay by default**, preventing duplicate output from steps that already ran.

Calling a third-party logger directly from Workflow code bypasses that replay-aware wrapper. Instead, configure the third-party implementation through the Temporal Client and continue using the Workflow logger accessor.

### Using the Logger Interface

The interface exposes four levels, from least to most severe:

```text
Debug → Info → Warn → Error
```

Get the logger near the start of the Workflow Definition. Each method accepts a message followed by optional **key-value pairs**:

```go
logger := workflow.GetLogger(ctx)

logger.Debug("Preparing to execute an Activity")
logger.Info("Calculated cost of order", "Tax", tax, "Total", total)
```

These excerpts assume the referenced business variables are already defined. Keep structured fields in key-value pairs so the configured logging implementation can process them consistently.

### Adding Shared Log Fields

Import **`go.temporal.io/sdk/log`** and use **`log.With`** to attach fields to every message from a derived logger:

```go
logger := log.With(workflow.GetLogger(ctx), "version", 4, "user", "jsmith")

logger.Debug("Preparing to execute an Activity")
logger.Info("Calculated cost of order", "Tax", tax, "Total", total)
```

Both messages include `version` and `user`; the second also includes `Tax` and `Total`. Deriving from `workflow.GetLogger(ctx)` preserves replay-aware logging.

### Activity Logging

Import **`go.temporal.io/sdk/activity`** and retrieve the logger using the Activity's `context.Context`:

```go
logger := activity.GetLogger(ctx)

logger.Info("Looking up customer in the database", "Key", customerID)
logger.Error("Database connection failed")
```

Completed Activities with recorded results are **not re-executed during Workflow replay**, so their logger does not need replay suppression. Activity retries are actual new attempts and can legitimately produce repeated log messages.

Using the SDK accessor provides a consistent logging API and automatically includes execution context, such as Activity Type and Task Queue, as structured fields.

### Limitations of the Default Logger

The SDK's default logger is intentionally basic:

- Writes to **standard output** of the process using it. Workflow and Activity messages appear in the Worker terminal unless that output is collected or redirected.
- Emits all four levels without a configurable minimum-level threshold.
- Offers limited formatting and destination control, with execution metadata included alongside the message.

The course describes metadata appearing before the message, but formatting is version-dependent. In this project's **Go SDK v1.37.0**, the [default logger implementation](https://github.com/temporalio/sdk-go/blob/v1.37.0/internal/log/default_logger.go) prints the level and message before the key-value fields.

### Integrating Another Logging Implementation

Provide a custom implementation through **`client.Options.Logger`**. This illustrative excerpt assumes `NewAlternateLogger` is an application-provided constructor returning an SDK-compatible logger:

```go
customLogger := NewAlternateLogger()
c, err := client.Dial(client.Options{
	Logger: customLogger,
})
if err != nil {
	return err
}
defer c.Close()
```

Configure destinations, formatting, and minimum log levels through the chosen logging package as in any other Go application. The supplied logger must implement the SDK's **`log.Logger`** interface: `Debug`, `Info`, `Warn`, and `Error`, each accepting a message and variadic key-value fields.

If a logging package does not implement that interface, use an **adapter**. Temporal's [Zap adapter example](https://github.com/temporalio/samples-go/tree/main/zapadapter) demonstrates integrating Zap. Workflow code should still call `workflow.GetLogger(ctx)`, and Activity code should call `activity.GetLogger(ctx)`.

## Accessing Results

### Asynchronous Execution Requests

Workflows can run for months or years, and individual Activities can also be long-running. Their execution is **asynchronous to the code that requests it**: requesting execution does not run the definition locally or wait for it to finish.

From client code:

```go
run, err := c.ExecuteWorkflow(ctx, options, MyWorkflow, input)
if err != nil {
	return err
}
```

From Workflow code, using a context with Activity options configured:

```go
future := workflow.ExecuteActivity(ctx, MyActivity, input)
```

The two return different kinds of handles:

| API | Returned handle | What the call waits for |
| :--- | :--- | :--- |
| `c.ExecuteWorkflow` | `client.WorkflowRun`, a future-like handle | The server's response to the start request, **not Workflow completion**. |
| `workflow.ExecuteActivity` | `workflow.Future` | Returns a Future without waiting for Activity completion; the SDK emits a scheduling Command through Workflow Task processing. |

The course calls these “non-blocking” to emphasize that they do not wait for execution results. A client start request still involves a network round trip and can fail before returning a handle.

### Waiting on Execution Results

Call **`Get`** to wait for a result or error. Pass a pointer to a variable matching the execution's output type so the SDK can decode the successful result into it.

For a Workflow result in client code:

```go
var result string
if err := run.Get(ctx, &result); err != nil {
	return err
}
```

For an Activity result in Workflow code, requesting execution and waiting are commonly chained:

```go
var result string
if err := workflow.ExecuteActivity(ctx, MyActivity, input).Get(ctx, &result); err != nil {
	return err
}
```

The chained expression waits because of **`Get`**, not because `ExecuteActivity` itself waits for completion. While retryable Activity failures are being retried, the Future remains pending.

In Workflow code, `Get` **yields through the SDK's deterministic scheduler** while waiting; it does not hold an operating-system thread blocked for the Activity's duration. Always check the error before treating the result as successful. A wait can also end because its context is canceled or, for client calls, reaches a deadline.

### Deferring Access to Execution Results

For **independent Activities**, request all executions before waiting on their results. This allows them to run concurrently without adding ordinary Go goroutines to Workflow code:

```go
futureA := workflow.ExecuteActivity(ctx, MyActivityA, inputA)
futureB := workflow.ExecuteActivity(ctx, MyActivityB, inputB)
futureC := workflow.ExecuteActivity(ctx, MyActivityC, inputC)

var resultA, resultB, resultC string
errA := futureA.Get(ctx, &resultA)
errB := futureB.Get(ctx, &resultB)
errC := futureC.Get(ctx, &resultC)
```

This excerpt assumes each Activity returns a string. Handle `errA`, `errB`, and `errC` before using the corresponding results.

**Waiting in A–B–C order does not force the Activities to execute in that order.** All three have already been scheduled; B and C can finish while the Workflow waits for A.

If each Activity takes approximately `T` and sufficient Worker and downstream capacity is available:

| Scheduling approach | Approximate total Activity time |
| :--- | :--- |
| Schedule A, wait; schedule B, wait; schedule C, wait | `3 × T` |
| Schedule A, B, and C; then wait for results | `T`, plus scheduling overhead |

Use this pattern only when the Activities do not depend on one another's results or require ordered side effects. Requesting concurrent execution enables parallel processing; available capacity determines the actual speedup.

## Using Appropriate Timeouts

### Match Timeouts to the Actual Work

Copying working examples can speed up development, but **their timeout values are not production defaults**. Reassess timeouts whenever an Activity's work changes—for example, when a simple greeting becomes a remote-service call, file-processing operation, or database query.

A five-second timeout may suit a greeting that normally completes in a fraction of a second, but not a substantially longer operation:

```go
options := workflow.ActivityOptions{
	StartToCloseTimeout: 5 * time.Second,
}
ctx = workflow.WithActivityOptions(ctx, options)

var result string
err := workflow.ExecuteActivity(ctx, MyActivity, input).Get(ctx, &result)
```

This excerpt assumes `time` and the SDK's `workflow` package are imported. Check `err` before using `result`.

### Choosing a Start-to-Close Timeout

**Start-to-Close Timeout bounds one Activity attempt**, measured from when a Worker starts its Activity Task. Set it **slightly longer than the slowest successful attempt you reasonably expect**, allowing for normal latency variation.

| Timeout choice | Consequence |
| :--- | :--- |
| Too short | Legitimate attempts time out, triggering unnecessary retries and potentially repeated external side effects. |
| Too long | A crashed Activity Worker can go undetected longer, delaying a retry and recovery and reducing effective throughput. |
| Matched to expected duration | Normal attempts have enough time to finish, while lost attempts are detected within a useful recovery window. |

Temporal cannot rely on a crashed Worker to report failure. Expiration of the attempt's timeout allows the server to detect missing completion and retry according to the Retry Policy. For **long-running Activities**, heartbeats with a **Heartbeat Timeout** can detect a lost or stalled Worker sooner than a long Start-to-Close Timeout alone.

Start-to-Close does **not** include time waiting in the Task Queue or bound the entire retry sequence. Use **Schedule-to-Close Timeout** when the Activity needs an overall deadline spanning queueing and retries.

A timeout also does not guarantee that an external operation has stopped. The original operation may still complete after Temporal has scheduled a retry, so side effects must remain idempotent or deduplicated.

## What Is a Timer?

### Durable Delays

A **durable Timer** introduces a delay in a Workflow Execution. Code awaiting the Timer pauses until it fires, while other independent Workflow work can continue.

**Timers are maintained by the Temporal Cluster**, rather than by a sleeping Worker thread. Waiting yields through the Workflow SDK, so no Worker thread needs to remain occupied for the duration of the delay. A Worker crash does not erase the Timer; a suitable Worker can resume the Workflow after it fires.

### Duration and Precision

A Timer is scheduled with a **fixed duration**, which can range from seconds to several years.

Avoid relying on **sub-second Timer precision**. Although an API may accept a duration shorter than one second, Worker–cluster round trips and scheduling latency affect when Workflow code actually resumes. Durable Timers are suitable for reliable delays, not precise real-time scheduling.

## Use Cases for Timers

Timers introduce delays when the business process requires time to pass before the next step.

| Use case | Example |
| :--- | :--- |
| **Fixed intervals** | A customer-onboarding Workflow sends email reminders one day, one week, and one month after signup. Calculate each delay relative to the intended reminder time, rather than accidentally adding those intervals together. |
| **Dynamically calculated intervals** | A previous Activity returns a value that determines how long the Workflow should wait before scheduling the next Activity. The delay can vary between iterations. |
| **Allowing time for offline steps** | A payment-processing Workflow waits to allow a paper check to clear or paperwork to arrive by courier or mail. |

A Timer indicates that **time has elapsed**, not that an external step has completed. If proceeding requires confirmation that a check cleared or documents arrived, obtain that confirmation separately before continuing.

## Pausing Workflow Execution for a Specified Duration

Use **`workflow.Sleep`** to wait for a durable Timer before continuing along the current Workflow code path. It is the **Workflow-safe replacement for Go's `time.Sleep`**.

```go
if err := workflow.Sleep(ctx, 10*time.Second); err != nil {
	return err
}
```

This excerpt assumes a Workflow returning only `error`, with `time` and the SDK's `workflow` package imported. For a Workflow returning a result and an error, return the appropriate zero-value result alongside the error.

`workflow.Sleep` takes two arguments:

1. **`ctx`:** The `workflow.Context` supplied to the Workflow, or a context derived from it.
2. **Duration:** How long the Timer should wait, here ten seconds.

The call yields through the Workflow SDK until the Timer **fires or is canceled**. A normal expiration returns `nil`; cancellation returns an error that the Workflow should handle. It does not occupy a Worker thread during the delay, and unrelated Workflow work can continue.

Use `workflow.Sleep`, not `time.Sleep`, inside Workflow code so the delay is managed durably by Temporal and remains compatible with History Replay.

## Running Code at a Specific Point in the Future

Use **`workflow.NewTimer`** to set a durable Timer and defer waiting for it. It is the **Workflow-safe counterpart to Go's `time.NewTimer`**.

```go
timerFuture := workflow.NewTimer(ctx, 5*time.Second)
logger.Info("The timer was set")

if err := timerFuture.Get(ctx, nil); err != nil {
	return err
}
logger.Info("The timer has fired")
```

This excerpt assumes a Workflow returning only `error` and a logger obtained through `workflow.GetLogger(ctx)`.

Like `workflow.Sleep`, `workflow.NewTimer` accepts a **`workflow.Context`** and a **duration**. Unlike `Sleep`, it returns a **`workflow.Future` immediately**, letting subsequent code run before waiting.

- The Future becomes ready when the Timer **fires or is canceled**.
- **`Get(ctx, nil)`** yields until it is ready. Pass `nil` because a Timer has no result payload to decode.
- Check the returned error: a canceled Timer must not be treated as having fired successfully.

| API | Setting and waiting |
| :--- | :--- |
| `workflow.Sleep` | Sets the Timer and waits in one call. |
| `workflow.NewTimer` | Sets the Timer and returns a Future; waiting is a separate operation. |

Use `NewTimer` when the Workflow needs to perform other orchestration between setting the Timer and awaiting it. The duration is relative; the code after a successful `Get` runs once a Worker processes the Timer's firing, rather than at a guaranteed exact wall-clock instant.

## What Happens to a Timer if the Worker Crashes?

**A Timer continues counting down even when no Workers are running.** It is maintained by the Temporal Cluster, not by the Worker process that requested it.

The execution flow is:

1. Workflow code calls `workflow.Sleep` or `workflow.NewTimer`.
2. The Worker reports a Command requesting a Timer with the specified duration, and the cluster persists it.
3. When the Timer expires, the cluster records its firing and schedules a Workflow Task so execution can continue.
4. A suitable Worker processes the task, reconstructing Workflow state from history if needed, and resumes the code awaiting the Timer.

Once the cluster has persisted the Timer, a Worker crash **does not restart its duration**. Consider a ten-second Timer whose Worker crashes three seconds after it starts:

| Worker recovery | Timer behavior |
| :--- | :--- |
| Worker restarts two seconds after the crash | About five seconds remain. The Workflow waits for the original Timer to fire. |
| Worker restarts twenty minutes after the crash | The Timer has already fired. The Worker can continue past the wait without starting another ten-second delay. |

The cluster fires the Timer at expiration regardless of Worker availability, but **Workflow code cannot resume until a suitable Worker processes the resulting task**. Worker outages and task-processing latency can therefore make the observed wait longer than the requested duration.

Design the business logic to tolerate delayed continuation—for example, by checking whether a reminder or follow-up action is still appropriate after a long outage. A durable Timer preserves the scheduled delay; it does not guarantee that application code runs at an exact instant.

## Hands-On Exercise #1: Observing Durable Execution

### Files and Objectives

This exercise adds Workflow and Activity logging, inserts a Timer between translation calls, and demonstrates that a second Worker can complete an execution after the original Worker stops.

The exercise code is in **`02-durable-execution/`**. It includes Workflow and Activity logging and a ten-second Timer between translation calls.

| File | Role |
| :--- | :--- |
| `02-durable-execution/workflow.go` | `SayHelloGoodbye` translates Hello, waits ten seconds, then translates Goodbye. |
| `02-durable-execution/activity.go` | `TranslateTerm` calls the translation service and logs its work. |
| `02-durable-execution/shared.go` | Input/result structs and Task Queue constant `durable-exec-tasks`. |
| `02-durable-execution/microservice/translation-service.go` | HTTP translation service on port `9998`; logs each successful translation. |
| `02-durable-execution/worker/main.go` | Registers the Workflow and Activity and polls the shared Task Queue. |
| `02-durable-execution/start/main.go` | Starts `durable-execution-workflow` using name and language arguments, then waits for and prints the result. |

### Parts A–C: Logging and Timer

Review these implementation steps in the local files:

**A. Workflow logging (`workflow.go`):**

1. Obtain a logger with `workflow.GetLogger(ctx)` near the start of the Workflow.
2. Log Workflow invocation at **Info** level, including the input name as a key-value pair.
3. Before each Activity call, log at **Debug** level which term will be translated and its language code.

**B. Activity logging (`activity.go`):**

1. Import `go.temporal.io/sdk/activity` and obtain the logger with `activity.GetLogger(ctx)`.
2. Log invocation at **Info** level, including the term and language code.
3. Log the successful translation at **Debug** level, including the translated value.
4. Optionally log errors at **Error** level before returning them; the local implementation includes these statements.

**C. Timer (`workflow.go`):**

After constructing `helloMessage` and before preparing the Goodbye Activity, log **`Sleeping between translation calls`** at Debug level and call `workflow.Sleep(ctx, 10*time.Second)`. This creates a window in which to stop the Worker. When writing or extending this code, handle the error from `workflow.Sleep` so cancellation does not fall through to the next Activity.

### Part D: Observe Worker Recovery

Run all Makefile commands from the course directory, **`02-exploring-durable-execution`**. Read the complete sequence first so you can react during the ten-second delay.

1. Stop any existing exercise Workers before starting the two below. Stop the previous translation service if it is using port `9998`, so this exercise's service can run and show its translation logs.
2. Start the Docker-based Temporal development server:

   ```bash
   make temporal-up
   ```

3. In **terminal 1**, start the translation service and leave it running:

   ```bash
   make 02-durable-execution-service
   ```

4. In **terminal 2**, start the first Worker:

   ```bash
   make 02-durable-execution-worker
   ```

5. In **terminal 3**, start a second Worker with the same command:

   ```bash
   make 02-durable-execution-worker
   ```

6. In **terminal 4**, start the Workflow using your name and a supported language code:

   ```bash
   make 02-durable-execution-start NAME=Andy LANG=sk
   ```

   The default name is `Andy` and the default language is French (`fr`). Supported codes are `de`, `es`, `fr`, `lv`, `mi`, `sk`, `tr`, and `zu`.

7. Watch both Worker terminals. As soon as one prints **`Sleeping between translation calls`**, press **Ctrl-C in that terminal** to stop that Worker. Leave the other Worker running.
8. Observe the remaining Worker continue execution after the Timer fires. Allow for task routing and recovery latency; the continuation need not occur at exactly ten seconds.
9. The starter terminal should print both translated messages, confirming completion despite the original Worker stopping. For `Andy` and `sk`, expect:

   ```json
   {
     "HelloMessage": "Ahoj, Andy",
     "GoodbyeMessage": "Zbohom, Andy"
   }
   ```

Ctrl-C requests Worker shutdown; this exercise demonstrates recovery after losing that Worker process, rather than simulating an abrupt machine failure.

### Inspect the Results

- **Service terminal:** Hello should have been translated only once for this run. Its completed Activity result was reused during replay, rather than calling the service again. Goodbye is translated when execution continues.
- **Worker terminals:** Compare original execution logs with the surviving Worker's continuation. Replay-aware Workflow logging suppresses duplicate messages from replayed steps.
- **Starter terminal:** Both messages should be returned successfully.
- **Web UI:** Open http://localhost:8233, select the `default` namespace, and inspect `durable-execution-workflow`. Look for the completed Hello Activity, `TimerStarted`, `TimerFired`, the Goodbye Activity, and Workflow completion.

Inspect the same execution from the CLI:

```bash
make 02-durable-execution-history
make 02-durable-execution-history-detailed
```

The starter reuses `durable-execution-workflow` as the Workflow ID. For repeated completed runs, inspect the latest Run ID; an already-running execution with that ID may be returned instead of starting a new one.

### Optional: Integrate Zap

Use Temporal's [Zap adapter sample](https://github.com/temporalio/samples-go/tree/main/zapadapter) as a guide to configuring Zap and supplying an SDK-compatible logger through `client.Options.Logger`. Continue using `workflow.GetLogger` and `activity.GetLogger` in application code. This is an optional self-study extension.
