package main

import (
	"fmt"
	"time"

	"github.com/mickeey2525/timeconvert/cmd"
)

func main() {

	layout := "01/02/2006 15:04:05"
	t, err := time.Parse(layout, "02/28/2016 09:03:46")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix())
	cmd.Execute()
}
