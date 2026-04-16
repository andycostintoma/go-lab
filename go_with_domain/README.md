# Go With Domain

Go With Domain is a DDD- and architecture-focused study area for writing Go services with stronger boundaries, clearer models, and better long-term maintainability. The emphasis is not on fancy abstractions for their own sake, but on the design pressure that appears once business logic, persistence, testing, and service boundaries start pulling in different directions.

## Main pieces

- `Go_with_Domain.md` for the main notes and diagrams
- `media/` for the supporting figures used throughout the notes
- `wild-workouts-go-ddd-example/` as the practical companion project from Three Dots Labs

The added Wild Workouts submodule gives the notes a real codebase to compare against while reading through the architectural ideas.

## What it covers

The material is centered on business-application structure rather than syntax drills. The notes dig into:

- shaping domain logic so that rules stay close to the business model
- separating transport, application, domain, and infrastructure concerns
- testing strategy for service code and integration boundaries
- persistence patterns and how not to let storage concerns dominate the model
- refactoring toward cleaner architecture in stages instead of rewriting everything at once
- CQRS and the places where read and write concerns benefit from different shapes

It is very much about trade-offs: where DDD helps, where it adds cost, and how to apply the ideas in Go without losing the language’s preference for straightforward code.

## Why it matters

This area matters because many Go services start simple and then accumulate pressure from real business requirements. At that point, naming, boundaries, testing, and data flow start mattering more than framework mechanics. Go With Domain focuses on that transition.

It is a good bridge between idiomatic Go and larger architectural thinking. The notes capture the reasoning, and the Wild Workouts example gives those ideas a concrete shape in a non-trivial application.

## Working style

Read `Go_with_Domain.md` first for the architecture walkthrough, then move through the `wild-workouts-go-ddd-example` submodule when you want to see how the same ideas look in a fuller application. The area works best when the notes and the example are read side by side.
