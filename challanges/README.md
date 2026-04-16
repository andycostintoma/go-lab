# Challanges

Challanges is a compact Go problem-solving area built around short, focused exercises rather than one long application. The material is small in size, but it covers exactly the kind of thinking that makes Go work feel sharper: splitting work across goroutines, controlling concurrency, structuring transformations cleanly, and proving behavior with tests.

## Structure

- `concurrency/` contains concurrent price-fetching exercises implemented with different coordination strategies
- `data-transformations/` contains smaller business-logic exercises such as order and shipping-cost calculations

Each sub-area has its own Go module, which keeps the feedback loop quick and the scope deliberately tight.

## What the exercises emphasize

The concurrency work is centered on familiar backend concerns:

- controlling fan-out work without flooding downstream systems
- limiting parallelism with semaphores
- distributing work with worker pools
- writing code that stays understandable even when goroutines are involved

The data-transformation exercises focus on a different but equally useful skill set:

- shaping business rules into small pure functions
- turning raw inputs into derived outputs cleanly
- testing transformations without framework noise
- practicing package structure and basic design decisions in miniature

## Why it matters

The value here is repetition on small problems. Larger projects are useful for architecture, but they often hide the individual skills that need deliberate practice. Challanges is where those skills stay visible: synchronization, work distribution, predictable transformations, and test-first validation of logic.

That makes the folder a good sharpening area between bigger application projects. It is the kind of material that improves how quickly you can reason about small units of Go code without needing an entire service around them.

## Working style

The workflow is intentionally lightweight:

```bash
go test ./...
```

Use each exercise as a small lab: read the problem, inspect the tests, implement the behavior, and compare approaches. The folder is most useful when treated as deliberate practice rather than just reference material.
