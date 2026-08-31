---
name: geelark-phone
version: 1.0.0
description: "GeeLark cloud phone management: list, create, start, stop, delete, update, clone, screenshot, GPS, shell, ADB, file, library, app, analytics, webhook, OEM, automation, and more."
metadata:
  requires:
    bins: ["geelark-cli"]
  cliHelp: "geelark-cli phone --help"
---

# phone — Cloud Phone Management

**CRITICAL — MUST read [`../geelark-shared/SKILL.md`](../geelark-shared/SKILL.md) first for authentication and configuration handling.**

## Command Overview

```bash
geelark-cli phone <command> [flags]
```

### Phone Operations

| Command | Description |
|---------|-------------|
| `list` | List cloud phones with filters |
| `create` | Batch create cloud phones |
| `simple-create` | Quick create a single cloud phone |
| `start` | Batch start cloud phones |
| `stop` | Batch stop cloud phones |
| `restart` | Restart a cloud phone |
| `delete` | Batch delete cloud phones |
| `status` | Query cloud phone status |
| `update` | Update cloud phone info |
| `clone` | Clone a cloud phone |
| `new-one` | One-click new machine (reset identity) |
| `brand-list` | List supported phone brands/models |
| `brand-team-list` | List team-uploaded brands/models |
| `transfer` | Transfer phones to another account |
| `move-group` | Move phones to a group |
| `screenshot` | Take a screenshot |
| `screenshot-result` | Get screenshot task result |
| `get-gps` | Get GPS info |
| `set-gps` | Set GPS location |
| `set-root` | Enable/disable root |
| `get-device-id` | Get device ID (Android_ID) |
| `send-sms` | Send SMS to a cloud phone |
| `set-net-type` | Set network type (Wi-Fi/Mobile) |
| `hide-accessibility` | Hide accessibility from apps |
| `import-contacts` | Import contacts |
| `import-contacts-result` | Get import contacts result |
| `net-config-get` | Get network config (blacklist) |
| `net-config-set` | Set network config (blacklist) |

### Sub-Commands

| Command | Description |
|---------|-------------|
| `adb` | ADB management (get-info, set-status) |
| `shell` | Execute shell commands (exec) |
| `file` | File management (upload-temp, upload-to-phone, upload-status, keybox-upload, keybox-result) |
| `library` | Library/material management (material-create, material-search, material-delete, tag-create, tag-search, tag-delete, tag-set) |
| `app` | Application management (shop-list, install, uninstall, start, stop, list, installable-list, upload, upload-status, batch, team-app) |
| `analytics` | Analytics account management (accounts-list, add-accounts, simple-add-account, update-account, delete-account, data) |
| `webhook` | Webhook management (get, set) |
| `oem` | OEM/White Label customization |
| `automation` | Automation task management (57 commands across platforms) |

## Which Command to Use

| Want to | Command |
|---------|---------|
| View cloud phones | [`list`](references/phone/geelark-phone-list.md) |
| Create a phone | [`simple-create`](references/phone/geelark-phone-simple-create.md) (recommended) or [`create`](references/phone/geelark-phone-create.md) (batch) |
| Start/stop phones | [`start`](references/phone/geelark-phone-start.md) / [`stop`](references/phone/geelark-phone-stop.md) |
| Restart a phone | [`restart`](references/phone/geelark-phone-restart.md) |
| Delete phones | [`delete`](references/phone/geelark-phone-delete.md) |
| Check phone status | [`status`](references/phone/geelark-phone-status.md) |
| Modify phone info | [`update`](references/phone/geelark-phone-update.md) |
| Clone a phone | [`clone`](references/phone/geelark-phone-clone.md) |
| Reset phone identity | [`new-one`](references/phone/geelark-phone-new-one.md) |
| Take a screenshot | [`screenshot`](references/phone/geelark-phone-screenshot.md) + [`screenshot-result`](references/phone/geelark-phone-screenshot-result.md) |
| Live stream a video | [`video-push`](references/phone/geelark-phone-video-push.md) / [`video-push-stop`](references/phone/geelark-phone-video-push-stop.md) / [`video-push-result`](references/phone/geelark-phone-video-push-result.md) |
| Set GPS | [`set-gps`](references/phone/geelark-phone-set-gps.md) |
| Get GPS | [`get-gps`](references/phone/geelark-phone-get-gps.md) |
| Execute shell | [`shell exec`](references/shell/geelark-phone-shell-exec.md) |
| ADB connection | [`adb get-info`](references/adb/geelark-phone-adb-get-info.md) / [`adb set-status`](references/adb/geelark-phone-adb-set-status.md) |
| Upload file | [`file`](references/file/geelark-phone-file.md) |
| Manage apps | [`app`](references/app/geelark-phone-app.md) |
| Analytics accounts | [`analytics`](references/analytics/geelark-phone-analytics.md) |
| Set webhook | [`webhook set`](references/webhook/geelark-phone-webhook-set.md) |
| OEM customization | [`oem customization`](references/oem/geelark-phone-oem-customization.md) |
| Automation tasks | [`automation`](references/automation/SKILL.md) |

## Typical Scenarios

```bash
# List phones
geelark-cli phone list --page 1 --page-size 10

# Create a phone
geelark-cli phone simple-create --region sgp --mobile-type "Android 12" --profile-name "myPhone"

# Start phones
geelark-cli phone start --ids "id1,id2"

# Take screenshot
geelark-cli phone screenshot --id "phone_id"
geelark-cli phone screenshot-result --task-id "task_id"

# Set GPS
geelark-cli phone set-gps --data "[{\"id\":\"phone_id\",\"latitude\":1.302,\"longitude\":103.875}]"

# Execute shell
geelark-cli phone shell exec --id "phone_id" --cmd "pm list packages"
```

## Notes

- **Cloud phone must be started** before operations like screenshot, shell, ADB, send-sms, set-root
- **Cloud phone must be stopped** before delete
- **Batch limits**: start/stop max 200, delete max 100, create max 100
- **Android version support varies**: some features only support specific Android versions
- **Do not call update API while starting a cloud phone**

## Error Codes

| Code | Description |
|------|-------------|
| 42001 | Cloud phone does not exist |
| 42002 | Cloud phone is not running |
| 43004 | Cloud phone has expired |
| 43005 | Cloud phone is executing a task |
| 43006 | Cloud phone is occupied by remote |
| 43009 | Cloud phone is started (cannot delete) |
| 43010 | Cloud phone is starting (cannot delete) |
| 43015 | Cloud phone does not support this operation |
| 44001 | Pro plan required for batch creation |
| 44002 | Plan environment limit reached |
| 44004 | Daily creation limit reached |
| 45002 | Cloud phone proxy is unavailable |
| 45003 | Proxy region not allowed |
| 45004 | Proxy check failed |
| 47002 | Cloud phone resources insufficient |

Full error codes: https://open.geelark.com/api/cloud-phone-error-codes
