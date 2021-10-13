package cmd

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

var taskRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart a task in a connector",
	RunE:  taskRestartCmdDo,
}

func init() {
	taskCmd.AddCommand(taskRestartCmd)
	setConnectorNameFlag(taskRestartCmd)
	setTaskIDFlag(taskRestartCmd)
}

func taskRestartCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connector, err := getConnectorName(cmd)
	if err != nil {
		return err
	}

	taskID, err := getTaskID(cmd)
	if err != nil {
		return err
	}

	if err := connect.RestartTask(endpoint, connector, taskID); err != nil {
		return err
	}
	fmt.Println("Restart task")
	return nil
}
