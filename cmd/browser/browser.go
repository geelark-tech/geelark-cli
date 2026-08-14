package browser

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/geelark-tech/geelark-cli/internal/client"
	"github.com/spf13/cobra"
)

type clientFactory func() (*client.Client, error)

// NewCmd creates the browser command group.
func NewCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "browser",
		Short: "Browser management (local API)",
		Long: `Manage GeeLark browsers via the local API.
Requires the GeeLark client to be running and logged in.
Default endpoint: http://localhost:40185 (configurable via 'config init --browser-base-url').`,
	}

	cmd.AddCommand(newStatusCmd(newClient))
	cmd.AddCommand(newListCmd(newClient))
	cmd.AddCommand(newCreateCmd(newClient))
	cmd.AddCommand(newSimpleCreateCmd(newClient))
	cmd.AddCommand(newEditCmd(newClient))
	cmd.AddCommand(newDeleteCmd(newClient))
	cmd.AddCommand(newStartCmd(newClient))
	cmd.AddCommand(newStopCmd(newClient))
	cmd.AddCommand(newCheckStatusCmd(newClient))
	cmd.AddCommand(newCloneCmd(newClient))
	cmd.AddCommand(newClearCacheCmd(newClient))
	cmd.AddCommand(newMoveGroupCmd(newClient))
	cmd.AddCommand(newTransferCmd(newClient))
	cmd.AddCommand(newGetCookieCmd(newClient))
	cmd.AddCommand(newGetBookmarkCmd(newClient))
	cmd.AddCommand(newSetBookmarkCmd(newClient))
	cmd.AddCommand(newGetKernelsCmd(newClient))
	cmd.AddCommand(newUpdateKernelsCmd(newClient))
	cmd.AddCommand(newExtGroupListCmd(newClient))
	cmd.AddCommand(newAutomationCmd(newClient))

	return cmd
}

