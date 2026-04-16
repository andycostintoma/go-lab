# Go Lab

Go Lab is the umbrella overview for the Go-focused material in this repo. It mixes language fundamentals, architecture notes, smaller backend services, and larger systems-learning projects, so the best way to use it is to treat this README as the map and each linked area README as the real entry point.

## Areas

- [`challanges`](challanges/README.md) for compact practice exercises around concurrency and data transformations
- [`event_driven_architecture_in_golang`](event_driven_architecture_in_golang/README.md) for notes, diagrams, and chapter code on event-driven systems
- [`go_with_domain`](go_with_domain/README.md) for DDD, clean architecture, CQRS, and the Wild Workouts companion project
- [`learning_go`](learning_go/README.md) for chapter-by-chapter Go fundamentals, concept notes, and example code
- [`reeling_it`](reeling_it/README.md) for a PostgreSQL-backed movie app with accounts, collections, and passkey auth
- [`ride_sharing_app`](ride_sharing_app/README.md) for a distributed Uber-style backend with messaging, infrastructure, and deployment material
- [`ultimate_go`](ultimate_go/README.md) for design-guideline and engineering-judgment notes built around the Ultimate Go material
- [`workout_tracker`](workout_tracker/README.md) for a compact authenticated workout-tracking API with migrations and local Docker setup

## How To Use It

Start with the area that matches the kind of work you want to study:

- fundamentals and language behavior: `learning_go`
- engineering judgment and code quality: `ultimate_go`
- architecture and service boundaries: `go_with_domain`
- asynchronous systems and messaging: `event_driven_architecture_in_golang` or `ride_sharing_app`
- smaller backend application references: `workout_tracker` and `reeling_it`
- short deliberate-practice exercises: `challanges`

## Project Mix

The repo is intentionally uneven in a good way. Some folders are notes-first study areas, some are code-first backend projects, and some combine both. That mix makes `go-lab` useful as both a study archive and a practical reference shelf for real Go work.
