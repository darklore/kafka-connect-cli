package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newTopicsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topics",
		Short: "List connector topic names",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}
			connector, err := cmd.Flags().GetString("connector")
			if err != nil {
				return err
			}

			topics, err := connect.ListTopics(endpoint, connector)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(topics); err != nil {
				return err
			}

			return nil
		},
	}

	setConnectorFlag(cmd)
	return cmd
}
