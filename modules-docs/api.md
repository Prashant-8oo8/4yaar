# HTTP API Design

## Conventions

- Base path: `/api/v1`.
- JSON request/response bodies unless a file stream is required.
- Authentication through the selected OIDC/session strategy.
- Consistent error format.
- UTC timestamps.
- Cursor pagination for large or frequently changing lists.
- Idempotency keys for operations where client retries could create duplicates.

## Example endpoint groups

| Method | Path | Purpose |
|---|---|---|
| GET | `/api/v1/me` | Current application profile |
| GET / POST | `/api/v1/workspaces` | List or create workspaces |
| GET / PATCH | `/api/v1/workspaces/{workspaceId}` | Read or update a workspace |
| GET / POST | `/api/v1/workspaces/{workspaceId}/members` | List or add members where allowed |
| POST | `/api/v1/workspaces/{workspaceId}/invitations` | Create an invitation |
| GET / POST | `/api/v1/workspaces/{workspaceId}/channels` | List or create channels |
| GET / POST | `/api/v1/channels/{channelId}/messages` | List or create messages |
| GET / POST | `/api/v1/projects` | List or create projects |
| GET / POST | `/api/v1/projects/{projectId}/tasks` | List or create tasks |
| GET / PATCH | `/api/v1/tasks/{taskId}` | Read or update a task |
| POST / DELETE | `/api/v1/tasks/{taskId}/dependencies` | Add/remove dependency edges |
| POST | `/api/v1/files/uploads` | Start a file upload |
| GET | `/api/v1/files/{fileId}/download` | Authorize and start a download |

These are illustrative routes. Finalize route naming and payloads in the API specification before implementation.

## Authorization

Every endpoint must check both identity and resource access. A valid token alone does not mean the user can access a workspace, channel, project, task, or file. Return `404` instead of `403` where the product's privacy policy intentionally hides resource existence.

## Pagination

Use stable cursor pagination for message history and other large lists. Define cursor ordering, tie-breakers, and behavior when records are deleted. Avoid unbounded list responses.

## Errors

Use a consistent response such as:

```json
{
  "error": {
    "code": "permission_denied",
    "message": "You do not have access to this resource",
    "request_id": "request-id"
  }
}
```

Do not expose SQL statements, stack traces, internal hostnames, or tokens.

## Versioning and documentation

Document request/response schemas, authentication, authorization requirements, pagination, error codes, and examples. Add contract tests for important endpoints and maintain backward compatibility where clients may update independently.
