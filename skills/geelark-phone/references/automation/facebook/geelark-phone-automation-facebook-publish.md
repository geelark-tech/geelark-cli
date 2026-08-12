# phone automation facebook-publish

Facebook post content with video.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--title <text>` | Title (max 200 chars, required) |
| `--video <csv>` | Comma-separated video URLs, max 10 (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |
| `--need-share-link` | Whether to retrieve the sharing link (default false) |

## Example

```bash
geelark-cli phone automation facebook-publish --id "557536075321468390" --schedule-at 1741846843 --title "My post" --video "https://material.geelark.com/a.mp4"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
