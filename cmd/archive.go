package cmd

import (
	"github.com/spf13/cobra"
)

var remove bool

// archiveCmd represents the archive command
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Display archived items",
	Run: func(_ *cobra.Command, _ []string) {
		archive()
	},
}

// archive displays Item(s) by date or removes archived tasks permanently.
func archive() {
	if remove {
		taskBook.Remove()
		return
	}
	taskBook.DisplayByDate(true)
}

func init() {
	rootCmd.AddCommand(archiveCmd)
	archiveCmd.Flags().BoolVar(&remove, "remove", false, "Remove archived tasks permanently")
}
