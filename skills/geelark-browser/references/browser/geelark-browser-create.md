# browser create

Create a new browser environment with full configuration via JSON. Supports account, proxy, fingerprint, simulation settings, etc.

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON object with browser creation parameters (required) |

### JSON Parameters (in --data)

| Parameter | Required | Type | Description |
|-----------|----------|------|-------------|
| `serialName` | Yes | string | Browser name, max 100 chars |
| `browserOs` | Yes | integer | OS: 1=Win, 2=Mac |
| `groupId` | No | string | Group ID |
| `tagIds` | No | array[string] | Tag IDs |
| `remark` | No | string | Remark, max 1500 chars |
| `cookie` | No | string | Cookie (JSON or Netscape format) |
| `accountPlatform` | No | string | Platform URL (only tiktok.com, facebook.com) |
| `accountUsername` | No | string | Platform username |
| `accountPassword` | No | string | Platform password |
| `accountTOTPSecret` | No | string | 2FA key |
| `openLastPage` | No | integer | Restore last visit: 1=Yes, 2=No |
| `openSpecPage` | No | integer | Open specified URL: 1=Yes, 2=No |
| `openTabs` | No | string | Specified URLs, separated by `;` |
| `browserUa` | No | string | UserAgent (random if empty) |
| `simulateConfig` | No | object | Simulation configuration |
| `proxyId` | No | string | Proxy ID (preferred) |
| `proxyConfig` | No | object | Proxy configuration |
| `browserStartArg` | No | string | Startup parameters |
| `browserKernelVer` | No | string | Browser kernel version: 134,138,142,143,144,145,146,147,148,149,150,auto (default auto) |
| `extGroup` | No | string | Extension category ID. If omitted or empty, the team extensions will be used. Query available IDs via [`ext-group-list`](geelark-browser-ext-group-list.md). |

### SimulateConfig Fields

| Parameter | Required | Type | Description |
|-----------|----------|------|-------------|
| `timeZone.switcher` | Yes | integer | 1=IP-based, 2=Custom, 3=Local |
| `timeZone.value` | Yes | string | Custom timezone value |
| `language.switcher` | Yes | integer | 1=IP-based, 2=Custom |
| `language.value` | Yes | string | Custom language |
| `resolution.switcher` | Yes | integer | 1=Random, 2=Custom, 3=Default |
| `resolution.value` | Yes | string | Custom resolution |
| `webRtc.switcher` | Yes | integer | 1=Privacy, 2=Replace, 3=Real, 4=Disable |
| `geoLocation.switcher` | Yes | integer | 1=Ask, 2=Disable, 3=Allow |
| `geoLocation.baseOnIp` | Yes | bool | Match based on IP |
| `geoLocation.longitude` | Yes | integer | Longitude |
| `geoLocation.latitude` | Yes | integer | Latitude |
| `geoLocation.accuracy` | Yes | integer | Accuracy (meters) |
| `canvas.switcher` | Yes | integer | 1=Noise, 2=Real |
| `webglImage.switcher` | Yes | integer | 1=Noise, 2=Real |
| `hardware.switcher` | Yes | integer | 1=Default, 2=Enabled, 3=Disabled |
| `audioContext.switcher` | Yes | integer | 1=Noise, 2=Disabled |
| `mediaDevice.switcher` | Yes | integer | 1=Noise, 2=Disabled |
| `clientRects.switcher` | Yes | integer | 1=Noise, 2=Disabled |
| `speechVoise.switcher` | Yes | integer | 1=Noise, 2=Off |
| `hardwareConcurrency` | Yes | integer | Hardware concurrency |
| `memoryDevice` | Yes | integer | Device memory |
| `doNotTrack` | Yes | integer | 0=Default, 1=On, 2=Off |
| `bluetooth.switcher` | Yes | integer | 1=Privacy, 2=True |
| `battery.switcher` | Yes | integer | 1=Privacy, 2=True |
| `portScanProtection.switcher` | Yes | integer | 1=Enable, 2=Disable |
| `portScanProtection.value` | Yes | string | Allowed ports, comma-separated |

### ProxyConfig Fields

| Parameter | Required | Type | Description |
|-----------|----------|------|-------------|
| `typeId` | Yes | integer | Proxy type: -1=Direct, 1=socks5, 2=http, 3=https, 20=IPIDEA, 21=IPHTML, 22=kookeey, 23=Lumatuo |
| `server` | No | string | Proxy host |
| `port` | No | integer | Proxy port |
| `username` | No | string | Proxy username |
| `password` | No | string | Proxy password |
| `country` | No | string | Dynamic proxy country |
| `region` | No | string | Dynamic proxy region |
| `city` | No | string | Dynamic proxy city |
| `useProxyCfg` | No | bool | Use configured dynamic proxy |
| `protocol` | No | integer | Dynamic proxy protocol: 1=socks5, 2=http |

## Examples

```bash
# Create with basic config
geelark-cli browser create --data '{"serialName":"myBrowser","browserOs":1}'

# Create with account and proxy
geelark-cli browser create --data '{"serialName":"test","browserOs":2,"accountPlatform":"https://www.tiktok.com/","accountUsername":"user","accountPassword":"pass","proxyConfig":{"typeId":-1}}'

# Create with full simulation config
geelark-cli browser create --data '{"serialName":"full","browserOs":1,"simulateConfig":{"timeZone":{"switcher":2,"value":"GMT-12:00 Etc/GMT+12"},"language":{"switcher":2,"value":"Albanian"},"resolution":{"switcher":2,"value":"750*1334"},"webRtc":{"switcher":1},"geoLocation":{"switcher":1,"baseOnIp":true},"canvas":{"switcher":1},"webglImage":{"switcher":1},"hardware":{"switcher":1},"audioContext":{"switcher":1},"mediaDevice":{"switcher":1},"clientRects":{"switcher":1},"speechVoise":{"switcher":1},"hardwareConcurrency":26,"memoryDevice":8,"doNotTrack":2,"bluetooth":{"switcher":1},"battery":{"switcher":1},"portScanProtection":{"switcher":1,"value":"80"}}}'
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Browser environment ID |

## Error Codes

| Code | Description |
|------|-------------|
| 44002 | Plan environment limit reached |
| 44003 | User environment limit reached |
| 44004 | Daily creation limit reached |
| 45006 | Incorrect proxy information |
| 45003 | Proxy not allowed |
| 45004 | Proxy verification failed |
| 43028 | Sub-user lacks environment group permissions |
