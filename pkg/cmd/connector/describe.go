package connector

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func newDescribeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe [connector]",
		Short: "Describe information of connector",
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
				return fmt.Errorf("failed to get connect client: %w", err)
			}

			connectorName := args[0]

			connector, err := client.GetConnector(connectorName)
			if err != nil {
				return fmt.Errorf("failed to get connector: %w", err)
			}

			if err := json.NewEncoder(os.Stdout).Encode(connector); err != nil {
				return fmt.Errorf("failed to encode connector to JSON: %w", err)
			}

			return nil
		},
	}

	return cmd
}
