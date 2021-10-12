package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Commands for connector tasks",
}

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks in a connector",
	RunE:  taskListCmdDo,
}

func init() {
	connectorCmd.AddCommand(taskCmd)

	taskCmd.AddCommand(taskListCmd)
	setConnectorNameFlag(taskListCmd)
}

func taskListCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connector, err := getConnectorName(cmd)
	if err != nil {
		return err
	}

	tasks, err := connect.ListTasks(endpoint, connector)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(tasks); err != nil {
		return err
	}

	return nil
}
