# phone video-push-result

Query the execution result of a video push task.

## Key Flags

| Flag | Description |
|------|-------------|
| `--task-id <text>` | Push task ID returned by [`video-push`](geelark-phone-video-push.md) (required) |

## Examples

```bash
geelark-cli phone video-push-result --task-id "2093234110406909952"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `status` | integer | Push status: 0=executing, 1=succeeded, 2=failed |
