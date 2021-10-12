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
	"fmt"
	"sort"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: configCmdDo,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		connectors, err := connect.GetConnectorNames(endpoint)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		}

		return connectors, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	connectorCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func configCmdDo(_ *cobra.Command, args []string) error {
	name := args[0]
	config, err := connect.GetConnectorConfig(endpoint, name)
	if err != nil {
		return err
	}

	// sort config map by config name
	keys := make([]string, 0, len(*config))
	for key := range *config {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %s\n", k, (*config)[k])
	}

	return nil
}
