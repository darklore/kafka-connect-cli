package main

import (
	"fmt"
	"os"

	"github.com/darklore/kafka-connect-cli/pkg/cmd"
)

func main() {
	c := cmd.NewCmd()
	if err := c.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
