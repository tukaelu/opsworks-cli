package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

var revision string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("ow: OpsWorks CLI Tool version %s (rev:%s)", version, revision))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
