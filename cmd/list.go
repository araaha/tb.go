package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List items by board",
	Example: "tb list @recipes @netflix",
	Run: func(_ *cobra.Command, args []string) {
		list(args)
	},
	ValidArgsFunction: func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllBoard(false), cobra.ShellCompDirectiveNoFileComp
	},
}

// list displays board by args. If args is empty, it displays every board.
func list(args []string) {
	if len(args) == 0 {
		taskBook.DisplayByBoard()
		return
	}
	taskBook.DisplayByBoardList(args)
}

func init() {
	rootCmd.AddCommand(listCmd)
}
