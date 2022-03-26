package main

import (
	"log"
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
			log.Fatal("Failed to get build info.")
		}
		version = buildInfo.Main.Version
	}

	c := cmd.NewCmd(version)
	if err := c.Execute(); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
}
