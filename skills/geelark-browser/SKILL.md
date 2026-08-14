---
name: geelark-browser
version: 1.0.0
description: "GeeLark browser management: list, create, start, stop, delete, edit, clone, clear cache, cookies, bookmarks, kernels, transfer, move group, automation tasks, and more."
metadata:
  requires:
    bins: ["geelark-cli"]
  cliHelp: "geelark-cli browser --help"
---

# browser — Browser Environment Management

**CRITICAL — MUST read [`../geelark-shared/SKILL.md`](../geelark-shared/SKILL.md) first for authentication and configuration handling.**

> **Note**: Browser commands use the **local API** (default endpoint: `http://localhost:40185`). The GeeLark client must be running and logged in. Configure endpoint via `geelark-cli config init --browser-base-url`.

## Command Overview

```bash
geelark-cli browser <command> [flags]
```

### Browser Operations

| Command | Description |
|---------|-------------|
| `api-status` | Check API interface availability |
| `list` | List browser environments with filters |
| `create` | Create a new browser (full JSON config) |
| `simple-create` | Quick create a single browser |
| `edit` | Edit browser configuration |
| `delete` | Delete browsers by IDs |
| `start` | Launch a browser |
| `stop` | Close a browser |
| `check-status` | Check browser startup status |
| `clone` | Clone a browser |
| `clear-cache` | Clear browser cache |
| `move-group` | Move browsers to a group |
| `transfer` | Transfer browsers to another account |
| `get-cookie` | Get browser cookies |
| `get-bookmark` | Get browser bookmarks |
| `set-bookmark` | Set browser bookmarks |
| `get-kernels` | List available browser kernels |
| `update-kernels` | Download/update a browser kernel |

### Sub-Commands

| Command | Description |
|---------|-------------|
| `automation` | Automation task management (20 commands across platforms) |

## Typical Scenarios

```bash
# Check API availability
geelark-cli browser api-status

# List browsers
geelark-cli browser list --page 1 --page-size 10

# Create a browser
geelark-cli browser simple-create --serial-name "myBrowser" --browser-os 1

# Start a browser
geelark-cli browser start --id "browser_id"

# Check browser status
geelark-cli browser check-status --id "browser_id"

# Stop a browser
geelark-cli browser stop --id "browser_id"

# Delete browsers
geelark-cli browser delete --ids "id1,id2"
```

## Notes

- **Local API only**: Browser commands communicate with the local GeeLark client, not the cloud API
- **Exception — `ext-group-list`**: This command uses the **cloud API** (`openapi.geelark.com`) to query team extension categories. It does not require the local client.
- **Client must be running**: Ensure the GeeLark client is open and logged in
- **Default endpoint**: `http://localhost:40185` (configurable via `config init --browser-base-url`)
- **Browser must be closed** before delete and clear-cache operations
- **Batch limits**: delete max 100, transfer max 200, clone max 100

## Error Codes

Full error codes: https://open.geelark.com/api/browser-error-codes
