# Ultimate Go

## 01. Design Guidelines Intro

### Design Guidelines Intro

- Ultimate Go starts with engineering judgment, not syntax.
- Change how you think about software construction, not just the syntax you can type.
- Software quality comes from understanding trade-offs and visible cost.
- Go is designed to keep enough cost visible that developers can reason about behavior.
- The machine is the model.

### Prepare Your Mind

- Prefer less code. Less code usually means fewer bugs, fewer moving parts, and smaller mental models.
- Large programs and large abstraction towers are not achievements by themselves.
- Prefer thin, precise abstraction layers over broad generic ones.
- Prefer guidelines over rigid rules. Rules are often applied blindly; guidelines explain why one choice fits better than another.
- Rules still matter, but rules have costs and must pull their weight.
- Follow standards consistently, but do not idolize them to the point that judgment disappears.
- Semantics should make ownership and behavior visible.
- The hardware is the platform. Keep the machine model in mind when reasoning about software.
- Every decision has a cost. If the cost is unclear, the decision is guesswork.
- Ask harder engineering questions than "does it compile": is it good, efficient, correct, on time, and worth its cost.
- Technology changes quickly; minds tend to change slowly. Learning Go means adjusting the mental model.
- Have a point of view about software design, but keep refining it as the model gets better.
- Introspection and self-review matter; part of the work is examining how you think, not just what you type.
- Champion quality, efficiency, and simplicity together.
- Learn to read code before trying to become a better writer of code.
- Reading code is where style, judgment, and taste are actually learned.
- Write code that other developers can safely maintain later.
- Large codebases fail first in human comprehension, not only in machine execution.
- Legacy systems are usually written by capable developers under constraints that later readers can no longer see.
- When the mental model breaks down, refactoring is usually needed.
- Debuggers are exploration tools, not substitutes for readable code.
- Heavy debugger dependence usually signals that logs, mental models, or both are too weak.

- [Overview and design philosophy](gotraining/topics/go/README.md)

### Productivity vs. Performance

- Go tries to balance productivity and performance instead of forcing a hard trade.
- The useful question is not whether Go wins language shootouts, but whether the program is fast enough for the problem.
- Hardware progress now depends heavily on software using cores effectively, so engineering judgment has to include concurrency and hardware sympathy.
- Performance work without measurement is speculation.
- Fast enough is the target; productivity still matters.
- The old fallback of "write it in C if performance matters" is the trade Go is trying to soften.
- In practice, most Go code should be fast enough before any special tuning.
- A useful standard is code that is straightforward enough that you can predict reasonably well how it will run on the machine.
- Build air conditioners: systems that keep working day after day without drama.

- [Productivity and performance guidance](gotraining/topics/go/README.md)

### Correctness vs. Performance

- Optimize for correctness first.
- Preferred order: correct, clear, concise, then fast.
- Prototype, review, and refactor. Good code is usually produced in drafts.
- Treat code as drafts that are revised into production quality.
- Prototype in the concrete first. Push contracts and abstractions later, after the working shape is understood.
- There is no royal road to correctness. It comes from invariants, testing, code reviews, and revision.
- Simple, direct solutions are usually stronger than clever ones.
- Dependable software matters more than flashy solutions.
- Early micro-optimization usually fights integrity and readability.

- [Correctness and performance guidance](gotraining/topics/go/README.md)

### Code Reviews

#### Integrity

- Integrity means reliability at both the micro and macro levels.
- At the micro level, code allocates memory, reads memory, and writes memory.
- Those operations should be accurate, consistent, and efficient.
- At the macro level, software is mostly data transformation: data comes in, is transformed, and goes out.
- Functions can often be read as input, transformation, and output.
- If you do not understand the data, you do not understand the problem.
- Write less code because every extra line is another place for bugs to hide.
- Writing less code also reduces testing surface and reduces the chances that shortcuts and dirty tricks creep in later.
- Error handling is part of the main path of correct software.
- Failure is expected. Systems should detect failure and recover from it.
- Integrity is not free. The usual trade is a little performance in exchange for reliability.
- Tedious error handling is often justified because many critical failures come from bad-path behavior, not happy-path behavior.
- A large share of critical distributed-system failures come from bad error handling rather than from the nominal path.
- Personal responsibility matters here: ask questions early, surface mistakes quickly, and do not protect bad decisions just because you made them.

#### Readability

