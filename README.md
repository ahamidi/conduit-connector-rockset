# Conduit Connector for Rockset
[Conduit](https://conduit.io) for [Rockset](https://rockset.com).

Records will be written to the configured Rockset collection. As the entire OpenCDC record will be written, the nested
payload will be base64 encoded.

Here is an example query to pull out only the payload from the record and decode it:
You can use the following SQL as the ingest transformation to pull out only the payload (after decoding it):
```sql
SELECT
    TRY_CAST(FROM_BASE64(_input.payload.after) as string) AS payload
FROM _input
```

## How to build?
Run `make build` to build the connector.

## Testing
todo

## Source
todo

### Configuration
todo

## Destination
A destination connector pushes data from upstream resources to a Rockset Collection via Conduit.

### Configuration

| name         | description                            | required | default value |
|--------------|----------------------------------------|----------|---------------|
| `region`     | Rockset region.                        | no       | us-west-2     |
| `collection` | Rockset collection.                    | true     |               |
| `workspace`  | Rockset Workspace.                     | true     |               |
| `api_key`    | API Key for accessing the Rockset API. | true     |               |

## Known Issues & Limitations
* Does not honor rate limiting response from Rockset API

## Planned work
- [ ] Add Source Connector
- [ ] Tests
- [ ] Handle rate limit responses from Rockset API
