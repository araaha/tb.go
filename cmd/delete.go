package cmd

import (
	"fmt"
	"strconv"

	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete item",
	Run: func(_ *cobra.Command, args []string) {
		delete(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllID(false), cobra.ShellCompDirectiveNoFileComp
	},
}

// delete places Item(s) into the archive
func delete(args []string) {
	var ids []string

	if err := validArgs(args, true); err != nil {
		fmt.Println(err)
		return
	}

	for _, arg := range args {
		id, _ := strconv.Atoi(arg)
		taskBook.Delete(id)

		ids = append(ids, arg)
	}

	fmt.Println(tb.ItemDeleted(ids))
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