- Readability includes visible cost.
- A reader should be able to look at code and predict roughly what work it causes.
- Hidden constructor, destructor, and operator costs are readability problems because the line of code no longer tells the truth about the work being done.
- Code must never lie.
- Readability is a team concern, not an individual style preference.
- Code should be understandable by the average developer on the team, not only the strongest specialist.
- Senior engineers should make the average developer more effective, not write code that only seniors can decode.
- Readability is partly subjective, but cost visibility is a measurable part of it.
- Optimize for code that is easy to understand later, not merely easy to type now.
- Readability has layers: line, function, package, and ecosystem. Cost should stay visible at every layer.
- Make things easy to understand, not merely easy to do.

#### Simplicity

- Simplicity follows integrity and readability.
- Simplicity is not the same as hiding everything behind abstraction.
- Good abstractions hide complexity without hiding cost.
- Simplicity is usually achieved through repeated refactoring, not on the first draft.
- First make the code work, then make the cost readable, then simplify what can safely be simplified.
- Encapsulation is hard. Precise abstractions help; vague abstractions usually make systems harder to reason about.
- Simplicity is hard to design and complicated to build.

- [Design philosophy and review priorities](gotraining/topics/go/README.md)
- `Integrity -> Readability -> Simplicity -> Performance`

### If Performance Matters

- Four common reasons software is slower than it should be: external latency, internal latency, data access patterns, and algorithmic efficiency.
- External latency includes network calls, system calls, and cross-process work. This is often the dominant cost.
- Internal latency includes GC, synchronization, and orchestration overhead.
- Data access patterns matter because layout and traversal affect the hardware directly.
- Algorithmic efficiency matters, but in many systems it is not the first thing to investigate.
- Performance arguments without measurements are still guesses.
- Never guess. Measure relevant workloads, profile the real bottleneck, and only then decide what is critical.
- Around `10,000` lines of code is a practical mental-model threshold for one developer.
- Once software grows past the size a single person can comfortably model, hidden cost and weak abstractions start compounding quickly.
- Use the debugger as a last-resort exploration tool, not as a substitute for clear code, logs, and mental models.

- [Performance guidance](gotraining/topics/go/README.md)

## 02. Memory And Data Semantics

### Memory & Data Semantics

- Type and memory layout are engineering concerns, not just syntax concerns.
- Type tells you both size and representation.
- The major semantic choice in Go is value semantics versus pointer semantics.
- Escape analysis and garbage collection explain much of Go's allocation cost model.

### Variables

- Type is life.
- Memory starts as bytes. Type gives those bytes size, representation, and meaning.
- Type tells you how much memory a value uses and what that memory means.
- Three broad categories matter here: built-in types, reference types, and user-defined struct types.
- Prefer `int` unless a specific width is required.
- `int` is the natural machine-sized integer for the current architecture, which is why it is the default choice most of the time.
- Prefer sized integers like `int32` or `uint64` only when width or protocol compatibility actually matters.
- A `word` is the machine-sized unit large enough to hold either an integer or an address on the current architecture.
- `float64` is a useful teaching type because it makes both width and representation obvious.
- Zero value is a core integrity feature in Go.
- Prefer `var` for zero-value construction.
- Prefer `:=` for non-zero initialization.
- Use that split for readability and consistency: zero-value declarations look one way, initialized values look another.
- Prefer conversion over casting. Conversion creates a new value with the target type rather than reinterpreting memory.
- Reinterpretation is dangerous because it can separate the bits from the meaning the type system is supposed to protect.
- A string value is a two-word data structure: pointer plus length.
- The zero-value string is a nil pointer plus a length of `0`.
- `value` means what is in the box, `address` means where the box is, and `data` can refer to both together.

```go
var count int
name := "gopher"
ratio := float64(10)
```

- Strings are fixed-size headers describing string data, not raw byte arrays in user code.

- [Variables README](gotraining/topics/go/language/variables/README.md)
- [Declare and initialize variables](gotraining/topics/go/language/variables/example1/example1.go)

### Struct Types

- Structs are composite types built from other existing types.
- Structs model the real data of the problem.
- If you do not understand the fields and their meaning, you do not understand the problem.
- Structs are your domain model, and they are also physical layouts in memory.
- Struct layout includes alignment and compiler-inserted padding.
- Alignment exists because hardware reads are cheaper when values sit on natural word boundaries.
- Crossing boundaries can require extra work, which is why 2-byte, 4-byte, and 8-byte alignment rules matter.
- Field sizes do not necessarily add up to final struct size because padding may be inserted for efficient access.
- Final struct size is also aligned to the largest field requirement, so end padding can appear too.
- Optimize struct layout for readability first.
- Reorder fields to reduce padding only when profiling or memory pressure justifies it.
- If fields are ordered largest-to-smallest without a proven need, that is usually an optimization review discussion, not an automatic win.
- Named types remain distinct even when their underlying structure matches.
- Assignment between named types requires explicit conversion.
- Anonymous structs and other literal types get more assignment flexibility because they are unnamed literal forms.
- Identical and compatible named structs are still different concrete types.

