# Docker and Docker Compose

## Purpose

Docker provides repeatable build artifacts. Docker Compose is the recommended early-development tool for running local dependencies consistently.

## Containerization guidelines

- Use multi-stage builds for Go services.
- Use a small runtime image and a non-root user where possible.
- Pin base image versions and update them deliberately.
- Do not copy `.env` files, local credentials, source-control metadata, or test artifacts into production images.
- Add `.dockerignore`.
- Expose only the required port.
- Send logs to stdout/stderr.
- Handle SIGTERM and shut down gracefully.
- Configure health checks where supported by the runtime/orchestrator.

## Local services

A Compose environment may include:
- PostgreSQL
- Neo4j
- Redis
- Keycloak
- Mailpit
- The Go API and React application when useful

Use named volumes for local database persistence. Document how to reset development data; make destructive reset commands explicit.

## Configuration

Use a local `.env.example` containing placeholder values, and keep `.env` ignored by Git. The example file must not contain real passwords or keys. Production secrets should be injected by the deployment platform or secret manager.

## Image build

Build images in CI from a known commit. Tag images with immutable identifiers such as the commit SHA and optionally a release version. Avoid relying only on a mutable `latest` tag for deployment.

## Local checks

- Services start with documented commands.
- Health/readiness endpoints report dependency status appropriately.
- Migrations can be run repeatably.
- Container logs contain actionable errors without secrets.
- Volumes and reset instructions are documented.

Add exact Compose commands once the repository's Compose file and service names are finalized.
