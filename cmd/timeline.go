package cmd

import (
	"github.com/spf13/cobra"
)

// timelineCmd represents the timeline command
var timelineCmd = &cobra.Command{
	Use:   "timeline",
	Short: "Display timeline view",
	Run:   timeline,
}

// timeline displays Items by date
func timeline(_ *cobra.Command, _ []string) {
	taskBook.DisplayByDate(false)
}

func init() {
	rootCmd.AddCommand(timelineCmd)
}
