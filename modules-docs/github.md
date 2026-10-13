# GitHub Integration

## Purpose

GitHub integration can connect a workspace project to repositories and surface useful repository activity, such as pull requests, issues, and commit links. Build it after core workspace, project, and permission flows are stable.

## Integration scope

Possible initial capabilities:
- Connect a GitHub App installation to a workspace.
- Select repositories a workspace may link to projects.
- Store repository metadata and installation identifiers.
- Receive webhooks for selected events.
- Display pull requests and issue references inside related projects.
- Link tasks to branches, commits, pull requests, and issues.

Avoid beginning with write access to repositories. Start with the minimum read/event permissions needed and add write features only with a clear product requirement.

## Prefer a GitHub App

Use a GitHub App rather than asking users to paste personal access tokens. A GitHub App supports installation-scoped permissions and short-lived installation access tokens.

## Installation flow

1. Workspace admin starts the GitHub App installation flow.
2. GitHub redirects to the configured callback.
3. Backend validates the callback/state and records the installation ID.
4. Admin selects repositories available to the workspace.
5. Backend stores repository IDs and mapping metadata in PostgreSQL.
6. Access is revalidated when listing or acting on linked repository data.

Do not trust a repository name alone as a unique identifier; store GitHub's stable repository ID and owner/name for display.

## Webhook security

- Verify webhook signatures using the configured secret and constant-time comparison.
- Reject invalid signatures.
- Deduplicate deliveries using the delivery ID.
- Acknowledge quickly and process heavier work asynchronously.
- Validate event type and payload shape.
- Do not treat webhook payloads as instructions to execute arbitrary commands.
- Log delivery IDs and safe error information, not secrets.

## Data model ideas

- `github_installations`: internal ID, workspace ID, provider installation ID, state, timestamps.
- `github_repositories`: stable repository ID, installation ID, owner/name, default branch, visibility metadata.
- `project_repositories`: project-to-repository mapping.
- `github_webhook_deliveries`: delivery ID, event type, processing state, timestamps.

Use migrations and unique constraints to prevent duplicate installations and webhook deliveries.

## Permissions and secrets

Request the minimum App permissions. Store private keys and webhook secrets in a secret manager. Do not expose installation tokens to the browser or log them. On uninstall or permission changes, disable the integration and invalidate local assumptions.

## Tests

Test signature verification, duplicate deliveries, malformed payloads, installation revocation, repository access boundaries, and retry behavior.
