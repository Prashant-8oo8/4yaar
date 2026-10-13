# Observability

## Goals

Make it possible to understand system health, diagnose errors, and measure user-facing performance without exposing sensitive data.

## Signals

- **Logs:** structured events with timestamp, severity, service, request ID, and safe context.
- **Metrics:** request rate, error rate, latency, queue depth, database pool saturation, WebSocket connection count, reconnect rate, and worker failures.
- **Traces:** distributed spans for HTTP requests, database operations, gRPC calls, and asynchronous event processing.

Use OpenTelemetry instrumentation where practical. Export to a supported backend; Prometheus and Grafana are suitable for metrics and dashboards, with CloudWatch available for AWS logs/metrics.

## Recommended metrics

- HTTP request count, status, and latency percentiles.
- gRPC call count, failures, and latency.
- PostgreSQL connection pool usage and query latency.
- Redis latency and connection failures.
- Outbox backlog and oldest pending event age.
- Email delivery attempts and failures.
- Neo4j query latency and projection lag.
- WebSocket active connections, disconnects, and event delivery failures.
- File upload/download failures.
- Container restarts and resource saturation.

## Logging rules

Never log passwords, access tokens, cookies, private keys, full signed URLs, or unnecessary message/file contents. Sanitize user-controlled strings to reduce log injection. Apply retention and access controls to logs.

## Alerting

Create alerts around user impact and sustained failure, not every transient error. Examples include elevated API error rates, growing outbox backlog, repeated worker failure, database exhaustion, failed backups, and unusual authentication failures.

## Health endpoints

Provide a lightweight liveness endpoint and a readiness endpoint that reflects whether the service can accept traffic. Avoid returning secret configuration or overly detailed dependency diagnostics to unauthenticated callers.

## Operational practice

Document dashboard links, alert ownership, escalation, common failure modes, and recovery steps. Test that telemetry still works after deployment changes.
