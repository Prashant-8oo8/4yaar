# Storage: PostgreSQL, Redis, and Object Storage

## Purpose

Use each storage system for the kind of data it handles well. Avoid treating every store as an interchangeable database.

## PostgreSQL: source of truth

Store durable application records such as:

- Users and application profiles
- Workspaces, teams, and memberships
- Invitations and role assignments
- Channels, messages, reactions, and read cursors
- Projects, tasks, labels, comments, and dependency edges
- File metadata and access-control relationships
- Notification records and delivery status
- GitHub installation and repository mapping metadata
- Outbox events and audit records

Use UUIDs or another consistent identifier strategy. Define foreign keys, unique constraints, indexes, and deletion behavior intentionally.

## Schema design guidance

- Include `created_at` and `updated_at` where useful.
- Use UTC timestamps.
- Add indexes for common access patterns, such as workspace membership lookup, channel message pagination, task-by-project, and outbox pending status.
- Use cursor pagination for long message timelines.
- Use database constraints to prevent duplicate membership and duplicate reactions where appropriate.
- Define soft-delete versus hard-delete behavior and retention rules explicitly.
- Store JSON only for flexible fields; keep important queryable relationships in normal columns/tables.

## Transactions

Use transactions for changes that must be atomic in PostgreSQL, such as creating a task and its initial audit/outbox event. Do not keep transactions open while sending email or making remote API calls.

## Migrations

Use a migration tool and commit migration files. Migrations should be reviewed, ordered, and safe to run in CI. For production, prefer expand-and-contract changes when old and new application versions may run at the same time. Document backup and rollback considerations; a destructive migration may not be safely reversible.

## Redis

Use Redis for ephemeral or performance-oriented data:

- Presence and typing indicators with TTLs
- Rate-limit counters
- Short-lived cache entries
- Pub/Sub fan-out between realtime gateway instances
- Temporary coordination data where loss is acceptable

Do not store the only copy of a message, task, invitation, or critical notification in Redis. Redis Pub/Sub is not a durable queue. Use a transactional outbox plus a durable processing mechanism for events that must not be lost.

## Object storage

Store file bytes in private Amazon S3 buckets or a compatible local service. PostgreSQL should hold metadata such as object key, owner, workspace, content type, size, checksum, and visibility/ACL references.

## Data consistency

PostgreSQL is authoritative. Update Neo4j through an outbox/event worker, and update Redis as a derived/ephemeral state. Every derived store should have a rebuild or expiry strategy.

## Backups and retention

- Enable managed database backups and point-in-time recovery where available.
- Define retention for messages, files, audit records, and outbox rows.
- Test restores periodically; a configured backup that has never been restored is not a verified recovery plan.
- Encrypt data in transit and at rest.
- Restrict database credentials by environment and least privilege.
