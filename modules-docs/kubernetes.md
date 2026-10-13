# Kubernetes and Helm

## Purpose

Kubernetes runs containerized application workloads, manages service discovery, supports rollout strategies, and restarts failed containers. It should be introduced after the application can be built and tested reliably with Docker.

## Workload choices

- `Deployment` for stateless API, frontend, realtime gateway, and workers.
- `Service` for stable in-cluster networking.
- `Ingress` or a managed gateway for external HTTP/WebSocket traffic.
- `ConfigMap` for non-sensitive configuration.
- Kubernetes `Secret` or an external secret integration for sensitive values; use cloud secret management where possible.
- `Job` for controlled one-off tasks such as migrations, with careful release ordering.
- `HorizontalPodAutoscaler` when metrics and workload behavior justify autoscaling.

## Health checks

- **Startup probe:** allows slow initialization before liveness checks begin.
- **Readiness probe:** indicates whether a pod should receive traffic.
- **Liveness probe:** detects a stuck process that should be restarted.

Do not make liveness fail simply because a remote dependency is temporarily unavailable; that can create restart storms.

## Security

- Run containers as non-root.
- Set CPU/memory requests and limits.
- Use least-privilege service accounts and RBAC.
- Restrict network access where practical.
- Avoid embedding credentials in Helm values committed to Git.
- Use image scanning and trusted registries.
- Apply TLS at the ingress/gateway.
- Separate development, staging, and production environments.

## Helm

Keep templates reusable but avoid over-generalizing the chart. Supply environment-specific values outside the base chart. Validate chart rendering in CI and review any changes that affect permissions, storage, ingress, or secrets.

## Stateful dependencies

Prefer managed PostgreSQL, Redis, and object storage in production unless there is a clear reason to operate them in-cluster. If running stateful systems in Kubernetes, plan persistent storage, backups, upgrades, monitoring, and disaster recovery separately.

## Deployments

Use rolling updates with appropriate readiness checks. Define resource limits and graceful termination behavior, especially for WebSocket connections. Plan how connected clients reconnect during rollout.

## Operational checklist

- Health endpoints work.
- Logs and metrics are collected.
- Secrets are not committed.
- Rollback procedure is documented.
- Database migrations are backward-compatible with rolling deployments.
- WebSocket reconnection works during pod termination.
