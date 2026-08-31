# phone video-push

Upload a video file to a cloud phone and start live streaming push. MP4 format is recommended.

> Only supported on Android 12/13/14/15/16 devices.

## Key Flags

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID (required) |
| `--file-url <text>` | Video file URL, MP4 recommended (required) |
| `--is-loop <n>` | Whether to loop the stream: 1=yes, 0=no (required) |

## Examples

```bash
geelark-cli phone video-push --id "628708996374619612" --file-url "https://material.geelark.com/open-upload/xxx.mp4" --is-loop 1
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskId` | string | Push task ID — pass to [`video-push-result`](geelark-phone-video-push-result.md) to query the result |

## Related Commands

- [`video-push-stop`](geelark-phone-video-push-stop.md) — stop the ongoing push
- [`video-push-result`](geelark-phone-video-push-result.md) — query push task result
