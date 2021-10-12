package cmd

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List connectors' name",
	RunE: func(cmd *cobra.Command, args []string) error {
		connectors, err := connect.GetConnectorNames(endpoint)
		if err != nil {
			return err
		}

		for _, c := range connectors {
			fmt.Println(c)
		}

		return nil
	},
}

func init() {
	connectorCmd.AddCommand(listCmd)
}
