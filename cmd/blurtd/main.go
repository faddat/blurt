package main

import (
	"os"

	"github.con/faddat/blurt/cmd/blurtd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
