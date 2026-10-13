# Contributing to 4Yaar

First of all, thank you for your interest in contributing to this project! Every contribution, whether it is a bug fix, feature, documentation improvement, test, or suggestion, is valuable.

Our goal is to build a secure, reliable, maintainable, and collaborative platform for developers. We welcome contributions from people of all experience levels who are willing to learn, share knowledge, and work respectfully with others.

## 1. Code of Conduct and Professional Ethics

All contributors are expected to follow these principles:

### Respect and Inclusivity
- Treat every contributor with respect, regardless of their experience, background, identity, or skill level.
- Be polite, patient, and constructive in discussions, issues, pull requests, and code reviews.
- Welcome different perspectives and provide feedback without personal attacks, harassment, discrimination, or insulting language.
- Help beginners learn instead of criticizing them for asking questions.

### Honesty and Integrity
- Submit work that you understand and can explain.
- Do not claim credit for someone else's work.
- Acknowledge original authors and respect applicable licenses.
- Never fabricate test results, benchmarks, research findings, or implementation details.
- Clearly disclose borrowed code, third-party dependencies, and relevant AI-assisted contributions when required by project policy or applicable licenses.
- Do not misrepresent unfinished work as completed or tested.

### Responsibility and Accountability
- Take responsibility for your changes and their potential impact.
- Communicate blockers, limitations, and mistakes honestly.
- Respond to review feedback professionally.
- Correct problems introduced by your changes when reasonably possible.
- Avoid making unnecessary changes outside the scope of your contribution.

### Collaboration and Knowledge Sharing
- Explain the reasoning behind significant technical decisions.
- Prefer constructive suggestions over criticism alone.
- Share useful documentation and reproducible examples.
- Give credit to collaborators and recognize their contributions.
- Resolve disagreements through respectful discussion and evidence.

### Privacy and Confidentiality
- Never expose passwords, access tokens, API keys, private user information, or confidential project data.
- Do not publish private discussions, screenshots, logs, or personal information without appropriate authorization.
- Use synthetic or anonymized data in examples and tests.
- Follow the project's security and data-handling requirements.

Violations of these principles may result in review restrictions or other proportionate action by project maintainers.

## 2. Ways to Contribute

You can contribute by:
- Fixing bugs and improving existing features.
- Developing frontend components using React and TypeScript.
- Developing backend APIs and services using Go and Chi.
- Improving real-time collaboration using WebSockets.
- Working on PostgreSQL, Neo4j, Redis, or file storage integrations.
- Improving authentication, authorization, and application security.
- Writing unit, integration, and end-to-end tests.
- Improving documentation, developer experience, and deployment configuration.
- Reporting reproducible bugs and suggesting useful features.

For significant changes, open an issue or discuss the proposal with maintainers before beginning implementation.

## 3. Getting Started

### Prerequisites

Install the tools required for the part of the project you intend to work on:

- Git
- Go
- Node.js and npm
- PostgreSQL
- Docker and Docker Compose, where required
- Additional services such as Redis, Neo4j, and Keycloak when working on related modules

Check the project's `README.md` and `modules-docs/` directory for the current setup instructions.

### Fork and Clone

1. Fork the repository.
2. Clone your fork.
3. Add the original repository as an upstream remote if applicable.
4. Create a dedicated branch for your work.

Example:

```bash
git clone <your-fork-url>
cd dev-collab

git remote add upstream <original-repository-url>

git checkout -b feature/project-members
```

Replace the placeholder URLs with the actual repository URLs. If you have already configured the upstream remote, do not add it again.

## 4. Git Branch and Commit Guidelines

Use clear, descriptive branch names:

- `feature/project-members`
- `fix/websocket-reconnect`
- `docs/update-setup-guide`
- `test/project-service`
- `refactor/api-middleware`

Prefer small, focused commits with meaningful messages.

Examples:

```text
feat: add project member invitations
fix: handle expired authentication tokens
test: add project service unit tests
docs: improve local development setup
refactor: simplify websocket connection handling
```

Where practical, follow the Conventional Commits format:

`type: short description`

Common types include `feat`, `fix`, `docs`, `test`, `refactor`, `perf`, and `chore`.

Do not commit generated build files, local environment files, credentials, or unrelated changes.

## 5. Coding Standards

### Go Backend

- Follow standard Go formatting using `gofmt`.
- Use clear package names and descriptive identifiers.
- Handle errors explicitly; do not silently ignore failures.
- Pass `context.Context` through operations that need cancellation or deadlines.
- Keep HTTP handlers focused on request validation and response handling.
- Place business rules in service or domain layers.
- Keep database access in repositories or clearly defined data-access packages.
- Use parameterized SQL queries to prevent SQL injection.
- Avoid unnecessary global state and tightly coupled dependencies.
- Add tests for new business logic and bug fixes.

Run:

```bash
gofmt -w .
go vet ./...
go test ./...
```

### React and TypeScript Frontend

- Use TypeScript types and interfaces where they improve clarity.
- Prefer small, reusable components with clear responsibilities.
- Follow existing naming conventions and project linting rules.
- Avoid `any` unless there is a justified reason.
- Keep API communication and UI rendering concerns reasonably separated.
- Handle loading, empty, success, and error states.
- Ensure keyboard accessibility and understandable form validation.
- Do not expose server secrets or privileged credentials in frontend code.
- Add component tests for important UI behavior.

Typical commands, depending on the configured npm scripts:

```bash
npm run lint
npm run test
npm run build
```

Run these from the frontend application directory. The actual script names must match `package.json`.

### Database and Storage

- Use version-controlled database migrations.
- Avoid destructive schema changes without an appropriate migration and review plan.
- Use transactions where multiple changes must succeed or fail together.
- Validate input and enforce appropriate constraints.
- Avoid storing sensitive information unnecessarily.
- Keep file metadata and file contents in their appropriate storage systems.
- Test migrations and database operations against a suitable test environment.

