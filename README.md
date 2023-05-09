# Conduit Connector for Rockset
[Conduit](https://conduit.io) for [Rockset](https://rockset.com).

Records will be written to the configured Rockset collection. As the entire OpenCDC record will be written, the nested
payload will be base64 encoded.

Here is an example query to pull out only the payload from the record and decode it:
```sql
SELECT
    CAST(FROM_BASE64(test.payload.after) as string)
FROM
    commons.test
LIMIT
    10
```
**Note:** The above query assumes the collection is named `commons` and the collection is named `test`.

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
* Complete OpenCDC record is written to Rockset. This includes the nested payload which is base64 encoded.

## Planned work
- [ ] Add Source Connector
- [ ] Tests
- [ ] Handle rate limit responses from Rockset API
