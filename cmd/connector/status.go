package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get a connector's status",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}
			connector, err := cmd.Flags().GetString("connector")
			if err != nil {
				return err
			}

			status, err := connect.GetConnectorStatus(endpoint, connector)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(status); err != nil {
				return err
			}

			return nil
		},
	}
	setConnectorFlag(cmd)
	return cmd
}
