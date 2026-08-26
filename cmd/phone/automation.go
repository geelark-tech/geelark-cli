package phone

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func newAutomationCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "automation",
		Short: "Cloud phone automation tasks",
		Long: `Manage cloud phone automation tasks via the remote API.
Includes task management, custom tasks, and platform-specific automation
(TikTok, Facebook, Instagram, YouTube, X/Twitter, Reddit, Threads, Pinterest, Google, Shein, and more).`,
	}

	cmd.AddCommand(newTaskQueryCmd(newClient))
	cmd.AddCommand(newTaskHistoryCmd(newClient))
	cmd.AddCommand(newTaskDetailCmd(newClient))
	cmd.AddCommand(newTaskCancelCmd(newClient))
	cmd.AddCommand(newTaskRestartCmd(newClient))
	cmd.AddCommand(newAddTaskCmd(newClient))
	cmd.AddCommand(newCustomTaskAddCmd(newClient))
	cmd.AddCommand(newTaskFlowListCmd(newClient))
	cmd.AddCommand(newTaskFlowImportCmd(newClient))
	cmd.AddCommand(newTaskFlowExportCmd(newClient))
	cmd.AddCommand(newTikTokLoginCmd(newClient))
	cmd.AddCommand(newTikTokStarCmd(newClient))
	cmd.AddCommand(newTikTokStarAsiaCmd(newClient))
	cmd.AddCommand(newTikTokCommentCmd(newClient))
	cmd.AddCommand(newTikTokCommentAsiaCmd(newClient))
	cmd.AddCommand(newTikTokFollowCmd(newClient))
	cmd.AddCommand(newTikTokFollowAsiaCmd(newClient))
	cmd.AddCommand(newTikTokEditProfileCmd(newClient))
	cmd.AddCommand(newTikTokMessageCmd(newClient))
	cmd.AddCommand(newTikTokMessageAsiaCmd(newClient))
	cmd.AddCommand(newTikTokHideCmd(newClient))
	cmd.AddCommand(newTikTokHideAsiaCmd(newClient))
	cmd.AddCommand(newTikTokDeleteCmd(newClient))
	cmd.AddCommand(newTikTokDeleteAsiaCmd(newClient))
	cmd.AddCommand(newTikTokDeleteCommentCmd(newClient))
	cmd.AddCommand(newTikTokDeleteCommentAsiaCmd(newClient))
	cmd.AddCommand(newFacebookLoginCmd(newClient))
	cmd.AddCommand(newFacebookPublishCmd(newClient))
	cmd.AddCommand(newFacebookAutoCommentCmd(newClient))
	cmd.AddCommand(newFacebookMaintenanceCmd(newClient))
	cmd.AddCommand(newFacebookPubReelsCmd(newClient))
	cmd.AddCommand(newFacebookReelsActiveCmd(newClient))
	cmd.AddCommand(newFacebookMessageCmd(newClient))
	cmd.AddCommand(newInstagramLoginCmd(newClient))
	cmd.AddCommand(newInstagramWarmupCmd(newClient))
	cmd.AddCommand(newInstagramPubReelsCmd(newClient))
	cmd.AddCommand(newInstagramPubReelsImagesCmd(newClient))
	cmd.AddCommand(newInstagramEditProfileCmd(newClient))
	cmd.AddCommand(newInstagramMessageCmd(newClient))
	cmd.AddCommand(newInstagramFollowAccountCmd(newClient))
	cmd.AddCommand(newInstagramAiCommentCmd(newClient))
	cmd.AddCommand(newYouTubePubVideoCmd(newClient))
	cmd.AddCommand(newYouTubePubShortCmd(newClient))
	cmd.AddCommand(newYouTubeMaintenanceCmd(newClient))
	cmd.AddCommand(newYouTubeEditProfileCmd(newClient))
	cmd.AddCommand(newXPublishCmd(newClient))
	cmd.AddCommand(newRedditWarmupCmd(newClient))
	cmd.AddCommand(newRedditVideoCmd(newClient))
	cmd.AddCommand(newRedditImageCmd(newClient))
	cmd.AddCommand(newThreadsImageCmd(newClient))
	cmd.AddCommand(newThreadsVideoCmd(newClient))
	cmd.AddCommand(newPinterestImageCmd(newClient))
	cmd.AddCommand(newPinterestVideoCmd(newClient))
	cmd.AddCommand(newGoogleLoginCmd(newClient))
	cmd.AddCommand(newGoogleAppDownloadCmd(newClient))
	cmd.AddCommand(newGoogleAppBrowserCmd(newClient))
	cmd.AddCommand(newSheinLoginCmd(newClient))
	cmd.AddCommand(newImportContactsTaskCmd(newClient))
	cmd.AddCommand(newKeyboxUploadCmd(newClient))
	cmd.AddCommand(newFileUploadCmd(newClient))
	cmd.AddCommand(newMultiPlatformVideoCmd(newClient))

	return cmd
}

