package connector

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
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
			cfg, err := util.GetOpenApiClientConfig(cmd)
			if err != nil {
				return err
			}

			connector := args[0]
			includeTasks, err := cmd.Flags().GetBool("include-tasks")
			if err != nil {
				return err
			}
			onlyFailed, err := cmd.Flags().GetBool("only-failed")
			if err != nil {
				return err
			}

			if err := connect.RestartConnectorOpenApi(cfg, connector, includeTasks, onlyFailed); err != nil {
				return err
			}

			fmt.Println("restart called")
			return nil
		},
	}

	cmd.Flags().Bool("include-tasks", false, "Restart task instances too")
	cmd.Flags().Bool("only-failed", false, "Restart only FAILED instances")
	return cmd
}
