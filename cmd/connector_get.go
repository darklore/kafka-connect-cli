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
)

func init() {
	connectorCmd.AddCommand(getCmd)

	setConnectorNameFlag(getCmd)
}

func setConnectorNameFlag(cmd *cobra.Command) error {
	flagName := "connector"
	cmd.Flags().StringP(flagName, "n", "", "connector name")
	if err := cmd.MarkFlagRequired(flagName); err != nil {
		return err
	}
	cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		endpoint, err := getEndpoint(cmd)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}

		connectors, err := connect.GetConnectorNames(endpoint)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}

		return connectors, cobra.ShellCompDirectiveNoFileComp
	})
	return nil
}

func getConnectorName(cmd *cobra.Command) (string, error) {
	connector, err := cmd.Flags().GetString("connector")
	if err != nil {
		return "", err
	}
	return connector, nil
}

func getCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connectorName, err := getConnectorName(cmd)
	if err != nil {
		return err
	}

	connector, err := connect.GetConnector(endpoint, connectorName)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(connector); err != nil {
		return err
	}

	return nil
}
