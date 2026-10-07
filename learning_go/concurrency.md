# Chapter 12: Concurrency in Go

## When to Use Concurrency

### Concurrency vs. Parallelism
* **Concurrency**: Structural breakdown of a program into independent execution units that safely share/exchange data.
* **Parallelism**: Simultaneous physical execution on multiple hardware CPU cores.
* **Amdahl's Law**: The non-parallelizable sequential part of an algorithm limits the maximum potential speedup, regardless of how many cores or threads are added.

### Guidelines for Concurrency
* **Concurrency is not free**: Launching goroutines and passing channel messages incurs memory & scheduling overhead.
* **I/O-Bound vs. In-Memory**:
  * **In-Memory**: Highly optimized CPU operations; concurrency often slows down execution due to channel overhead.
  * **I/O-Bound**: Reading/writing disk, database, or network requests. Concurrency is ideal here because waiting on I/O frees up CPU time for other tasks.
* **Rule of Thumb**: Write serial code first. Benchmark before introducing concurrency to verify actual performance gains.

---

## Goroutines

### Process, Thread, vs. Goroutine
* **Process**: OS-managed instance of a running program with private memory space.
* **Thread**: OS-scheduled unit of execution within a process; shares process memory.
* **Goroutine**: Go runtime-managed lightweight execution unit mapped onto OS threads (M:N scheduling).

### Advantages over OS Threads
* **Fast Creation**: No OS kernel resource allocation.
* **Small Stack**: Starts at ~2 KB (dynamic growth/shrinkage) vs. 1–2 MB fixed OS thread stacks.
* **Low Context-Switch Overhead**: Managed entirely in user space without kernel syscalls.
* **Smart Scheduler**: Integrated with Network Poller (auto-pauses on I/O) and Garbage Collector.

### Concurrent Worker Handoff

When performing work concurrently, business logic should remain isolated from channel communications. Worker goroutines consume inputs from an input channel, run pure business logic, and send results into an output channel.

```go
package main

import "fmt"

func process(val int) int {
    return val * 2
}

const numGoroutines = 5

func processConcurrently(inVals []int) []int {
    in := make(chan int, numGoroutines)
    out := make(chan int, numGoroutines)

    for i := 0; i < numGoroutines; i++ {
        go func() {
            for val := range in {
                out <- process(val)
            }
        }()
    }

    go func() {
        for _, v := range inVals {
            in <- v
        }
        close(in)
    }()

    outVals := make([]int, 0, len(inVals))
    for i := 0; i < len(inVals); i++ {
        outVals = append(outVals, <-out)
    }

    return outVals
}

func main() {
    x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    result := processConcurrently(x)
    fmt.Println(result)
}
```

#### Explanation
1. **API Concurrency Hygiene**: `processConcurrently` accepts `[]int` and returns `[]int`. Callers do not manage channels or goroutines.
2. **Pure Business Logic**: `process()` takes `int` and returns `int`, remaining unaware of concurrency.
3. **Closure Wrapper**: Anonymous goroutines wrap `process()`, handling channel reading and writing.

---

## Channels

### Directional Channels
* Reference type created via `make(chan T, [capacity])`. Zero value is `nil`.
* **Single Consumption**: Every value written to a channel is read by exactly one receiving goroutine.
* **Directional Signatures**:
  * `<-chan T` (Read-Only)
  * `chan<- T` (Write-Only)

### Reading, Writing, and Buffering

Unbuffered channels (`cap = 0`) have no memory storage slots and perform direct memory-to-memory copies between two meeting goroutines. Buffered channels (`cap > 0`) store values in an internal buffer up to their capacity limit.

```go
// 1. Buffered (cap = 1): WORKS on a single goroutine
ch1 := make(chan int, 1)
ch1 <- 42
v := <-ch1

// 2. Unbuffered (cap = 0): DEADLOCKS on a single goroutine!
ch2 := make(chan int)
ch2 <- 42 // Blocks immediately waiting for a receiver goroutine
```

#### Explanation
1. **`make(chan int, 1)`**: Has 1 storage slot in memory. The single goroutine writes `42` into the buffer without blocking, then reads it out.
2. **`make(chan int)`**: Has 0 storage slots. Writing `42` pauses the goroutine immediately waiting for a receiver. Because no other goroutine exists, the Go runtime crashes with a deadlock error.

### Using for-range and Channels

The `for-range` loop can iterate over a channel, reading values sequentially until the channel is closed and all queued items are drained.

```go
for v := range ch {
    // processes v until ch is closed and drained
}
```

