package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// connectorCmd represents the connector command
var connectorCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Commands for connectors",
	RunE:  connectorCreateCmdDo,
}

func init() {
	connectorCmd.AddCommand(connectorCreateCmd)

	connectorCreateCmd.Flags().String("connector-config", "", "File path to connector config file")
	connectorCreateCmd.MarkFlagRequired("connector-config")
	connectorCreateCmd.MarkFlagFilename("connector-config", "json")
}

func connectorCreateCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	config, err := cmd.Flags().GetString("connector-config")
	if err != nil {
		return err
	}

	configFile, err := os.Open(config)
	if err != nil {
		return err
	}
	defer configFile.Close()

	connector, err := connect.CreateConnector(endpoint, configFile)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(connector); err != nil {
		return err
	}

	return nil
}
