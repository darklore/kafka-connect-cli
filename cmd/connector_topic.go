package cmd

import (
	"encoding/json"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

// listTopicCmd represents the listTopic command
var listTopicCmd = &cobra.Command{
	Use:   "topics",
	Short: "List connector topic names",
	RunE:  listTopicCmdDo,
}

func init() {
	connectorCmd.AddCommand(listTopicCmd)
	setConnectorNameFlag(listTopicCmd)
}

func listTopicCmdDo(cmd *cobra.Command, args []string) error {
	endpoint, err := getEndpoint(cmd)
	if err != nil {
		return err
	}

	connector, err := getConnectorName(cmd)
	if err != nil {
		return err
	}

	topics, err := connect.ListTopics(endpoint, connector)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(topics); err != nil {
		return err
	}

	return nil
}
