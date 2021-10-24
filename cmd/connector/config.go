package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Get a connector's config",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}
			connector, err := cmd.Flags().GetString("connector")
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
		},
	}

	setConnectorFlag(cmd)
	return cmd
}
