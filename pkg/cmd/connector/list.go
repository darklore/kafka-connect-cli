package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List connectors' name",
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := util.GetOpenApiClientConfig(cmd)
			if err != nil {
				return err
			}

			connectors, err := connect.ListConnectors(cfg)
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
