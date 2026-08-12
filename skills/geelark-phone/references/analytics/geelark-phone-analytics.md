# phone analytics

Analytics account management for TikTok, YouTube, Instagram, and Reddit. Requires Pro plan for data queries.

## Command Overview

```bash
geelark-cli phone analytics <command> [flags]
```

| Command | Description |
|---------|-------------|
| `accounts-list` | List analytics accounts |
| `add-accounts` | Batch add analytics accounts |
| `simple-add-account` | Quick add a single analytics account |
| `update-account` | Update an analytics account |
| `delete-account` | Delete an analytics account |
| `data` | Get analytics account data |
| `tags-list` | List analytics tags |
| `tags-create` | Create an analytics tag |
| `tags-update` | Update an analytics tag |
| `tags-delete` | Delete analytics tags |

### Channel Values

| Value | Platform |
|-------|----------|
| 0 | TikTok |
| 1 | YouTube |
| 2 | Instagram |
| 4 | Reddit |

## accounts-list

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number |
| `--page-size <n>` | Page size, 1-100 |
| `--channel <n>` | Platform filter |
| `--account <text>` | Account name filter |
| `--user-account <text>` | Operator account email |

```bash
geelark-cli phone analytics accounts-list --page 1 --page-size 10
geelark-cli phone analytics accounts-list --channel 0 --account "tk_acc"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `items[]` | array | Account data |
| `items[].id` | string | Account ID |
| `items[].account` | string | Account name |
| `items[].channel` | integer | Platform: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit |
| `items[].remark` | string | Remark |
| `items[].operator` | string | Username of the last operator |
| `items[].created_time` | integer | Creation timestamp (seconds) |
| `items[].updated_time` | integer | Last update timestamp (seconds) |
| `items[].tags` | array[object] | Account tag list; empty array `[]` when no tags |
| `items[].tags[].id` | string | Tag ID |
| `items[].tags[].name` | string | Tag name |
| `items[].tags[].color` | integer | Tag color index |

## add-accounts

Batch add analytics accounts. Max 200 per request.

| Flag | Description |
|------|-------------|
| `--channel <n>` | Platform (required) |
| `--data <json>` | JSON array; each element contains `account` (required, max 64 chars), `remark` (optional), `tagIds` (optional, array of tag IDs, max 20) |

```bash
geelark-cli phone analytics add-accounts --channel 0 --data "[{\"account\":\"acc1\",\"remark\":\"my note\",\"tagIds\":[\"tag1\"]}]"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `bizCode` | integer | Business status: 0=all successful, 1=account limit exceeded, 2=partially successful with failures |
| `successCount` | integer | Number of successfully added accounts |
| `failCount` | integer | Number of failed additions |
| `repeatCount` | integer | Number of duplicate additions |

## simple-add-account

Quick add a single analytics account.

| Flag | Description |
|------|-------------|
| `--channel <n>` | Platform (required) |
| `--account <text>` | Account name, max 64 chars (required) |
| `--remark <text>` | Remark/note |
| `--tag-ids <csv>` | Comma-separated tag IDs, max 20 after deduplication |

```bash
geelark-cli phone analytics simple-add-account --channel 0 --account "myAccount"
geelark-cli phone analytics simple-add-account --channel 1 --account "ytAcc" --tag-ids "tag1,tag2"
```

### Response Fields

