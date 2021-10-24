package cmd

import (
	"log"
	"runtime/debug"

	"github.com/darklore/kafka-connect-cli/cmd/connector"
	"github.com/darklore/kafka-connect-cli/cmd/plugins"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kafka-connect-cli",
		Short: "Command line tool arounc kakfa connect REST interface",
	}

	cmd.PersistentFlags().StringP("endpoint", "e", "http://localhost:8083", "Kafka connect REST endpoint")

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("Failed to get build info.")
	}
	cmd.Version = buildInfo.Main.Version

	// subcommands
	cmd.AddCommand(newVersionCmd())

	cmd.AddCommand(newWorkerCmd())
	cmd.AddCommand(connector.NewCmd())
	cmd.AddCommand(plugins.NewConnectorPluginCmd())

	return cmd
}
