package plugins

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newConnectorPluginListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List connector plugins",
		RunE: func(cmd *cobra.Command, args []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}

			plugins, err := connect.ListConnectorPlugin(endpoint)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(plugins); err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
