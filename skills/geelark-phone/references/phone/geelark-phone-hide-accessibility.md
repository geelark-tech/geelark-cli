# phone hide-accessibility

Hide the cloud phone accessibility service from specified apps. Supports Android 12/13/14/15/16. Overwrites previous configuration.

## Key Flags

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Cloud phone IDs (required) |
| `--pkg-name <csv>` | App package names (required). On Android 14/16, pass the app's own package name. On Android 12/13/15, pass the package name of the app you want to hide from (e.g. com.zhiliaoapp.musically for TikTok). |

## Examples

```bash
geelark-cli phone hide-accessibility --ids "id1,id2" --pkg-name "com.zhiliaoapp.musically"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `failDetails[]` | array | Failed details |
| `failDetails[].id` | string | Cloud phone ID |
| `failDetails[].code` | integer | Error code |
| `failDetails[].msg` | string | Error message |

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 43037 | Does not support this device |
