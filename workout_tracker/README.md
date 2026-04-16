# Workout Tracker

Workout Tracker is a Go backend API for user accounts, token-based authentication, and workout management. It is a compact service project, but it already contains the core pieces that make a backend feel real: routing, auth middleware, persistence, migrations, protected CRUD operations, and a set of concrete requests for testing the API by hand.

## What it includes

- user registration via `POST /users`
- token creation via `POST /tokens/authentication`
- protected workout endpoints for create, read, update, and delete
- internal application wiring under `internal/`
- database migrations and local container setup
- curl notes for exercising the API by hand

The route setup already outlines a useful service shape: public health and auth endpoints, plus authenticated workout operations guarded by middleware.

## Main files and folders

- `main.go` starts the HTTP server and application wiring
- `internal/routes/` defines the API surface
- `internal/middleware/` handles authentication and user protection
- `migrations/` contains the database schema changes
- `docker-compose.yml` supports local infrastructure
- `post_notes.txt` and `tokens.txt` capture testing and workflow notes

The notes files are especially useful here because they show the exact payloads used to create users, authenticate, and create or update workouts. That makes the folder feel closer to a working backend lab than a bare code dump.

## Why it matters

Workout Tracker matters because it stays focused on the common backend concerns that show up early in service work: user creation, authentication, protected routes, structured handlers, and persistence-backed CRUD flows. It is small enough to understand quickly, but it already captures the recurring pieces that appear in larger APIs.

That makes it a solid reference for the stage between toy handlers and a more complex production service.

## Local development

Run the backing services with the included Docker setup, then start the Go server and use the example requests in `post_notes.txt` to exercise the API end to end. The folder is well suited to iterative backend practice because the routes, auth flow, and sample payloads are already spelled out.
