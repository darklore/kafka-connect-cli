package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"

	"github.com/spf13/cobra"
)

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config [connector]",
		Short: "Get a connector's config",
		Args:  cobra.ExactValidArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				return util.ValidConnectorArgs(cmd)
			}
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			endpoint, err := util.GetEndpoint(cmd)
			if err != nil {
				return err
			}

			connector := args[0]

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

	return cmd
}
