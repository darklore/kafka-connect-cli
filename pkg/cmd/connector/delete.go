package connector

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a connector",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}

			connector, err := cmd.Flags().GetString("connector")
			if err != nil {
				return err
			}

			if err := connect.DeleteConnector(endpoint, connector); err != nil {
				return err
			}

			fmt.Println("Deleted.")
			return nil
		},
	}

	setConnectorFlag(cmd)
	return cmd
}
