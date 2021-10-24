package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Commands for connectors",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}

			fileName, err := cmd.Flags().GetString("connector-config")
			if err != nil {
				return err
			}

			configFile, err := os.Open(fileName)
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
		},
	}

	cmd.Flags().String("connector-config", "", "File path to connector config file")
	cmd.MarkFlagRequired("connector-config")
	cmd.MarkFlagFilename("connector-config", "json")
	return cmd
}
