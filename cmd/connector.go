package cmd

import (
	"github.com/spf13/cobra"
)

// connectorCmd represents the connector command
var connectorCmd = &cobra.Command{
	Use:   "connector",
	Short: "Commands for connectors",
}

func init() {
	rootCmd.AddCommand(connectorCmd)
}
