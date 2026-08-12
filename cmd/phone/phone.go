package phone

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/geelark-tech/geelark-cli/internal/client"
	"github.com/geelark-tech/geelark-cli/internal/output"
	"github.com/spf13/cobra"
)

type clientFactory func() (*client.Client, error)

// NewCmd creates the phone command group.
func NewCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "phone",
		Short: "Cloud phone management",
		Long:  "Manage GeeLark cloud phones — list, create, start, stop, delete, update, clone, screenshot, set GPS, and more.",
	}

	cmd.AddCommand(newListCmd(newClient))
	cmd.AddCommand(newStartCmd(newClient))
	cmd.AddCommand(newStopCmd(newClient))
	cmd.AddCommand(newRestartCmd(newClient))
	cmd.AddCommand(newDeleteCmd(newClient))
	cmd.AddCommand(newStatusCmd(newClient))
	cmd.AddCommand(newCreateCmd(newClient))
	cmd.AddCommand(newUpdateCmd(newClient))
	cmd.AddCommand(newCloneCmd(newClient))
	cmd.AddCommand(newScreenshotCmd(newClient))
	cmd.AddCommand(newScreenshotResultCmd(newClient))
	cmd.AddCommand(newGetGPSCmd(newClient))
	cmd.AddCommand(newSetGPSCmd(newClient))
	cmd.AddCommand(newResetCmd(newClient))
	cmd.AddCommand(newSetRootCmd(newClient))
	cmd.AddCommand(newGetDeviceIDCmd(newClient))
	cmd.AddCommand(newSendSMSCmd(newClient))
	cmd.AddCommand(newBrandListCmd(newClient))
	cmd.AddCommand(newBrandTeamListCmd(newClient))
	cmd.AddCommand(newTransferCmd(newClient))
	cmd.AddCommand(newSetNetTypeCmd(newClient))
	cmd.AddCommand(newHideAccessibilityCmd(newClient))
	cmd.AddCommand(newMoveGroupCmd(newClient))
	cmd.AddCommand(newImportContactsCmd(newClient))
	cmd.AddCommand(newImportContactsResultCmd(newClient))
	cmd.AddCommand(newNetConfigGetCmd(newClient))
	cmd.AddCommand(newNetConfigSetCmd(newClient))
	cmd.AddCommand(newSimpleCreateCmd(newClient))
	cmd.AddCommand(newAutomationCmd(newClient))
	cmd.AddCommand(newLibraryCmd(newClient))
	cmd.AddCommand(newFileCmd(newClient))
	cmd.AddCommand(newShellCmd(newClient))
	cmd.AddCommand(newADBCmd(newClient))
	cmd.AddCommand(newWebhookCmd(newClient))
	cmd.AddCommand(newOEMCmd(newClient))
	cmd.AddCommand(newAnalyticsCmd(newClient))
	cmd.AddCommand(newAppCmd(newClient))

	return cmd
}

func newListCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int
	var chargeMode, openStatus int
	var ids, serialName, remark, groupName string
	var tags, proxyIds, serialNos string
	var chargeModeSet, openStatusSet bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all cloud phones",
		Long:  "Retrieve the list of cloud phones with optional filters.",
		Example: `  geelark-cli phone list --page 1 --page-size 10
  geelark-cli phone list --ids "id1,id2"
  geelark-cli phone list --serial-name "test"
  geelark-cli phone list --tags "tag1,tag2"
  geelark-cli phone list --charge-mode 1 --open-status 1
  geelark-cli phone list --proxy-ids "proxy1,proxy2"
  geelark-cli phone list --serial-nos "238,239"`,
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
			if chargeModeSet {
				body["chargeMode"] = chargeMode
			}
			if openStatusSet {
				body["openStatus"] = openStatus
			}
			if proxyIds != "" {
				body["proxyIds"] = strings.Split(proxyIds, ",")
			}
			if serialNos != "" {
				body["serialNos"] = strings.Split(serialNos, ",")
			}

			result, err := c.PostAndPrint("/open/v1/phone/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Number of records per page (max 100)")
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs (max 100, ignores page/pageSize)")
	cmd.Flags().StringVar(&serialName, "serial-name", "", "Filter by cloud phone name")
	cmd.Flags().StringVar(&remark, "remark", "", "Filter by remark")
	cmd.Flags().StringVar(&groupName, "group-name", "", "Filter by group name")
	cmd.Flags().StringVar(&tags, "tags", "", "Comma-separated tag names")
	cmd.Flags().IntVar(&chargeMode, "charge-mode", -1, "Charge mode: 0=pay-per-minute, 1=monthly subscription")
	cmd.Flags().IntVar(&openStatus, "open-status", -1, "Power state: 0=off, 1=on")
	cmd.Flags().StringVar(&proxyIds, "proxy-ids", "", "Comma-separated proxy IDs (max 10)")
	cmd.Flags().StringVar(&serialNos, "serial-nos", "", "Comma-separated cloud phone serial numbers (max 100)")

	// Track whether charge-mode / open-status were explicitly set
	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		chargeModeSet = cmd.Flags().Changed("charge-mode")
		openStatusSet = cmd.Flags().Changed("open-status")
		return nil
	}

	return cmd
}

