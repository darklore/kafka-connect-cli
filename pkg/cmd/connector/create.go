package connector

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [connector config json]",
		Short: "Create a new connector",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		ValidArgsFunction: func(_ *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				return []string{"json"}, cobra.ShellCompDirectiveFilterFileExt
			}
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := util.GetConnectClient(cmd)
			if err != nil {
				return fmt.Errorf("failed to get connect client: %w", err)
			}

			fileName := args[0]

			configFile, err := os.Open(fileName)
			if err != nil {
				return fmt.Errorf("failed to open config file: %w", err)
			}
			defer configFile.Close()

			connector, err := client.CreateConnector(configFile)
			if err != nil {
				return fmt.Errorf("failed to create connector: %w", err)
			}

			if err := json.NewEncoder(os.Stdout).Encode(connector); err != nil {
				return fmt.Errorf("failed to encode connector to JSON: %w", err)
			}

			return nil
		},
	}

	return cmd
}
