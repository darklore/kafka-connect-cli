/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
	worker, err := connect.GetWorker(endpoint)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(worker); err != nil {
		return err
	}

	return nil
}
