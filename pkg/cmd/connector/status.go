package connector

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func newStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status [connector]",
		Short: "Get a connector's status",
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

			connector := args[0]

			status, err := client.GetConnectorStatus(connector)
			if err != nil {
				return fmt.Errorf("failed to get connector status: %w", err)
			}

			if err := json.NewEncoder(os.Stdout).Encode(status); err != nil {
				return fmt.Errorf("failed to encode status to JSON: %w", err)
			}

			return nil
		},
	}

	return cmd
}
