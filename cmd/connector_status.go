package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get a connector's status",
	RunE:  statusCmdDo,
}

func init() {
	connectorCmd.AddCommand(statusCmd)

	setConnectorNameFlag(statusCmd)
}

func statusCmdDo(_ *cobra.Command, args []string) error {
	status, err := connect.GetConnectorStatus(endpoint, connectorName)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(status); err != nil {
		return err
	}

	return nil
}
