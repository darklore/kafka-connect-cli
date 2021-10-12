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
}

func init() {
	connectorCmd.AddCommand(deleteCmd)

	setConnectorNameFlag(deleteCmd)
}

func deleteCmdDo(_ *cobra.Command, args []string) error {
	err := connect.DeleteConnector(endpoint, connectorName)
	if err != nil {
		return err
	}
	return nil
}
