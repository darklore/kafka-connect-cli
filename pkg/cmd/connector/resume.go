package connector

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func newResumeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resume [connector]",
		Short: "Resume a paused connector",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				return util.ValidConnectorArgs(cmd)
			}
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := util.GetConnectClient(cmd)
			if err != nil {
				return err
			}

			connector := args[0]

			if err := client.ResumeConnector(connector); err != nil {
				return err
			}

			fmt.Println("resume called")
			return nil
		},
	}

	return cmd
}