```go
type user struct {
	name string
	age  int
}
```

- After struct layout, the explanation pivots into the runtime model that makes frames and semantics meaningful.
- `P` is the logical processor, `M` is the operating system thread, and `G` is the goroutine.
- `G` behaves like an application-level thread executing instructions sequentially.
- A `G` cannot execute by itself. It needs an `M`, and that `M` needs a `P` attached so Go code can run.
- Each `G` gets its own contiguous stack.
- That stack starts small, around `2 KB`, and grows as needed.
- Every function call carves another frame from the same goroutine stack.
- The frame is the local sandbox for that function's reads, writes, locals, and parameters.
- Restricting direct mutation to the current frame is part of how the model protects integrity.

```text
physical processor
  -> cores
    -> hardware threads
      -> OS schedules M here

P = logical processor
M = operating system thread
G = goroutine

        runnable G values
      [ G ] [ G ] [ G ]
          \    |    /
           \   |   /
             [ P ]
               |
             attached
               |
             [ M ]
               |
          scheduled by OS
               |
      hardware thread / core

Each G has its own stack:

G
|
+-- stack
    +-- frame: runtime startup work
    +-- frame: main()
    +-- frame: increment(count)
```

- [Struct types README](gotraining/topics/go/language/struct_types/README.md)
- [Declare and initialize structs](gotraining/topics/go/language/struct_types/example1/example1.go)
- [Anonymous structs](gotraining/topics/go/language/struct_types/example2/example2.go)
- [Named vs unnamed types](gotraining/topics/go/language/struct_types/example3/example3.go)
- [Alignment and padding](gotraining/topics/go/language/struct_types/advanced/example1/example1.go)
- [Goroutines README](gotraining/topics/go/concurrency/goroutines/README.md)
- [Goroutines and parallelism](gotraining/topics/go/concurrency/goroutines/example3/example3.go)

### Pointers-Part 1 (Pass by Values)

- Everything in Go is passed by value.
- A function call creates a new stack frame.
- Each frame acts like a sandbox for direct reads and writes.
- Passing a value means copying that value into the next frame.
- Value semantics mean each function works on its own copy of the data.
- Main benefits: isolation, easier reasoning, fewer side effects, strong integrity, and often better locality later when performance matters.
- Main cost: copying and possible reconciliation when multiple operations need shared state.
- The inefficiency is not only extra copying; it can also be the extra code needed to merge independent results back together.
- Values are copied into the next frame, so mutation stays local unless sharing is introduced deliberately.
- Function parameters are the mechanical way input gets into the next frame so the transformation can happen there.
- When the callee returns, the caller still has its original value because the mutation happened only inside the callee's frame.

```go
func increment(n int) {
	n++
}
```

- [Pass by value](gotraining/topics/go/language/pointers/example1/example1.go)

### Pointers-Part 2 (Sharing Data)

- Pointer semantics still pass by value, but the copied value is an address.
- A pointer gives access to data outside the current frame.
- Direct access stays inside the current frame. Pointers create indirect access to data outside it.
- Main benefit: efficient sharing of one underlying piece of data.
- Main cost: side effects, because multiple parts of the program can mutate shared state.
- Default to value semantics first.
- Go is a value-semantics-first language. Move to pointer semantics only when value semantics stop being reasonable or practical.
- Move to pointer semantics when sharing is required or value copying is no longer reasonable.
- A pointer type is a literal type formed as `*T`.
- A pointer is not just an address; it is the address of a specific concrete type, because the program still needs to know how to interpret and mutate the memory safely.
- Passing a pointer is still pass-by-value; the copied value is the address.

```go
func incrementInPlace(n *int) {
	(*n)++
}
```

- Use `var` for zero-value construction when binding to a variable.
- Use literal construction for non-zero initialization.
- Prefer `var u user` over `u := user{}` when the intent is zero value.
- Use value construction for named variables so sharing stays explicit at the point where `&` appears.
- Avoid pointer-semantic construction into a named variable when a value will do; it hides cost and sharing too early.
- Construction in a direct `return` or inside a function call is the main exception, because readability is still preserved at the use site.
- Factory functions advertise semantics in their return type: returning `user` gives the caller a copy, returning `*user` gives shared access.

```go
func newUser() user { return user{name: "bill"} }
func newUserPtr() *user { return &user{name: "bill"} }
```

- Return type tells the caller whether the value is copied out or shared.
- `createUserV1` in the example below returns a value and the caller receives its own copy.
- `createUserV2` returns `&u`, which makes the sharing direction visible and explains why `u` must be heap-allocated.
- If construction were written as `u := &user{...}` and then just `return u`, the return line would hide the allocation and lose readability.
- Sharing down the call stack can still stay on the stack.
- Sharing up the call stack is the pattern that forces heap construction in these examples.

