package cmd

import (
	"fmt"

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
	topics, err := connect.ListTopics(endpoint, connectorName)
	if err != nil {
		return err
	}
	for _, t := range topics {
		fmt.Println(t)
	}

	return nil
}
