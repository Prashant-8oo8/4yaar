# Backend

## Stack

- Go for application services.
- Chi for HTTP routing and middleware.
- PostgreSQL for durable records.
- Neo4j driver for graph queries.
- Redis client for cache and event fan-out.
- gRPC and Protocol Buffers for selected internal calls.
- OpenTelemetry for traces and metrics.

## Suggested package layout

```text
services/api/
├── cmd/api/main.go
├── internal/
│   ├── config/
│   ├── http/
│   │   ├── middleware/
│   │   ├── handlers/
│   │   └── responses/
│   ├── auth/
│   ├── workspace/
│   ├── messaging/
│   ├── projects/
│   ├── files/
│   ├── notifications/
│   ├── integrations/
│   ├── storage/postgres/
│   ├── storage/neo4j/
│   └── events/
├── migrations/
├── go.mod
└── README.md
```

Adapt this layout to the actual repository; avoid empty packages that do not yet have a clear responsibility.

## Handler pattern

A handler should parse and validate input, obtain the authenticated principal from context, call application logic, and map the result to an HTTP response. It should not contain complex domain rules or handwritten SQL spread across route functions.

## API conventions

- Prefix endpoints with `/api/v1`.
- Use JSON for ordinary request/response payloads.
- Return appropriate status codes: `400` for invalid input, `401` for unauthenticated requests, `403` for denied access, `404` for missing or intentionally hidden resources, `409` for conflicts, and `500` for unexpected failures.
- Use a consistent error body, for example `{"error":{"code":"invalid_input","message":"The request is invalid","request_id":"..."}}`.
- Add pagination to list endpoints.
- Use UTC timestamps in ISO 8601 format.
- Do not expose internal stack traces or database errors to clients.

## Configuration

Read configuration from environment variables or a typed configuration file supplied at runtime. Validate required settings at startup. Keep safe local defaults only for non-sensitive values. Do not bake secrets into the binary or container image.

## Database access

- Use parameterized queries.
- Apply schema changes through migrations.
- Use transactions for multi-record changes that must succeed together.
- Set connection pool limits and request timeouts.
- Avoid holding transactions open while making network calls.
- Use optimistic concurrency or explicit locking for conflict-prone updates.

## Authentication and authorization

Authentication answers who the user is; authorization answers whether they may perform an action on a resource. Every handler must authorize the operation using the authenticated identity and resource membership. Do not trust user IDs or role names sent in request bodies as proof of identity.

## Reliability

- Set HTTP read, write, and idle timeouts.
- Limit request body sizes.
- Add rate limiting to sensitive and expensive endpoints.
- Use bounded retries only for transient failures.
- Make worker operations idempotent.
- Use graceful shutdown and stop accepting new work before closing dependencies.

## Testing

Add unit tests for domain logic and handlers, integration tests against real dependencies where behavior depends on the database or protocol, and contract tests for APIs and protobuf services.
