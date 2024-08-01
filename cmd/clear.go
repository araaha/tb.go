package cmd

import (
	"fmt"
	"strconv"

	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Delete all checked items",
	Run: func(_ *cobra.Command, _ []string) {
		clear()
	},
}

// clear updates every completed Task as archived
func clear() {
	var ids []string

	for _, item := range taskBook.Items {
		if task, ok := item.(*tb.Task); ok {
			if task.IsComplete && !task.IsArchive {

				taskBook.UpdateTask(task.ID, func(*tb.Task) *tb.Task {
					task.IsArchive = !task.IsArchive
					return task
				})

				ID := strconv.Itoa(task.ID)
				ids = append(ids, ID)
			}
		}
	}

	fmt.Println(tb.ItemDeleted(ids))
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
