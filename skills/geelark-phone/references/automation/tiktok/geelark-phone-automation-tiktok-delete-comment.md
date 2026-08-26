# phone automation tiktok-delete-comment

Delete TikTok comments by keywords. Supports standard and Asia regions.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--keywords <csv>` | Comma-separated keywords, max 100 items, 100 chars each (required) |
| `--name <text>` | Task name (max 32 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation tiktok-delete-comment --id "557536075321468390" --schedule-at 1741846843 --keywords "hello,world"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
