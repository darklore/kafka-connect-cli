package tasks

import (
	"fmt"

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
	return cmd
}

// TODO refactor
func setConnectorFlag(cmd *cobra.Command) {
	flagName := "connector"
	cmd.Flags().StringP(flagName, "n", "", "connector name")
	cmd.MarkFlagRequired(flagName)

	cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}

		connectors, err := connect.ListConnectors(endpoint)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}

		return connectors, cobra.ShellCompDirectiveNoFileComp
	})
}

func setTaskIDFlag(cmd *cobra.Command) error {
	flagName := "task"
	cmd.Flags().StringP(flagName, "t", "", "task ID")
	if err := cmd.MarkFlagRequired(flagName); err != nil {
		return err
	}
	cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		connector, err := cmd.Flags().GetString("connector")
		if err != nil || connector == "" {
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

func getTaskID(cmd *cobra.Command) (string, error) {
	id, err := cmd.Flags().GetString("task")
	if err != nil {
		return "", err
	}
	return id, nil
}
