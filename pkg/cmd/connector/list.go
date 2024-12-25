package connector

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List connectors' name",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := util.GetConnectClient(cmd)
			if err != nil {
				return fmt.Errorf("failed to get connect client: %w", err)
			}

			connectors, err := client.ListConnectors()
			if err != nil {
				return fmt.Errorf("failed to list connectors: %w", err)
			}

			if err := json.NewEncoder(os.Stdout).Encode(connectors); err != nil {
				return fmt.Errorf("failed to encode connectors to JSON: %w", err)
			}

			return nil
		},
	}

	return cmd
}
