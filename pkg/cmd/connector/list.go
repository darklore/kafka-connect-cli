package connector

import (
	"encoding/json"
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
				return err
			}

			connectors, err := client.ListConnectors()
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(connectors); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
