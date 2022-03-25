package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List connectors' name",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}

			connectors, err := connect.ListConnectorNames(endpoint)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(connectors); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
