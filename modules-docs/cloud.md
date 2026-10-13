# Cloud Deployment on AWS

## Purpose

Use AWS managed services where practical to reduce operational overhead and improve security, backup, and monitoring capabilities.

## Suggested service mapping

| Need | AWS service |
|---|---|
| Kubernetes | Amazon EKS |
| Container registry | Amazon ECR |
| Relational database | Amazon RDS for PostgreSQL |
| Redis-compatible cache | Amazon ElastiCache |
| File/object storage | Amazon S3 |
| Email | Amazon SES |
| Secrets | AWS Secrets Manager or Systems Manager Parameter Store |
| Logs and metrics | Amazon CloudWatch, with Prometheus/Grafana as selected |
| Identity | Keycloak hosted and operated securely, or a separately selected managed identity service if requirements change |
| DNS/TLS entry | Route 53 plus a suitable load balancer/certificate service |

This is a suggested mapping, not a requirement to use every service from day one.

## Network design

- Place application workloads and data services in an appropriately designed VPC.
- Keep databases and caches private.
- Expose only the intended public entry point.
- Restrict security groups to necessary source/destination ports.
- Use private endpoints where practical.
- Design outbound access for image pulls, email, GitHub APIs, and required updates.
- Decide how operators access private resources securely.

## Identity and secrets

Use IAM roles for workloads where supported rather than long-lived access keys. Store database credentials, signing secrets, webhook secrets, and provider credentials in a secret manager. Rotate credentials and audit access.

## Data protection

- Enable encryption at rest and in transit.
- Configure PostgreSQL backups and point-in-time recovery.
- Apply private bucket policies and block public access on S3.
- Define object lifecycle and retention policies.
- Test restoration and recovery, not only backup creation.
- Document RPO (acceptable data loss) and RTO (acceptable recovery time).

## Cost controls

- Set budgets and billing alerts.
- Right-size database and Kubernetes capacity.
- Use autoscaling only with suitable metrics and limits.
- Apply S3 lifecycle policies.
- Remove unused development resources.
- Track data transfer and managed-service costs.

## Deployment sequence

1. Provision network and IAM foundations.
2. Create ECR repositories.
3. Provision RDS, Redis, buckets, secrets, and monitoring.
4. Create or configure EKS and ingress/TLS.
5. Deploy the application with Helm.
6. Run migrations through a controlled release step.
7. Verify readiness, logs, metrics, authentication, uploads, email, and WebSocket reconnect.
8. Test rollback and database restore procedures.

## Environments

Keep development, staging, and production isolated. Use separate credentials and data. Avoid copying production data into development. Define who can approve production changes and access sensitive resources.
