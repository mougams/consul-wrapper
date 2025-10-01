package main

import (
	"log"

	"github.com/tionebsalocin/consul-wrapper/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("command execution failed: %v", err)
	}
}
