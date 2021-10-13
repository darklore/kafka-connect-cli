package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// connectorCmd represents the connector command
var connectorCmd = &cobra.Command{
	Use:   "connector",
	Short: "Commands for connectors",
}

var connectorListCmd = &cobra.Command{
	Use:   "list",
	Short: "List connectors' name",
	RunE:  connectorListCmdDo,
}

func init() {
	rootCmd.AddCommand(connectorCmd)
	connectorCmd.AddCommand(connectorListCmd)
}

func connectorListCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connectors, err := connect.GetConnectorNames(endpoint)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(connectors); err != nil {
		return err
	}

	return nil
}
