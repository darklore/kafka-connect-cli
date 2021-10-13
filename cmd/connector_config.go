package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get a connector's config",
	RunE:  configCmdDo,
}

func init() {
	connectorCmd.AddCommand(configCmd)
	setConnectorNameFlag(configCmd)
}

func configCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connector, err := getConnectorName(cmd)
	if err != nil {
		return err
	}

	config, err := connect.GetConnectorConfig(endpoint, connector)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(config); err != nil {
		return err
	}

	return nil
}
