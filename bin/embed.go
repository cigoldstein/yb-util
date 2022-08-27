package bin

import (
	_ "embed"
	"fmt"
	"github.com/amenzhinsky/go-memexec"
	"main/log"
)

var logger = log.Log()

//go:embed yugatool
var ytEmbed []byte

//go:embed yugaware-client
var ywcEmbed []byte

func ec(err error) {
	if err != nil {
		logger.Error(err)
	}
}

func ExecYt(args string) {

	yt, err := memexec.New(ytEmbed)
	ec(err)

	defer yt.Close()

	cmd := yt.Command("cluster_info")
	combinedOutput, _ := cmd.CombinedOutput()

	fmt.Println(string(combinedOutput))
}

func ExecYwc(args string) {

	ywc, err := memexec.New(ywcEmbed)
	ec(err)

	defer ywc.Close()

	cmd := ywc.Command(args)
	combinedOutput, _ := cmd.CombinedOutput()

	fmt.Println(string(combinedOutput))
}
