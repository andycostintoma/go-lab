# Event Driven Architecture In Golang

Event Driven Architecture In Golang is one of the more systems-heavy areas in `go-lab`. It combines long-form architecture notes, diagrams, chapter summaries, and a substantial chapter-organized codebase to study how event-driven systems behave in Go once the problems move beyond simple synchronous APIs.

## Main pieces

- `Event_Driven_Architecture_In_Golang.md` for the full notes
- `Summary.md` for condensed takeaways and diagrams
- `media/` for the chapter images used by the notes
- `code/` for the practical implementations, organized across modular monoliths, domain events, event sourcing, asynchronous connections, event-carried state transfer, and workflow chapters

Inside `code/`, the material progresses through stages such as:

- `01_Designing_And_Planning`
- `02_Modular_Monolith`
- `03_Domain_Events`
- `04_Event_Sourcing`
- `05_Asynchronous_Connections`
- `06_Event_Carried_State_Transfer`
- `07_Message_Workflows_Orchestrated_Saga`
- `08_Transactional_Messaging`
- `09_Testing`
- `10_Deploying`
- `11_Monitoring_Observability`

## What it covers

The material starts with the core mental model of events as immutable facts, then expands into the harder design questions that appear in real systems:

- thin versus rich events
- brokered communication versus direct calls
- event notification, event-carried state transfer, and event sourcing
- eventual consistency and dual-write problems
- workflows, asynchronous coordination, and system boundaries
- testing, deployment, and operational concerns in distributed systems
- the way services change shape once message flow becomes a first-class concern

The notes and code reinforce each other. The markdown explains the trade-offs and vocabulary, while the chapter implementations show what those choices look like in real Go services.

## Why it matters

This area matters because event-driven architecture is easy to talk about abstractly and much harder to reason about concretely. The combination of diagrams, summaries, and code makes it easier to see where the real friction lives: consistency, replayability, communication flow, state propagation, and the operational cost of looser coupling.

It is especially useful as a reference when moving from ordinary CRUD services into systems where communication patterns and state propagation become part of the core design.

## Working style

Start with `Summary.md` if you want the compressed concepts first, then move into `Event_Driven_Architecture_In_Golang.md` for the fuller walkthrough. After that, use the matching folders in `code/` to study how the architectural ideas show up in actual Go implementations.
