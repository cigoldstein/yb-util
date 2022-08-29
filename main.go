package main

import (
	"time"

	"main/cmd"
	"main/log"
)

var logger = log.Log()
var start = time.Now()

func main() {
	cmd.Execute()
}
