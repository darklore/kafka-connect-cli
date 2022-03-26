package util

import (
	"log"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func ValidConnectorArgs(cmd *cobra.Command) ([]string, cobra.ShellCompDirective) {
	endpoint, err := GetEndpoint(cmd)
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, ""))
		return []string{}, cobra.ShellCompDirectiveError
	}

	connectors, err := connect.ListConnectorNames(endpoint)
	if err != nil {
		log.Println(err)
		return []string{}, cobra.ShellCompDirectiveError
	}

	return connectors, cobra.ShellCompDirectiveNoFileComp
}

func AddEndpointFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("endpoint", "e", "http://localhost:8083", "Kafka connect REST endpoint")
}

func GetEndpoint(cmd *cobra.Command) (string, error) {
	return cmd.InheritedFlags().GetString("endpoint")
}

func AddOutputFormatFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("output", "o", "json", "Output format")
}

func GetOutputFormat(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("output")
}