#### Explanation
1. **Single Variable**: Channel iteration yields only the value `v` (no index).
2. **Exit Conditions**: Loop terminates automatically when the channel is closed and drained.
3. **Deadlock Warning**: If the sender fails to `close(ch)`, the loop blocks forever waiting for values.

### Closing a Channel
* **Closing Operations**: `close(ch)` closes an open channel. Attempting to write to or close a closed channel causes a panic.
* **Reading Closed Channels**: Never blocks. Returns remaining buffered items in FIFO order, then yields the zero value.
* **Comma-Ok Idiom (`v, ok := <-ch`)**: `ok == true` means open; `ok == false` means closed and drained.
* **Ownership**: Responsibility for closing lies strictly with the writing goroutine.
* **Garbage Collection**: Unreferenced channels are automatically garbage-collected whether open or closed. Closing is only required to signal receivers.

### Understanding How Channels Behave

#### The 3 Master Rules for Channel Behavior
1. **Open Channels (Normal)**: Read blocks if empty; Write blocks if full.
2. **Closed Channels (Safety Contract)**: Write or Close panics. Read never blocks (drains buffer, then yields zero-value with `ok = false`).
3. **Nil Channels (`var ch chan T`)**: Read or Write blocks forever (hangs). Close panics.

#### Channel Behavior Matrix

| Operation | Unbuffered, Open | Unbuffered, Closed | Buffered, Open | Buffered, Closed | Nil Channel |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Read (`<-ch`)** | Block until write | Return zero value (`ok == false`) | Block if buffer empty | Return buffered items; zero value when empty | Hang forever |
| **Write (`ch <- v`)** | Block until read | PANIC | Block if buffer full | PANIC | PANIC |
| **Close (`close(ch)`)** | Works | PANIC | Works (buffered items readable) | PANIC | PANIC |

---

## select

### Core Mechanics & Properties
* **Multiplexing Channels**: Allows a goroutine to wait on multiple channel reads and writes simultaneously.
* **Random Case Selection**: If multiple cases are ready, `select` picks one randomly. Order does not matter, preventing starvation.
* **Deadlock Avoidance**: Prevents circular waiting by proceeding with whichever channel operation is ready first.

### Resolving Channel Deadlocks

When two goroutines try to send to unbuffered channels before receiving, they reach a circular deadlock where neither can proceed.

```go
// Deadlocking Goroutines:
package main

import "fmt"

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        inGoroutine := 1
        ch1 <- inGoroutine
        fromMain := <-ch2
        fmt.Println("goroutine:", inGoroutine, fromMain)
    }()

    inMain := 2
    ch2 <- inMain
    fromGoroutine := <-ch1
    fmt.Println("main:", inMain, fromGoroutine)
}
```

```go
// Fixing Deadlock with select:
package main

import "fmt"

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        v := 1
        ch1 <- v
        v2 := <-ch2
        fmt.Println(v, v2)
    }()

    v := 2
    var v2 int
    select {
    case ch2 <- v:
    case v2 = <-ch1:
    }

    fmt.Println(v, v2)
}
```

#### Explanation
1. **The Deadlock**: `main` pauses waiting to send to `ch2`, while `goroutine` pauses waiting to send to `ch1`. Neither can move forward, causing a runtime crash.
2. **The select Fix**: `main` checks if it can write to `ch2` OR read from `ch1`. Because `goroutine` is actively sending to `ch1`, `case v2 = <-ch1` is ready immediately and executes.

### Non-Blocking Operations (default Clause)
* Adding `default` makes `select` non-blocking.
* **CPU Spinning Warning**: Avoid putting `default` inside a `for-select` loop without throttling—it will consume 100% CPU.

---

## Concurrency Practices and Patterns

### Keep Your APIs Concurrency-Free

Concurrency is an internal implementation detail. Hiding channels and mutexes allows refactoring internal concurrency logic without breaking API callers.

```go
// Bad API (forces caller to manage channels):
func ProcessData(in <-chan int) <-chan int

// Good API (concurrency encapsulated inside function):
func ProcessData(inVals []int) []int
```

#### Explanation
1. **Bad API**: Exposes channels in function parameters, forcing caller goroutines to handle channel creation, buffer size, and channel closing.
2. **Good API**: Accepts standard slices and returns standard slices. Internal concurrency is completely hidden from the caller.

### Goroutines, for Loops, and Varying Variables

Capturing a loop variable inside a goroutine closure behaves differently depending on the Go version due to loop variable allocation rules.

