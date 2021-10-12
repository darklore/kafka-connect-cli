package cmd

import (
	"fmt"

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

func deleteCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connector, err := getConnectorName(cmd)
	if err != nil {
		return err
	}

	if err := connect.DeleteConnector(endpoint, connector); err != nil {
		return err
	}

	fmt.Println("Deleted.")
	return nil
}
