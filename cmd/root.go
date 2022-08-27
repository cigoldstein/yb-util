package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/bin"
	"main/log"
	"os"
)

var logger = log.Log()

var filesFlag []string
var caseNumFlag int = 0
var emailFlag string
var dropzoneIdFlag string

var rootCmd = &cobra.Command{
	Use:   "yb_log_uploader",
	Short: "utility to upload logs to yugabyte support",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var yugatoolCmd = &cobra.Command{
	Use:   "db",
	Short: "yugatool",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Running yugatool")
		bin.ExecYt("")
	},
}

var ywClientCmd = &cobra.Command{
	Use:   "api",
	Short: "Interact with the Yugaware API",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Running yugaware-client")
		bin.ExecYwc("")
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

	rootCmd.AddCommand(yugatoolCmd)
	rootCmd.AddCommand(ywClientCmd)

	//rootCmd.Flags().StringSliceVarP(&filesFlag, "files", "f", nil, "List of files to upload")
	//rootCmd.Flags().IntVarP(&caseNumFlag, "case_num", "c", 0, "Zendesk case number to attach files to (required)")
	//rootCmd.Flags().StringVarP(&emailFlag, "email", "e", "", "Email address of submitter (required)")
	//rootCmd.Flags().StringVar(&dropzoneIdFlag, "dropzone_id", "S4dsLt2meOtq1iWgBhqJYsuEe2nzvYuv03j_Y6LqhY0", "Override default dropzone ID")
	//
	//rootCmd.MarkFlagRequired("case_num")
	//rootCmd.MarkFlagRequired("email")
	//rootCmd.Flags().MarkHidden("dropzone_id")

}

func main() {
}
