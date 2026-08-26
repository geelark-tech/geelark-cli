# geelark-phone automation

Cloud phone automation task management. Create, query, cancel, and retry RPA tasks across platforms (TikTok, Facebook, Instagram, YouTube, X, Reddit, Threads, Pinterest, Google, SHEIN) and custom task flows.

## Command Overview

| Command | Description |
|---------|-------------|
| `task-query` | Query tasks by IDs |
| `task-history` | Batch query task history (past 7 days) |
| `task-detail` | Query task detail with logs |
| `task-cancel` | Cancel waiting/in-progress tasks |
| `task-restart` | Retry failed/cancelled tasks |
| `add-custom-task` | Create custom automation task |
| `task-flow-list` | Query custom task flows |
| `task-flow-import` | Import/update custom task flow |
| `task-flow-export` | Export custom task flow |
| `add-task` | Add TikTok automation task (video/warmup/image set) |
| `tiktok-login` | TikTok auto login |
| `tiktok-star` | TikTok random star (like) |
| `tiktok-star-asia` | TikTok random star — Asia |
| `tiktok-comment` | TikTok AI random comment |
| `tiktok-comment-asia` | TikTok AI random comment — Asia |
| `tiktok-follow` | TikTok random follow |
| `tiktok-follow-asia` | TikTok random follow — Asia |
| `tiktok-edit-profile` | TikTok profile editing |
| `tiktok-message` | TikTok private message |
| `tiktok-message-asia` | TikTok private message — Asia |
| `tiktok-hide` | Hide all TikTok videos |
| `tiktok-hide-asia` | Hide all TikTok videos — Asia |
| `tiktok-delete` | Delete all TikTok videos |
| `tiktok-delete-asia` | Delete all TikTok videos — Asia |
| `facebook-login` | Facebook auto login |
| `facebook-publish` | Facebook publish video post |
| `facebook-auto-comment` | Facebook auto comment |
| `facebook-maintenance` | Facebook account maintenance |
| `facebook-pub-reels` | Facebook publish Reels |
| `facebook-reels-active` | Facebook Reels maintenance |
| `facebook-message` | Facebook private message |
| `instagram-login` | Instagram auto login |
| `instagram-warmup` | Instagram AI warmup |
| `instagram-pub-reels` | Instagram publish Reels video |
| `instagram-pub-reels-images` | Instagram publish Reels image set |
| `instagram-edit-profile` | Instagram profile editing |
| `instagram-message` | Instagram private message |
| `youtube-pub-video` | YouTube publish video |
| `youtube-pub-short` | YouTube publish Short |
| `youtube-maintenance` | YouTube account maintenance |
| `youtube-edit-profile` | Edit YouTube profile |
| `youtube-comment` | YouTube AI random comment |
| `youtube-message` | YouTube private message |
| `x-publish` | X (Twitter) publish tweet |
| `reddit-warmup` | Reddit AI warmup |
| `reddit-video` | Reddit publish video |
| `reddit-image` | Reddit publish image post |
| `threads-image` | Threads publish image post |
| `threads-video` | Threads publish video |
| `pinterest-image` | Pinterest publish image Pin |
| `pinterest-video` | Pinterest publish video Pin |
| `google-login` | Google auto login |
| `google-app-download` | Download apps on Google Play |
| `google-app-browser` | Open app on Google for browsing |
| `shein-login` | SHEIN auto login |
| `import-contacts` | Batch import contacts |
| `keybox-upload` | Upload Keybox |
| `file-upload` | Batch upload files |
| `multi-platform-video` | Multichannel video distribution |

## Common Flags

All automation task creation commands share these common optional flags:

| Flag | Description |
|------|-------------|
| `--name <text>` | Task name (max 128 chars) |
| `--remark <text>` | Remark (max 200 chars) |

All automation task creation commands require:

| Flag | Description |
|------|-------------|
| `--id <text>` | Cloud phone ID |
| `--schedule-at <n>` | Schedule time, second-level timestamp |

## Task Status

| Status | Value | Description |
|--------|-------|-------------|
| Waiting | 1 | Task is queued |
| In progress | 2 | Task is running |
| Completed | 3 | Task finished successfully |
| Failed | 4 | Task failed |
| Cancelled | 7 | Task was cancelled |

## References

### Task Management
- [task-query](task/geelark-phone-automation-task-query.md)
- [task-history](task/geelark-phone-automation-task-history.md)
- [task-detail](task/geelark-phone-automation-task-detail.md)
- [task-cancel](task/geelark-phone-automation-task-cancel.md)
- [task-restart](task/geelark-phone-automation-task-restart.md)

### Custom Task
- [add-custom-task](custom-task/geelark-phone-automation-add-custom-task.md)
- [task-flow-list](custom-task/geelark-phone-automation-task-flow-list.md)
- [task-flow-import](custom-task/geelark-phone-automation-task-flow-import.md)
- [task-flow-export](custom-task/geelark-phone-automation-task-flow-export.md)