func newStartCmd(newClient clientFactory) *cobra.Command {
	var ids string
	var width, center, energySavingMode int
	var materialTagIds string
	var widthSet, centerSet, energySavingSet bool

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start cloud phones",
		Long: `Batch start cloud phones by IDs.
Supports display width, centering, energy-saving mode, and material tag filters.`,
		Example: `  geelark-cli phone start --ids "id1,id2,id3"
  geelark-cli phone start --ids "id1" --width 480 --center 1 --energy-saving 1
  geelark-cli phone start --ids "id1" --material-tag-ids "tagId1,tagId2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"ids": strings.Split(ids, ","),
			}
			if widthSet {
				body["width"] = width
			}
			if centerSet {
				body["center"] = center
			}
			if energySavingSet {
				body["energySavingMode"] = energySavingMode
			}
			if materialTagIds != "" {
				body["materialTagIds"] = strings.Split(materialTagIds, ",")
			}

			result, err := c.PostAndPrint("/open/v1/phone/start", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs, max 200 (required)")
	cmd.Flags().IntVar(&width, "width", 336, "Cloud phone display width in px (200-600, default 336)")
	cmd.Flags().IntVar(&center, "center", 1, "Whether display is centered: 0=no, 1=yes (default 1)")
	cmd.Flags().IntVar(&energySavingMode, "energy-saving", 0, "Energy-saving mode: 0=disabled, 1=enabled (auto shutdown after 30min idle)")
	cmd.Flags().StringVar(&materialTagIds, "material-tag-ids", "", "Comma-separated material tag IDs (max 10, requires OEM)")
	_ = cmd.MarkFlagRequired("ids")

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		widthSet = cmd.Flags().Changed("width")
		centerSet = cmd.Flags().Changed("center")
		energySavingSet = cmd.Flags().Changed("energy-saving")
		return nil
	}

	return cmd
}

func newStopCmd(newClient clientFactory) *cobra.Command {
	var ids string

	cmd := &cobra.Command{
		Use:     "stop",
		Short:   "Stop cloud phones",
		Long:    "Batch shut down cloud phones by IDs (max 200).",
		Example: `  geelark-cli phone stop --ids "id1,id2,id3"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"ids": strings.Split(ids, ","),
			}

			result, err := c.PostAndPrint("/open/v1/phone/stop", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs, max 200 (required)")
	_ = cmd.MarkFlagRequired("ids")

	return cmd
}