```go
// Bug in Go <= 1.21 (all goroutines read last 'v'):
for _, v := range a {
    go func() { ch <- v * 2 }()
}

// Fix for Go <= 1.21 (explicit parameter passing):
for _, v := range a {
    go func(val int) { ch <- val * 2 }(v)
}
```

#### Explanation
1. **Go 1.21 and earlier**: `for` loops reused a single memory location for `v`. By the time goroutines ran, `v` held the final loop item, causing all goroutines to print the same value.
2. **Fix**: Passing `v` explicitly as a function parameter creates a separate variable copy per goroutine.
3. **Go 1.22+**: Automatically allocates a new variable instance per iteration, making closures safe out-of-the-box.

### Always Clean Up Your Goroutines

Goroutines blocked forever on unread or unwritten channels are never garbage-collected by the Go runtime, permanently leaking stack memory.

```go
func countTo(max int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < max; i++ {
            ch <- i
        }
        close(ch)
    }()
    return ch
}
```

#### Explanation
1. **The Leak**: If a caller breaks early out of `for i := range countTo(10)`, no goroutine is left to read from `ch`.
2. **Memory Impact**: The `countTo` goroutine stays blocked at `ch <- i` forever. Its stack memory remains allocated indefinitely.

### Use the Context to Terminate Goroutines

Passing a `context.Context` to async workers allows callers to signal termination when work is canceled or finished early.

```go
func countTo(ctx context.Context, max int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := 0; i < max; i++ {
            select {
            case <-ctx.Done():
                return
            case ch <- i:
            }
        }
    }()
    return ch
}

// Caller Usage:
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

#### Explanation
1. **`<-ctx.Done()`**: Listens for a cancellation signal on the context.
2. **Clean Exit**: When the caller breaks early and executes `defer cancel()`, `<-ctx.Done()` becomes ready, causing the worker goroutine to return immediately and release its stack memory.

### Know When to Use Buffered and Unbuffered Channels

Buffered channels are appropriate when you know the exact number of launched worker goroutines ($N$) and want workers to write output and terminate without waiting for a reader.

```go
func processChannel(ch chan int) []int {
    const conc = 10
    results := make(chan int, conc)

    for i := 0; i < conc; i++ {
        go func() {
            v := <-ch
            results <- process(v)
        }()
    }

    var out []int
    for i := 0; i < conc; i++ {
        out = append(out, <-results)
    }
    return out
}
```

#### Explanation
1. **Buffer Capacity**: `make(chan int, 10)` guarantees 10 available storage slots for 10 workers.
2. **Non-Blocking Write**: Each worker writes its result into the buffer slot and exits immediately without waiting for `processChannel` to read the value.

### Implement Backpressure

Under heavy traffic spikes, accepting unlimited concurrent requests crashes downstream services and exhausts server memory. Backpressure sets a concurrency limit, rejecting excess work immediately with `HTTP 429 Too Many Requests`.

```go
type PressureGauge struct {
    ch chan struct{}
}

func New(limit int) *PressureGauge {
    return &PressureGauge{
        ch: make(chan struct{}, limit),
    }
}

