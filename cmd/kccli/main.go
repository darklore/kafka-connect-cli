package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/darklore/kafka-connect-cli/pkg/cmd"
)

var (
	version = ""
)

func main() {
	if version == "" {
		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			slog.Error("Failed to get build info.")
		}
		version = buildInfo.Main.Version
	}

	c := cmd.NewCmd(version)
	if err := c.Execute(); err != nil {
		slog.Error(fmt.Sprintf("%+v", err))
		os.Exit(1)
	}
}