Same as [`add-accounts`](#add-accounts).

## update-account

| Flag | Description |
|------|-------------|
| `--id <text>` | Account ID (required) |
| `--account <text>` | New account name, max 64 chars |
| `--channel <n>` | New platform |
| `--remark <text>` | New remark |
| `--tag-ids <csv>` | Comma-separated tag IDs (max 20); pass empty string to clear tags |

```bash
geelark-cli phone analytics update-account --id "id" --account "newName"
geelark-cli phone analytics update-account --id "id" --tag-ids "tag1,tag2"
geelark-cli phone analytics update-account --id "id" --tag-ids ""
```

### Response Fields

Success only (no `data` field). Standard envelope with `code=0` indicates success.

## delete-account

| Flag | Description |
|------|-------------|
| `--channel <n>` | Platform (required) |
| `--account <text>` | Account name (required) |

```bash
geelark-cli phone analytics delete-account --channel 0 --account "myAccount"
```

### Response Fields

Success only (no `data` field). Standard envelope with `code=0` indicates success.

## data

Query analytics account data (play count, follower count, etc.). Requires Pro plan.

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number |
| `--page-size <n>` | Page size, 1-100 |
| `--channel <n>` | Platform filter |
| `--account <text>` | Account name filter |
| `--data-date <n>` | Search date timestamp (seconds) |
| `--created-id <text>` | User ID who added the account |

```bash
geelark-cli phone analytics data --page 1 --page-size 10 --channel 0
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `pageSize` | integer | Page size |
| `items[]` | array | Account data |
| `items[].id` | string | Account ID |
| `items[].channel` | integer | Platform |
| `items[].account` | string | Account name |
| `items[].playCount` | integer | Play count (-1 = not yet updated) |
| `items[].followerCount` | integer | Follower count (-1 = not yet updated) |
| `items[].diggCount` | integer | Like/digg count (-1 = not yet updated) |
| `items[].commentCount` | integer | Comment count (-1 = not yet updated) |
| `items[].collectCount` | integer | Collect count (-1 = not yet updated) |
| `items[].shareCount` | integer | Share count (-1 = not yet updated) |
| `items[].dataDate` | integer | Data date timestamp (-1 = not yet updated) |
| `items[].addAccDate` | integer | Account added timestamp (seconds) |
| `items[].remark` | string | Remark |
| `items[].createdId` | string | User ID who added the account |
| `items[].username` | string | Username who added the account |

### Error Codes

| Code | Description |
|------|-------------|
| 43002 | Please upgrade to Pro plan to use this feature |

## tags-list

List analytics tags with optional fuzzy search.

| Flag | Description |
|------|-------------|
| `--name <text>` | Tag name (fuzzy search) |
| `--page <n>` | Page number (default 1) |
| `--page-size <n>` | Page size, max 200 (default 200) |

```bash
geelark-cli phone analytics tags-list --page 1 --page-size 200
geelark-cli phone analytics tags-list --name "Important"
```

### Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `pageSize` | integer | Page size |
| `list[]` | array | Tag list |
| `list[].id` | string | Tag ID |
| `list[].name` | string | Tag name |
| `list[].color` | integer | Tag color value |

## tags-create

Create an analytics tag. Tag name must be unique within the team.

| Flag | Description |
|------|-------------|
| `--name <text>` | Tag name, max 100 chars, must be unique (required) |
| `--color <n>` | Tag color value (default 0) |

```bash
geelark-cli phone analytics tags-create --name "Important Account" --color 1
```

### Response Fields

Success only (no `data` field). Standard envelope with `code=0` indicates success.

### Error Codes

| Code | Description |
|------|-------------|
| 40004 | Invalid request parameters |
| 43021 | Tag with the same name already exists |

## tags-update

Update an analytics tag.

| Flag | Description |
|------|-------------|
| `--id <text>` | Tag ID (required) |
| `--name <text>` | Tag name, max 100 chars (required) |
| `--color <n>` | Tag color value (default 0) |

```bash
geelark-cli phone analytics tags-update --id "tag_id" --name "Core Account" --color 2
```

### Response Fields

Success only (no `data` field). Standard envelope with `code=0` indicates success.

### Error Codes

| Code | Description |
|------|-------------|
| 40004 | Invalid request parameters |
| 43021 | Tag with the same name already exists |
| 43022 | Tag not found |

## tags-delete

Delete one or more analytics tags. Associated account links are also removed.

| Flag | Description |
|------|-------------|
| `--ids <csv>` | Comma-separated tag IDs (required) |

```bash
geelark-cli phone analytics tags-delete --ids "tag1,tag2"
```

### Response Fields

Success only (no `data` field). Standard envelope with `code=0` indicates success.

### Error Codes

| Code | Description |
|------|-------------|
| 40004 | Invalid request parameters |
| 43022 | Tag not found |
