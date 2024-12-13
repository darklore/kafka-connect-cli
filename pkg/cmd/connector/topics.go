package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newTopicsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topics [connector]",
		Short: "List connector topic names",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				return util.ValidConnectorArgs(cmd)
			}
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := util.GetOpenApiClientConfig(cmd)
			if err != nil {
				return err
			}

			connector := args[0]

			topics, err := connect.ListTopicsOpenApi(cfg, connector)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(topics); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
