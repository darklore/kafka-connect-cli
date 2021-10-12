package cmd

import (
	"fmt"

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
	connectors, err := connect.GetConnectorNames(endpoint)
	if err != nil {
		return err
	}

	for _, c := range connectors {
		fmt.Println(c)
	}

	return nil
}
