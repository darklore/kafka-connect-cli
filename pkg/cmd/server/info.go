package server

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newServerInfoCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info",
		Short: "Get a connect server information",
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := util.GetOpenApiClientConfig(cmd)
			if err != nil {
				return err
			}

			worker, err := connect.GetServerInfo(cfg)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(worker); err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
