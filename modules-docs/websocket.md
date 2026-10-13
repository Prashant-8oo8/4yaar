# WebSockets and Real-Time Events

## Purpose

WebSockets provide a persistent, bidirectional connection for live chat messages, typing indicators, presence, and selected task updates. They do not replace the durable HTTP API or database.

## Connection lifecycle

1. Client opens a WSS connection.
2. Server authenticates the connection.
3. Server records connection state and associates it with the authenticated user.
4. Client requests or joins permitted channels.
5. Server authorizes each subscription.
6. Client and server exchange versioned event envelopes.
7. On disconnect, ephemeral connection state expires or is cleaned up.
8. Client reconnects with backoff and requests missed durable events.

## Event envelope

Use a stable envelope, for example:

```json
{
  "event_id": "uuid",
  "type": "message.created",
  "version": 1,
  "occurred_at": "2026-01-01T12:00:00Z",
  "payload": {
    "message_id": "uuid",
    "channel_id": "uuid"
  }
}
```

The values are illustrative. Define exact schemas in code and document compatibility rules. Do not broadcast private message content to users who are not authorized to see it.

## Message send flow

- Validate message size and format.
- Derive sender identity from the authenticated connection, never from a client-supplied `sender_id`.
- Authorize access to the target channel.
- Persist the message in PostgreSQL.
- Commit an outbox event in the same transaction if reliable publication is required.
- Publish the event and deliver it to authorized connections.
- Return or emit a stable message ID so retries can be deduplicated.

## Redis and multiple instances

Redis Pub/Sub can fan events across gateway instances but does not guarantee replay or durable delivery. Use the outbox for reliable event creation and a durable broker or replay mechanism if the system requires delivery guarantees. Even with event fan-out, clients should recover from gaps through persisted message history.

## Presence and typing

Presence and typing indicators are ephemeral. Use TTLs and refresh them while a connection is active. Do not write every typing event to PostgreSQL. Define how multiple browser tabs/devices contribute to a user's presence.

## Reconnect and ordering

- Use exponential backoff with jitter.
- Track the last known durable message/event cursor.
- Fetch missed history after reconnect.
- Deduplicate by event ID.
- Define whether ordering is guaranteed per channel and how concurrent sends are ordered.
- Do not assume a network reconnect can restore events that existed only in Pub/Sub.

## Security

- Require WSS outside local development.
- Validate the `Origin` header against the expected site origins.
- Enforce message-size limits and per-user rate limits.
- Authorize every subscription and operation.
- Close idle or invalid connections safely.
- Avoid access tokens in query strings.
- Revoke or update subscriptions when access is removed.

## Tests

Test unauthorized subscriptions, malformed messages, oversized payloads, reconnect behavior, duplicate delivery, multiple gateway instances, and history synchronization.
