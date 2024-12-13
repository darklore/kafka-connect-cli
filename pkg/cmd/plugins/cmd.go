package plugins

import (
	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func NewConnectorPluginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugins",
		Short: "Commands for connector plugin",
	}

	util.AddEndpointHostFlag(cmd)
	util.AddEndpointSchemeFlag(cmd)
	cmd.AddCommand(newConnectorPluginListCmd())
	return cmd
}
