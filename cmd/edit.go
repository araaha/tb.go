package cmd

import (
	"fmt"
	"strconv"
	"strings"

	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit item description",
	Example: "tb edit @19 Approve pull request",
	Run: func(_ *cobra.Command, args []string) {
		edit(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllID(true), cobra.ShellCompDirectiveNoFileComp
	},
}

// edit updates the description of an Item(s)
func edit(args []string) {
	l := len(args)
	switch {
	case l == 0:
		fmt.Println(tb.MissingID())
		return
	case l == 1:
		if strings.HasPrefix(args[0], "@") {
			fmt.Println(tb.MissingDesc())
		}
		fmt.Println(tb.MissingID())
		return
	case l > 2:
		fmt.Println(tb.InvalidIDArgNumber())
		return
	}

	tmpID, tmpDesc := args[0], args[1]
	if !strings.HasPrefix(tmpID, "@") {
		if !strings.HasPrefix(tmpDesc, "@") {
			fmt.Println(tb.MissingID())
			return
		}
		tmpID, tmpDesc = tmpDesc, tmpID
	}
	desc := tmpDesc
	if len(desc) == 0 {
		fmt.Println(tb.MissingDesc())
		return
	}

	id, err := strconv.Atoi(tmpID[1:])
	if err != nil {
		fmt.Println(tb.InvalidID(tmpID[1:]))
		return
	}

	_, item := taskBook.GetIndexAndItemByID(id)
	if item == nil {
		fmt.Println(tb.InvalidID(id))
		return
	}
	if item.GetBaseItem().IsArchive {
		fmt.Println(tb.ItemAlreadyArchived())
		return
	}

	taskBook.Update(id, func(item tb.Item) tb.Item {
		item.GetBaseItem().Description = desc
		return item
	})

	fmt.Println(tb.ItemEdited(id))
}

func init() {
	rootCmd.AddCommand(editCmd)
}