- [Pointers README](gotraining/topics/go/language/pointers/README.md)

- [Sharing data I](gotraining/topics/go/language/pointers/example2/example2.go)
- [Sharing data II](gotraining/topics/go/language/pointers/example3/example3.go)
- [Escape analysis and factory-return semantics](gotraining/topics/go/language/pointers/example4/example4.go)

### Pointers-Part 3 (Escape Analysis)

- Construction does not determine stack vs heap placement. Sharing direction does.
- Escape analysis decides whether a value can stay on the stack or must move to the heap.
- Escape analysis is compile-time static analysis by the compiler.
- The decision is driven by sharing behavior, not by where the code text appears.
- An allocation means heap construction, not just "memory was used".
- Sharing down the call stack can often remain on the stack.
- Sharing data so it outlives the frame usually forces heap allocation.
- Construction by itself tells you very little; sharing direction is the stronger signal.
- If the only sharing shown is `return &u`, that line alone already signals heap construction.
- If construction is hidden as `u := &user{...}`, the later return loses that cost signal.
- Heap allocation is a cost signal because heap objects bring GC into the picture.
- Stack values are preferred when possible because stacks are self-cleaning and usually friendlier to performance.
- Heap-backed values can still look like ordinary values in source code even though the runtime must access them through an address underneath.
- `go build -gcflags -m=2` is the right tool for understanding why a value allocates after profiling shows that allocation matters.
- Profilers tell you what allocates; compiler escape output explains why.

- [Escape analysis](gotraining/topics/go/language/pointers/example4/example4.go)

### Pointers-Part 3 (Stack Growth)

- Goroutines start with small stacks and those stacks grow as needed.
- Initial goroutine stacks are small, around `2 KB`, so stack growth is expected rather than exceptional.
- Small starting stacks are what make very large goroutine counts practical.
- Stack growth is done by moving to a larger contiguous stack and adjusting internal references.
- Stack growth usually doubles the current stack size into a new contiguous allocation.
- Existing frame data is copied into that new stack segment, so stack values can physically move.
- Because goroutine stacks can move, stack memory cannot be safely shared arbitrarily.
- Internal stack pointers for the same goroutine can be fixed up during growth.
- Cross-goroutine stack pointers would require global pointer repair across many stacks, which is exactly what the runtime avoids.
- Sharing between goroutines therefore pushes values toward the heap because stacks can move.
- Stack growth cost exists, but it is a small integrity cost, not usually the dominant reason software is slow.
- Go tries to keep values on the stack when it safely can because stacks are self-cleaning and generally friendlier to performance.

- [Stack growth](gotraining/topics/go/language/pointers/example5/example5.go)

### Pointers-Part 3 (Garbage Collection)

- Stack memory is reclaimed naturally when frames disappear.
- Heap memory is managed by the garbage collector.
- Go's GC is a concurrent mark-and-sweep collector tuned around latency, throughput, and practical cloud resource use.
- All collectors do the same broad job: identify heap values that are still reachable, identify those that are not, and reclaim the unreachable memory for reuse.
- Focus on GC behavior and cost, not collector-brand trivia.
- The pacer decides when a collection should start and tries to begin as late as possible while still finishing before the heap budget is exhausted.
- With the default `GOGC=100`, heap growth is budgeted relative to the currently live heap rather than treated as an unlimited pool.
- A simple mental model is `live heap + equal-sized gap` when `GOGC=100`.
- If `2 MB` remain live after collection, the runtime budgets roughly another `2 MB` of gap before the next cycle.
- If `3 MB` remain live after the next cycle, the budget becomes roughly `6 MB` total heap.
- The first collection begins around a `4 MB` live heap, and later targets are computed from the marked-live heap after the previous cycle.
- GC pressure is how quickly that post-collection gap fills back up.
- Filling the gap faster means collection has to start again sooner.
- GC runs in three phases: mark start, concurrent marking, and mark termination.
- The mark start and mark termination phases stop the world briefly.
- One early stop-the-world cost is turning on the write barrier so concurrent heap mutation can continue with integrity.
- Healthy goroutines help here because the runtime needs goroutines to reach safe points; long tight loops without function calls delay that coordination.
- The stop-the-world target is small, but total GC time still matters because throughput drops across the full cycle.

