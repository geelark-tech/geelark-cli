# browser edit

Update an existing browser environment configuration via JSON.

## Key Flags

| Flag | Description |
|------|-------------|
| `--data <json>` | JSON object with browser update parameters (required) |

### JSON Parameters (in --data)

| Parameter | Required | Type | Description |
|-----------|----------|------|-------------|
| `id` | Yes | string | Browser ID |
| `serialName` | No | string | Browser name, max 100 chars |
| `groupId` | No | string | Group ID (empty = ungrouped) |
| `tagIds` | No | array[string] | Tag IDs (empty = no tags) |
| `remark` | No | string | Remark, max 1500 chars |
| `cookie` | No | string | Cookie (JSON or Netscape format) |
| `accountPlatform` | No | string | Platform URL |
| `accountUsername` | No | string | Platform username |
| `accountPassword` | No | string | Platform password |
| `accountTOTPSecret` | No | string | 2FA key |
| `openLastPage` | No | integer | Restore last visit: 1=Yes, 2=No |
| `openSpecPage` | No | integer | Open specified URL: 1=Yes, 2=No |
| `openTabs` | No | string | Specified URLs, separated by `;` |
| `browserOs` | No | integer | OS: 1=Win, 2=Mac |
| `browserUa` | No | string | UserAgent |
| `simulateConfig` | No | object | Simulation configuration (same as create) |
| `proxyId` | No | string | Proxy ID |
| `proxyConfig` | No | object | Proxy configuration (same as create) |
| `browserStartArg` | No | string | Startup parameters |
| `extGroup` | No | string | Extension category ID. If omitted or passed as an empty string, the team extensions will be used. |

> Omitted fields will not be changed. For `simulateConfig` and `proxyConfig` structure, see [create](geelark-browser-create.md).

## Examples

```bash
# Update browser name
geelark-cli browser edit --data '{"id":"browser_id","serialName":"newName"}'

# Update proxy
geelark-cli browser edit --data '{"id":"browser_id","proxyConfig":{"typeId":-1}}'

# Update account info
geelark-cli browser edit --data '{"id":"browser_id","accountPlatform":"https://www.tiktok.com/","accountUsername":"user","accountPassword":"pass"}'
```

## Response Fields

> The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

On success, `code` is `0` and `msg` is `"success"`. No additional `data` fields.

## Error Codes

| Code | Description |
|------|-------------|
| 45006 | Incorrect proxy information |
| 45003 | Proxy not allowed |
| 45004 | Proxy verification failed |
| 43028 | Sub-user lacks environment group permissions |
