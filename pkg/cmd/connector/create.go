package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [connector config json]",
		Short: "Create a new connector",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				return []string{"json"}, cobra.ShellCompDirectiveFilterFileExt
			}
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := util.GetOpenApiClientConfig(cmd)
			if err != nil {
				return err
			}

			fileName := args[0]

			configFile, err := os.Open(fileName)
			if err != nil {
				return err
			}
			defer configFile.Close()

			connector, err := connect.CreateConnector2(cfg, configFile)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(connector); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
