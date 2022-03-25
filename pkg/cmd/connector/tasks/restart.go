package tasks

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newRestartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart [connector] [task id]",
		Short: "Restart a task in a connector",
		Args:  cobra.ExactValidArgs(2),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
			switch len(args) {
			case 0:
				return util.ValidConnectorArgs(cmd)
			case 1:
				connector := args[0]
				return validateTaskArgs(cmd, connector)
			default:
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			endpoint, err := util.GetEndpoint(cmd)
			if err != nil {
				return err
			}

			connector := args[0]
			taskID := args[1]

			if err := connect.RestartTask(endpoint, connector, taskID); err != nil {
				return err
			}
			fmt.Println("Restart task")
			return nil
		},
	}

	return cmd
}
