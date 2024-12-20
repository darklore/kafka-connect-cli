package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"

	"github.com/spf13/cobra"
)

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config [connector]",
		Short: "Get a connector's config",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				return util.ValidConnectorArgs(cmd)
			}
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := util.GetConnectClient(cmd)
			if err != nil {
				return err
			}

			connector := args[0]

			config, err := client.GetConnectorConfig(connector)
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
