package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "command",
		Short: "Hello world",
		SilenceUsage: true,
	}

	cmd.AddCommand(gnote())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
