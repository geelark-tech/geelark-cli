# browser ext-group-list

List browser extension categories for the current team. Returns the available `extGroup` IDs that can be used when creating or editing a browser.

> **Note**: Unlike most browser commands, this one uses the **cloud API** (`openapi.geelark.com`), not the local browser API. It does not require the local GeeLark client to be running.

## Key Flags

| Flag | Description |
|------|-------------|
| `--page <n>` | Page number, min 1 (default 1) |
| `--page-size <n>` | Page size, 1-100 (default 10) |

## Examples

```bash
# List extension categories
geelark-cli browser ext-group-list --page 1 --page-size 10
```

## Response Fields

> Below is the `data` inner structure. The full response is wrapped in the standard envelope (`traceId` / `code` / `msg` / `data`), see [geelark-shared](../../../geelark-shared/SKILL.md#api-response-format).

| Field | Type | Description |
|-------|------|-------------|
| `total` | integer | Total count |
| `page` | integer | Current page |
| `pageSize` | integer | Page size |
| `items[]` | array | Extension category list |
| `items[].id` | string | Extension category ID — pass this as `extGroup` in [create](geelark-browser-create.md) / [simple-create](geelark-browser-simple-create.md) / [edit](geelark-browser-edit.md) |
| `items[].name` | string | Extension category name |

## Usage with Other Commands

The returned `id` values are used as the `extGroup` parameter:

- [`create`](geelark-browser-create.md) — `--data '{"extGroup":"<id>"}'`
- [`simple-create`](geelark-browser-simple-create.md) — `--ext-group "<id>"`
- [`edit`](geelark-browser-edit.md) — `--data '{"extGroup":"<id>"}'`

If `extGroup` is omitted or empty, the team extensions are used.