```text
10k requests
1000 GC
2 ms pace

                 4 MB heap
        +------------------------+
        |       2 MB gap         |
        +------------------------+
        |      2 MB live         |
        +------------------------+

GC targets:

10-30 micro   first stop-the-world to get goroutines
              to a safe point and turn on write barrier

300 micro     total GC time for the whole cycle

80 micro      possible stop-the-world budget across
STW           the collection

Important questions:

- How fast does the 2 MB gap fill?
- How many GC cycles are needed for the same work?
- How much throughput is lost during each 300 micro cycle?
```

- Stop-the-world time is only one latency number.
- Total GC time matters too because the program is not running at full throttle for the duration of the collection.
- During concurrent marking, the collector takes roughly `25%` of the available CPU capacity.
- On a program that could otherwise run `4` goroutines in parallel, GC can effectively reduce application work to about `3` goroutines during that period.
- If allocation pressure is high enough that the collector may miss its target, the runtime can trigger mark assist.
- Mark assist means an allocating goroutine is forced to do GC work itself instead of only application work.
- That is another way throughput drops even when the program is still technically running concurrently.
- Healthy goroutines matter here too: a cooperating runtime works better when goroutines make function calls regularly instead of sitting in long tight loops.
- Garbage collection is still a net win because it removes a large amount of cognitive load from concurrent software.
- The engineering job is not to eliminate GC, but to be sympathetic to it.
- Three practical GC questions matter:
  - how much stop-the-world time is required
  - how long the full collection takes
  - how many collections are required to complete a fixed amount of work
- Fewer collections for the same workload usually means less cumulative latency.
- Simply making the heap much larger is not the main win.
- A larger gap may reduce collection frequency while also increasing scanning work because there can be more live heap values to walk.
- A smaller heap can be helpful because fewer heap values often means less mark work per collection.
- The stronger target is the same workload with fewer allocated bytes and fewer live heap objects.
- Profiling memory should look at both allocated bytes and allocation count because both affect GC cost.
- Reducing non-productive allocations helps the gap fill more slowly.
- Reducing object count helps because the collector has fewer live objects to scan and mark.

- [Pointers README](gotraining/topics/go/language/pointers/README.md)

### Constants

- Constants are not read-only variables.
- Constants exist at compile time.
- Constants can be typed or constants of a kind.
- Literal values are unnamed constants of a kind.
- Model untyped constants as kinds rather than ordinary runtime values.
- Untyped constants act more like kinds and can be promoted by the compiler when safe.
- This forms a parallel compile-time numeric system with much more precision than ordinary runtime numeric values.
- Typed constants obey the constraints of their declared type.
- Kind promotion rules matter: float kind can dominate int kind in constant expressions before a final concrete type is chosen.
- Kind can also promote to a concrete type when an operation requires like types.
- Constant promotion is why literal math feels flexible while ordinary variables still require explicit conversions.
- `iota` is useful for enumerations and bitmask-style constants.
- Untyped constants can carry far more precision than ordinary runtime numeric values.
- A typed constant is bound by the limits of its declared type. A kind-based constant can remain valid until it is forced into a concrete runtime type.

```go
const minutes = 5
const timeout = minutes * 60
```

- Very large constants can still compile while they remain compile-time values; the failure only appears when forcing them into a runtime type that cannot hold them.

- [Constants README](gotraining/topics/go/language/constants/README.md)
- [Constant basics](gotraining/topics/go/language/constants/example1/example1.go)
- [Kind-based constant system](gotraining/topics/go/language/constants/example2/example2.go)
- [Iota](gotraining/topics/go/language/constants/example3/example3.go)
- [Implicit conversion](gotraining/topics/go/language/constants/example4/example4.go)

### Garbage Collection Addendum Part 1

- Stop-the-world time matters, but total GC time matters too because the program is not running at full throughput during the whole cycle.
- During concurrent marking, GC also consumes part of the program's CPU capacity, so throughput drops even while application work continues.
- A practical rule of thumb is that concurrent marking can take roughly a quarter of available CPU capacity.
- On a program that could otherwise run `4` goroutines in parallel, that can mean effectively running only `3` for application work during much of the collection.
- Garbage collection therefore creates internal latency even when stop-the-world pauses stay small.
- A useful way to think about GC cost is not just "how long did the pause last" but also "how long were we not running at full throttle".

- [Pointers README](gotraining/topics/go/language/pointers/README.md)

### Garbage Collection Addendum Part 2

- Mark assist means allocating goroutines may be forced to do GC work themselves so the collector can keep up.
- GC pressure depends on allocation rate and object count.
- Object count matters because marking work depends on how many live heap objects must be traversed.
- Byte volume matters because more bytes per unit of work fill the heap gap faster and trigger more frequent collections.
- The practical question is not "can GC be made invisible" but "how many collections are required to process a fixed amount of work, and how expensive is each one".
- The important view is behavior over time: how often collection runs, how long it takes, and how much throughput is lost while it runs.
- For a fixed workload, fewer collections usually means less cumulative latency.
- Reducing two heap objects to one can matter because the collector now has one fewer live object to scan.

