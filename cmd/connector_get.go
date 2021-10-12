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
	"errors"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information of connector",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	RunE: getCmdDo,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		connectors, err := connect.GetConnectorNames(endpoint)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}

		return connectors, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	connectorCmd.AddCommand(getCmd)
}

func getCmdDo(_ *cobra.Command, args []string) error {
	name := args[0]
	connector, err := connect.GetConnector(endpoint, name)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(connector); err != nil {
		return err
	}

	return nil
}
