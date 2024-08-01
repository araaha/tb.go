package cmd

import (
	"fmt"
	"strconv"
	"strings"

	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move item between boards",
	Run:   move,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllID(true), cobra.ShellCompDirectiveNoFileComp
	},
}

// move updates the board(s) an Item belongs to
func move(_ *cobra.Command, args []string) {
	var ids []string
	var boards []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "@") {
			ids = append(ids, arg)
		} else {
			if arg == "" {
				fmt.Println(tb.MissingBoards())
				return
			}
			boards = append(boards, "@"+arg)
		}
	}

	if len(ids) == 0 {
		fmt.Println(tb.MissingID())
		return
	}
	if len(ids) > 1 {
		fmt.Println(tb.InvalidIDArgNumber())
		return
	}
	if len(boards) == 0 {
		fmt.Println(tb.MissingBoards())
		return
	}

	id, err := strconv.Atoi(ids[0][1:])
	if err != nil {
		fmt.Println(tb.InvalidID(ids[0][1:]))
		return
	}

	_, item := taskBook.GetIndexAndItemByID(id)
	if item == nil {
		fmt.Println(tb.InvalidID(id))
		return
	}

	taskBook.Update(id, func(item tb.Item) tb.Item {
		item.GetBaseItem().Boards = boards
		return item
	})

	fmt.Println(tb.ItemMoved(id, boards))
}

func init() {
	rootCmd.AddCommand(moveCmd)
}
