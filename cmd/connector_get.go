package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get information of connector",
		RunE:  getCmdDo,
	}
	connectorName string
)

func init() {
	connectorCmd.AddCommand(getCmd)

	setConnectorNameFlag(getCmd)
}

func setConnectorNameFlag(cmd *cobra.Command) error {
	cmd.Flags().StringVarP(&connectorName, "name", "n", "", "connector name")
	if err := cmd.MarkFlagRequired("name"); err != nil {
		return err
	}
	cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		connectors, err := connect.GetConnectorNames(endpoint)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}

		return connectors, cobra.ShellCompDirectiveNoFileComp
	})
	return nil
}

func getCmdDo(_ *cobra.Command, args []string) error {
	connector, err := connect.GetConnector(endpoint, connectorName)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(connector); err != nil {
		return err
	}

	return nil
}
