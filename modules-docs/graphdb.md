# Graph Database: Neo4j

## Purpose

Neo4j supports relationship-oriented queries that are awkward or expensive to express through repeated relational joins. It is not the primary source of truth for application records.

## Candidate graph model

Nodes:
- `User`
- `Workspace`
- `Team`
- `Project`
- `Task`
- `Channel`
- `Repository` (if GitHub integration is enabled)

Relationships may include:
- `(User)-[:MEMBER_OF]->(Workspace)`
- `(User)-[:PART_OF]->(Team)`
- `(User)-[:WORKS_ON]->(Project)`
- `(Task)-[:BELONGS_TO]->(Project)`
- `(Task)-[:ASSIGNED_TO]->(User)`
- `(Task)-[:DEPENDS_ON]->(Task)`
- `(User)-[:PARTICIPATES_IN]->(Channel)`
- `(Repository)-[:CONNECTED_TO]->(Project)`

Store stable application IDs as properties and index/uniquely constrain them. Keep sensitive personal information out of the graph unless a query genuinely requires it.

## Example: traverse task dependencies

```cypher
MATCH (blocked:Task {id: $taskId})
      -[:DEPENDS_ON*1..5]->(dependency:Task)
RETURN DISTINCT dependency.id AS id,
                dependency.title AS title
```

This example finds dependencies reachable within five edges. Decide whether the relationship direction means “this task depends on that task” and document it consistently. Validate the requested task's workspace access before returning results.

## Example: collaborators on a project

```cypher
MATCH (u:User)-[:WORKS_ON]->(p:Project {id: $projectId})
RETURN DISTINCT u.id AS userId
```

The graph should be queried only after the backend verifies the caller can view the project.

## Synchronization strategy

1. Commit the authoritative change and an outbox event in the same PostgreSQL transaction.
2. Worker reads pending events.
3. Worker applies an idempotent Neo4j upsert/delete.
4. Worker marks the event processed only after success.
5. Retry transient errors with bounded backoff and record persistent failures.
6. Provide reconciliation tools to compare/rebuild the graph from PostgreSQL.

Do not assume a single atomic transaction covers PostgreSQL and Neo4j. Graph results may be briefly stale; identify which user-facing workflows can tolerate that.

## Security and operations

- Use parameterized Cypher queries.
- Use a dedicated database user with minimum required permissions.
- Keep credentials in environment-specific secret storage.
- Apply connection and query timeouts.
- Avoid unbounded variable-length traversals.
- Monitor query latency, failures, memory, and index usage.
- Back up or rebuild the graph according to its role as a derived projection.

## Tests

Test dependency direction, cycles if prohibited, maximum traversal depth, duplicate event delivery, delete handling, and access checks around graph query endpoints.
