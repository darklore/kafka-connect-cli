package tasks

import (
	"fmt"
	"strconv"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func newRestartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart [connector] [task id]",
		Short: "Restart a task in a connector",
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
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
			client, err := util.GetConnectClient(cmd)
			if err != nil {
				return err
			}

			connector := args[0]
			taskID, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				return err
			}

			if err := client.RestartTask(connector, int32(taskID)); err != nil {
				return err
			}
			fmt.Println("Restart task")
			return nil
		},
	}

	return cmd
}
