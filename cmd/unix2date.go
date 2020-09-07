/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// unix2dateCmd represents the unix2date command
var unix2dateCmd = &cobra.Command{
	Use:   "unix2date",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

func init() {
	rootCmd.AddCommand(unix2dateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unix2dateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unix2dateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func convertUTC(utime int64) time.Time {
	return time.Unix(utime, 0).UTC()
}
