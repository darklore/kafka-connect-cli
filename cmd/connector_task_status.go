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

		connector, err := getConnectorName(cmd)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		if connector == "" {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		endpoint, err := getEndpoint(cmd)
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
	})
	return nil
}

func taskStatusCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connector, err := getConnectorName(cmd)
	if err != nil {
		return err
	}

	tasks, err := connect.GetTaskStatus(endpoint, connector, taskID)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(tasks); err != nil {
		return err
	}

	return nil
}