func newDeleteCmd(newClient clientFactory) *cobra.Command {
	var ids string

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete cloud phones",
		Long:    "Batch delete cloud phones by IDs (max 100). Cloud phones must be stopped first.",
		Example: `  geelark-cli phone delete --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"ids": strings.Split(ids, ","),
			}

			result, err := c.PostAndPrint("/open/v1/phone/delete", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs, max 200 (required)")
	_ = cmd.MarkFlagRequired("ids")

	return cmd
}

func newRestartCmd(newClient clientFactory) *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "restart",
		Short: "Restart a cloud phone",
		Long:  "Restart a cloud phone. Ensure the cloud phone startup callback has been received before calling.",
		Example: `  geelark-cli phone restart --id "631490227545875981"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id": id,
			}

			result, err := c.PostAndPrint("/open/v1/phone/restart", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newStatusCmd(newClient clientFactory) *cobra.Command {
	var ids string

	cmd := &cobra.Command{
		Use:     "status",
		Short:   "Query cloud phone status",
		Long:    "Retrieve the status of cloud phones by IDs (max 100).",
		Example: `  geelark-cli phone status --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"ids": strings.Split(ids, ","),
			}

			result, err := c.PostAndPrint("/open/v1/phone/status", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs, max 100 (required)")
	_ = cmd.MarkFlagRequired("ids")

	return cmd
}

func newCreateCmd(newClient clientFactory) *cobra.Command {
	var mobileType string
	var chargeMode int
	var region string
	var dataJSON string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new cloud phones",
		Long: `Create new cloud phones with specified parameters.

The --data flag accepts a JSON array of environment parameters.`,
		Example: `  geelark-cli phone create --region "sgp" --mobile-type "Android 12" --data "[{\"profileName\":\"myPhone\",\"proxyInformation\":\"socks5://user:pass@1.2.3.4:1080\"}]"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			if region != "cn" && region != "sgp" && region != "us" {
				return fmt.Errorf("invalid --region %q: must be one of cn, sgp, us", region)
			}

			var envData []interface{}
			if err := json.Unmarshal([]byte(dataJSON), &envData); err != nil {
				return fmt.Errorf("invalid --data JSON: %w", err)
			}

			body := map[string]interface{}{
				"mobileType": mobileType,
				"data":       envData,
			}
			if chargeMode >= 0 {
				body["chargeMode"] = chargeMode
			}
			body["region"] = region

			result, err := c.PostAndPrint("/open/v1/phone/addNew", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&mobileType, "mobile-type", "Android 12", "Cloud phone type (Android 9/10/11/12/13/14/15/16)")
	cmd.Flags().IntVar(&chargeMode, "charge-mode", 0, "Billing mode: 0=on-demand, 1=monthly")
	cmd.Flags().StringVar(&region, "region", "", "Region: cn, sgp, us (required)")
	cmd.Flags().StringVar(&dataJSON, "data", "", "JSON array of environment parameters (required)")
	_ = cmd.MarkFlagRequired("data")
	_ = cmd.MarkFlagRequired("region")

	return cmd
}

func newUpdateCmd(newClient clientFactory) *cobra.Command {
	var id, name, remarkStr, groupID, proxyID string
	var tagIDs string
	var dataJSON string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update cloud phone information",
		Long:  "Modify cloud phone name, remark, tags, proxy, group, etc.",
		Example: `  geelark-cli phone update --id "phone_id" --name "new name" --remark "new remark"
  geelark-cli phone update --id "phone_id" --tag-ids "tag1,tag2" --group-id "group_id"
  geelark-cli phone update --id "phone_id" --data "{\"proxyConfig\":{\"typeId\":1,\"server\":\"1.2.3.4\",\"port\":1080,\"username\":\"u\",\"password\":\"p\"}}"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id": id,
			}
			if name != "" {
				body["name"] = name
			}
			if remarkStr != "" {
				body["remark"] = remarkStr
			}
			if groupID != "" {
				body["groupID"] = groupID
			}
			if tagIDs != "" {
				body["tagIDs"] = strings.Split(tagIDs, ",")
			}
			if proxyID != "" {
				body["proxyId"] = proxyID
			}
			if dataJSON != "" {
				var extra map[string]interface{}
				if err := json.Unmarshal([]byte(dataJSON), &extra); err != nil {
					return fmt.Errorf("invalid --data JSON: %w", err)
				}
				for k, v := range extra {
					body[k] = v
				}
			}

			result, err := c.PostAndPrint("/open/v1/phone/detail/update", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "New cloud phone name")
	cmd.Flags().StringVar(&remarkStr, "remark", "", "New cloud phone remark")
	cmd.Flags().StringVar(&groupID, "group-id", "", "New group ID")
	cmd.Flags().StringVar(&tagIDs, "tag-ids", "", "Comma-separated tag IDs")
	cmd.Flags().StringVar(&proxyID, "proxy-id", "", "Proxy ID")
	cmd.Flags().StringVar(&dataJSON, "data", "", "Additional JSON data (e.g. proxyConfig)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newCloneCmd(newClient clientFactory) *cobra.Command {
	var envID, groupID string
	var amount int
	var cloneName, cloneRemark, cloneTag, cloneProxy, cloneNetType bool

	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone a cloud phone",
		Long: `Generate new cloud phones by cloning an existing one.
Retains country, timezone, language, and GPS information.
Applications and data will be cleared.`,
		Example: `  geelark-cli phone clone --env-id "phone_id" --amount 2
  geelark-cli phone clone --env-id "phone_id" --amount 1 --group-id "group_id" --clone-name --clone-proxy
  geelark-cli phone clone --env-id "phone_id" --amount 3 --clone-name --clone-remark --clone-tag --clone-proxy --clone-net-type`,
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
			if cloneNetType {
				body["cloneNetType"] = true
			}

			result, err := c.PostAndPrint("/open/v1/phone/clone", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envID, "env-id", "", "Cloud phone ID to clone (required)")
	cmd.Flags().IntVar(&amount, "amount", 1, "Number of clones (1-100)")
	cmd.Flags().StringVar(&groupID, "group-id", "", "Target group ID (ungrouped if not specified)")
	cmd.Flags().BoolVar(&cloneName, "clone-name", false, "Clone the name")
	cmd.Flags().BoolVar(&cloneRemark, "clone-remark", false, "Clone the remark")
	cmd.Flags().BoolVar(&cloneTag, "clone-tag", false, "Clone the tags")
	cmd.Flags().BoolVar(&cloneProxy, "clone-proxy", false, "Clone the proxy")
	cmd.Flags().BoolVar(&cloneNetType, "clone-net-type", false, "Clone the network type")
	_ = cmd.MarkFlagRequired("env-id")

	return cmd
}

func newScreenshotCmd(newClient clientFactory) *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "screenshot",
		Short: "Take a screenshot of a cloud phone",
		Long: `Get a screenshot from a running cloud phone.
Returns a task ID that can be used to query the screenshot result via callback.`,
		Example: `  geelark-cli phone screenshot --id "phone_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id": id,
			}

			result, err := c.PostAndPrint("/open/v1/phone/screenShot", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newScreenshotResultCmd(newClient clientFactory) *cobra.Command {
	var taskID string

	cmd := &cobra.Command{
		Use:   "screenshot-result",
		Short: "Get screenshot task result",
		Long: `Query the status of a cloud phone screenshot task.
After requesting a screenshot, use this to get the result within 30 minutes.
Status: 0=failed, 1=in progress, 2=succeeded, 3=execution failed.`,
		Example: `  geelark-cli phone screenshot-result --task-id "task_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"taskId": taskID,
			}

			result, err := c.PostAndPrint("/open/v1/phone/screenShot/result", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&taskID, "task-id", "", "Screenshot task ID (required, returned by screenshot command)")
	_ = cmd.MarkFlagRequired("task-id")

	return cmd
}

func newResetCmd(newClient clientFactory) *cobra.Command {
	var id, mobileType string
	var changeBrandModel, keepNetType, keepPhoneNumber, keepRegion, keepLanguage bool

	cmd := &cobra.Command{
		Use:   "new-one",
		Short: "One-click new machine (reset cloud phone identity)",
		Long: `Generate a new cloud phone identity. Applications and data will be cleared.
Optionally preserve network type, phone number, region, and language.
By default, brand/model is randomized.`,
		Example: `  geelark-cli phone new-one --id "phone_id"
  geelark-cli phone new-one --id "phone_id" --keep-region --keep-language --keep-phone-number
  geelark-cli phone new-one --id "phone_id" --no-change-brand
  geelark-cli phone new-one --id "phone_id" --mobile-type "Android 16"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"id": id,
			}
			if cmd.Flags().Changed("no-change-brand") {
				body["changeBrandModel"] = !changeBrandModel
			}
			if keepNetType {
				body["keepNetType"] = true
			}
			if keepPhoneNumber {
				body["keepPhoneNumber"] = true
			}
			if keepRegion {
				body["keepRegion"] = true
			}
			if keepLanguage {
				body["keepLanguage"] = true
			}
			if mobileType != "" {
				body["mobileType"] = mobileType
			}

			result, err := c.PostAndPrint("/open/v2/phone/newOne", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().BoolVar(&changeBrandModel, "no-change-brand", false, "Do NOT randomize brand/model (keep unchanged)")
	cmd.Flags().BoolVar(&keepNetType, "keep-net-type", false, "Preserve network connection type")
	cmd.Flags().BoolVar(&keepPhoneNumber, "keep-phone-number", false, "Preserve phone number")
	cmd.Flags().BoolVar(&keepRegion, "keep-region", false, "Preserve region (otherwise follows proxy)")
	cmd.Flags().BoolVar(&keepLanguage, "keep-language", false, "Preserve language (otherwise defaults to English)")
	cmd.Flags().StringVar(&mobileType, "mobile-type", "", "Change mobile type: Android 9, 10, 11, 12, 13, 14, 15, 16")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newGetGPSCmd(newClient clientFactory) *cobra.Command {
	var ids string

	cmd := &cobra.Command{
		Use:     "get-gps",
		Short:   "Get GPS information of cloud phones",
		Long:    "Query the GPS information (latitude and longitude) of cloud phones.",
		Example: `  geelark-cli phone get-gps --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"ids": strings.Split(ids, ","),
			}

			result, err := c.PostAndPrint("/open/v1/phone/gps/get", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs (required)")
	_ = cmd.MarkFlagRequired("ids")

	return cmd
}

func newSetGPSCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string

	cmd := &cobra.Command{
		Use:   "set-gps",
		Short: "Set GPS for cloud phones",
		Long: `Set/update the GPS information of cloud phones (batch supported).
Longitude range: [-180.0, 180.0], Latitude range: [-90.0, 90.0]
Not supported on Android 16.`,
		Example: `  geelark-cli phone set-gps --data "[{\"id\":\"phone_id\",\"latitude\":1.302,\"longitude\":103.875}]"
  geelark-cli phone set-gps --data "[{\"id\":\"id1\",\"latitude\":1.302,\"longitude\":103.875},{\"id\":\"id2\",\"latitude\":11.302,\"longitude\":104.875}]"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			var list []interface{}
			if err := json.Unmarshal([]byte(dataJSON), &list); err != nil {
				return fmt.Errorf("invalid --data JSON: %w", err)
			}

			body := map[string]interface{}{
				"list": list,
			}

			result, err := c.PostAndPrint("/open/v1/phone/gps/set", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&dataJSON, "data", "", `JSON array of GPS data [{"id":"...","latitude":0.0,"longitude":0.0}] (required)`)
	_ = cmd.MarkFlagRequired("data")

	// suppress unused import warning
	_ = output.PrintSuccess

	return cmd
}

func newSetRootCmd(newClient clientFactory) *cobra.Command {
	var ids string
	var open bool

	cmd := &cobra.Command{
		Use:   "set-root",
		Short: "Set root status on cloud phones",
		Long: `Enable or disable root on cloud phones.
Supports Android 12/13/14/15/16. Cloud phone must be started first.`,
		Example: `  geelark-cli phone set-root --ids "id1,id2" --open
  geelark-cli phone set-root --ids "id1" --open=false`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"ids":  strings.Split(ids, ","),
				"open": open,
			}
			result, err := c.PostAndPrint("/open/v1/root/setStatus", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs (required)")
	cmd.Flags().BoolVar(&open, "open", true, "Enable (true) or disable (false) root")
	_ = cmd.MarkFlagRequired("ids")
	return cmd
}

func newGetDeviceIDCmd(newClient clientFactory) *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:     "get-device-id",
		Short:   "Get cloud phone device ID",
		Long:    "Get the cloud phone unique hardware device ID (Android_ID / serialno).",
		Example: `  geelark-cli phone get-device-id --id "phone_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"id": id}
			result, err := c.PostAndPrint("/open/v1/phone/serialNum/get", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func newSendSMSCmd(newClient clientFactory) *cobra.Command {
	var id, phoneNumber, text string

	cmd := &cobra.Command{
		Use:   "send-sms",
		Short: "Send SMS to a cloud phone",
		Long: `Send SMS to a cloud phone. Cloud phone must be started first.
Supports Android 12/13/14/15.`,
		Example: `  geelark-cli phone send-sms --id "phone_id" --phone-number "+17723504471" --text "your code: 6666"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"id":          id,
				"phoneNumber": phoneNumber,
				"text":        text,
			}
			result, err := c.PostAndPrint("/open/v1/phone/sendSms", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "Phone number with country code (required)")
	cmd.Flags().StringVar(&text, "text", "", "SMS content (required)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("phone-number")
	_ = cmd.MarkFlagRequired("text")
	return cmd
}

func newBrandListCmd(newClient clientFactory) *cobra.Command {
	var androidVer int

	cmd := &cobra.Command{
		Use:     "brand-list",
		Short:   "List cloud phone brands and models",
		Long:    "Get the list of supported phone brands and models for a given Android version.",
		Example: `  geelark-cli phone brand-list --android-ver 12`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"androidVer": androidVer}
			result, err := c.PostAndPrint("/open/v1/phone/brand/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&androidVer, "android-ver", 12, "Android version (10-16)")
	_ = cmd.MarkFlagRequired("android-ver")
	return cmd
}

func newBrandTeamListCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize, androidVer int

	cmd := &cobra.Command{
		Use:     "brand-team-list",
		Short:   "List team-uploaded cloud phone brands and models",
		Long:    "Get the list of custom uploaded team phone brands and models for a given Android version.",
		Example: `  geelark-cli phone brand-team-list --android-ver 10 --page 1 --page-size 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"page":       page,
				"pageSize":   pageSize,
				"androidVer": androidVer,
			}
			result, err := c.PostAndPrint("/open/v1/phone/brand/teamList", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number, min 1")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Page size, 1-100")
	cmd.Flags().IntVar(&androidVer, "android-ver", 12, "Android version (9/10/11/12/13/15)")
	_ = cmd.MarkFlagRequired("android-ver")
	return cmd
}

func newTransferCmd(newClient clientFactory) *cobra.Command {
	var ids, account, transferOption string

	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "Transfer cloud phones to another account",
		Long:  "Transfer cloud phones to a target account. Max 200 phones per request.",
		Example: `  geelark-cli phone transfer --ids "id1,id2" --account "user@geelark.com"
  geelark-cli phone transfer --ids "id1" --account "user@geelark.com" --transfer-option "name,proxy,tag,remark"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"ids":     strings.Split(ids, ","),
				"account": account,
			}
			if transferOption != "" {
				body["transferOption"] = strings.Split(transferOption, ",")
			}
			result, err := c.PostAndPrint("/open/v1/phone/transfer", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs, max 200 (required)")
	cmd.Flags().StringVar(&account, "account", "", "Target account email (required)")
	cmd.Flags().StringVar(&transferOption, "transfer-option", "", "Comma-separated transfer options: name,proxy,tag,remark,files")
	_ = cmd.MarkFlagRequired("ids")
	_ = cmd.MarkFlagRequired("account")
	return cmd
}

func newSetNetTypeCmd(newClient clientFactory) *cobra.Command {
	var id, wifiId string
	var netType int

	cmd := &cobra.Command{
		Use:   "set-net-type",
		Short: "Set cloud phone network type",
		Long: `Set the network connection mode of a cloud phone.
0=Wi-Fi, 1=Mobile. Only supported on Android 12/13/15.`,
		Example: `  geelark-cli phone set-net-type --id "phone_id" --net-type 0
  geelark-cli phone set-net-type --id "phone_id" --net-type 0 --wifi-id "TP-link"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"id":      id,
				"netType": netType,
			}
			if wifiId != "" {
				body["wifiId"] = wifiId
			}
			result, err := c.PostAndPrint("/open/v1/phone/net/set", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().IntVar(&netType, "net-type", 0, "Network type: 0=Wi-Fi, 1=Mobile")
	cmd.Flags().StringVar(&wifiId, "wifi-id", "", "Wi-Fi name, max 16 chars")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("net-type")
	return cmd
}

func newHideAccessibilityCmd(newClient clientFactory) *cobra.Command {
	var ids, pkgName string

	cmd := &cobra.Command{
		Use:   "hide-accessibility",
		Short: "Hide accessibility in apps",
		Long: `Hide the cloud phone accessibility service from specified apps.
Supports Android 12/13/15. Overwrites previous configuration.`,
		Example: `  geelark-cli phone hide-accessibility --ids "id1,id2" --pkg-name "com.zhiliaoapp.musically"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"ids":     strings.Split(ids, ","),
				"pkgName": strings.Split(pkgName, ","),
			}
			result, err := c.PostAndPrint("/open/v1/phone/hideAccessibility", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated cloud phone IDs (required)")
	cmd.Flags().StringVar(&pkgName, "pkg-name", "", "Comma-separated app package names (required)")
	_ = cmd.MarkFlagRequired("ids")
	_ = cmd.MarkFlagRequired("pkg-name")
	return cmd
}

func newMoveGroupCmd(newClient clientFactory) *cobra.Command {
	var envIds, groupID string

	cmd := &cobra.Command{
		Use:   "move-group",
		Short: "Move cloud phones to a group",
		Long:  `Move cloud phones to a specified group. Pass group-id "0" to move to ungrouped. Max 100 phones.`,
		Example: `  geelark-cli phone move-group --env-ids "id1,id2" --group-id "group_id"
  geelark-cli phone move-group --env-ids "id1" --group-id "0"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"envIds":  strings.Split(envIds, ","),
				"groupId": groupID,
			}
			result, err := c.PostAndPrint("/open/v1/phone/moveGroup", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&envIds, "env-ids", "", "Comma-separated cloud phone IDs, max 100 (required)")
	cmd.Flags().StringVar(&groupID, "group-id", "", "Target group ID, use '0' for ungrouped (required)")
	_ = cmd.MarkFlagRequired("env-ids")
	_ = cmd.MarkFlagRequired("group-id")
	return cmd
}

func newImportContactsCmd(newClient clientFactory) *cobra.Command {
	var id, dataJSON string

	cmd := &cobra.Command{
		Use:     "import-contacts",
		Short:   "Import contacts to a cloud phone",
		Long:    "Batch import contacts to a cloud phone. Returns a task ID for querying the result.",
		Example: `  geelark-cli phone import-contacts --id "phone_id" --data "[{\"firstName\":\"Jay\",\"mobile\":\"13288888888\"}]"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			var contacts []interface{}
			if err := json.Unmarshal([]byte(dataJSON), &contacts); err != nil {
				return fmt.Errorf("invalid --data JSON: %w", err)
			}
			body := map[string]interface{}{
				"id":       id,
				"contacts": contacts,
			}
			result, err := c.PostAndPrint("/open/v1/phone/importContacts", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&dataJSON, "data", "", `JSON array of contacts [{"firstName":"...","mobile":"..."}] (required)`)
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("data")
	return cmd
}

func newImportContactsResultCmd(newClient clientFactory) *cobra.Command {
	var taskID string

	cmd := &cobra.Command{
		Use:   "import-contacts-result",
		Short: "Get import contacts task result",
		Long: `Query the status of a contact import task.
Status: 1=in progress, 2=successful, 3=failed. Valid for 1 hour after creation.`,
		Example: `  geelark-cli phone import-contacts-result --task-id "task_id"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"taskId": taskID}
			result, err := c.PostAndPrint("/open/v1/phone/importContactsResult", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&taskID, "task-id", "", "Import contacts task ID (required)")
	_ = cmd.MarkFlagRequired("task-id")
	return cmd
}

func newNetConfigGetCmd(newClient clientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "net-config-get",
		Short:   "Get cloud phone network config",
		Long:    "Get cloud phone network settings including access blacklist.",
		Example: `  geelark-cli phone net-config-get`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostAndPrint("/open/v1/phone/netConfig/get", nil)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func newNetConfigSetCmd(newClient clientFactory) *cobra.Command {
	var blackList string

	cmd := &cobra.Command{
		Use:   "net-config-set",
		Short: "Set cloud phone network config",
		Long: `Modify cloud phone network settings including access blacklist.
Max 3 blacklisted domains. Supports Android 9/10/11/12/13/15.`,
		Example: `  geelark-cli phone net-config-set --blacklist "a.com,b.com"
  geelark-cli phone net-config-set --blacklist ""`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{}
			if blackList != "" {
				body["blackList"] = strings.Split(blackList, ",")
			} else {
				body["blackList"] = []string{}
			}
			result, err := c.PostAndPrint("/open/v1/phone/netConfig/set", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&blackList, "blacklist", "", "Comma-separated blacklisted domains (max 3, empty to clear)")
	return cmd
}

func newSimpleCreateCmd(newClient clientFactory) *cobra.Command {
	var mobileType, region, profileName, proxyInformation string
	var profileGroup, profileTags, profileNote, phoneNumber, phoneName string
	var chargeMode, proxyQueryChannel, proxyNumber, netType int
	var chargeModeSet, proxyQueryChannelSet, proxyNumberSet, netTypeSet bool

	cmd := &cobra.Command{
		Use:   "simple-create",
		Short: "Quick create a single cloud phone",
		Long: `Simplified command to create a single cloud phone with flat flags.
Use 'create' for batch creation or advanced parameters.`,
		Example: `geelark-cli phone simple-create --region "sgp" --mobile-type "Android 12" --profile-name "myPhone" --proxy-information "socks5://user:pass@1.2.3.4:1080"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			if region != "cn" && region != "sgp" && region != "us" {
				return fmt.Errorf("invalid --region %q: must be one of cn, sgp, us", region)
			}

			if proxyInformation == "" && !proxyNumberSet {
				return fmt.Errorf("at least one of --proxy-information or --proxy-number is required")
			}
			if proxyNumberSet && proxyNumber < 0 {
				return fmt.Errorf("invalid --proxy-number %d: must be a non-negative integer", proxyNumber)
			}

			envItem := map[string]interface{}{
				"profileName": profileName,
			}
			if proxyInformation != "" {
				envItem["proxyInformation"] = proxyInformation
			}
			if proxyQueryChannelSet {
				envItem["proxyQueryChannel"] = proxyQueryChannel
			}
			if proxyNumberSet {
				envItem["proxyNumber"] = proxyNumber
			}
			if profileGroup != "" {
				envItem["profileGroup"] = profileGroup
			}
			if profileTags != "" {
				envItem["profileTags"] = strings.Split(profileTags, ",")
			}
			if profileNote != "" {
				envItem["profileNote"] = profileNote
			}
			if netTypeSet {
				envItem["netType"] = netType
			}
			if phoneNumber != "" {
				envItem["phoneNumber"] = phoneNumber
			}
			if phoneName != "" {
				envItem["phoneName"] = phoneName
			}

			body := map[string]interface{}{
				"mobileType": mobileType,
				"data":       []interface{}{envItem},
			}
			if chargeModeSet {
				body["chargeMode"] = chargeMode
			}
			body["region"] = region

			result, err := c.PostAndPrint("/open/v1/phone/addNew", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&profileName, "profile-name", "", "Cloud phone name (required)")
	cmd.Flags().StringVar(&mobileType, "mobile-type", "Android 12", "Cloud phone type (Android 9/10/11/12/13/14/15/16)")
	cmd.Flags().IntVar(&chargeMode, "charge-mode", 0, "Billing mode: 0=on-demand, 1=monthly")
	cmd.Flags().StringVar(&region, "region", "", "Region: cn, sgp, us (required)")
	cmd.Flags().StringVar(&proxyInformation, "proxy-information", "", "Proxy info, e.g. socks5://user:pass@host:port")
	cmd.Flags().IntVar(&proxyQueryChannel, "proxy-query-channel", 2, "Proxy detection channel: 1=ip-api, 2=IP2Location")
	cmd.Flags().IntVar(&proxyNumber, "proxy-number", 0, "Serial number of an added proxy")
	cmd.Flags().StringVar(&profileGroup, "profile-group", "", "Group name (auto-created if not exists)")
	cmd.Flags().StringVar(&profileTags, "profile-tags", "", "Comma-separated tag names (auto-created if not exists)")
	cmd.Flags().StringVar(&profileNote, "profile-note", "", "Remark/note for the cloud phone")
	cmd.Flags().IntVar(&netType, "net-type", -1, "Network type: 0=Wi-Fi, 1=Mobile (Android 12/13/15 only)")
	cmd.Flags().StringVar(&phoneNumber, "phone-number", "", "Custom phone number (auto-generated if empty)")
	cmd.Flags().StringVar(&phoneName, "phone-name", "", "Device name (auto-generated if empty, not supported on Android 9/11)")
	_ = cmd.MarkFlagRequired("profile-name")
	_ = cmd.MarkFlagRequired("region")

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		chargeModeSet = cmd.Flags().Changed("charge-mode")
		proxyQueryChannelSet = cmd.Flags().Changed("proxy-query-channel")
		proxyNumberSet = cmd.Flags().Changed("proxy-number")
		netTypeSet = cmd.Flags().Changed("net-type")
		return nil
	}

	return cmd
}
