# Security Requirements

## Security principles

- Deny by default.
- Authenticate every protected operation.
- Authorize against the specific resource and workspace.
- Use least privilege for users, services, databases, cloud roles, and CI jobs.
- Protect secrets and personal data.
- Log security-relevant actions without recording credentials or unnecessary private content.

## Application security

- Validate all input on the server.
- Use parameterized SQL and Cypher.
- Limit request sizes and execution time.
- Use rate limits on login, invitation, messaging, upload, and expensive query paths.
- Return safe error messages; never expose stack traces to clients.
- Use CSRF protections where cookie-based authentication requires them.
- Set appropriate browser security headers.
- Render user-generated content safely.

## Authorization

Apply access checks to REST endpoints, gRPC methods, WebSocket channel subscriptions, message sends, file operations, project/task queries, and integration data. Test horizontal privilege escalation across workspaces and vertical escalation between roles.

## Secrets

Never commit `.env`, private keys, tokens, database passwords, or webhook secrets. Use a secret manager in deployed environments, restrict access, rotate credentials, and remove leaked secrets rather than merely deleting them from the latest commit.

## Transport and storage

Use TLS for external and internal network traffic where supported. Encrypt managed storage and backups. Keep object-storage buckets private and use short-lived signed URLs. Avoid exposing database ports publicly.

## Audit events

Record sensitive actions such as role changes, invitations, integration installation/removal, file permission changes, and administrative operations. Include actor, action, target, timestamp, and outcome, while excluding secrets.

## Dependency and supply-chain security

- Keep dependencies updated.
- Review vulnerability alerts.
- Scan container images.
- Pin build tooling.
- Restrict CI permissions.
- Protect release branches and deployment credentials.

## Incident readiness

Document how to rotate credentials, revoke sessions, disable an integration, restore backups, and investigate suspicious access. Define a process for reporting and triaging vulnerabilities.
