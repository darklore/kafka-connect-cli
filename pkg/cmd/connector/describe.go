package connector

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd/util"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func newDescribeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe [connector]",
		Short: "Describe information of connector",
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

			connectorName := args[0]

			connector, err := connect.GetConnector(endpoint, connectorName)
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
