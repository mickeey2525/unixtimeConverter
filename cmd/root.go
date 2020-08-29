package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// RootCmd is root command
var rootCmd = &cobra.Command{
	Use:   "unix2time",
	Short: "unixtime timestamp converter to timestamp",
	Run: func(cmd *cobra.Command, args []string) {
		var unixtime int64
		if len(args) == 0 {
			unixtime = time.Now().Unix()
		} else {
			var err error
			unixtime, err = strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("failed to parse error happend: %s", err)
			}
		}
		tm := convertUTC(unixtime)
		fmt.Println(tm)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func convertUTC(utime int64) time.Time {
	return time.Unix(utime, 0).UTC()
}
