# browser list

List browser environments with optional filters. Supports pagination and multiple filter criteria.

## Key Flags

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number, min 1 (default 1) |
| `--page-size <n>` | Page size, 1-100 (default 10) |
| `--ids <csv>` | Filter by browser IDs (max 100, ignores page/pageSize) |
| `--serial-name <text>` | Filter by browser name |
| `--remark <text>` | Filter by remark |
| `--group-name <text>` | Filter by group name |
| `--tags <csv>` | Filter by tag names |

## Examples

```bash
# List all browsers
geelark-cli browser list --page 1 --page-size 10

# Filter by name
geelark-cli browser list --serial-name "myBrowser"

# Filter by IDs
geelark-cli browser list --ids "id1,id2"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `pageSize` | integer | Page size |
| `items[]` | array | Browser list |
| `items[].id` | string | Browser ID |
| `items[].serialName` | string | Browser name |
| `items[].serialNo` | string | Browser number |
| `items[].remark` | string | Remark |
| `items[].group` | object | Group info |
| `items[].group.id` | string | Group ID |
| `items[].group.name` | string | Group name |
| `items[].group.remark` | string | Group remark |
| `items[].tags[]` | array | Tag list |
| `items[].tags[].name` | string | Tag name |
| `items[].extGroup` | string | Extension category ID (empty = team extensions). Query available IDs via [`ext-group-list`](geelark-browser-ext-group-list.md). |
| `items[].proxy` | object | Proxy info |
| `items[].proxy.type` | string | Proxy type |
| `items[].proxy.server` | string | Proxy server |
| `items[].proxy.port` | integer | Proxy port |
| `items[].proxy.username` | string | Proxy username |
| `items[].proxy.password` | string | Proxy password |
| `items[].accountInfo` | object | Account info |
| `items[].accountInfo.url` | string | Platform URL |
| `items[].accountInfo.userName` | string | Platform username |
| `items[].accountInfo.passWord` | string | Platform password |
| `items[].accountInfo.totpSecret` | string | 2FA key |
| `items[].accountInfo.openLastPage` | integer | Restore last visit: 1=Yes, 2=No |
| `items[].accountInfo.openSpecPage` | integer | Open specified URL: 1=Yes, 2=No |
| `items[].accountInfo.openSiteUrl` | bool | Open platform page |
| `items[].accountInfo.autoOpenUrls` | array[string] | Specified URLs |
| `items[].simulateInfo` | object | Simulation info |
| `items[].simulateInfo.os` | integer | OS: 1=Win, 2=Mac |
| `items[].simulateInfo.vendor` | integer | Browser type: 1=Chrome |
| `items[].simulateInfo.mixtureKey` | string | Fingerprint algorithm ID |
| `items[].simulateInfo.ua` | string | User agent |
| `items[].simulateInfo.uaVersion` | string | Browser version |
| `items[].simulateInfo.timeZone` | object | Time zone |
| `items[].simulateInfo.webRtc` | object | WebRTC |
| `items[].simulateInfo.geoLocation` | object | Geolocation |
| `items[].simulateInfo.language` | object | Language |
| `items[].simulateInfo.resolution` | object | Resolution |
| `items[].simulateInfo.font` | object | Font |
| `items[].simulateInfo.canvas` | object | Canvas |
| `items[].simulateInfo.webglImage` | object | WebGL Image |
| `items[].simulateInfo.webglMetadata` | object | WebGL Metadata |
| `items[].simulateInfo.hardware` | object | Hardware acceleration |
| `items[].simulateInfo.audioContext` | object | AudioContext |
| `items[].simulateInfo.mediaDevice` | object | Media device |
| `items[].simulateInfo.clientRects` | object | ClientRects |
| `items[].simulateInfo.speechVoise` | object | SpeechVoices |
| `items[].simulateInfo.hardwareConcurrency` | integer | Hardware concurrency |
| `items[].simulateInfo.memoryDevice` | integer | Device memory |
| `items[].simulateInfo.doNotTrack` | integer | Do Not Track: 0=Default, 1=On, 2=Off |
| `items[].simulateInfo.bluetooth` | object | Bluetooth |
| `items[].simulateInfo.battery` | object | Battery |
| `items[].simulateInfo.portScanProtection` | object | Port scan protection |
