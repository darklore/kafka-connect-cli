package connector

import (
	"github.com/darklore/kafka-connect-cli/cmd/connector/tasks"
	"github.com/darklore/kafka-connect-cli/pkg/kafka/connect"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connector",
		Short: "Commands for connectors",
	}

	// add subcommands
	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newConfigCmd())
	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newDeleteCmd())
	cmd.AddCommand(newGetCmd())
	cmd.AddCommand(newPauseCmd())
	cmd.AddCommand(newRestartCmd())
	cmd.AddCommand(newResumeCmd())
	cmd.AddCommand(newStatusCmd())
	cmd.AddCommand(newTopicsCmd())

	cmd.AddCommand(tasks.NewTasksCmd())
	return cmd
}

// TODO refactor
func setConnectorFlag(cmd *cobra.Command) {
	flagName := "connector"
	cmd.Flags().StringP(flagName, "n", "", "connector name")
	cmd.MarkFlagRequired(flagName)

	cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}

		connectors, err := connect.ListConnectors(endpoint)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveError
		}

		return connectors, cobra.ShellCompDirectiveNoFileComp
	})
}
