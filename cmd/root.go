package cmd

import (
	"fmt"
	"os"

	"github.com/blang/vfs"
	"github.com/spf13/cobra"
	ytCmd "github.com/yugabyte/yb-tools/yugatool/cmd"
	ywcCmd "github.com/yugabyte/yb-tools/yugaware-client/cmd"

	"main/log"
)

var logger = log.Log()

var rootCmd = &cobra.Command{
	Use:   "yb-util",
	Short: "Collection of yugabyte utilities",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {

	// import as type "cobra.Command" from github repos
	importYtCmd := ytCmd.RootInit(vfs.OS())
	importYwcCmd := ywcCmd.RootInit()

	// add new cobra commands to our root cobra commands
	rootCmd.AddCommand(importYtCmd)
	rootCmd.AddCommand(importYwcCmd)

}
