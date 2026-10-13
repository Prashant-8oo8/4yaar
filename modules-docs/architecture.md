# Architecture

## Purpose

This document defines the major components, their responsibilities, communication paths, and data ownership boundaries.

## High-level components

- **Web client:** React + TypeScript application running in the browser.
- **HTTP API:** Go + Chi routes for authenticated request/response operations.
- **Application modules:** workspace, membership, messaging, projects/tasks, files, notifications, and integrations.
- **Realtime gateway:** WebSocket connections, authorization, event delivery, and reconnect support.
- **Internal RPC:** gRPC contracts for communication between services where RPC boundaries are useful.
- **PostgreSQL:** durable source of truth for transactional application data.
- **Neo4j:** graph projection used for relationship-heavy queries.
- **Redis:** ephemeral cache, presence, rate-limit counters, and cross-instance event fan-out.
- **Keycloak:** identity provider for OIDC login and token issuance.
- **S3:** object storage for file contents.
- **Email provider:** SES or SMTP-compatible provider for outbound email.

## Recommended initial deployment

Begin with a modular monolith to keep development and deployment manageable. Use Go packages with explicit interfaces and clear ownership. Keep the realtime gateway and notification worker as separate processes only when they need independent scaling or lifecycle management.

Do not create a separate microservice for every noun. A service boundary introduces deployment, networking, observability, failure handling, and data-consistency costs.

## Request flow

1. Browser sends an HTTPS request with an access token.
2. Middleware validates the token's issuer, audience, signature, and expiry.
3. The handler checks the user's membership and permission for the target resource.
4. The application module performs the operation using PostgreSQL transactions as appropriate.
5. An outbox event is committed with the business change when reliable asynchronous processing is needed.
6. Workers publish events to Redis, update Neo4j projections, or send email.
7. The API returns a response with a stable error format.

## Realtime message flow

1. Client connects over WSS and authenticates using the selected secure token strategy.
2. Server validates the identity and authorizes channel access.
3. Client sends a message command with a client-generated idempotency key where appropriate.
4. Server validates and persists the message in PostgreSQL.
5. Server publishes a message-created event.
6. Redis distributes the event to realtime gateway instances.
7. Gateways deliver the event to authorized connected clients.
8. On reconnect, the client requests messages after its last known cursor.

Redis Pub/Sub alone is not durable. Use a transactional outbox and a durable broker or replayable event log if losing a publication would violate a business requirement.

## Data ownership

| Data | Authoritative store | Notes |
|---|---|---|
| Workspace and membership | PostgreSQL | Enforce tenant and role rules in application logic and database constraints where possible |
| Messages and reactions | PostgreSQL | Persist before broadcasting |
| Projects and tasks | PostgreSQL | Transactional updates and dependency validation |
| Relationship projection | Neo4j | Rebuildable from authoritative data/events |
| Presence and typing | Redis / memory | Ephemeral; expiry is expected |
| File bytes | S3 | Use private buckets and signed URLs |
| File metadata and ACLs | PostgreSQL | Authorize before issuing upload/download access |
| Identity credentials | Keycloak | Application stores identity subject/reference |

## Cross-database consistency

PostgreSQL and Neo4j do not share an ordinary atomic transaction. Write the authoritative change and an outbox record in the same PostgreSQL transaction. A worker processes the outbox and updates Neo4j idempotently. Track retries and failures, and provide a reconciliation/rebuild procedure.

## Failure handling

- Database unavailable: fail requests that require durable data; do not pretend writes succeeded.
- Redis unavailable: preserve durable writes where possible, degrade ephemeral features, and retry event publication through durable processing.
- Neo4j unavailable: allow unrelated core workflows to continue; graph queries may be unavailable or temporarily stale.
- Email provider unavailable: queue and retry email instead of rolling back a completed invitation.
- Object storage unavailable: return a clear file-operation error and avoid creating false metadata records.
- WebSocket disconnect: reconnect with backoff and synchronize missed events from persisted history.

## Design rules

- Keep domain logic out of HTTP handlers.
- Use request IDs and structured logs across component boundaries.
- Make asynchronous handlers idempotent.
- Version public APIs and protobuf contracts deliberately.
- Document service ownership and event schemas.
- Use timeouts and bounded retries for network calls.
