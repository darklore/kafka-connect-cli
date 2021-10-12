package cmd

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information of connector",
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	RunE: getCmdDo,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		connectors, err := connect.GetConnectorNames(endpoint)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}

		return connectors, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	connectorCmd.AddCommand(getCmd)
}

func getCmdDo(_ *cobra.Command, args []string) error {
	name := args[0]
	connector, err := connect.GetConnector(endpoint, name)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(connector); err != nil {
		return err
	}

	return nil
}
