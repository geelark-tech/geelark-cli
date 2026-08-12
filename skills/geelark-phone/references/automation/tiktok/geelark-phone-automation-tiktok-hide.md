# phone automation tiktok-hide

Hide TikTok videos. Supports standard and Asia regions.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |
| `--number <n>` | Number of videos to hide, range 0-999; 0 or unset = hide all |

## Example

```bash
geelark-cli phone automation tiktok-hide --id "557536075321468390" --schedule-at 1741846843
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
