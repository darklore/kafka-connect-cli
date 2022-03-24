package connector

import (
	"github.com/darklore/kafka-connect-cli/pkg/cmd/connector/tasks"
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
	cmd.AddCommand(newUpdateCmd())

	cmd.AddCommand(tasks.NewTasksCmd())
	return cmd
}
