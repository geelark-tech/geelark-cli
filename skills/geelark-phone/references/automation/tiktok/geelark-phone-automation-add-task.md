# phone automation add-task

Add a TikTok automation task (publish video, warmup, or publish image set). For warmup tasks, call directly. For video/image tasks, upload materials first via `file upload-temp`.

## Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON object with task parameters (required). Structure varies by task type |

## Top-Level Parameters

These fields go in the root of the `--data` JSON object:

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `taskType` | Yes | integer | Task type: `1` = Publish video, `2` = Warmup, `3` = Publish image set |
| `planName` | No | string | Task plan name; auto-generated if not provided |
| `remark` | No | string | Remarks, up to 200 characters |
| `list` | Yes | array | Task parameter array; max 100 items per call |

## taskType=1: Publish Video

`list[]` item fields:

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `scheduleAt` | Yes | integer | Scheduled time, second-level timestamp. If less than current time, defaults to now |
| `envId` | Yes | string | Cloud phone ID |
| `video` | Yes | string | Video URL. |
| `videoDesc` | No | string | Video description, max 4000 characters |
| `productId` | No | string | Product ID for shopping link |
| `productTitle` | No | string | Product display title |
| `refVideoId` | No | string | Similar/reference video ID |
| `maxTryTimes` | No | integer | Max auto-retry count, 0-3, default 3 |
| `timeoutMin` | No | integer | Timeout in minutes, 30-80, default 80 |
| `sameVideoVolume` | No | integer | Same video volume, 0-100 |
| `sourceVideoVolume` | No | integer | Original video volume, 0-100 |
| `markAI` | No | bool | Label AI-generated content, default false |
| `cover` | No | string | Video cover image URL. |
| `needShareLink` | No | bool | Whether to retrieve sharing link, default false |
| `randomUse` | No | bool | Whether to browse randomly before posting, default false |
| `browseTime` | No | integer | Random browsing time in minutes, required when `randomUse` is true, 1-20 |
| `gameName` | No | string | Game name, max 256 chars |
| `location` | No | string | Location, max 500 chars |
| `index` | No | integer | Product selection index when multiple products share the same name, 1-100; 0 or unset = first product |

Example:

```json
{
  "planName": "testAdd",
  "taskType": 1,
  "list": [{
    "scheduleAt": 1718744459,
    "envId": "123456654321",
    "video": "https://demo.geelark.com/open-upload/DhRP36s3.mp4",
    "videoDesc": "My video description",
    "maxTryTimes": 3,
    "timeoutMin": 80,
    "needShareLink": true,
    "gameName": "game",
    "location": "1600 Pennsylvania Avenue NW, Washington, DC 20500, USA"
  }]
}
```

## taskType=2: Warmup

`list[]` item fields:

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `scheduleAt` | Yes | integer | Scheduled time, second-level timestamp. If less than current time, defaults to now |
| `envId` | Yes | string | Cloud phone ID |
| `action` | Yes | string | Warmup action: `search profile`, `search video`, or `browse video` |
| `keywords` | No | array[string] | Search keywords; required for `search profile` / `search video`, optional for `browse video` |
| `duration` | Yes | integer | Browsing duration in minutes |

Example:

```json
{
  "planName": "testAdd",
  "taskType": 2,
  "list": [{
    "scheduleAt": 1718744459,
    "envId": "123456654321",
    "action": "search video",
    "keywords": ["hi"],
    "duration": 10
  }]
}
```

## taskType=3: Publish Image Set

`list[]` item fields:

| Field | Required | Type | Description |
|-------|----------|------|-------------|
| `scheduleAt` | Yes | integer | Scheduled time, second-level timestamp. If less than current time, defaults to now |
| `envId` | Yes | string | Cloud phone ID |
| `images` | Yes | array[string] | Image URLs. |
| `videoDesc` | No | string | Description, max 4000 characters |
| `videoId` | No | string | Same/reference video ID |
| `videoTitle` | No | string | Gallery title, max 90 characters |
| `productId` | No | string | Product ID for shopping link |
| `productTitle` | No | string | Product display title |
| `maxTryTimes` | No | integer | Max auto-retry count, 0-3, default 3 |
| `timeoutMin` | No | integer | Timeout in minutes, 30-80, default 80 |
| `sameVideoVolume` | No | integer | Same video volume, 0-100 |
| `markAI` | No | bool | Label AI-generated content, default false |
| `needShareLink` | No | bool | Whether to retrieve sharing link, default false |
| `randomUse` | No | bool | Whether to browse randomly before posting, default false |
| `browseTime` | No | integer | Random browsing time in minutes, required when `randomUse` is true, 1-20 |
| `randomBgm` | No | bool | Randomly match background music, default false |
| `index` | No | integer | Product selection index when multiple products share the same name, 1-100; 0 or unset = first product |

Example:

```json
{
  "planName": "testAdd",
  "taskType": 3,
  "list": [{
    "scheduleAt": 1718744459,
    "envId": "123456654321",
    "images": ["https://demo.geelark.com/open-upload/img1.jpg", "https://demo.geelark.com/open-upload/img2.jpg"],
    "videoDesc": "My image set description",
    "videoTitle": "My Gallery",
    "randomBgm": true
  }]
}
```

## CLI Example

```bash
geelark-cli phone automation add-task --data '{"taskType":2,"list":[{"scheduleAt":1718744459,"envId":"123456654321","action":"browse video","duration":10}]}'
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `taskIds[]` | array[string] | Array of created task IDs |

### Error Codes

| Code | Description |
|------|-------------|
| 41000 | Insufficient task credits |
| 41001 | Balance not enough |
| 43004 | Cloud phone has expired |
| 43018 | Monthly cloud phone not bound to monthly device |
| 48004 | App required by task does not meet requirements |
