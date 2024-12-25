package connector

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func newRestartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart [connector]",
		Short: "Restart a connector",
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
			includeTasks, err := cmd.Flags().GetBool("include-tasks")
			if err != nil {
				return fmt.Errorf("failed to get include-tasks flag: %w", err)
			}
			onlyFailed, err := cmd.Flags().GetBool("only-failed")
			if err != nil {
				return fmt.Errorf("failed to get only-failed flag: %w", err)
			}

			if err := client.RestartConnector(connector, includeTasks, onlyFailed); err != nil {
				return fmt.Errorf("failed to restart connector: %w", err)
			}

			fmt.Println("restart called")
			return nil
		},
	}

	cmd.Flags().Bool("include-tasks", false, "Restart task instances too")
	cmd.Flags().Bool("only-failed", false, "Restart only FAILED instances")
	return cmd
}
