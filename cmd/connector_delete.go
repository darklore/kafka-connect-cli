package cmd

import (
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a connector",
	RunE:  deleteCmdDo,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		connectors, err := connect.GetConnectorNames(endpoint)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}

		return connectors, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	connectorCmd.AddCommand(deleteCmd)
}

func deleteCmdDo(_ *cobra.Command, args []string) error {
	name := args[0]
	err := connect.DeleteConnector(endpoint, name)
	if err != nil {
		return err
	}
	return nil
}
