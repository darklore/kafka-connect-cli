package tasks

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [connector]",
		Short: "List tasks in a connector",
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
				return err
			}

			connector := args[0]

			tasks, err := client.GetTaskConfigs(connector)
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
