# browser simple-create

Quick create a single browser with just the required fields. Use `create` for full configuration via JSON.

## Key Flags

| Flag | Description |
|------|-------------|
| `--serial-name <text>` | Browser name, max 100 chars (required) |
| `--browser-os <n>` | OS: 1=Win, 2=Mac (required) |
| `--browser-kernel-ver <text>` | Browser kernel version: 134,138,142,143,144,145,146,147,148,149,150,auto (default auto) |
| `--ext-group <text>` | Extension category ID (empty = team extensions). Query available IDs via [`ext-group-list`](geelark-browser-ext-group-list.md). |

## Examples

```bash
# Create a Windows browser
geelark-cli browser simple-create --serial-name "myBrowser" --browser-os 1

# Create a Mac browser
geelark-cli browser simple-create --serial-name "macBrowser" --browser-os 2

# Create with a specific kernel version
geelark-cli browser simple-create --serial-name "myBrowser" --browser-os 1 --browser-kernel-ver "149"

# Create with an extension category
geelark-cli browser simple-create --serial-name "myBrowser" --browser-os 1 --ext-group "497548067550006541"
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format). Response structure is identical to [`create`](geelark-browser-create.md).

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Browser environment ID |

## Error Codes

| Code | Description |
|------|-------------|
| 44002 | Plan environment limit reached |
| 44003 | User environment limit reached |
| 44004 | Daily creation limit reached |
