package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [connector name] [connector config json]",
		Short: "Update or create a connector",
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
			switch len(args) {
			case 0:
				return util.ValidConnectorArgs(cmd)
			case 1:
				return []string{"json"}, cobra.ShellCompDirectiveFilterFileExt
			default:
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := util.GetConnectClient(cmd)

			if err != nil {
				return err
			}

			connectorName := args[0]
			fileName := args[1]

			configFile, err := os.Open(fileName)
			if err != nil {
				return err
			}
			defer configFile.Close()

			connector, err := client.UpdateConnector(connectorName, configFile)
			if err != nil {
				return err
			}

			if err := json.NewEncoder(os.Stdout).Encode(connector); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
