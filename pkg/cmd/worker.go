package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newWorkerCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "worker",
		Short: "Get a connect worker information",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
			if err != nil {
				return err
			}

			worker, err := connect.GetWorker(endpoint)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(worker); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
