# Reeling It

Reeling It is a full Go movie application with a PostgreSQL-backed API, server-rendered routes, frontend assets, account features, and passkey-based authentication. It reads much more like a real product slice than a small tutorial app.

## What it includes

- movie discovery endpoints for top, random, search, genres, and individual movie details
- account registration and authentication flows
- authenticated favorites and watchlist collections
- passkey and WebAuthn support for modern login flows
- PostgreSQL-backed data access under `data/`
- HTTP handlers, models, logging, and static frontend assets

The API surface already covers a healthy amount of real application behavior: catalog browsing, account management, protected collections, and alternate authentication mechanisms beyond a simple username/password flow.

## Main files and folders

- `main.go` wires the data layer, handlers, routes, and WebAuthn configuration
- `handlers/` contains the HTTP surface for movies, accounts, and passkeys
- `data/` contains the persistence layer for movies, users, and passkeys
- `public/` contains the frontend assets
- `models/` holds the core application types
- `logger/` contains the custom logging setup
- `docker-compose.yml`, `init.sql`, and `assets/` support local setup and seeded data

The `main.go` wiring shows how the application stitches together database access, custom handlers, static file serving, account flows, and WebAuthn setup in a single Go service.

## Why it matters

Reeling It matters because it combines several concerns that are usually split across separate examples: media content APIs, accounts, protected user-specific state, persistence, SSR-style routes, and modern authentication. That makes it a useful reference for how a mid-sized Go web application starts to fit together once it grows beyond a single endpoint and a simple database model.

It is especially useful as a reference project for HTTP handler structure, data access organization, and the point where authentication features start affecting the overall application shape.

## Local development

The application expects PostgreSQL via `DATABASE_URL`, and the included Docker setup helps provide the database layer for local development. Start there, then run the Go service and exercise the movie, account, and passkey flows through the exposed endpoints and frontend routes.
