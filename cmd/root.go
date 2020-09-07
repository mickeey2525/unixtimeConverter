package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is root command
var rootCmd = &cobra.Command{
	Use:   "unix2time",
	Short: "unixtime timestamp converter to timestamp",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unix2time was called")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
