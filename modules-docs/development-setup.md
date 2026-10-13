# Development Setup

## Prerequisites

Install supported versions of:
- Go
- Node.js and the project's chosen package manager
- Docker and Docker Compose
- Git
- Protocol Buffers tooling if gRPC generation is used
- AWS CLI and Terraform only when working on cloud infrastructure

Pin or document versions once the repository's toolchain is selected.

## Environment configuration

Create a local `.env` from `.env.example`. Use placeholders in the example and never commit the real `.env`. Required configuration may include:

- HTTP listen address and port
- PostgreSQL connection URL
- Redis address
- Neo4j URI and credentials
- Keycloak issuer/client configuration
- Object storage endpoint, bucket, and region
- Email provider settings
- GitHub App settings if integration is enabled
- Logging and telemetry configuration

Names and required values must match the actual application configuration code.

## Suggested local workflow

1. Clone the repository.
2. Copy `.env.example` to `.env` and fill in local-only values.
3. Start dependencies using Docker Compose.
4. Wait for health checks to pass.
5. Apply database migrations.
6. Start the Go API.
7. Start the React frontend.
8. Run tests and lint checks before opening a pull request.

Add exact commands to this file as scripts and service names are committed. Do not publish guessed commands that do not match the repository.

## Local email

Use Mailpit or another local capture server so test emails do not reach real recipients. Include the local UI address in the actual Compose documentation after it is configured.

## Troubleshooting

- Check container logs and health status first.
- Verify ports are not already in use.
- Confirm environment variable names match the code.
- Check database migrations and credentials.
- Confirm the Keycloak issuer URL is reachable from the process that validates tokens.
- Verify that browser API and WebSocket URLs match the local deployment.
- Do not disable authentication or TLS checks as a permanent workaround.

## Clean reset

Document a separate explicit command for deleting local volumes and development data. Warn contributors before any destructive reset command.
