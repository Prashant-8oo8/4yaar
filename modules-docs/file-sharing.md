# File Sharing and Attachments

## Purpose

Support file attachments in messages and project/task contexts while keeping object contents out of the relational database.

## Storage model

- **S3:** stores bytes under generated object keys.
- **PostgreSQL:** stores metadata, owner, workspace, object key, size, content type, checksum, and access relationships.
- **API:** authorizes uploads and downloads and issues short-lived signed URLs.

Do not use a user-supplied filename as the object key. Treat the original filename as untrusted display metadata.

## Upload flow

1. Client requests an upload for a specific workspace/resource.
2. Backend authenticates the user and checks permission.
3. Backend validates allowed size and content type policy.
4. Backend creates a pending metadata record and a unique object key.
5. Backend returns a short-lived signed upload URL or accepts a streamed upload.
6. Client uploads bytes directly to S3 when using signed URLs.
7. Backend verifies completion, size, and other required metadata.
8. Metadata is marked ready only after validation succeeds.

For high-risk deployments, scan uploaded files for malware and consider quarantine before making them available.

## Download flow

1. Client requests a file by metadata ID.
2. Backend checks current authorization and workspace/resource access.
3. Backend issues a short-lived signed URL or streams the file.
4. Client downloads from private storage.

A signed URL is a bearer capability: anyone holding it may use it until expiry. Keep expiry short and avoid logging full signed URLs.

## Security

- Keep buckets private and block public access.
- Generate random object keys.
- Enforce maximum file sizes and allowed file types.
- Do not trust file extensions or browser-supplied content types.
- Set safe `Content-Disposition` and content-type headers.
- Escape filenames in the UI.
- Prevent cross-workspace file access.
- Define deletion, retention, and orphan-object cleanup.
- Encrypt data at rest and in transit.

## Testing

Test authorization for cross-workspace access, expired signed URLs, oversized uploads, incomplete uploads, duplicate completion callbacks, deletion cleanup, and malicious filenames.
