package plugins

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func newConnectorPluginListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List connector plugins",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := util.GetConnectClient(cmd)
			if err != nil {
				return err
			}

			plugins, err := client.ListConnectorPlugins()
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
