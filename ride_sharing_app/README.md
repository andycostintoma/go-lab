# Ride Sharing App

Ride Sharing App is a microservices learning project for building an Uber-style backend in Go. It combines architecture notes with a substantial implementation under `code/`, and it pushes beyond “multiple services” into the harder parts of distributed systems: asynchronous messaging, real-time communication, deployment, infrastructure, and cross-service coordination.

## Main pieces

- `Ride_Sharing_App.md` for the running architecture and implementation notes
- `media/` for the screenshots and diagrams used by the notes
- `code/` as a Git submodule pointing to the standalone repo: `https://github.com/andycostintoma/ride-sharing-app`

The nested `code/README.md` now acts as the entry point for the standalone project repo while this folder keeps the higher-level notes and media around it.

## What it covers

The system evolves from straightforward service boundaries into a fuller distributed design with:

- an API gateway at the edge
- separate trip, driver, and payment services
- shared protobuf contracts and infrastructure code
- RabbitMQ-driven messaging and event flow
- WebSocket support for real-time communication
- Kubernetes-oriented local and production deployment paths
- frontend and infrastructure layers that sit around the services instead of being hand-waved away

The notes make that architectural progression explicit, while the codebase shows how the pieces are wired together in practice.

## Why it matters

This area matters because it does not stop at individual services. It focuses on service collaboration: trip creation, driver assignment, payment flow, notifications, broker durability, message handling, and real-time updates. In other words, it treats microservices as an interaction problem, not just a folder layout.

That makes it one of the stronger systems-learning projects in `go-lab`, especially if the goal is to understand how an application changes once messaging, deployment, and service coordination become central.

## Working style

Use `Ride_Sharing_App.md` as the guided walkthrough for the system design, then move into `code/` when you want the concrete implementation details. The folder works best when treated as both an architecture case study and an implementation reference.
