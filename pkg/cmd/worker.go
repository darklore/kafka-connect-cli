package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newWorkerCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "worker",
		Short: "Get a connect worker information",
		RunE: func(cmd *cobra.Command, _ []string) error {
			endpoint, err := util.GetEndpoint(cmd)
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

	util.AddEndpointFlag(cmd)
	return cmd
}
