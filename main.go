package main

import (
	_ "embed"
	"main/cmd"
	"main/log"
	"time"
)

var logger = log.Log()
var start = time.Now()

func ec(err error) {
	if err != nil {
		logger.Error(err)
	}
}

func main() {
	cmd.Execute()
}
