package connector

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// connectorRestartCmd represents the restart command
func newPauseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause",
		Short: "Pause a connector",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}

			connector, err := cmd.Flags().GetString("connector")
			if err != nil {
				return err
			}

			if err := connect.PauseConnector(endpoint, connector); err != nil {
				return err
			}

			fmt.Println("pause called")
			return nil
		},
	}

	setConnectorFlag(cmd)
	return cmd
}
