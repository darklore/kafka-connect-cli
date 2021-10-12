package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// connectorPluginCmd represents the connectorPlugin command
var connectorPluginCmd = &cobra.Command{
	Use:   "connector-plugin",
	Short: "SubCommand for connector plugin",
}

var connectorPluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List connector plugins",
	RunE:  listConnectorPluginCmdDo,
}

func init() {
	rootCmd.AddCommand(connectorPluginCmd)
	connectorPluginCmd.AddCommand(connectorPluginListCmd)
}

func listConnectorPluginCmdDo(cmd *cobra.Command, args []string) error {
	plugins, err := connect.ListConnectorPlugin(endpoint)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(plugins); err != nil {
		return err
	}

	return nil
}
