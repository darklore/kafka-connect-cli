package cmd

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// connectorRestartCmd represents the restart command
var connectorRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart a connector",
	RunE:  connectorRestartCmdDo,
}

func init() {
	connectorCmd.AddCommand(connectorRestartCmd)
	setConnectorNameFlag(connectorRestartCmd)
}

func connectorRestartCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connector, err := getConnectorName(cmd)
	if err != nil {
		return err
	}

	if err := connect.RestartConnector(endpoint, connector); err != nil {
		return err
	}

	fmt.Println("restart called")
	return nil
}