func (pg *PressureGauge) Process(f func()) error {
    select {
    case pg.ch <- struct{}{}:
        f()
        <-pg.ch
        return nil
    default:
        return errors.New("no more capacity")
    }
}
```

#### Explanation
1. **Zero-Memory Semaphore (`chan struct{}`)**: `struct{}` occupies 0 bytes of RAM, making `chan struct{}` the most memory-efficient token bucket mechanism.
2. **Acquiring & Releasing Tokens**: `pg.ch <- struct{}{}` puts a token into the 10-slot channel buffer before executing `f()`. Once `f()` finishes, `<-pg.ch` removes the token, freeing capacity for the next incoming request.
3. **Non-Blocking Rejection (`default`)**: If all 10 buffer slots are occupied, `pg.ch <- struct{}{}` cannot write. Instead of blocking or queuing, `select` drops into `default` instantly and returns a capacity error without wasting CPU or memory.

### Turn Off a case in a select

Reading from a closed channel in a `for-select` loop returns zero values immediately without blocking, causing `select` to spin continuously and waste CPU cycles.

```go
for count := 0; count < 2; {
    select {
    case v, ok := <-in1:
        if !ok {
            in1 = nil // Operations on nil channels block forever, turning off this case!
            count++
            continue
        }
        out = append(out, v)
    case v, ok := <-in2:
        if !ok {
            in2 = nil // Case disabled!
            count++
            continue
        }
        out = append(out, v)
    }
}
```

#### Explanation
1. **The Closed Channel Problem**: When `in1` closes, `<-in1` returns `(0, false)`. In a loop, `select` would continuously select this case forever.
2. **Setting to `nil` (`in = nil`)**: Operations on `nil` channels block forever. When `select` encounters a `nil` channel case, it ignores that branch in all future loop iterations.

### Time Out Code

When calling external operations or long-running workers, programs need a hard deadline to prevent waiting indefinitely if a worker hangs.

```go
func timeLimit[T any](worker func() T, limit time.Duration) (T, error) {
    out := make(chan T, 1)
    ctx, cancel := context.WithTimeout(context.Background(), limit)
    defer cancel()

    go func() {
        out <- worker()
    }()

    select {
    case result := <-out:
        return result, nil
    case <-ctx.Done():
        var zero T
        return zero, errors.New("work timed out")
    }
}
```

#### Explanation
1. **`context.WithTimeout` Deadline Timer**: Creates a context that automatically closes `ctx.Done()` when `limit` (2 seconds) elapses.
2. **1-Capacity Buffer Prevents Leaks**: `out := make(chan T, 1)` allows the worker goroutine to write its result and exit cleanly even if `timeLimit` timed out and returned early.
3. **`select` Race**: Multiplexes `<-out` against `<-ctx.Done()`. If `<-ctx.Done()` triggers first, it returns a timeout error immediately.

### Use WaitGroups

When multiple goroutines perform concurrent work, the parent goroutine needs a mechanism to wait for all of them to finish. `sync.WaitGroup` maintains an atomic counter tracking active goroutines.

#### Basic Usage Snippet

```go
var wg sync.WaitGroup
wg.Add(3)

go func() {
    defer wg.Done()
    doThing1()
}()
go func() {
    defer wg.Done()
    doThing2()
}()
go func() {
    defer wg.Done()
    doThing3()
}()

wg.Wait() // Blocks until counter hits 0
```

#### The Monitoring Goroutine Pattern Snippet (`processAndGather`)

When multiple worker goroutines write results to a single shared output channel, `close(out)` must happen **only once**, precisely after **all workers exit**. A dedicated monitoring goroutine solves this:

```go
func processAndGather[T, R any](in <-chan T, processor func(T) R, num int) []R {
    out := make(chan R, num)
    var wg sync.WaitGroup
    wg.Add(num)

    for i := 0; i < num; i++ {
        go func() {
            defer wg.Done()
            for v := range in {
                out <- processor(v)
            }
        }()
    }

    // Monitoring Goroutine: waits for workers, then closes output channel
    go func() {
        wg.Wait()
        close(out)
    }()

    var result []R
    for v := range out {
        result = append(result, v)
    }
    return result
}
```

#### Explanation
1. **WaitGroup Mechanics**:
   * Zero value of `sync.WaitGroup` is ready to use without initialization.
   * `wg.Add(n)` increments the counter. Call it before starting goroutines.
   * `defer wg.Done()` decrements the counter. Using `defer` guarantees execution even if a worker panics.
   * `wg.Wait()` blocks until the counter reaches 0.
2. **Crucial Rule on Passing WaitGroups**:
   * **Never pass `sync.WaitGroup` by value!** If passed by value, Go creates a copy, so `Done()` decrements the copy while `Wait()` blocks forever on the original. Always capture via closure or pass a pointer (`*sync.WaitGroup`).
3. **The Monitoring Goroutine Pattern**:
   * Spawning a separate background goroutine `go func() { wg.Wait(); close(out) }()` allows the main function to transition directly to `for v := range out`.
   * When all processing workers call `Done()`, `wg.Wait()` unblocks inside the monitoring goroutine, calling `close(out)`. This safely terminates the `for v := range out` loop.

### golang.org/x/sync/errgroup

The `errgroup` package (from `golang.org/x/sync/errgroup`) extends `sync.WaitGroup` by adding **error propagation** and **automatic context cancellation** across a group of concurrent workers.

```go
import (
    "context"
    "fmt"
    "golang.org/x/sync/errgroup"
)