- [Pointers README](gotraining/topics/go/language/pointers/README.md)

### Garbage Collection Addendum Part 3

- Non-moving heap objects matter because stable addresses simplify interop and avoid a different class of movement problems.
- Zero allocation is not the goal.
- The goal is fewer unnecessary allocations, fewer live objects to scan, and fewer unnecessary GC cycles for the same workload.
- Do not try to win by simply inflating the heap budget. A larger gap may reduce collection frequency while also making each collection do more scanning work.
- A smaller heap can help because less heap usually means less scanning work per collection.
- The real target is to do the same amount of work with fewer bytes allocated and fewer live objects left for the collector to walk.
- The better target is less memory and fewer heap objects per request, task, or other unit of work.
- Profiling should look at both allocated bytes and allocation count, because both affect GC cost.
- Garbage collection is a feature, not a failure. The job is to be sympathetic to it, not to fantasize about removing it.

- [Pointers README](gotraining/topics/go/language/pointers/README.md)

## 03. Data Structures

### Data Structures

- Arrays, slices, and maps are the core data structures in this section.
- Slices are the most important data structure in day-to-day Go.
- Data-oriented design is the framing idea: choose structures that fit the machine and the workload.
- Hardware sympathy matters because traversal cost is often dominated by data layout, not by the surface simplicity of the code.
- Go keeps the data-structure set small on purpose: arrays and slices keep contiguous-memory behavior visible, and maps also try to keep data as contiguous as practical internally.

- [Overview and design philosophy](gotraining/topics/go/README.md)

### Arrays-Part 1 (Mechanical Sympathy)

- Arrays allocate contiguous blocks of fixed-size memory.
- Arrays are the most important data structure relative to the hardware because they give direct control over spatial locality.
- Linear traversal through contiguous memory is friendly to cache lines, prefetchers, and TLBs.
- Traversal cost should be thought of as lost instruction budget, not only as abstract time.
- A rough teaching model here is `3 GHz * 4 instructions per cycle`, which makes `1 ns ~= 12` instructions.
- A main-memory reference can easily cost around `100 ns`, which is about `1200` lost instructions in that teaching model.
- Cache lines are typically `64 bytes`, so traversing linearly lets one fetch pay for nearby values too.
- Row-major traversal through a matrix is fast because it walks memory in the way the hardware expects.
- Column-major traversal over row-major storage is slower because it defeats cache-line locality and can create TLB pain as well.
- Linked lists are a useful contrast because pointer chasing destroys locality even when the loop body looks simple.
- The benchmark lesson here is relative, not absolute: benchmark on an idle machine and distrust numbers that look too precise on a busy one.
- Four lines of traversal code can still produce wildly different performance if one access pattern helps the hardware and the other fights it.
- The machine is still the model: similar source code can perform very differently when the access pattern changes.

- [Arrays README](gotraining/topics/go/language/arrays/README.md)

### Arrays-Part 2 (Semantics)

- Arrays are values.
- Array length is part of the type, so `[4]int` and `[5]int` are different types.
- Copying an array copies all of its elements.
- An array of five strings is still one contiguous fixed-size value.
- On the teaching machine model used here, each string header is two words, so an array of five strings is a contiguous block of ten words.
- Built-in types like numbers, strings, and booleans usually move through the program with value semantics.
- That same guidance applies to struct fields, not just parameters and returns.
- Strings are the key example because the string value itself is a small fixed-size header even though the bytes it describes live elsewhere.
- Assignment like `data[0] = "apple"` copies only the string value into that array slot: pointer plus length, not the string bytes themselves.
- That means the array keeps value semantics for its elements while the backing bytes for each string can still be shared efficiently.
- This is the blend being taught here:
  - value semantics for the fixed-size header
  - pointer semantics for the unknown-sized backing bytes
- The unknown-size part is the thing worth sharing.
- The fixed-size part is the thing cheap enough to copy freely.
- The main exception for built-in types is when null or absence semantics are required and a plain value is no longer enough.
- The win is balancing value and pointer semantics: copy the small fixed-size string headers freely, and share the unknown-sized backing bytes efficiently.

```text
array value: [5]string

index      0            1            2            3            4
      +------------+------------+------------+------------+------------+
ptr   |    *       |   nil      |   nil      |   nil      |   nil      |
len   |    5       |    0       |    0       |    0       |    0       |
      +------------+------------+------------+------------+------------+
           |
           v
        +-------+
bytes   | apple |
        +-------+

What gets copied into the array slot is the string header.
What stays shared is the backing byte data.
```

