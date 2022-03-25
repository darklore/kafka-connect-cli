package connector

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// connectorRestartCmd represents the restart command
func newPauseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause [connector]",
		Short: "Pause a connector",
		Args:  cobra.ExactValidArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				return util.ValidConnectorArgs(cmd)
			}
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			endpoint, err := util.GetEndpoint(cmd)
			if err != nil {
				return err
			}

			connector := args[0]

			if err := connect.PauseConnector(endpoint, connector); err != nil {
				return err
			}

			fmt.Println("pause called")
			return nil
		},
	}

	return cmd
}
