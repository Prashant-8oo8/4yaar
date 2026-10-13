# Email Notifications

## Purpose

Send transactional email for workspace invitations, email verification, password recovery (normally managed by the identity provider), mentions, and important security events.

## Provider

Use Amazon SES in AWS or an SMTP-compatible provider. Use Mailpit locally to capture messages without sending real email. Keep the provider behind an interface so it can be replaced without changing domain logic.

## Delivery architecture

For important email, write a notification/outbox record as part of the application transaction, then let a worker deliver it asynchronously. This prevents a provider outage from making a completed workspace invitation or message fail.

## Worker requirements

- Retry transient provider errors with exponential backoff and jitter.
- Cap retries and move persistent failures to a dead-letter/error state.
- Record provider message ID, attempt count, timestamps, and safe error classification.
- Make duplicate processing safe.
- Respect provider quotas and rate limits.
- Monitor queue depth, failure rate, and delivery latency.

## Templates

Keep templates version-controlled and separate from business logic. Include plain-text alternatives and accessible HTML. Escape user-controlled values to prevent HTML injection. Avoid including private message content or unnecessary personal data.

## Links and tokens

Invitation and verification links should use high-entropy, single-use, expiring tokens. Do not put credentials or reusable access tokens in email URLs. Validate the token server-side and make token redemption idempotent where appropriate.

## Configuration

Use environment-specific configuration for sender identity, provider endpoint, and credentials. Store production credentials in a secret manager. Verify domain ownership and configure SPF, DKIM, and DMARC before production sending.

## Testing

- Local messages are captured by Mailpit.
- Templates render required fields and escape untrusted values.
- Provider timeouts trigger retry behavior.
- Permanent failures are surfaced for operator review.
- Duplicate worker delivery does not create duplicate application side effects.