```go
numbers := [4]int{10, 20, 30, 40}

var data [5]string
data[0] = "apple"
```

- [Declare and iterate arrays](gotraining/topics/go/language/arrays/example1/example1.go)
- [Different type arrays](gotraining/topics/go/language/arrays/example2/example2.go)
- [Contiguous memory allocations](gotraining/topics/go/language/arrays/example3/example3.go)

### Arrays-Part 3 (Range Mechanics)

- `for range` exposes both value-semantic and pointer-semantic behavior.
- Value-semantic range over an array ranges over a copy of the array.
- Pointer-semantic range indexes directly into the original array.
- The key detail is not only that `v` is a copy; the ranged array itself is copied in the value form.
- If the original array is mutated during value-semantic range, the iteration keeps reading from the copied array.
- If the original array is mutated during pointer-semantic range, later iterations can observe the updated values.
- `for i := range array` is pointer-semantic in effect because indexing happens against the original array.
- `for i, v := range array` is value-semantic because the range expression creates the iteration copy.
- In the value form, `v` is effectively a copy of the copy, which is exactly the isolation value semantics are trying to buy.

```go
friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

for i, v := range friends {
	friends[1] = "Jack"
	_ = i
	_ = v
}
```

- [Array range mechanics](gotraining/topics/go/language/arrays/example4/example4.go)

### Slices-Part 1 (Declare, Length & Reference Types)

- Slices give dynamic behavior while preserving array-based locality.
- A slice is a three-word data structure: pointer, length, and capacity.
- Length is the number of elements currently readable and writable.
- Capacity is the number of elements available from the current pointer to the end of the backing array.
- Slices are reference types.
- Reference types still move across program boundaries with value semantics, but reads and writes happen through pointer semantics to shared backing data.
- The reference-type family here is slices, maps, channels, interfaces, and functions.
- `make` is used for slices, maps, and channels because those types need initialized runtime state or backing storage.
- Review rule: avoid pointers to slices, channels, and interfaces in normal code. Those types already contain the indirection they need.
- If `len` and `cap` differ, length still controls what can be read or written right now; capacity is reserved for future growth.

```go
users := make([]string, 0, 8)
```

- [Slices README](gotraining/topics/go/language/slices/README.md)
- [Declare and length](gotraining/topics/go/language/slices/example1/example1.go)
- [Reference types](gotraining/topics/go/language/slices/example2/example2.go)

### Slices-Part 2 (Appending Slices)

- Nil slice and empty slice are different semantics.
- A nil slice has nil pointer, length `0`, and capacity `0`.
- An empty slice also has length `0`, but it represents empty rather than nil semantics.
- `var` gives the zero value consistently, which is why it preserves nil-slice semantics cleanly.
- Empty literal construction can produce empty semantics instead of nil semantics, which is why `var` remains the safer zero-value signal.
- An empty slice still has a pointer value under the hood, commonly to the runtime's empty-struct sentinel.
- `append` receives its own copy of the slice header, may grow backing storage, and returns a new slice header.
- That makes `append` a value-semantic mutation API over a reference-type data structure.
- When `len == cap`, append must allocate a new backing array.
- Growth is aggressive for smaller slices and smoother for larger ones; the teaching heuristic here is doubling below about `1000` elements and smaller percentage growth after that.
- When growth happens, old backing arrays can become garbage.
- `append` is designed for general use as-is; do not invent custom append logic unless measurement proves a real need.
- Preallocate when you truly know the needed capacity.
- Do not preallocate on a whim. Integrity and readability come first, and the profiler can tell you when growth really matters.
- Do not preset length when the real intention is only to reserve capacity.
- If length is pre-set accidentally, append starts after that length instead of at index `0`, which is a classic bug.
- If exact size is known in advance, direct indexing can be simpler and more efficient than repeated append.
- Common leak suspects in Go include blocked goroutines, cache-like maps with keys never deleted, append patterns that retain old backing arrays, and missed `Close`-style calls.

```go
var data []string
data = append(data, "a", "b", "c")
```

- [Appending slices](gotraining/topics/go/language/slices/example4/example4.go)

### Slices-Part 3 (Taking Slices of Slices)

- Taking a slice of a slice creates a new slice header over the same backing array.
- The new header gets its own pointer, length, and capacity view.
- Capacity is recalculated from the new pointer position to the end of the backing array.
- A helpful mental model is `A : A+len`, then capacity runs from that new pointer to the end of the original backing array.
- Because reads and writes use pointer semantics, subslices can create side effects across sibling views.
- Mutating by index through a subslice mutates shared backing data.
- Appending through a subslice can also mutate shared backing storage when spare capacity still exists.
- Three-index slicing is the tool for clamping capacity when later append must not overwrite shared data.
- Three-index slicing is the way to force copy-on-write behavior on the next append by making `len == cap` in the new view.
- This pattern is useful for cheap views over larger buffers such as binary-protocol parsing.