func newTaskQueryCmd(newClient clientFactory) *cobra.Command {
	var ids string
	cmd := &cobra.Command{
		Use:     "task-query",
		Short:   "Query cloud phone tasks by IDs",
		Example: `  geelark-cli phone automation task-query --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostAndPrint("/open/v1/task/query", map[string]interface{}{"ids": strings.Split(ids, ",")})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated task IDs (required)")
	_ = cmd.MarkFlagRequired("ids")
	return cmd
}

func newTaskHistoryCmd(newClient clientFactory) *cobra.Command {
	var size int
	var lastId, ids string
	cmd := &cobra.Command{
		Use:     "task-history",
		Short:   "Batch query task history",
		Long:    "Query all tasks scheduled within the past 7 days.",
		Example: `  geelark-cli phone automation task-history --size 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{}
			if size > 0 {
				body["size"] = size
			}
			if lastId != "" {
				body["lastId"] = lastId
			}
			if ids != "" {
				body["ids"] = strings.Split(ids, ",")
			}
			result, err := c.PostAndPrint("/open/v1/task/historyRecords", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().IntVar(&size, "size", 0, "Number of records per page (max 100)")
	cmd.Flags().StringVar(&lastId, "last-id", "", "Last item ID from previous page for pagination")
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated task IDs (max 100)")
	return cmd
}

func newTaskDetailCmd(newClient clientFactory) *cobra.Command {
	var id, searchAfterJSON string
	cmd := &cobra.Command{
		Use:     "task-detail",
		Short:   "Query cloud phone task detail",
		Example: `  geelark-cli phone automation task-detail --id "1234567898"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id}
			if searchAfterJSON != "" {
				var sa interface{}
				if err := json.Unmarshal([]byte(searchAfterJSON), &sa); err != nil {
					return fmt.Errorf("invalid --search-after JSON: %w", err)
				}
				body["searchAfter"] = sa
			}
			result, err := c.PostAndPrint("/open/v1/task/detail", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Task ID (required)")
	cmd.Flags().StringVar(&searchAfterJSON, "search-after", "", "Log pagination parameter as JSON array")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func newTaskCancelCmd(newClient clientFactory) *cobra.Command {
	var ids string
	cmd := &cobra.Command{
		Use:     "task-cancel",
		Short:   "Cancel cloud phone tasks",
		Example: `  geelark-cli phone automation task-cancel --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostAndPrint("/open/v1/task/cancel", map[string]interface{}{"ids": strings.Split(ids, ",")})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated task IDs (required)")
	_ = cmd.MarkFlagRequired("ids")
	return cmd
}

func newTaskRestartCmd(newClient clientFactory) *cobra.Command {
	var ids string
	cmd := &cobra.Command{
		Use:     "task-restart",
		Short:   "Retry cloud phone tasks",
		Example: `  geelark-cli phone automation task-restart --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostAndPrint("/open/v1/task/restart", map[string]interface{}{"ids": strings.Split(ids, ",")})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated task IDs (required)")
	_ = cmd.MarkFlagRequired("ids")
	return cmd
}

func newAddTaskCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string
	cmd := &cobra.Command{
		Use:   "add-task",
		Short: "Add TikTok video/image/warmup task",
		Long: `Add a TikTok automation task (publish video, warmup, or publish image set).
Use --data to pass the full task configuration as JSON.`,
		Example: `  geelark-cli phone automation add-task --data "{\"taskType\":2,\"list\":[{\"scheduleAt\":1718744459,\"envId\":\"123456654321\",\"action\":\"browse video\",\"duration\":10}]}"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			var body interface{}
			if err := json.Unmarshal([]byte(dataJSON), &body); err != nil {
				return fmt.Errorf("invalid --data JSON: %w", err)
			}
			result, err := c.PostAndPrint("/open/v1/task/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&dataJSON, "data", "", "JSON object with task parameters (required)")
	_ = cmd.MarkFlagRequired("data")
	return cmd
}

func newCustomTaskAddCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	var flowID, paramMapJSON string
	cmd := &cobra.Command{
		Use:   "add-custom-task",
		Short: "Create a custom automation task",
		Long: `Create a custom cloud phone automation task using a task flow ID.
First call 'task-flow-list' to get available task flows, then create a task with the flow ID.`,
		Example: `  geelark-cli phone automation add-custom-task --id "557536075321468390" --schedule-at 1741846843 --flow-id "562316072435344885"
  geelark-cli phone automation add-custom-task --id "557536075321468390" --schedule-at 1741846843 --flow-id "562316072435344885" --param-map "{\"Title\":\"video\"}"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "flowId": flowID}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if paramMapJSON != "" {
				var pm interface{}
				if err := json.Unmarshal([]byte(paramMapJSON), &pm); err != nil {
					return fmt.Errorf("invalid --param-map JSON: %w", err)
				}
				body["paramMap"] = pm
			}
			result, err := c.PostAndPrint("/open/v1/task/rpa/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&flowID, "flow-id", "", "Task flow ID (required)")
	cmd.Flags().StringVar(&paramMapJSON, "param-map", "", "Task flow parameters as JSON string")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("flow-id")
	return cmd
}

func newTaskFlowListCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int
	cmd := &cobra.Command{
		Use:     "task-flow-list",
		Short:   "Query custom task flows",
		Example: `  geelark-cli phone automation task-flow-list --page 1 --page-size 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostAndPrint("/open/v1/task/flow/list", map[string]interface{}{"page": page, "pageSize": pageSize})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().IntVar(&page, "page", 1, "Page number (required)")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Page size, max 100 (required)")
	return cmd
}

func newTaskFlowImportCmd(newClient clientFactory) *cobra.Command {
	var id, gal string
	cmd := &cobra.Command{
		Use:     "task-flow-import",
		Short:   "Import or update a custom task flow",
		Example: `  geelark-cli phone automation task-flow-import --gal "{\"content\":{...}}"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"gal": gal}
			if id != "" {
				body["id"] = id
			}
			result, err := c.PostAndPrint("/open/v1/task/flow/import", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Custom task flow ID (for update)")
	cmd.Flags().StringVar(&gal, "gal", "", "Custom task flow data (required)")
	_ = cmd.MarkFlagRequired("gal")
	return cmd
}

func newTaskFlowExportCmd(newClient clientFactory) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:     "task-flow-export",
		Short:   "Export a custom task flow",
		Example: `  geelark-cli phone automation task-flow-export --id "612345671223083526"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostAndPrint("/open/v1/task/flow/export", map[string]interface{}{"id": id})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Custom task flow ID (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func newTikTokLoginCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, account, password, twoFAKey string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:   "tiktok-login",
		Short: "TikTok account login",
		Example: `  geelark-cli phone automation tiktok-login --id "557536075321468390" --schedule-at 1741846843 --account "test@gmail.com" --password "123456"
  geelark-cli phone automation tiktok-login --id "557536075321468390" --schedule-at 1741846843 --account "test@gmail.com" --password "123456" --two-fa-key "2FAKEY"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "account": account, "password": password}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if twoFAKey != "" {
				body["twoFAKey"] = twoFAKey
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokLogin", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&account, "account", "", "Account (max 64 chars, required)")
	cmd.Flags().StringVar(&password, "password", "", "Password (max 64 chars, required)")
	cmd.Flags().StringVar(&twoFAKey, "two-fa-key", "", "2FA Key")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("account")
	_ = cmd.MarkFlagRequired("password")
	return cmd
}

func newTikTokStarCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	var likeProbability int
	cmd := &cobra.Command{
		Use:     "tiktok-star",
		Short:   "TikTok random star (like)",
		Example: `  geelark-cli phone automation tiktok-star --id "557536075321468390" --schedule-at 1741846843 --like-probability 50`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if likeProbability > 0 {
				body["likeProbability"] = likeProbability
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokRandomStar", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&likeProbability, "like-probability", 0, "Probability of liking, 0-100, default 30")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newFacebookLoginCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, email, password string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "facebook-login",
		Short:   "Facebook auto login",
		Example: `  geelark-cli phone automation facebook-login --id "557536075321468390" --schedule-at 1741846843 --email "test@gmail.com" --password "123456"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "email": email, "password": password}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/faceBookLogin", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&email, "email", "", "Email (max 64 chars, required)")
	cmd.Flags().StringVar(&password, "password", "", "Password (max 64 chars, required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("email")
	_ = cmd.MarkFlagRequired("password")
	return cmd
}

func newFacebookPublishCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, title, video string
	var scheduleAt int64
	var needShareLink bool
	cmd := &cobra.Command{
		Use:     "facebook-publish",
		Short:   "Facebook post content",
		Example: `  geelark-cli phone automation facebook-publish --id "557536075321468390" --schedule-at 1741846843 --title "title" --video "https://material.geelark.com/a.mp4"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "title": title, "video": strings.Split(video, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if needShareLink {
				body["needShareLink"] = true
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/faceBookPublish", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&title, "title", "", "Title (max 200 chars, required)")
	cmd.Flags().StringVar(&video, "video", "", "Comma-separated video URLs, max 10 (required)")
	cmd.Flags().BoolVar(&needShareLink, "need-share-link", false, "Whether to retrieve the sharing link (default false)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("video")
	return cmd
}

func newFacebookAutoCommentCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, postAddress, comment, keyword string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "facebook-auto-comment",
		Short:   "Facebook auto comment",
		Example: `  geelark-cli phone automation facebook-auto-comment --id "557536075321468390" --schedule-at 1741846843 --post-address "https://abc.com" --comment "test1,test2" --keyword "k1,k2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "postAddress": postAddress, "comment": strings.Split(comment, ","), "keyword": strings.Split(keyword, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/faceBookAutoComment", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&postAddress, "post-address", "", "Post address (max 128 chars, required)")
	cmd.Flags().StringVar(&comment, "comment", "", "Comma-separated comments, max 10 (required)")
	cmd.Flags().StringVar(&keyword, "keyword", "", "Comma-separated keywords, max 10 (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("post-address")
	_ = cmd.MarkFlagRequired("comment")
	_ = cmd.MarkFlagRequired("keyword")
	return cmd
}

func newFacebookMaintenanceCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, keyword string
	var scheduleAt int64
	var browsePostsNum int
	cmd := &cobra.Command{
		Use:     "facebook-maintenance",
		Short:   "Facebook account maintenance",
		Example: `  geelark-cli phone automation facebook-maintenance --id "557536075321468390" --schedule-at 1741846843 --browse-posts-num 10 --keyword "k1,k2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "browsePostsNum": browsePostsNum, "keyword": strings.Split(keyword, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/faceBookActiveAccount", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&browsePostsNum, "browse-posts-num", 0, "Number of posts to browse, 1-100 (required)")
	cmd.Flags().StringVar(&keyword, "keyword", "", "Comma-separated keywords, max 10 (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("browse-posts-num")
	_ = cmd.MarkFlagRequired("keyword")
	return cmd
}

func newFacebookPubReelsCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, description, video, page string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "facebook-pub-reels",
		Short:   "Facebook publish Reels video",
		Example: `  geelark-cli phone automation facebook-pub-reels --id "557536075321468390" --schedule-at 1741846843 --description "desc" --video "https://material.geelark.com/a.mp4"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "description": description, "video": video}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if page != "" {
				body["page"] = page
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/faceBookPubReels", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&description, "description", "", "Caption (max 500 chars, required)")
	cmd.Flags().StringVar(&video, "video", "", "Video URL (required)")
	cmd.Flags().StringVar(&page, "page", "", "Page")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("description")
	_ = cmd.MarkFlagRequired("video")
	return cmd
}

func newFacebookReelsActiveCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, searchKeywords string
	var scheduleAt int64
	var videoNumber int
	cmd := &cobra.Command{
		Use:     "facebook-reels-active",
		Short:   "Facebook Reels maintenance",
		Example: `  geelark-cli phone automation facebook-reels-active --id "557536075321468390" --schedule-at 1741846843 --video-number 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "videoNumber": videoNumber}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if searchKeywords != "" {
				body["searchKeywords"] = strings.Split(searchKeywords, ",")
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/facebookReelsActive", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&videoNumber, "video-number", 0, "Estimated number of videos viewed (required)")
	cmd.Flags().StringVar(&searchKeywords, "search-keywords", "", "Comma-separated search keywords")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("video-number")
	return cmd
}

func newFacebookMessageCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, usernames, content string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "facebook-message",
		Short:   "Send private message on Facebook",
		Example: `  geelark-cli phone automation facebook-message --id "557536075321468390" --schedule-at 1741846843 --usernames "user1" --content "hello"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "usernames": strings.Split(usernames, ","), "content": content}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/faceBookMessage", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&usernames, "usernames", "", "Comma-separated usernames (required)")
	cmd.Flags().StringVar(&content, "content", "", "Message content, max 20000 chars (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("usernames")
	_ = cmd.MarkFlagRequired("content")
	return cmd
}

func newInstagramLoginCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, account, password string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "instagram-login",
		Short:   "Instagram auto login",
		Example: `  geelark-cli phone automation instagram-login --id "557536075321468390" --schedule-at 1741846843 --account "test@gmail.com" --password "123456"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "account": account, "password": password}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/instagramLogin", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&account, "account", "", "Account (max 64 chars, required)")
	cmd.Flags().StringVar(&password, "password", "", "Password (max 64 chars, required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("account")
	_ = cmd.MarkFlagRequired("password")
	return cmd
}

func newInstagramWarmupCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, keyword string
	var scheduleAt int64
	var browseVideo int
	cmd := &cobra.Command{
		Use:     "instagram-warmup",
		Short:   "Instagram AI account warmup",
		Example: `  geelark-cli phone automation instagram-warmup --id "557536075321468390" --schedule-at 1741846843 --browse-video 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if browseVideo > 0 {
				body["browseVideo"] = browseVideo
			}
			if keyword != "" {
				body["keyword"] = keyword
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/instagramWarmup", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&browseVideo, "browse-video", 0, "Number of videos viewed, 1-100")
	cmd.Flags().StringVar(&keyword, "keyword", "", "Search keyword")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newInstagramPubReelsCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, description, video, sameStyleUrl string
	var scheduleAt int64
	var sameStyleVoice, originalVoice int
	var aiTag, needShareLink bool
	cmd := &cobra.Command{
		Use:     "instagram-pub-reels",
		Short:   "Instagram publish Reels video",
		Example: `  geelark-cli phone automation instagram-pub-reels --id "557536075321468390" --schedule-at 1741846843 --description "desc" --video "https://material.geelark.com/a.mp4"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "description": description, "video": strings.Split(video, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if sameStyleUrl != "" {
				body["sameStyleUrl"] = sameStyleUrl
			}
			if sameStyleVoice > 0 {
				body["sameStyleVoice"] = sameStyleVoice
			}
			if originalVoice > 0 {
				body["originalVoice"] = originalVoice
			}
			if aiTag {
				body["aiTag"] = true
			}
			if needShareLink {
				body["needShareLink"] = true
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/instagramPubReels", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&description, "description", "", "Caption (max 2200 chars, required)")
	cmd.Flags().StringVar(&video, "video", "", "Comma-separated video URLs, max 10 (required)")
	cmd.Flags().StringVar(&sameStyleUrl, "same-style-url", "", "Same style URL")
	cmd.Flags().IntVar(&sameStyleVoice, "same-style-voice", 0, "Same style volume, 0-100")
	cmd.Flags().IntVar(&originalVoice, "original-voice", 0, "Original volume, 0-100")
	cmd.Flags().BoolVar(&aiTag, "ai-tag", false, "AI tag, defaults to false")
	cmd.Flags().BoolVar(&needShareLink, "need-share-link", false, "Whether to retrieve sharing link")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("description")
	_ = cmd.MarkFlagRequired("video")
	return cmd
}

func newInstagramPubReelsImagesCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, description, image, sameStyleUrl string
	var scheduleAt int64
	var aiTag, publishPost, needShareLink bool
	cmd := &cobra.Command{
		Use:     "instagram-pub-reels-images",
		Short:   "Instagram publish Reels image",
		Example: `  geelark-cli phone automation instagram-pub-reels-images --id "557536075321468390" --schedule-at 1741846843 --description "desc" --image "https://material.geelark.com/a.jpg"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "image": strings.Split(image, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if description != "" {
				body["description"] = description
			}
			if sameStyleUrl != "" {
				body["sameStyleUrl"] = sameStyleUrl
			}
			if aiTag {
				body["aiTag"] = true
			}
			if publishPost {
				body["publishPost"] = true
			}
			if needShareLink {
				body["needShareLink"] = true
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/instagramPubReelsImages", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&description, "description", "", "Caption (max 2200 chars)")
	cmd.Flags().StringVar(&image, "image", "", "Comma-separated image URLs, max 10 (required)")
	cmd.Flags().StringVar(&sameStyleUrl, "same-style-url", "", "Same style URL")
	cmd.Flags().BoolVar(&aiTag, "ai-tag", false, "AI tag, defaults to false")
	cmd.Flags().BoolVar(&publishPost, "publish-post", false, "Posting a POST request, defaults to false")
	cmd.Flags().BoolVar(&needShareLink, "need-share-link", false, "Whether to retrieve sharing link")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("image")
	return cmd
}

func newInstagramEditProfileCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	var profilePicture, nickname, username, biography, linkURL, linkTitle string
	cmd := &cobra.Command{
		Use:     "instagram-edit-profile",
		Short:   "Edit Instagram profile",
		Example: `  geelark-cli phone automation instagram-edit-profile --id "557536075321468390" --schedule-at 1741846843 --nickname "myName"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if profilePicture != "" {
				body["profilePicture"] = []string{profilePicture}
			}
			if nickname != "" {
				body["nickname"] = nickname
			}
			if username != "" {
				body["username"] = username
			}
			if biography != "" {
				body["biography"] = biography
			}
			if linkURL != "" {
				body["linkURL"] = linkURL
			}
			if linkTitle != "" {
				body["linkTitle"] = linkTitle
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/instagramEdit", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&profilePicture, "profile-picture", "", "Avatar URL")
	cmd.Flags().StringVar(&nickname, "nickname", "", "Nickname")
	cmd.Flags().StringVar(&username, "username", "", "Username")
	cmd.Flags().StringVar(&biography, "biography", "", "Biography")
	cmd.Flags().StringVar(&linkURL, "link-url", "", "Link URL")
	cmd.Flags().StringVar(&linkTitle, "link-title", "", "Link title")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newInstagramMessageCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, usernames, content string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "instagram-message",
		Short:   "Send private message on Instagram",
		Example: `  geelark-cli phone automation instagram-message --id "557536075321468390" --schedule-at 1741846843 --usernames "user1" --content "hello"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "usernames": strings.Split(usernames, ","), "content": content}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/instagramMessage", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&usernames, "usernames", "", "Comma-separated usernames (required)")
	cmd.Flags().StringVar(&content, "content", "", "Message content, max 1000 chars (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("usernames")
	_ = cmd.MarkFlagRequired("content")
	return cmd
}

func newInstagramFollowAccountCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, usernames string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "instagram-follow-account",
		Short:   "Follow Instagram accounts",
		Example: `  geelark-cli phone automation instagram-follow-account --id "557536075321468390" --schedule-at 1741846843 --usernames "ins1,ins2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "username": strings.Split(usernames, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/instagramFollowAccount", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&usernames, "usernames", "", "Comma-separated usernames, max 100, each max 30 chars (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("usernames")
	return cmd
}

func newInstagramAiCommentCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	var useAi bool
	var randomRate int
	cmd := &cobra.Command{
		Use:   "instagram-ai-comment",
		Short: "Instagram AI random comment",
		Example: `  geelark-cli phone automation instagram-ai-comment --id "557536075321468390" --schedule-at 1741846843 --random-rate 50
  geelark-cli phone automation instagram-ai-comment --id "557536075321468390" --schedule-at 1741846843 --use-ai --random-rate 30`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "randomRate": randomRate}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if cmd.Flags().Changed("use-ai") {
				body["useAi"] = useAi
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/instagramAiComment", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().BoolVar(&useAi, "use-ai", false, "Whether to use AI for comments (default false)")
	cmd.Flags().IntVar(&randomRate, "random-rate", 0, "Random probability, 0-100 (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("random-rate")
	return cmd
}

func newYouTubePubVideoCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, title, description, video string
	var scheduleAt int64
	var isDisclosureMandatory bool
	cmd := &cobra.Command{
		Use:     "youtube-pub-video",
		Short:   "YouTube publish video",
		Example: `  geelark-cli phone automation youtube-pub-video --id "557536075321468390" --schedule-at 1741846843 --title "title" --description "desc" --video "https://material.geelark.com/a.mp4"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "title": title, "description": description, "video": video}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if isDisclosureMandatory {
				body["isDisclosureMandatory"] = true
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/youtubePubVideo", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&title, "title", "", "Title (max 100 chars, required)")
	cmd.Flags().StringVar(&description, "description", "", "Description (max 5000 chars, required)")
	cmd.Flags().StringVar(&video, "video", "", "Video URL (required)")
	cmd.Flags().BoolVar(&isDisclosureMandatory, "is-disclosure-mandatory", false, "Whether to force disclosure (default false)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("description")
	_ = cmd.MarkFlagRequired("video")
	return cmd
}

func newYouTubePubShortCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, title, video, sameStyleUrl string
	var scheduleAt int64
	var sameStyleVoice, originalVoice int
	var isDisclosureMandatory bool
	cmd := &cobra.Command{
		Use:     "youtube-pub-short",
		Short:   "YouTube publish Short",
		Example: `  geelark-cli phone automation youtube-pub-short --id "557536075321468390" --schedule-at 1741846843 --title "title" --video "https://material.geelark.com/a.mp4" --same-style-voice 50 --original-voice 50`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "title": title, "video": video, "sameStyleVoice": sameStyleVoice, "originalVoice": originalVoice}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if sameStyleUrl != "" {
				body["sameStyleUrl"] = sameStyleUrl
			}
			if isDisclosureMandatory {
				body["isDisclosureMandatory"] = true
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/youtubePubShort", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&title, "title", "", "Title (max 100 chars, required)")
	cmd.Flags().StringVar(&video, "video", "", "Video URL (required)")
	cmd.Flags().StringVar(&sameStyleUrl, "same-style-url", "", "Same style URL (max 500 chars)")
	cmd.Flags().IntVar(&sameStyleVoice, "same-style-voice", 0, "Same style volume, 0-100 (required)")
	cmd.Flags().IntVar(&originalVoice, "original-voice", 0, "Original voice volume, 0-100 (required)")
	cmd.Flags().BoolVar(&isDisclosureMandatory, "is-disclosure-mandatory", false, "Whether to force disclosure, defaults to false")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("video")
	_ = cmd.MarkFlagRequired("same-style-voice")
	_ = cmd.MarkFlagRequired("original-voice")
	return cmd
}

func newYouTubeMaintenanceCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	var browseVideoNum int
	var keyword string
	cmd := &cobra.Command{
		Use:   "youtube-maintenance",
		Short: "YouTube account maintenance",
		Example: `  geelark-cli phone automation youtube-maintenance --id "557536075321468390" --schedule-at 1741846843 --browse-video-num 10 --keyword "k1,k2"
  geelark-cli phone automation youtube-maintenance --id "557536075321468390" --schedule-at 1741846843 --browse-video-num 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "browseVideoNum": browseVideoNum}
			if keyword != "" {
				body["keyword"] = strings.Split(keyword, ",")
			}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/youTubeActiveAccount", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&browseVideoNum, "browse-video-num", 0, "Number of videos to browse, 1-100 (required)")
	cmd.Flags().StringVar(&keyword, "keyword", "", "Comma-separated keywords, max 10")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("browse-video-num")
	return cmd
}

func newYouTubeEditProfileCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, profileName, handle, description string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "youtube-edit-profile",
		Short:   "Edit YouTube profile",
		Example: `  geelark-cli phone automation youtube-edit-profile --id "557536075321468390" --schedule-at 1741846843 --profile-name "myName"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if profileName != "" {
				body["profileName"] = profileName
			}
			if handle != "" {
				body["handle"] = handle
			}
			if description != "" {
				body["description"] = description
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/youtubeEdit", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&profileName, "profile-name", "", "Profile name (max 50 chars)")
	cmd.Flags().StringVar(&handle, "handle", "", "Handle/identifier name (max 100 chars)")
	cmd.Flags().StringVar(&description, "description", "", "Description (max 1000 chars)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newXPublishCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, description, video string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "x-publish",
		Short:   "Publish content on X (Twitter)",
		Example: `  geelark-cli phone automation x-publish --id "557536075321468390" --schedule-at 1741846843 --description "desc" --video "https://material.geelark.com/a.mp4"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "description": description, "video": strings.Split(video, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/xPublish", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&description, "description", "", "Caption (max 280 chars, required)")
	cmd.Flags().StringVar(&video, "video", "", "Comma-separated video URLs, max 1 (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("description")
	_ = cmd.MarkFlagRequired("video")
	return cmd
}

func newRedditWarmupCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, keyword, keywords string
	var scheduleAt int64
	var duration int
	cmd := &cobra.Command{
		Use:   "reddit-warmup",
		Short: "Reddit AI account warmup",
		Example: `  geelark-cli phone automation reddit-warmup --id "557536075321468390" --schedule-at 1741846843 --duration 30
  geelark-cli phone automation reddit-warmup --id "557536075321468390" --schedule-at 1741846843 --duration 30 --keywords "cat,dog"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "duration": duration}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if keywords != "" {
				body["keywords"] = strings.Split(keywords, ",")
			} else if keyword != "" {
				body["keyword"] = keyword
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/redditWarmup", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&duration, "duration", 0, "Duration in minutes (required)")
	cmd.Flags().StringVar(&keywords, "keywords", "", "Comma-separated search keywords, max 100 (preferred over --keyword)")
	cmd.Flags().StringVar(&keyword, "keyword", "", "Search keyword (deprecated, use --keywords instead)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("duration")
	return cmd
}

func newRedditVideoCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, title, description, video, community string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "reddit-video",
		Short:   "Publish video on Reddit",
		Example: `  geelark-cli phone automation reddit-video --id "557536075321468390" --schedule-at 1741846843 --title "title" --video "https://material.geelark.com/a.mp4" --community "cat"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "title": title, "video": strings.Split(video, ","), "community": community}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if description != "" {
				body["description"] = description
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/redditVideo", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&title, "title", "", "Title (required)")
	cmd.Flags().StringVar(&description, "description", "", "Description")
	cmd.Flags().StringVar(&video, "video", "", "Comma-separated video URLs (required)")
	cmd.Flags().StringVar(&community, "community", "", "Community (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("video")
	_ = cmd.MarkFlagRequired("community")
	return cmd
}

func newRedditImageCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, title, description, images, community string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "reddit-image",
		Short:   "Publish pictures and texts on Reddit",
		Example: `  geelark-cli phone automation reddit-image --id "557536075321468390" --schedule-at 1741846843 --title "title" --images "https://material.geelark.com/a.jpg" --community "cat"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "title": title, "images": strings.Split(images, ","), "community": community}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if description != "" {
				body["description"] = description
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/redditImage", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&title, "title", "", "Title (required)")
	cmd.Flags().StringVar(&description, "description", "", "Description")
	cmd.Flags().StringVar(&images, "images", "", "Comma-separated image URLs (required)")
	cmd.Flags().StringVar(&community, "community", "", "Community (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("images")
	_ = cmd.MarkFlagRequired("community")
	return cmd
}

func newThreadsImageCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, topic, title, images string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "threads-image",
		Short:   "Publish pictures and texts on Threads",
		Example: `  geelark-cli phone automation threads-image --id "557536075321468390" --schedule-at 1741846843 --title "title" --images "https://material.geelark.com/a.jpg"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "title": title, "images": strings.Split(images, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if topic != "" {
				body["topic"] = topic
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/threadsImage", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&topic, "topic", "", "Topic")
	cmd.Flags().StringVar(&title, "title", "", "Title (max 500 chars, required)")
	cmd.Flags().StringVar(&images, "images", "", "Comma-separated image URLs (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("images")
	return cmd
}

func newThreadsVideoCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, topic, title, video string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "threads-video",
		Short:   "Publish video on Threads",
		Example: `  geelark-cli phone automation threads-video --id "557536075321468390" --schedule-at 1741846843 --title "title" --video "https://material.geelark.com/a.mp4"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "title": title, "video": strings.Split(video, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if topic != "" {
				body["topic"] = topic
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/threadsVideo", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&topic, "topic", "", "Topic")
	cmd.Flags().StringVar(&title, "title", "", "Title (max 500 chars, required)")
	cmd.Flags().StringVar(&video, "video", "", "Comma-separated video URLs (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("video")
	return cmd
}

func newPinterestImageCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, title, description, images, link string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "pinterest-image",
		Short:   "Publish pictures and texts on Pinterest",
		Example: `  geelark-cli phone automation pinterest-image --id "557536075321468390" --schedule-at 1741846843 --title "title" --description "desc" --images "https://material.geelark.com/a.jpg"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "title": title, "description": description, "images": strings.Split(images, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if link != "" {
				body["link"] = link
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/pinterestImage", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&title, "title", "", "Title (max 100 chars, required)")
	cmd.Flags().StringVar(&description, "description", "", "Description (max 800 chars, required)")
	cmd.Flags().StringVar(&images, "images", "", "Comma-separated image URLs (required)")
	cmd.Flags().StringVar(&link, "link", "", "Link")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("description")
	_ = cmd.MarkFlagRequired("images")
	return cmd
}

func newPinterestVideoCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, title, description, video, link string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "pinterest-video",
		Short:   "Publish video on Pinterest",
		Example: `  geelark-cli phone automation pinterest-video --id "557536075321468390" --schedule-at 1741846843 --title "title" --description "desc" --video "https://material.geelark.com/a.mp4"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "title": title, "description": description, "video": strings.Split(video, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if link != "" {
				body["link"] = link
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/pinterestVideo", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&title, "title", "", "Title (max 100 chars, required)")
	cmd.Flags().StringVar(&description, "description", "", "Description (max 800 chars, required)")
	cmd.Flags().StringVar(&video, "video", "", "Comma-separated video URLs (required)")
	cmd.Flags().StringVar(&link, "link", "", "Link")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("description")
	_ = cmd.MarkFlagRequired("video")
	return cmd
}

func newGoogleLoginCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, email, password, code2fa string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "google-login",
		Short:   "Google auto login",
		Example: `  geelark-cli phone automation google-login --id "557536075321468390" --schedule-at 1741846843 --email "test@gmail.com" --password "123456"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "email": email, "password": password}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if code2fa != "" {
				body["code2fa"] = code2fa
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/googleLogin", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&email, "email", "", "Email (max 64 chars, required)")
	cmd.Flags().StringVar(&password, "password", "", "Password (max 64 chars, required)")
	cmd.Flags().StringVar(&code2fa, "code-2fa", "", "2FA code")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("email")
	_ = cmd.MarkFlagRequired("password")
	return cmd
}

func newGoogleAppDownloadCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, appName string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "google-app-download",
		Short:   "Download apps on Google Play",
		Example: `  geelark-cli phone automation google-app-download --id "557536075321468390" --schedule-at 1741846843 --app-name "TikTok"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "appName": appName}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/googleAppDownload", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&appName, "app-name", "", "App name (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("app-name")
	return cmd
}

func newGoogleAppBrowserCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, appName, description string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "google-app-browser",
		Short:   "Open the app on Google for browsing",
		Example: `  geelark-cli phone automation google-app-browser --id "557536075321468390" --schedule-at 1741846843 --app-name "TikTok"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "appName": appName}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if description != "" {
				body["description"] = description
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/googleAppBrowser", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&appName, "app-name", "", "App name (required)")
	cmd.Flags().StringVar(&description, "description", "", "Describe your experience")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("app-name")
	return cmd
}

func newSheinLoginCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, email, password string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "shein-login",
		Short:   "SHEIN auto login",
		Example: `  geelark-cli phone automation shein-login --id "557536075321468390" --schedule-at 1741846843 --email "test@gmail.com" --password "123456"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "email": email, "password": password}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/sheinLogin", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&email, "email", "", "Email (max 64 chars, required)")
	cmd.Flags().StringVar(&password, "password", "", "Password (max 64 chars, required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("email")
	_ = cmd.MarkFlagRequired("password")
	return cmd
}

func newImportContactsTaskCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, contactsJSON string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "import-contacts",
		Short:   "Batch import contacts to cloud phone",
		Example: `  geelark-cli phone automation import-contacts --id "557536075321468390" --schedule-at 1741846843 --contacts "[{\"firstName\":\"jay\",\"mobile\":\"13288888888\"}]"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			var contacts interface{}
			if err := json.Unmarshal([]byte(contactsJSON), &contacts); err != nil {
				return fmt.Errorf("invalid --contacts JSON: %w", err)
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "contacts": contacts}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/importContacts", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&contactsJSON, "contacts", "", "Contacts array as JSON string (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("contacts")
	return cmd
}

func newKeyboxUploadCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, files string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "keybox-upload",
		Short:   "Upload Keybox to the cloud phone",
		Example: `  geelark-cli phone automation keybox-upload --id "557536075321468390" --schedule-at 1741846843 --files "https://material.geelark.com/a.xml"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "files": strings.Split(files, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/keyboxUpload", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&files, "files", "", "Comma-separated file URLs, max 100 (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("files")
	return cmd
}

func newFileUploadCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, files string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "file-upload",
		Short:   "Upload files to the cloud phone in batches",
		Example: `  geelark-cli phone automation file-upload --id "557536075321468390" --schedule-at 1741846843 --files "https://material.geelark.com/a.mp4"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "files": strings.Split(files, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/fileUpload", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&files, "files", "", "Comma-separated file URLs, max 100 (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("files")
	return cmd
}

func newMultiPlatformVideoCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, title, video string
	var tiktokTitle, youtubeTitle, instagramTitle string
	var tiktokRecreateLink, youtubeRecreateLink, instagramRecreateLink string
	var scheduleAt int64
	var sameStyleVoice, originalVoice int
	cmd := &cobra.Command{
		Use:   "multi-platform-video",
		Short: "Multichannel video distribution (TikTok/Instagram Reels/YouTube Shorts)",
		Long: `Multichannel video distribution (TikTok/Instagram Reels/YouTube Shorts).
Use --title for a shared title across all platforms, or use per-platform titles
(--tiktok-title, --youtube-title, --instagram-title) for individual titles.
Use --tiktok-recreate-link / --youtube-recreate-link / --instagram-recreate-link
to recreate the same style, along with --same-style-voice and --original-voice.`,
		Example: `  geelark-cli phone automation multi-platform-video --id "557536075321468390" --schedule-at 1741846843 --title "title" --video "https://material.geelark.com/a.mp4"
  geelark-cli phone automation multi-platform-video --id "557536075321468390" --schedule-at 1741846843 --tiktok-title "tt" --youtube-title "yt" --instagram-title "ig" --video "https://material.geelark.com/a.mp4" --tiktok-recreate-link "https://example.com" --same-style-voice 50 --original-voice 50`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "video": strings.Split(video, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if title != "" {
				body["title"] = title
			}
			if tiktokTitle != "" {
				body["tiktokTitle"] = tiktokTitle
			}
			if youtubeTitle != "" {
				body["youtubeTitle"] = youtubeTitle
			}
			if instagramTitle != "" {
				body["instagramTitle"] = instagramTitle
			}
			if tiktokRecreateLink != "" {
				body["tiktokRecreateLink"] = tiktokRecreateLink
			}
			if youtubeRecreateLink != "" {
				body["youtubeRecreateLink"] = youtubeRecreateLink
			}
			if instagramRecreateLink != "" {
				body["instagramRecreateLink"] = instagramRecreateLink
			}
			if sameStyleVoice > 0 {
				body["sameStyleVoice"] = sameStyleVoice
			}
			if originalVoice > 0 {
				body["originalVoice"] = originalVoice
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/multiPlatformVideoDistribution", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&title, "title", "", "Shared title for all platforms (max 100 chars)")
	cmd.Flags().StringVar(&tiktokTitle, "tiktok-title", "", "TikTok title (max 4000 chars)")
	cmd.Flags().StringVar(&youtubeTitle, "youtube-title", "", "YouTube title (max 100 chars)")
	cmd.Flags().StringVar(&instagramTitle, "instagram-title", "", "Instagram title (max 2200 chars)")
	cmd.Flags().StringVar(&video, "video", "", "Comma-separated video URLs, max 10 (required)")
	cmd.Flags().StringVar(&tiktokRecreateLink, "tiktok-recreate-link", "", "TikTok style link")
	cmd.Flags().StringVar(&youtubeRecreateLink, "youtube-recreate-link", "", "YouTube style link")
	cmd.Flags().StringVar(&instagramRecreateLink, "instagram-recreate-link", "", "Instagram style link")
	cmd.Flags().IntVar(&sameStyleVoice, "same-style-voice", 0, "Same style volume, 0-100")
	cmd.Flags().IntVar(&originalVoice, "original-voice", 0, "Original voice volume, 0-100")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("video")
	return cmd
}

func newTikTokStarAsiaCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	var likeProbability int
	cmd := &cobra.Command{
		Use:     "tiktok-star-asia",
		Short:   "TikTok random star (Asia)",
		Example: `  geelark-cli phone automation tiktok-star-asia --id "557536075321468390" --schedule-at 1741846843`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if likeProbability > 0 {
				body["likeProbability"] = likeProbability
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokRandomStarAsia", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&likeProbability, "like-probability", 0, "Probability of liking, 0-100, default 30")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newTikTokCommentCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, comment, imageUrl string
	var scheduleAt int64
	var useAi, commentProbability int
	var links, searchKeywords string
	var likeVideo bool
	cmd := &cobra.Command{
		Use:     "tiktok-comment",
		Short:   "TikTok AI comment",
		Example: `  geelark-cli phone automation tiktok-comment --id "557536075321468390" --schedule-at 1741846843 --use-ai 2 --comment "test"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "useAi": useAi, "comment": comment}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if links != "" {
				body["links"] = strings.Split(links, ",")
			}
			if commentProbability > 0 {
				body["commentProbability"] = commentProbability
			}
			if searchKeywords != "" {
				body["searchKeywords"] = strings.Split(searchKeywords, ",")
			}
			if likeVideo {
				body["likeVideo"] = true
			}
			if imageUrl != "" {
				body["imageUrl"] = imageUrl
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokRandomComment", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&useAi, "use-ai", 0, "Use AI: 1=AI (Pro only), 2=manual comment (required)")
	cmd.Flags().StringVar(&comment, "comment", "", "Comment content (max 500 chars, required when use-ai=2)")
	cmd.Flags().StringVar(&links, "links", "", "Comma-separated specified links")
	cmd.Flags().IntVar(&commentProbability, "comment-probability", 0, "Comment probability, 0-100, default 30")
	cmd.Flags().StringVar(&searchKeywords, "search-keywords", "", "Comma-separated search keywords")
	cmd.Flags().BoolVar(&likeVideo, "like-video", false, "Whether to like, defaults to false")
	cmd.Flags().StringVar(&imageUrl, "image-url", "", "Comment image URL, max 500 chars (effective when use-ai=2)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("use-ai")
	return cmd
}

func newTikTokCommentAsiaCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, comment, imageUrl string
	var scheduleAt int64
	var useAi, commentProbability int
	var links, searchKeywords string
	var likeVideo bool
	cmd := &cobra.Command{
		Use:     "tiktok-comment-asia",
		Short:   "TikTok AI comment (Asia)",
		Example: `  geelark-cli phone automation tiktok-comment-asia --id "557536075321468390" --schedule-at 1741846843 --use-ai 2 --comment "test"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "useAi": useAi, "comment": comment}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if links != "" {
				body["links"] = strings.Split(links, ",")
			}
			if commentProbability > 0 {
				body["commentProbability"] = commentProbability
			}
			if searchKeywords != "" {
				body["searchKeywords"] = strings.Split(searchKeywords, ",")
			}
			if likeVideo {
				body["likeVideo"] = true
			}
			if imageUrl != "" {
				body["imageUrl"] = imageUrl
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokRandomCommentAsia", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&useAi, "use-ai", 0, "Use AI: 1=AI (Pro only), 2=manual comment (required)")
	cmd.Flags().StringVar(&comment, "comment", "", "Comment content (max 500 chars, required when use-ai=2)")
	cmd.Flags().StringVar(&links, "links", "", "Comma-separated specified links")
	cmd.Flags().IntVar(&commentProbability, "comment-probability", 0, "Comment probability, 0-100, default 30")
	cmd.Flags().StringVar(&searchKeywords, "search-keywords", "", "Comma-separated search keywords")
	cmd.Flags().BoolVar(&likeVideo, "like-video", false, "Whether to like, defaults to false")
	cmd.Flags().StringVar(&imageUrl, "image-url", "", "Comment image URL, max 500 chars (effective when use-ai=2)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("use-ai")
	return cmd
}

func newTikTokFollowCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	var followProbability int
	cmd := &cobra.Command{
		Use:     "tiktok-follow",
		Short:   "TikTok random follow",
		Example: `  geelark-cli phone automation tiktok-follow --id "557536075321468390" --schedule-at 1741846843 --follow-probability 30`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "followProbability": followProbability}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokRandomFollow", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&followProbability, "follow-probability", 0, "Follow probability, 0-100 (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("follow-probability")
	return cmd
}

func newTikTokFollowAsiaCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	var followProbability int
	cmd := &cobra.Command{
		Use:     "tiktok-follow-asia",
		Short:   "TikTok random follow (Asia)",
		Example: `  geelark-cli phone automation tiktok-follow-asia --id "557536075321468390" --schedule-at 1741846843 --follow-probability 30`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "followProbability": followProbability}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokRandomFollowAsia", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&followProbability, "follow-probability", 0, "Follow probability, 0-100 (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("follow-probability")
	return cmd
}

func newTikTokEditProfileCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, avatar, nickName, bio, site, username string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "tiktok-edit-profile",
		Short:   "TikTok profile edit",
		Example: `  geelark-cli phone automation tiktok-edit-profile --id "557536075321468390" --schedule-at 1741846843 --nick-name "test"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if avatar != "" {
				body["avatar"] = avatar
			}
			if nickName != "" {
				body["nickName"] = nickName
			}
			if bio != "" {
				body["bio"] = bio
			}
			if site != "" {
				body["site"] = site
			}
			if username != "" {
				body["username"] = username
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokEdit", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&avatar, "avatar", "", "Avatar URL (1:1 aspect ratio)")
	cmd.Flags().StringVar(&nickName, "nick-name", "", "Nickname (max 30 chars)")
	cmd.Flags().StringVar(&bio, "bio", "", "Bio (max 160 chars)")
	cmd.Flags().StringVar(&site, "site", "", "Website URL (must start with http/https)")
	cmd.Flags().StringVar(&username, "username", "", "Username")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newTikTokMessageCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, usernames, content string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "tiktok-message",
		Short:   "Send private message on TikTok",
		Example: `  geelark-cli phone automation tiktok-message --id "557536075321468390" --schedule-at 1741846843 --usernames "user1" --content "hello"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "usernames": strings.Split(usernames, ","), "content": content}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokMessage", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&usernames, "usernames", "", "Comma-separated usernames (required)")
	cmd.Flags().StringVar(&content, "content", "", "Message content, max 6000 chars (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("usernames")
	_ = cmd.MarkFlagRequired("content")
	return cmd
}

func newTikTokMessageAsiaCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, usernames, content string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "tiktok-message-asia",
		Short:   "Send private message on TikTok (Asia)",
		Example: `  geelark-cli phone automation tiktok-message-asia --id "557536075321468390" --schedule-at 1741846843 --usernames "user1" --content "hello"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "usernames": strings.Split(usernames, ","), "content": content}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokMessageAsia", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&usernames, "usernames", "", "Comma-separated usernames (required)")
	cmd.Flags().StringVar(&content, "content", "", "Message content, max 6000 chars (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("usernames")
	_ = cmd.MarkFlagRequired("content")
	return cmd
}

func newTikTokHideCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	var number int
	cmd := &cobra.Command{
		Use:     "tiktok-hide",
		Short:   "Hide TikTok videos",
		Example: `  geelark-cli phone automation tiktok-hide --id "557536075321468390" --schedule-at 1741846843`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if number > 0 {
				body["number"] = number
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokHide", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&number, "number", 0, "Number of videos to hide, range 0-999; 0 or unset = hide all")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newTikTokHideAsiaCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	var number int
	cmd := &cobra.Command{
		Use:     "tiktok-hide-asia",
		Short:   "Hide TikTok videos (Asia)",
		Example: `  geelark-cli phone automation tiktok-hide-asia --id "557536075321468390" --schedule-at 1741846843`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			if number > 0 {
				body["number"] = number
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokHideAsia", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().IntVar(&number, "number", 0, "Number of videos to hide, range 0-999; 0 or unset = hide all")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newTikTokDeleteCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "tiktok-delete",
		Short:   "Delete all TikTok videos",
		Example: `  geelark-cli phone automation tiktok-delete --id "557536075321468390" --schedule-at 1741846843`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokDel", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newTikTokDeleteAsiaCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "tiktok-delete-asia",
		Short:   "Delete all TikTok videos (Asia)",
		Example: `  geelark-cli phone automation tiktok-delete-asia --id "557536075321468390" --schedule-at 1741846843`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokDelAsia", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 128 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	return cmd
}

func newTikTokDeleteCommentCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, keywords string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "tiktok-delete-comment",
		Short:   "Delete TikTok comments",
		Example: `  geelark-cli phone automation tiktok-delete-comment --id "557536075321468390" --schedule-at 1741846843 --keywords "hello,world"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "keywords": strings.Split(keywords, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokDeleteComment", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 32 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&keywords, "keywords", "", "Comma-separated keywords, max 100 items, 100 chars each (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("keywords")
	return cmd
}

func newTikTokDeleteCommentAsiaCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark, keywords string
	var scheduleAt int64
	cmd := &cobra.Command{
		Use:     "tiktok-delete-comment-asia",
		Short:   "Delete TikTok comments (Asia)",
		Example: `  geelark-cli phone automation tiktok-delete-comment-asia --id "557536075321468390" --schedule-at 1741846843 --keywords "hello,world"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id, "scheduleAt": scheduleAt, "keywords": strings.Split(keywords, ",")}
			if name != "" {
				body["name"] = name
			}
			if remark != "" {
				body["remark"] = remark
			}
			result, err := c.PostAndPrint("/open/v1/rpa/task/tiktokDeleteCommentAsia", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Task name (max 32 chars)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark (max 200 chars)")
	cmd.Flags().Int64Var(&scheduleAt, "schedule-at", 0, "Schedule time, second-level timestamp (required)")
	cmd.Flags().StringVar(&keywords, "keywords", "", "Comma-separated keywords, max 100 items, 100 chars each (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("schedule-at")
	_ = cmd.MarkFlagRequired("keywords")
	return cmd
}
