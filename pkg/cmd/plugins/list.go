package plugins

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newConnectorPluginListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List connector plugins",
		RunE: func(cmd *cobra.Command, _ []string) error {
			// endpoint, err := util.GetEndpoint(cmd)
			cfg, err := util.GetOpenApiClientConfig(cmd)
			if err != nil {
				return err
			}

			plugins, err := connect.ListConnectorPluginOpenApi(cfg)
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
