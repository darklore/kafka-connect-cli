package tasks

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status [connector] [task id]",
		Short: "Get a task status in a connector",
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
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}

			connector := args[0]
			taskID := args[1]

			tasks, err := connect.GetTaskStatus(endpoint, connector, taskID)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(tasks); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
