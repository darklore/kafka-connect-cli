package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// listTaskCmd represents the listTask command
var taskStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get a task status in a connector",
	RunE:  taskStatusCmdDo,
}

var taskID int

func init() {
	taskCmd.AddCommand(taskStatusCmd)
	setConnectorNameFlag(taskStatusCmd)
	setTaskIDFlag(taskStatusCmd)
}

func setTaskIDFlag(cmd *cobra.Command) error {
	flagName := "task"
	cmd.Flags().IntVarP(&taskID, flagName, "t", 0, "task ID")
	if err := cmd.MarkFlagRequired(flagName); err != nil {
		return err
	}
	cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if connectorName == "" {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		tasks, err := connect.ListTasks(endpoint, connectorName)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		IDs := make([]string, 0, len(tasks))
		for _, t := range tasks {
			IDs = append(IDs, fmt.Sprintf("%d", t.ID.ID))
		}

		return IDs, cobra.ShellCompDirectiveNoFileComp
	})
	return nil
}

func taskStatusCmdDo(_ *cobra.Command, args []string) error {
	tasks, err := connect.GetTaskStatus(endpoint, connectorName, taskID)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(tasks); err != nil {
		return err
	}

	return nil
}
