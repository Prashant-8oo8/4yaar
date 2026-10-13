# Roadmap

This roadmap is a suggested sequence, not a promise of dates. Adjust it based on available time, implementation progress, and the size of the first release.

## Phase 1: Foundation

- Repository layout, configuration, logging, and error conventions.
- Go HTTP API and React application.
- PostgreSQL migrations and repository layer.
- Keycloak OIDC login.
- User profiles, workspaces, membership, and roles.
- Unit and API tests.

**Exit criteria:** users can sign in and workspace access is enforced.

## Phase 2: Projects and tasks

- Projects, tasks, status transitions, assignment, labels, due dates, and comments.
- Dependency validation and authorization.
- Pagination and audit events.

**Exit criteria:** members can manage project work securely.

## Phase 3: Messaging

- Channels and direct messages.
- Persisted message history and pagination.
- WebSocket gateway and event schemas.
- Presence, typing, reconnect, and missed-history synchronization.
- Rate limiting and authorization tests.

**Exit criteria:** multiple users can collaborate in real time without losing durable message history.

## Phase 4: Graph and internal RPC

- Neo4j projection from PostgreSQL outbox events.
- Dependency traversal and selected relationship queries.
- gRPC service contracts and code generation.
- Retry, idempotency, and projection-rebuild procedures.

**Exit criteria:** graph results are correct, access-controlled, and recoverable.

## Phase 5: Files, email, and GitHub

- Private object storage and signed URLs.
- Invitation and notification email worker.
- GitHub App installation and webhook validation.
- Integration-specific permission and retry tests.

**Exit criteria:** external operations are secure and failure-tolerant.

## Phase 6: Delivery and operations

- Docker images and Compose setup.
- CI checks and integration tests.
- OpenTelemetry, metrics, dashboards, and alerts.
- Kubernetes manifests/Helm chart.
- AWS infrastructure using Terraform.
- Backup, restore, and rollback runbooks.

**Exit criteria:** deployment is repeatable, observable, and recoverable.

## Future enhancements

- Full-text search.
- GitHub/GitLab activity views.
- Collaborative document editing.
- Code review workflows.
- Advanced notifications and user preferences.
- Audit export and retention controls.
- Recommendation features based on explicit, privacy-aware collaboration data.

## Scope control

Do not add every infrastructure component at the beginning. Build a working vertical slice first, validate the architecture with real use, then introduce additional services where they solve a demonstrated need.
