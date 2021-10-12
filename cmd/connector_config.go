package cmd

import (
	"fmt"
	"sort"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get a connector's config",
	RunE:  configCmdDo,
}

func init() {
	connectorCmd.AddCommand(configCmd)
	setConnectorNameFlag(configCmd)
}

func configCmdDo(_ *cobra.Command, args []string) error {
	config, err := connect.GetConnectorConfig(endpoint, connectorName)
	if err != nil {
		return err
	}

	// sort config map by config name
	keys := make([]string, 0, len(*config))
	for key := range *config {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %s\n", k, (*config)[k])
	}

	return nil
}
