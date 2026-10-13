# Testing Strategy

## Goals

Tests should verify business behavior, authorization boundaries, data integrity, API contracts, and deployment readiness. A passing unit-test suite alone does not prove that the whole system works.

## Test layers

### Unit tests
Use Go's `testing` package for domain logic, validation, permission rules, and pure functions. Mock only boundaries where a real dependency is not needed.

### HTTP/API tests
Test request validation, authentication middleware, status codes, error shapes, pagination, and authorization using test handlers and controlled dependencies.

### Integration tests
Run against real PostgreSQL, Redis, Neo4j, or gRPC servers where correctness depends on their actual behavior. Testcontainers can help create isolated dependencies in CI.

### Frontend tests
Test important components, form validation, loading/error states, keyboard behavior, and realtime event handling.

### End-to-end tests
Automate essential journeys: login, create workspace, invite member, create project/task, send a message, reconnect, and download an authorized attachment.

### Contract tests
Validate REST schemas and protobuf compatibility. Test webhook payload validation and signature verification.

## Essential security cases

- Unauthenticated requests are rejected.
- Users cannot read or mutate another workspace's private data.
- Guests cannot perform admin operations.
- A client cannot spoof sender identity.
- Invalid, expired, wrong-audience, or wrong-issuer tokens are rejected.
- WebSocket subscriptions and file downloads enforce authorization.
- SQL/Cypher input is parameterized.
- Webhook signatures and invitation tokens are validated.

## Reliability cases

- Duplicate outbox processing is idempotent.
- Database failures do not produce false success responses.
- Redis disconnects do not erase durable messages.
- WebSocket reconnect fetches missed history.
- Email provider failures are retried and surfaced.
- Neo4j projection can be rebuilt from authoritative records/events.
- Graceful shutdown does not silently abandon accepted work.

## CI requirements

A pull request should run formatting, static analysis, unit tests, relevant integration tests, frontend checks, and contract generation/verification. Build the application and container image. Add vulnerability and dependency checks appropriate to the repository.

## Test data

Use synthetic data, not production user data. Tests should clean up after themselves or use isolated databases. Never put secrets in test fixtures or logs.

## Coverage

Use coverage as a diagnostic, not as the only quality metric. Prioritize permissions, state transitions, failure paths, and core domain behavior over a high percentage of trivial lines.
