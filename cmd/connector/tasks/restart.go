package tasks

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newRestartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart",
		Short: "Restart a task in a connector",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}

			connector, err := cmd.Flags().GetString("connector")
			if err != nil {
				return err
			}

			taskID, err := cmd.Flags().GetString("task")
			if err != nil {
				return err
			}

			if err := connect.RestartTask(endpoint, connector, taskID); err != nil {
				return err
			}
			fmt.Println("Restart task")
			return nil
		},
	}

	setConnectorFlag(cmd)
	setTaskIDFlag(cmd)
	return cmd
}