### APIs and Real-Time Communication

- Keep API responses and error handling consistent.
- Validate all incoming data on the server.
- Check authentication and authorization for protected operations.
- Validate WebSocket messages and enforce message-size limits.
- Avoid broadcasting private information to unauthorized users.
- Maintain compatibility or document breaking changes to API contracts and Protocol Buffer definitions.

## 6. Testing Requirements

Testing is an essential part of every contribution.

Before submitting a pull request:

- Run the relevant unit tests.
- Run integration tests when database or external-service behavior changes.
- Run frontend tests for UI changes.
- Run end-to-end tests for important user journeys when applicable.
- Add regression tests for bug fixes.
- Report any tests you could not run and explain why.

Never claim that tests passed unless you actually ran them and verified the results.

Contributions should not reduce test coverage unnecessarily or weaken existing security checks merely to make tests pass.

## 7. Pull Request Process

1. Sync your branch with the latest relevant changes.
2. Review your own diff for mistakes, debugging code, secrets, and unrelated edits.
3. Run the appropriate tests and quality checks.
4. Push your branch to your fork.
5. Open a pull request against the appropriate target branch.
6. Explain what changed, why it changed, and how it was tested.
7. Link related issues and include screenshots for significant UI changes where useful.
8. Respond respectfully to review feedback and update your changes as needed.

### Pull Request Checklist

Before requesting review, confirm that:

- [ ] The change has a clear purpose.
- [ ] The implementation follows project conventions.
- [ ] Tests were added or updated where appropriate.
- [ ] Relevant tests and checks were run.
- [ ] Documentation was updated if behavior or setup changed.
- [ ] No credentials, secrets, or private user data were committed.
- [ ] No unrelated files or generated artifacts were included.
- [ ] Breaking changes and known limitations are documented.
- [ ] Third-party code and dependencies are appropriately licensed.

Maintainers may request changes before approving or merging a pull request.

## 8. Code Review Ethics

Code review should improve the quality of the project while supporting contributors.

Reviewers should:
- Focus on correctness, security, maintainability, performance, and readability.
- Explain the reasoning behind requested changes.
- Distinguish critical issues from optional suggestions.
- Avoid hostile, dismissive, or personal comments.
- Acknowledge good solutions and useful improvements.
- Disclose relevant conflicts of interest when appropriate.

Contributors should:
- Consider feedback carefully and ask questions when something is unclear.
- Disagree respectfully and provide technical reasoning.
- Avoid taking technical feedback personally.
- Request clarification when requirements are ambiguous.
- Resolve discussions collaboratively instead of attempting to win arguments.

## 9. Security and Responsible Disclosure

Security vulnerabilities must be handled carefully.

- Do not publish exploitable vulnerability details in a public issue before maintainers have had an opportunity to assess the risk.
- Report sensitive vulnerabilities through the repository's designated private security reporting channel.
- Never test against systems, accounts, or data without authorization.
- Do not commit credentials, private keys, tokens, or production configuration.
- Do not weaken authentication, authorization, validation, or logging safeguards to bypass a failing test.
- Include security-focused tests for changes involving permissions, sessions, uploads, or real-time events when applicable.

If a secret is accidentally committed, notify maintainers immediately and rotate or revoke the exposed credential. Removing it from the latest commit alone may not remove it from repository history.

## 10. Dependencies and Third-Party Code

Before introducing a new dependency:

- Confirm that it is necessary and actively maintained where possible.
- Review its license, security posture, and compatibility.
- Prefer existing project dependencies when they already solve the problem adequately.
- Avoid unnecessary dependency duplication.
- Document important new dependencies and configuration requirements.

Do not copy code from proprietary or incompatible sources. Follow the applicable licenses and attribution requirements.

## 11. Documentation and Communication

Documentation is part of the contribution, not an optional afterthought.

Update the relevant documentation when changing:
- Installation or configuration.
- Environment variables.
- API endpoints or request formats.
- Database schemas or migrations.
- Authentication and authorization behavior.
- Testing commands.
- Deployment procedures.
- User-visible behavior.

Use clear language, reproducible examples, and accurate instructions. Never include real credentials in example configuration files.

## 12. Scope, Quality, and Maintainability

Prefer solutions that are understandable, testable, and appropriate for the project's current needs.

- Avoid premature optimization and unnecessary abstraction.
- Do not introduce microservices or complex infrastructure without a clear justification.
- Keep changes focused and reviewable.
- Consider backward compatibility and the effect on existing contributors.
- Document trade-offs when there is no obviously ideal solution.
- Avoid hiding technical debt behind misleading comments or documentation.

A smaller, well-tested contribution is often more valuable than a large, unreviewable change.

## 13. License and Ownership

By submitting a contribution, you confirm that you have the right to submit it and that it complies with the repository's applicable license and contribution terms.

Contributions remain subject to the project's existing license and any applicable contributor agreement. Do not assume that submitting a pull request transfers copyright or grants rights beyond the terms that apply to the repository.

## 14. Questions and Support

If you need help:

- Read the root `README.md` and relevant files in `modules-docs/`.
- Search existing issues and pull requests before opening a duplicate.
- Open an issue for a reproducible bug or a well-defined feature proposal.
- Use the repository's designated discussion channel if one exists.
- Share only the logs and configuration details needed to diagnose the problem, with sensitive information removed.

Please be patient while maintainers review issues and contributions.

## Final Note

The best contributions combine technical quality with honesty, respect, responsibility, and a willingness to learn.

**Be respectful. Be transparent. Test your work. Protect user data. Give credit. Help others succeed.**

Thank you for helping make this project better for everyone.
