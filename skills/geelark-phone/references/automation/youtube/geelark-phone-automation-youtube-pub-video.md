# phone automation youtube-pub-video

YouTube publish video.


## Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--schedule-at <n>` | Schedule time, second-level timestamp (required) |
| `--title <text>` | Title (max 100 chars, required) |
| `--description <text>` | Description (max 5000 chars, required) |
| `--video <text>` | Video URL (required) |
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |
| `--is-disclosure-mandatory` | Whether to force disclosure (default false) |

## Example

```bash
geelark-cli phone automation youtube-pub-video --id "557536075321468390" --schedule-at 1741846843 --title "My video" --description "Description" --video "https://material.geelark.com/a.mp4"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Created task ID |
