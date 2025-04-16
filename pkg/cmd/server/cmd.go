package server

import (
	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Commands for kafka connect server",
	}

	util.AddEndpointHostFlag(cmd)
	cmd.AddCommand(newServerInfoCmd())
	return cmd
}
