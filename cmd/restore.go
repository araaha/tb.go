package cmd

import (
	"fmt"
	"strconv"

	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:     "restore",
	Short:   "Restore items from archive",
	Run:     restore,
	Example: "tb restore 4 5",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllID(false, true), cobra.ShellCompDirectiveNoFileComp
	},
}

// restore restores archived items
func restore(_ *cobra.Command, args []string) {
	var ids []string

	if err := validArgs(args, false); err != nil {
		fmt.Println(err)
		return
	}

	for _, arg := range args {
		id, _ := strconv.Atoi(arg)
		taskBook.Update(id, func(item tb.Item) tb.Item {
			item.GetBaseItem().IsArchive = !item.GetBaseItem().IsArchive
			return item
		})

		ids = append(ids, arg)
	}

	fmt.Println(tb.MarkRestored(ids))
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
