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
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		connectors, err := connect.GetConnectorNames(endpoint)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}

		return connectors, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	connectorCmd.AddCommand(configCmd)
}

func configCmdDo(_ *cobra.Command, args []string) error {
	name := args[0]
	config, err := connect.GetConnectorConfig(endpoint, name)
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
