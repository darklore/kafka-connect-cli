package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafka-connect-cli",
	Short: "Command line tool around kafka connect REST interface",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("endpoint", "e", "http://localhost:8083", "Kafka connect REST endpoint")
}

func getEndpoint(cmd *cobra.Command) (string, error) {
	endpoint, err := cmd.Root().PersistentFlags().GetString("endpoint")
	if err != nil {
		return "", err
	}
	return endpoint, nil
}
