package plugins

import (
	"github.com/spf13/cobra"
)

func NewConnectorPluginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugins",
		Short: "Commands for connector plugin",
	}

	cmd.AddCommand(newConnectorPluginListCmd())
	return cmd
}