```go
slice2 := slice1[2:4]
slice2[0] = "CHANGED"
```

- [Taking slices of slices](gotraining/topics/go/language/slices/example3/example3.go)
- [Three-index slicing](gotraining/topics/go/language/slices/advanced/example1/example1.go)

### Slices-Part 4 (Slices & References)

- Taking the address of a slice element and then appending to the slice is a dangerous pattern.
- Append may replace the backing array.
- If that happens, the old element pointer still points into the old backing array instead of the new one.
- That creates stale references and subtle bugs that may only show up under production load.
- These bugs are especially nasty because they can disappear after restart and may be hard to reproduce in tests.
- An `append` outside a clearly bounded build, decode, or unmarshal phase is a review smell worth slowing down for.

```go
shareUser := &users[1]
users = append(users, user{})
shareUser.likes++
```

- [Slices and references](gotraining/topics/go/language/slices/example5/example5.go)

### Slices-Part 5 (Strings & Slices)

- Go strings are byte sequences encoded as UTF-8.
- Source files need to be UTF-8 too if string literals are going to mean what they look like.
- Ranging over a string yields runes, which are Unicode code points.
- A rune is not the same thing as a byte, and it is not always the same thing as a user-perceived character.
- Multi-byte code points are normal in UTF-8.
- The Chinese examples in this section make the point clearly because each rune occupies multiple bytes.
- UTF-8 encoded runes can use up to `4` bytes.
- String slicing operates on byte offsets, not rune counts.
- Copying bytes for a rune into a scratch buffer is one way to inspect the encoded representation directly.
- `copy` works with slices, which is why arrays often get sliced first when used as scratch destinations.
- Every array in Go is just a slice waiting to happen is the practical mental model here: arrays, strings, and slices all keep converging on view-over-contiguous-bytes thinking.

```go
for i, r := range s {
	_ = i
	_ = r
}
```

- [Strings and slices](gotraining/topics/go/language/slices/example6/example6.go)

### Slices-Part 6 (Range Mechanics)

- The same range distinction from arrays applies to slices.
- Value-semantic range copies the slice header for iteration.
- Pointer-semantic range works against the original slice header.
- Value-semantic range protects the iteration shape from later mutations to the original slice variable.
- Pointer-semantic range can panic if the original slice is shortened during iteration and later indexes no longer exist.
- When reviewing code, range behavior is really a question about whether the loop is ranging over a copied view or the live one.

```go
for _, v := range friends {
	friends = friends[:2]
	_ = v
}
```

- [Slice range mechanics](gotraining/topics/go/language/slices/example8/example8.go)
- [Efficient traversals](gotraining/topics/go/language/slices/example9/example9.go)

### Maps

- Maps store key-value pairs.
- Maps must be made before normal writes; the zero-value map is readable but not writable.
- `make` is the standard way to initialize maps for use, even though literal construction can also populate them.
- Reading a missing key returns the zero value for the map's value type.
- The `value, ok := m[key]` form distinguishes absence from present-with-zero-value.
- Zero-value reads are what make patterns like `m[key]++` convenient.
- Keys must be comparable.
- A simple rule is: if the type cannot participate in equality comparison, it cannot be a map key.
- Map iteration order is randomized intentionally.
- That randomness is there so code does not become dependent on a specific map iteration order.
- If deterministic order is needed, collect keys into a slice and sort the slice.
- Map elements are not addressable.
- To update a struct-like value in a map, read it into a local variable, mutate the local copy, then write it back.
- Maps are reference types: the map header is copied by value, while operations mutate shared map state underneath.
- That map header copy is already enough indirection; adding another pointer on top is usually unnecessary.
- Maps are not safe for concurrent read-write use without synchronization.

```go
scores := map[string]int{"anna": 21}
scores["anna"]++
```

- [Maps README](gotraining/topics/go/language/maps/README.md)
- [Declare, write, read, and delete](gotraining/topics/go/language/maps/example1/example1.go)
- [Absent keys](gotraining/topics/go/language/maps/example2/example2.go)
- [Map key restrictions](gotraining/topics/go/language/maps/example3/example3.go)
- [Map literals and range](gotraining/topics/go/language/maps/example4/example4.go)
- [Sorting maps by key](gotraining/topics/go/language/maps/example5/example5.go)
- [Taking an element's address](gotraining/topics/go/language/maps/example6/example6.go)
- [Maps are reference types](gotraining/topics/go/language/maps/example7/example7.go)
