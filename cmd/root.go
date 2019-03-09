package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "ow",
	Short:   "CLI for AWS OpsWorks",
	Version: fmt.Sprintf("%s (rev:%s)", version, revision),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
