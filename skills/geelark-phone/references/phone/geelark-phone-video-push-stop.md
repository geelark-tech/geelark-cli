# phone video-push-stop

Stop the ongoing live streaming push on a cloud phone.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |

## Examples

```bash
geelark-cli phone video-push-stop --id "628708996374619612"
```

## Response Fields

> This command returns only the standard envelope (`traceId` / `code` / `msg`) with no `data` field. A `code` of `0` indicates success. See [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).
