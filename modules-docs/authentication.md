# Authentication and Authorization

## Identity provider

Use Keycloak as the OpenID Connect (OIDC) identity provider. The application delegates credential handling and login flows to Keycloak rather than implementing password storage itself.

## Login flow

1. Browser starts the configured OIDC authorization flow.
2. User authenticates with Keycloak.
3. The application receives the authorization response through the configured client flow.
4. The client/backend validates tokens according to the selected architecture.
5. The application maps the identity's stable subject (`sub`) to an application user profile.
6. The backend authorizes each operation against workspace membership and resource permissions.

Use Authorization Code Flow with PKCE for browser-based clients. Decide deliberately between a backend-for-frontend session using secure cookies and a direct browser API-token approach. Document token storage, refresh, logout, and CSRF protections before production.

## Token validation

The API must validate, as applicable:

- Signature using trusted issuer keys.
- Exact expected issuer.
- Expected audience.
- Expiration and not-before timestamps.
- Allowed algorithms.
- Required scopes or roles where relevant.

Do not accept a token merely because it can be decoded. Cache public keys safely and handle key rotation.

## Application user model

Keep an application user record with a stable internal ID, identity-provider subject, display name, avatar/profile fields, timestamps, and account state. Do not use email as the only permanent identity key because email addresses can change.

## Authorization model

Suggested roles:

- `OWNER`: workspace lifecycle and ownership actions.
- `ADMIN`: member management and workspace configuration.
- `MEMBER`: ordinary collaboration and project operations.
- `GUEST`: restricted access to explicitly permitted resources.

Define the exact permission matrix before implementing role checks. Enforce authorization server-side for every REST endpoint, gRPC method, WebSocket subscription, file operation, and integration callback.

## Invitations

- Generate high-entropy, single-use invitation tokens.
- Store a hash of the token rather than the raw token when feasible.
- Bind the invitation to a workspace, intended role, expiry, and optionally an email address.
- Enforce expiry and one-time redemption.
- Do not allow a client to choose an elevated role during redemption.
- Audit creation, redemption, revocation, and expiration.

## WebSocket authentication

Authenticate during the handshake using a strategy compatible with the browser and deployment. Avoid putting long-lived bearer tokens in query strings because URLs may be logged. After connection, authorize each channel subscription and re-check permissions when membership changes.

## Account recovery and security

Use the identity provider's supported verification and password-reset features. Apply rate limits to login-related operations. Do not disclose whether an arbitrary email address has an account. Audit sensitive changes without logging secrets.

## Required tests

- Invalid, expired, wrong-issuer, and wrong-audience tokens are rejected.
- Users cannot access another workspace's private resources.
- Role changes take effect according to the documented session/authorization policy.
- Expired and already-used invitations are rejected.
- WebSocket subscriptions enforce the same access rules as the HTTP API.
