package cmd

import (
	"fmt"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// connectorRestartCmd represents the restart command
var connectorResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume a paused connector",
	RunE:  connectorResumeCmdDo,
}

func init() {
	connectorCmd.AddCommand(connectorResumeCmd)
	setConnectorNameFlag(connectorResumeCmd)
}

func connectorResumeCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connector, err := getConnectorName(cmd)
	if err != nil {
		return err
	}

	if err := connect.ResumeConnector(endpoint, connector); err != nil {
		return err
	}

	fmt.Println("resume called")
	return nil
}
