package cmd

import (
	"github.com/darklore/kafka-connect-cli/pkg/cmd/connector"
	"github.com/darklore/kafka-connect-cli/pkg/cmd/plugins"
	"github.com/darklore/kafka-connect-cli/pkg/cmd/server"
	"github.com/spf13/cobra"
)

func NewCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kccli",
		Short: "Command line tool arounc kakfa connect REST interface",
	}
	cmd.Version = version

	// subcommands
	cmd.AddCommand(newVersionCmd())

	cmd.AddCommand(server.NewServerCmd())
	cmd.AddCommand(connector.NewCmd())
	cmd.AddCommand(plugins.NewConnectorPluginCmd())

	return cmd
}
