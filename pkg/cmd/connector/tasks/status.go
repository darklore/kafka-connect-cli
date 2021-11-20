package tasks

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get a task status in a connector",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}

			connector, err := cmd.Flags().GetString("connector")
			if err != nil {
				return err
			}

			taskID, err := getTaskID(cmd)
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
		},
	}

	setConnectorFlag(cmd)
	setTaskIDFlag(cmd)
	return cmd
}
