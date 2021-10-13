package cmd

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// connectorRestartCmd represents the restart command
var connectorPauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause a connector",
	RunE:  connectorPauseCmdDo,
}

func init() {
	connectorCmd.AddCommand(connectorPauseCmd)
	setConnectorNameFlag(connectorPauseCmd)
}

func connectorPauseCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connector, err := getConnectorName(cmd)
	if err != nil {
		return err
	}

	if err := connect.PauseConnector(endpoint, connector); err != nil {
		return err
	}

	fmt.Println("pause called")
	return nil
}
