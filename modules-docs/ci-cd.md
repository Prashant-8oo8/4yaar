# CI/CD with GitHub Actions

## Goals

Automate quality checks, builds, security scans, and controlled deployment so that the same tested artifact can move through environments.

## Pull request workflow

Recommended steps:
1. Check formatting and linting.
2. Run Go unit tests and static analysis.
3. Run frontend type checks, linting, and tests.
4. Run integration tests for supported dependencies.
5. Validate protobuf generation/contracts and API schemas.
6. Build the frontend and backend.
7. Build container images.
8. Run dependency, secret, and image vulnerability checks.
9. Publish test and coverage reports where useful.

## Main-branch workflow

After required checks pass:
- Build immutable image tags based on commit SHA.
- Push images to ECR.
- Deploy to a development/staging environment.
- Run smoke tests.
- Require an approval gate before production where appropriate.
- Record the deployed commit and image digest.

## Secrets and permissions

- Use GitHub Actions secrets or OIDC federation to AWS.
- Prefer short-lived OIDC-based cloud credentials over long-lived AWS access keys.
- Give workflows only the permissions they need.
- Pin third-party actions to reviewed versions or immutable commit SHAs according to repository policy.
- Never print secrets in logs.

## Database migrations

Treat migrations as release artifacts. Validate them in CI, review destructive changes, and coordinate with rolling deployments. Prefer backward-compatible expand-and-contract migrations when old and new application versions may overlap.

## Rollback

Document whether rollback means redeploying a previous image, reverting a configuration change, or running a forward-fix migration. Database changes may not be reversible; backups and tested recovery remain necessary.

## Branch and release policy

Use pull requests and required checks. Define version tagging and release notes when the project reaches stable releases. Do not deploy unreviewed code directly from a developer workstation to production.
