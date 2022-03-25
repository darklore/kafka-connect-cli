package util

import (
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func ValidConnectorArgs(cmd *cobra.Command) ([]string, cobra.ShellCompDirective) {
	endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}

	connectors, err := connect.ListConnectors(endpoint)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}

	return connectors, cobra.ShellCompDirectiveNoFileComp
}
