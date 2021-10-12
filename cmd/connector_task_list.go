package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// listTaskCmd represents the listTask command
var listTaskCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks in a connector",
	RunE:  listTaskCmdDo,
}

func init() {
	taskCmd.AddCommand(listTaskCmd)
	setConnectorNameFlag(listTaskCmd)
}

func listTaskCmdDo(_ *cobra.Command, args []string) error {
	tasks, err := connect.ListTasks(endpoint, connectorName)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(tasks); err != nil {
		return err
	}

	return nil
}
