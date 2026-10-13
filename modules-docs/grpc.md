# gRPC and Protocol Buffers

## Purpose

Use gRPC for typed internal service-to-service communication where the boundary provides value. It is not necessary to route every internal function call through RPC; a modular monolith can call internal packages directly.

## Contract structure

Keep protobuf files versioned and organized by domain:

```text
proto/
├── workspace/v1/workspace.proto
├── messaging/v1/messaging.proto
├── project/v1/project.proto
└── notification/v1/notification.proto
```

Example:

```proto
syntax = "proto3";

package project.v1;

option go_package = "example.com/dev-collab/gen/project/v1;projectv1";

service ProjectService {
  rpc GetProject(GetProjectRequest) returns (GetProjectResponse);
}

message GetProjectRequest {
  string project_id = 1;
}

message GetProjectResponse {
  string project_id = 1;
  string name = 2;
}
```

Replace the example Go module path with the actual module path in `go.mod`.

## Contract rules

- Do not reuse field numbers after deleting a field; reserve removed numbers and names.
- Add fields in a backward-compatible way.
- Version packages when a breaking contract change is necessary.
- Define error mapping and deadlines.
- Propagate request IDs and tracing context.
- Do not expose internal-only services directly to the public internet.

## Authentication and authorization

Use TLS for networked gRPC connections. Authenticate callers using service identity or a documented token strategy. User context propagated from the API must be treated as untrusted unless the receiving service can verify its provenance. Each service must enforce its own resource permissions or rely on a clearly documented trusted boundary.

## Reliability

Set deadlines on RPCs, cap message sizes, handle cancellation, and avoid unbounded retries. Retry only operations that are safe to repeat or use idempotency keys. Use health checking and graceful shutdown.

## Code generation

Pin compatible `protoc`, Go, and gRPC plugin versions. Add generation commands to the repository's Makefile and verify generated contracts in CI. Decide whether generated code is committed or built during CI, and use one approach consistently.

## Tests

- Test request validation and service logic.
- Test backward compatibility of contracts where feasible.
- Test deadlines, cancellation, authentication, and error mapping.
- Run integration tests against a real local gRPC server.
