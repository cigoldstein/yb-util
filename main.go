package main

import (
	_ "embed"
	"main/log"
	"time"

	"github.com/amenzhinsky/go-memexec"
)

var logger = log.Log()

//go:embed bin/yugaware-client
var YwcEmbed []byte

//go:embed hello
var helloEmbed []byte

var start = time.Now()

func init() {

}

func ec(err error) {
	if err != nil {
		logger.Error(err)
	}
}

func main() {

	ywc, err := memexec.New(YwcEmbed)
	ec(err)

	defer ywc.Close()

	cmd := ywc.Command("login")
	combinedOutput, _ := cmd.CombinedOutput()

	logger.Info(string(combinedOutput))
	//
	//hello, err := memexec.New(helloEmbed)
	//ec(err)
	//cmd2 := hello.Command()
	//output, err := cmd2.CombinedOutput()
	//ec(err)
	//logger.Info(string(output))

	//logger.Info("Execution time: ", time.Since(start))

}