func newStatusCmd(newClient clientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "api-status",
		Short:   "Check API interface availability",
		Long:    "Check whether the local Browser API is available.",
		Example: `  geelark-cli browser api-status`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostBrowserAndPrint("/api/v1/status", nil)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func newListCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int
	var ids, serialName, remark, groupName, tags string

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List browsers",
		Long:    "Query the created browser environments with optional filters.",
		Example: `  geelark-cli browser list --page 1 --page-size 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{}
			if page > 0 {
				body["page"] = page
			}
			if pageSize > 0 {
				body["pageSize"] = pageSize
			}
			if ids != "" {
				body["ids"] = strings.Split(ids, ",")
			}
			if serialName != "" {
				body["serialName"] = serialName
			}
			if remark != "" {
				body["remark"] = remark
			}
			if groupName != "" {
				body["groupName"] = groupName
			}
			if tags != "" {
				body["tags"] = strings.Split(tags, ",")
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Page size (max 100)")
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated browser IDs (max 100)")
	cmd.Flags().StringVar(&serialName, "serial-name", "", "Filter by browser name")
	cmd.Flags().StringVar(&remark, "remark", "", "Filter by remark")
	cmd.Flags().StringVar(&groupName, "group-name", "", "Filter by group name")
	cmd.Flags().StringVar(&tags, "tags", "", "Comma-separated tag names")

	return cmd
}

func newCreateCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new browser",
		Long: `Create a new browser environment with full configuration via JSON.
Supports account, proxy, fingerprint, simulation settings, etc.`,
		Example: `  geelark-cli browser create --data "{\"serialName\":\"myBrowser\",\"browserOs\":1}"
  geelark-cli browser create --data "{\"serialName\":\"test\",\"browserOs\":2,\"accountPlatform\":\"https://www.tiktok.com/\",\"accountUsername\":\"user\",\"accountPassword\":\"pass\"}"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			var body interface{}
			if err := json.Unmarshal([]byte(dataJSON), &body); err != nil {
				return fmt.Errorf("invalid --data JSON: %w", err)
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/create", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&dataJSON, "data", "", "JSON object with browser creation parameters (required)")
	_ = cmd.MarkFlagRequired("data")
	return cmd
}

func newSimpleCreateCmd(newClient clientFactory) *cobra.Command {
	var serialName, browserKernelVer, extGroup string
	var browserOs int

	cmd := &cobra.Command{
		Use:   "simple-create",
		Short: "Quick create a single browser",
		Long: `Simplified command to create a single browser with just the required fields.
Use 'create' for full configuration via JSON.`,
		Example: `  geelark-cli browser simple-create --serial-name "myBrowser" --browser-os 1
  geelark-cli browser simple-create --serial-name "macBrowser" --browser-os 2
  geelark-cli browser simple-create --serial-name "myBrowser" --browser-os 1 --browser-kernel-ver "149" --ext-group "497548067550006541"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"serialName": serialName,
				"browserOs":  browserOs,
			}
			if browserKernelVer != "" {
				body["browserKernelVer"] = browserKernelVer
			}
			if extGroup != "" {
				body["extGroup"] = extGroup
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/create", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&serialName, "serial-name", "", "Browser environment name, max 100 chars (required)")
	cmd.Flags().IntVar(&browserOs, "browser-os", 1, "Operating system: 1=Win, 2=Mac (required)")
	cmd.Flags().StringVar(&browserKernelVer, "browser-kernel-ver", "", "Browser kernel version: 134,138,142,143,144,145,146,147,148,149,150,auto (default auto)")
	cmd.Flags().StringVar(&extGroup, "ext-group", "", "Extension category ID (empty = team extensions)")
	_ = cmd.MarkFlagRequired("serial-name")
	_ = cmd.MarkFlagRequired("browser-os")
	return cmd
}

func newEditCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string

	cmd := &cobra.Command{
		Use:     "edit",
		Short:   "Edit a browser",
		Long:    "Update an existing browser environment configuration via JSON.",
		Example: `  geelark-cli browser edit --data "{\"id\":\"browser_id\",\"serialName\":\"newName\"}"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			var body interface{}
			if err := json.Unmarshal([]byte(dataJSON), &body); err != nil {
				return fmt.Errorf("invalid --data JSON: %w", err)
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/update", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&dataJSON, "data", "", "JSON object with browser update parameters (required)")
	_ = cmd.MarkFlagRequired("data")
	return cmd
}

func newDeleteCmd(newClient clientFactory) *cobra.Command {
	var ids string

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete browsers",
		Long:    "Delete browser environments by IDs (max 100).",
		Example: `  geelark-cli browser delete --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"envIds": strings.Split(ids, ","),
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/delete", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated browser IDs (required)")
	_ = cmd.MarkFlagRequired("ids")
	return cmd
}

func newStartCmd(newClient clientFactory) *cobra.Command {
	var id, webhook string

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Launch a browser",
		Long: `Launch/open a browser environment by ID.
Optionally specify a webhook URL to receive a callback when the browser finishes starting.`,
		Example: `  geelark-cli browser start --id "browser_id"
  geelark-cli browser start --id "browser_id" --webhook "http://localhost:3001"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id}
			if webhook != "" {
				body["webhook"] = webhook
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/start", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Browser ID (required)")
	cmd.Flags().StringVar(&webhook, "webhook", "", "Callback URL to notify after browser startup completes")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func newStopCmd(newClient clientFactory) *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:     "stop",
		Short:   "Close a browser",
		Long:    "Close a running browser environment by ID.",
		Example: `  geelark-cli browser stop --id "browser_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/stop", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Browser ID (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func newCheckStatusCmd(newClient clientFactory) *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:     "check-status",
		Short:   "Check browser startup status",
		Long:    "Check whether a specific browser is running. Returns status (open/close) and debug port.",
		Example: `  geelark-cli browser check-status --id "browser_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/status", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Browser ID (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func newCloneCmd(newClient clientFactory) *cobra.Command {
	var envID, groupID string
	var amount int
	var cloneName, cloneRemark, cloneTag, cloneProxy, cloneCookie, cloneAccount bool

	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone a browser",
		Long:  "Generate new browsers by cloning an existing one with the same OS and advanced settings.",
		Example: `  geelark-cli browser clone --env-id "browser_id" --amount 2
  geelark-cli browser clone --env-id "browser_id" --amount 1 --clone-name --clone-proxy --clone-cookie`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"envId":  envID,
				"amount": amount,
			}
			if groupID != "" {
				body["groupId"] = groupID
			}
			if cloneName {
				body["cloneName"] = true
			}
			if cloneRemark {
				body["cloneRemark"] = true
			}
			if cloneTag {
				body["cloneTag"] = true
			}
			if cloneProxy {
				body["cloneProxy"] = true
			}
			if cloneCookie {
				body["cloneCookie"] = true
			}
			if cloneAccount {
				body["cloneAccount"] = true
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/clone", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envID, "env-id", "", "Browser ID to clone (required)")
	cmd.Flags().IntVar(&amount, "amount", 1, "Number of clones (1-100)")
	cmd.Flags().StringVar(&groupID, "group-id", "", "Target group ID")
	cmd.Flags().BoolVar(&cloneName, "clone-name", false, "Clone the name")
	cmd.Flags().BoolVar(&cloneRemark, "clone-remark", false, "Clone the remark")
	cmd.Flags().BoolVar(&cloneTag, "clone-tag", false, "Clone the tags")
	cmd.Flags().BoolVar(&cloneProxy, "clone-proxy", false, "Clone the proxy")
	cmd.Flags().BoolVar(&cloneCookie, "clone-cookie", false, "Clone the cookies")
	cmd.Flags().BoolVar(&cloneAccount, "clone-account", false, "Clone the account information")
	_ = cmd.MarkFlagRequired("env-id")
	return cmd
}

func newClearCacheCmd(newClient clientFactory) *cobra.Command {
	var ids, cacheTypes string

	cmd := &cobra.Command{
		Use:   "clear-cache",
		Short: "Clear browser cache",
		Long: `Clear local cache of browser environments. Ensure browsers are closed first.
Cache types: local_storage, indexeddb, extension_cache, cookie, history, image_file`,
		Example: `  geelark-cli browser clear-cache --ids "id1,id2" --type "cookie,history"
  geelark-cli browser clear-cache --ids "id1" --type "local_storage,indexeddb,extension_cache,cookie,history,image_file"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"ids":  strings.Split(ids, ","),
				"type": strings.Split(cacheTypes, ","),
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/deleteCache", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated browser IDs (required)")
	cmd.Flags().StringVar(&cacheTypes, "type", "", "Comma-separated cache types: local_storage,indexeddb,extension_cache,cookie,history,image_file (required)")
	_ = cmd.MarkFlagRequired("ids")
	_ = cmd.MarkFlagRequired("type")
	return cmd
}

func newMoveGroupCmd(newClient clientFactory) *cobra.Command {
	var envIds, groupID string

	cmd := &cobra.Command{
		Use:     "move-group",
		Short:   "Move browsers to a group",
		Long:    `Move browser environments to a specified group. Use "0" for ungrouped. Max 100.`,
		Example: `  geelark-cli browser move-group --env-ids "id1,id2" --group-id "group_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"envIds":  strings.Split(envIds, ","),
				"groupId": groupID,
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/moveGroup", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envIds, "env-ids", "", "Comma-separated browser IDs (required)")
	cmd.Flags().StringVar(&groupID, "group-id", "", "Target group ID, '0' for ungrouped (required)")
	_ = cmd.MarkFlagRequired("env-ids")
	_ = cmd.MarkFlagRequired("group-id")
	return cmd
}

func newTransferCmd(newClient clientFactory) *cobra.Command {
	var envIDs, username, transferOption string

	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "Transfer browsers to another account",
		Long: `Transfer browser environments to a target account. Max 200 per request.
transfer-option allowed values (comma-separated): name, proxy, tag, remark, files`,
		Example: `  geelark-cli browser transfer --env-ids "id1,id2" --username "user@geelark.com"
  geelark-cli browser transfer --env-ids "id1" --username "user@geelark.com" --transfer-option "name,proxy,tag,remark"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"envIds":   strings.Split(envIDs, ","),
				"username": username,
			}
			if transferOption != "" {
				body["transferOption"] = strings.Split(transferOption, ",")
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/transfer", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envIDs, "env-ids", "", "Comma-separated browser IDs, max 200 (required)")
	cmd.Flags().StringVar(&username, "username", "", "Target user account email (required)")
	cmd.Flags().StringVar(&transferOption, "transfer-option", "", "Comma-separated options: name, proxy, tag, remark, files")
	_ = cmd.MarkFlagRequired("env-ids")
	_ = cmd.MarkFlagRequired("username")
	return cmd
}

func newGetCookieCmd(newClient clientFactory) *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:     "get-cookie",
		Short:   "Get browser cookies",
		Long:    "Query the cookies of a browser environment.",
		Example: `  geelark-cli browser get-cookie --id "browser_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/getCookie", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Browser ID (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func newGetBookmarkCmd(newClient clientFactory) *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:     "get-bookmark",
		Short:   "Get browser bookmarks",
		Long:    "Query the bookmarks of a browser environment.",
		Example: `  geelark-cli browser get-bookmark --id "browser_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/getBookmark", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Browser ID (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func newSetBookmarkCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string

	cmd := &cobra.Command{
		Use:     "set-bookmark",
		Short:   "Set browser bookmarks",
		Long:    "Set/update bookmarks for a browser environment.",
		Example: `  geelark-cli browser set-bookmark --data "{\"id\":\"browser_id\",\"bookmarks\":[{\"name\":\"Google\",\"url\":\"https://google.com\"}]}"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			var body interface{}
			if err := json.Unmarshal([]byte(dataJSON), &body); err != nil {
				return fmt.Errorf("invalid --data JSON: %w", err)
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/setBookmark", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&dataJSON, "data", "", "JSON bookmark data (required)")
	_ = cmd.MarkFlagRequired("data")
	return cmd
}

func newGetKernelsCmd(newClient clientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "get-kernels",
		Short:   "List available browser kernels",
		Long:    "Query the list of available browser kernels.",
		Example: `  geelark-cli browser get-kernels`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/getKernelsList", nil)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func newUpdateKernelsCmd(newClient clientFactory) *cobra.Command {
	var kernelVersion string

	cmd := &cobra.Command{
		Use:     "update-kernels",
		Short:   "Download and update a browser kernel",
		Long:    "Download or update the specified browser kernel to the latest version.",
		Example: `  geelark-cli browser update-kernels --kernel-version "143"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"kernel_version": kernelVersion,
			}
			result, err := c.PostBrowserAndPrint("/api/v1/browser/updateKernels", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&kernelVersion, "kernel-version", "", "Browser kernel version, e.g. 143 (required)")
	_ = cmd.MarkFlagRequired("kernel-version")
	return cmd
}

func newExtGroupListCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int

	cmd := &cobra.Command{
		Use:   "ext-group-list",
		Short: "List browser extension categories",
		Long: `Query the browser extension categories for the current team.
Returns the available extGroup IDs that can be used when creating or editing a browser.

Note: This command is not the local browser API.`,
		Example: `  geelark-cli browser ext-group-list --page 1 --page-size 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{}
			if page > 0 {
				body["page"] = page
			}
			if pageSize > 0 {
				body["pageSize"] = pageSize
			}
			result, err := c.PostAndPrint("/open/v1/browser/extGroup/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number, min 1")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Page size, 1-100")
	return cmd
}
