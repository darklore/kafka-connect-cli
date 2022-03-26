package tasks

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func NewTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tasks",
		Short: "Commands for connector tasks",
	}

	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newRestartCmd())
	cmd.AddCommand(newStatusCmd())
	return cmd
}

func validateTaskArgs(cmd *cobra.Command, connector string) ([]string, cobra.ShellCompDirective) {
	endpoint, err := util.GetEndpoint(cmd)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	tasks, err := connect.ListTasks(endpoint, connector)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	IDs := make([]string, 0, len(tasks))
	for _, t := range tasks {
		IDs = append(IDs, fmt.Sprintf("%d", t.ID.ID))
	}

	return IDs, cobra.ShellCompDirectiveNoFileComp
}