### TikTok
- [add-task](tiktok/geelark-phone-automation-add-task.md)
- [tiktok-login](tiktok/geelark-phone-automation-tiktok-login.md)
- [tiktok-star](tiktok/geelark-phone-automation-tiktok-star.md) / [tiktok-star-asia](tiktok/geelark-phone-automation-tiktok-star-asia.md)
- [tiktok-comment](tiktok/geelark-phone-automation-tiktok-comment.md) / [tiktok-comment-asia](tiktok/geelark-phone-automation-tiktok-comment-asia.md)
- [tiktok-follow](tiktok/geelark-phone-automation-tiktok-follow.md) / [tiktok-follow-asia](tiktok/geelark-phone-automation-tiktok-follow-asia.md)
- [tiktok-edit-profile](tiktok/geelark-phone-automation-tiktok-edit-profile.md)
- [tiktok-message](tiktok/geelark-phone-automation-tiktok-message.md) / [tiktok-message-asia](tiktok/geelark-phone-automation-tiktok-message-asia.md)
- [tiktok-hide](tiktok/geelark-phone-automation-tiktok-hide.md) / [tiktok-hide-asia](tiktok/geelark-phone-automation-tiktok-hide-asia.md)
- [tiktok-delete](tiktok/geelark-phone-automation-tiktok-delete.md) / [tiktok-delete-asia](tiktok/geelark-phone-automation-tiktok-delete-asia.md)
- [tiktok-delete-comment](tiktok/geelark-phone-automation-tiktok-delete-comment.md) / [tiktok-delete-comment-asia](tiktok/geelark-phone-automation-tiktok-delete-comment-asia.md)

### Facebook
- [facebook-login](facebook/geelark-phone-automation-facebook-login.md)
- [facebook-publish](facebook/geelark-phone-automation-facebook-publish.md)
- [facebook-auto-comment](facebook/geelark-phone-automation-facebook-auto-comment.md)
- [facebook-maintenance](facebook/geelark-phone-automation-facebook-maintenance.md)
- [facebook-pub-reels](facebook/geelark-phone-automation-facebook-pub-reels.md)
- [facebook-reels-active](facebook/geelark-phone-automation-facebook-reels-active.md)
- [facebook-message](facebook/geelark-phone-automation-facebook-message.md)

### Instagram
- [instagram-login](instagram/geelark-phone-automation-instagram-login.md)
- [instagram-warmup](instagram/geelark-phone-automation-instagram-warmup.md)
- [instagram-pub-reels](instagram/geelark-phone-automation-instagram-pub-reels.md)
- [instagram-pub-reels-images](instagram/geelark-phone-automation-instagram-pub-reels-images.md)
- [instagram-edit-profile](instagram/geelark-phone-automation-instagram-edit-profile.md)
- [instagram-message](instagram/geelark-phone-automation-instagram-message.md)
- [instagram-follow-account](instagram/geelark-phone-automation-instagram-follow-account.md)
- [instagram-ai-comment](instagram/geelark-phone-automation-instagram-ai-comment.md)

### YouTube
- [youtube-pub-video](youtube/geelark-phone-automation-youtube-pub-video.md)
- [youtube-pub-short](youtube/geelark-phone-automation-youtube-pub-short.md)
- [youtube-maintenance](youtube/geelark-phone-automation-youtube-maintenance.md)
- [youtube-comment](youtube/geelark-phone-automation-youtube-comment.md)
- [youtube-message](youtube/geelark-phone-automation-youtube-message.md)

### X (Twitter)
- [x-publish](x/geelark-phone-automation-x-publish.md)

### Reddit
- [reddit-warmup](reddit/geelark-phone-automation-reddit-warmup.md)
- [reddit-video](reddit/geelark-phone-automation-reddit-video.md)
- [reddit-image](reddit/geelark-phone-automation-reddit-image.md)

### Threads
- [threads-image](threads/geelark-phone-automation-threads-image.md)
- [threads-video](threads/geelark-phone-automation-threads-video.md)

### Pinterest
- [pinterest-image](pinterest/geelark-phone-automation-pinterest-image.md)
- [pinterest-video](pinterest/geelark-phone-automation-pinterest-video.md)

### Google
- [google-login](google/geelark-phone-automation-google-login.md)
- [google-app-download](google/geelark-phone-automation-google-app-download.md)
- [google-app-browser](google/geelark-phone-automation-google-app-browser.md)

### SHEIN
- [shein-login](shein/geelark-phone-automation-shein-login.md)

### Other
- [import-contacts](other/geelark-phone-automation-import-contacts.md)
- [keybox-upload](other/geelark-phone-automation-keybox-upload.md)
- [file-upload](other/geelark-phone-automation-file-upload.md)
- [multi-platform-video](other/geelark-phone-automation-multi-platform-video.md)
