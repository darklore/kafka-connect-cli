package offsets

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func NewOffsetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "offsets",
		Short: "Commands for connector offfsets",
	}

	cmd.AddCommand(newGetOffsetsCmd())
	cmd.AddCommand(newResetOffsetCmd())
	return cmd
}

func newGetOffsetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "List connectors' offsets",
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

			offsets, err := client.GetOffsets(connector)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(offsets); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func newResetOffsetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset connectors' offsets",
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

			if err := client.ResetOffsets(connector); err != nil {
				return err
			}

			fmt.Println("reset called")
			return nil
		},
	}
	return cmd
}
