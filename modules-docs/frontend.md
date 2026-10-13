# Frontend

## Stack

- React
- TypeScript
- A maintained routing library
- A query/cache library for server state, if useful
- WebSocket client for realtime updates
- Accessible component and form patterns

Choose exact package versions when creating the application and commit the lockfile.

## Main screens

- Sign-in / authentication callback
- Workspace selector and workspace settings
- Members, teams, and invitations
- Channel list and message view
- Direct messages
- Projects and task boards
- Task details and dependencies
- Notifications
- Profile and account settings
- File attachment preview/download
- Integration settings for GitHub

## Client architecture

Separate route/page components, reusable UI components, API clients, typed data models, authentication state, and realtime connection logic. Keep authorization decisions on the server; hiding a button is not access control.

## API client

- Send requests to the configured API base URL.
- Handle `401` by following the application's token refresh or reauthentication flow.
- Display safe error messages and retain request IDs for support.
- Avoid logging access tokens or sensitive message content.
- Use cancellation/abort support for requests that are no longer needed.
- Implement pagination for long message and task lists.

## Realtime client

- Establish a WebSocket connection only after the user is authenticated.
- Use an explicit event envelope with a type and version.
- Reconnect with exponential backoff and jitter.
- Avoid infinite rapid reconnect loops.
- Track the last received message/event cursor.
- On reconnect, fetch missed history through the API.
- Deduplicate events by stable event/message ID.
- Treat presence and typing as ephemeral state.

## User experience and accessibility

- Provide loading, empty, error, and retry states.
- Support keyboard navigation and visible focus.
- Label form fields and communicate validation errors.
- Avoid relying on color alone to indicate task status.
- Confirm destructive actions where appropriate.
- Preserve unsent message drafts where practical.

## Security

- Do not put long-lived secrets in browser code.
- Do not store sensitive tokens in local storage unless the threat model explicitly accepts that risk.
- Prefer a secure, well-reviewed OIDC flow and an explicit token-storage strategy.
- Validate URLs and file types before rendering previews; never treat file names as trusted HTML.
- Render user-generated content safely and avoid unsafe HTML injection.

## Testing

Use unit tests for components and utilities, integration tests for important flows, and browser end-to-end tests for login, workspace membership, messaging, tasks, and file permissions.
