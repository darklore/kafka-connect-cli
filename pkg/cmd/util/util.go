package util

import (
	"log"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect/openapi"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func ValidConnectorArgs(cmd *cobra.Command) ([]string, cobra.ShellCompDirective) {
	cfg, err := GetOpenApiClientConfig(cmd)
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, ""))
		return []string{}, cobra.ShellCompDirectiveError
	}

	connectors, err := connect.ListConnectors(cfg)
	if err != nil {
		log.Println(err)
		return []string{}, cobra.ShellCompDirectiveError
	}

	return connectors, cobra.ShellCompDirectiveNoFileComp
}

func AddEndpointSchemeFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String("scheme", "http", "scheme of the endpoint url")
}

func AddEndpointHostFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String("host", "localhost:8083", "Kafka connect REST endpoint")
}

func AddEndpointFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("endpoint", "e", "http://localhost:8083", "Kafka connect REST endpoint")
}

func GetEndpointScheme(cmd *cobra.Command) (string, error) {
	return cmd.InheritedFlags().GetString("scheme")
}

func GetEndpointHost(cmd *cobra.Command) (string, error) {
	return cmd.InheritedFlags().GetString("host")
}

func GetEndpoint(cmd *cobra.Command) (string, error) {
	return cmd.InheritedFlags().GetString("endpoint")
}

func GetOpenApiClientConfig(cmd *cobra.Command) (*openapi.Configuration, error) {
	scheme, err := cmd.InheritedFlags().GetString("scheme")
	if err != nil {
		return nil, err
	}
	host, err := cmd.InheritedFlags().GetString("host")
	if err != nil {
		return nil, err
	}
	cfg := openapi.NewConfiguration()
	cfg.Host = host
	cfg.Scheme = scheme
	return cfg, nil
}

func AddOutputFormatFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("output", "o", "json", "Output format")
}

func GetOutputFormat(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("output")
}
