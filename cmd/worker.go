package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Get a connect worker information",
	RunE:  workerCmdRun,
}

func init() {
	rootCmd.AddCommand(workerCmd)
}

func workerCmdRun(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
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
}
