# phone automation tiktok-comment-asia

TikTok AI random comment — Asia region.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--use-ai <n>` | 1=AI comment (Pro only), 2=custom comment (required) |
| `--comment <text>` | Comment content (max 500 chars; required when use-ai=2) |
| `--links <csv>` | Comma-separated specified links |
| `--comment-probability <n>` | Comment probability, 0-100, default 30 |
| `--search-keywords <csv>` | Comma-separated search keywords |
| `--like-video` | Whether to like (default false) |
| `--image-url <text>` | Comment image URL, max 500 chars (effective when use-ai=2) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

## Example

```bash
geelark-cli phone automation tiktok-comment-asia --id "557536075321468390" --schedule-at 1741846843 --use-ai 2 --comment "Great video!"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