func fetchAll(urls []string) error {
    g, ctx := errgroup.WithContext(context.Background())

    for _, url := range urls {
        url := url
        g.Go(func() error {
            return fetchURL(ctx, url)
        })
    }

    if err := g.Wait(); err != nil {
        return fmt.Errorf("fetch failed: %w", err)
    }
    return nil
}
```

#### Explanation
1. **`errgroup.WithContext(parentCtx)`**: Returns a new `Group` instance `g` and a derived `ctx`.
2. **`g.Go(func() error)`**: Automatically manages internal WaitGroup counters (`Add`/`Done`).
3. **First Error Cancellation**: If **any** worker returns a non-nil `error`, `g` captures that error and **immediately cancels `ctx`** (`ctx.Done()`), signaling sibling workers to abort.
4. **`g.Wait()`**: Blocks until all worker goroutines complete, returning the **first non-nil error** returned by any worker.

### Run Code Exactly Once

Package `init()` functions should be reserved for immutable package-level state. When initialization is slow or expensive (and may not even be needed on every run), you should **lazy-load** data so it executes exactly once upon first use.

#### Classic `sync.Once` (`13_sync_once`)

```go
type SlowComplicatedParser interface {
    Parse(string) string
}

func initParser() SlowComplicatedParser {
    // Heavy setup and loading here...
    return &parserImpl{}
}

var parser SlowComplicatedParser
var once sync.Once

func Parse(dataToParse string) string {
    once.Do(func() {
        parser = initParser() // Executed ONLY on first call!
    })
    return parser.Parse(dataToParse)
}
```

#### Go 1.21+ Functional Helpers (`14_sync_value`)

Go 1.21 added generic helper functions (`sync.OnceFunc`, `sync.OnceValue`, `sync.OnceValues`) that return a wrapped function with automatic return value caching, eliminating package-level mutable state variables:

```go
// Wraps initParser and caches its single return value:
var initParserCached = sync.OnceValue(initParser)

func Parse(dataToParse string) string {
    parser := initParserCached() // Returns cached parser instance
    return parser.Parse(dataToParse)
}
```

#### Explanation
1. **Why Lazy Initialization?**: Setup code that takes hundreds of milliseconds or connects to external resources shouldn't slow down application startup. `sync.Once` defers execution until the exact moment the feature is first called.
2. **Classic `sync.Once` Mechanics**:
   * **Zero Value Useful**: Declaring `var once sync.Once` is ready to use immediately without calling a constructor.
   * **`once.Do(f)`**: Executes closure `f` on the first call. Subsequent calls block until `f` completes, then return immediately without running `f` again.
   * **Crucial Rule on Copies**: Never pass or copy a `sync.Once` instance, as each copy maintains its own internal state flag. Declaring `var once sync.Once` inside a function body is a bug (creates a fresh instance on every call).
3. **Go 1.21+ Functional Helpers**:
   * **`sync.OnceFunc(f)`**: Wraps a function that returns **0** values.
   * **`sync.OnceValue(f)`**: Wraps a function that returns **1** value (`T`), caching the return value.
   * **`sync.OnceValues(f)`**: Wraps a function that returns **2** values (`(T, error)`), caching both values.
   * **Cleaner Architecture**: Eliminates package-level mutable state variables (`parser`), replacing them with a thread-safe cached function handle (`initParserCached`).

### When to Use Mutexes Instead of Channels

While channels are ideal for orchestrating data pipelines and passing ownership of values between goroutines, **Mutexes** (`sync.Mutex` and `sync.RWMutex`) are better when managing **shared mutable state** (e.g. updating fields in an in-memory struct or map).

```go
type Scoreboard struct {
    sync.RWMutex
    scores map[string]int
}

func (s *Scoreboard) Update(name string, val int) {
    s.Lock()
    defer s.Unlock()
    s.scores[name] = val
}

func (s *Scoreboard) Read(name string) (int, bool) {
    s.RLock()
    defer s.RUnlock()
    val, ok := s.scores[name]
    return val, ok
}
```

#### Explanation

1. **Channels vs. Mutexes**:
   * **Channels**: Orchestrate data flow, pass ownership between goroutines, coordinate concurrent execution pipelines.
   * **Mutexes**: Protect shared in-memory data structures (maps, structs, caches) where multiple goroutines need to read and write to the exact same memory locations.

2. **`sync.Mutex` vs `sync.RWMutex`**:
   * **`sync.Mutex`**: Exclusive locking (`Lock()` / `Unlock()`). Only 1 goroutine (reader or writer) can hold the lock at a time.
   * **`sync.RWMutex`**: Reader/Writer locking (`RLock()` / `RUnlock()` for readers; `Lock()` / `Unlock()` for writers). Multiple goroutines can read concurrently as long as no writer holds the lock.

3. **Best Practices**:
   * Keep lock duration as short as possible.
   * Always use `defer s.Unlock()` (or `defer s.RUnlock()`) immediately after locking to ensure the lock is freed even if code panics.
   * Avoid calling external functions while holding a lock (prevents deadlocks).
