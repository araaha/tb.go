package cmd

import (
	"fmt"
	"strconv"

	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
)

// starCmd represents the star command
var starCmd = &cobra.Command{
	Use:     "star",
	Short:   "Star/unstar item",
	Example: "tb star 18 19",
	Run:     star,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllID(false), cobra.ShellCompDirectiveNoFileComp
	},
}

// star either stars or unstars an item(s)
func star(_ *cobra.Command, args []string) {
	var starItem []string
	var unstarItem []string

	if err := validArgs(args, true); err != nil {
		fmt.Println(err)
		return
	}

	for _, arg := range args {
		id, _ := strconv.Atoi(arg)
		_, item := taskBook.GetIndexAndItemByID(id)

		if item.GetBaseItem().IsStarred {
			unstarItem = append(unstarItem, arg)
		} else {
			starItem = append(starItem, arg)
		}
		taskBook.Update(id, func(Item tb.Item) tb.Item {
			item.GetBaseItem().IsStarred = !item.GetBaseItem().IsStarred
			return item
		})
	}

	fmt.Println(tb.MarkOrUnmarkAttribute(starItem, unstarItem, "Starred item", "Unstarred item"))

}

func init() {
	rootCmd.AddCommand(starCmd)
}
