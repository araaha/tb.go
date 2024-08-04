package cmd

import (
	"fmt"

	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
	"strconv"
)

// beginCmd represents the begin command
var beginCmd = &cobra.Command{
	Use:     "begin",
	Short:   "Begin/pause task",
	Example: "tb begin 6 8",
	Run: func(_ *cobra.Command, args []string) {
		begin(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllID(false, false), cobra.ShellCompDirectiveNoFileComp
	},
}

// begin either starts or pauses a Task(s)
func begin(args []string) {
	var startedTasks []string
	var pausedTasks []string

	if err := validArgs(args, true); err != nil {
		fmt.Println(err)
		return
	}

	for _, arg := range args {
		id, _ := strconv.Atoi(arg)
		_, item := taskBook.GetIndexAndItemByID(id)
		if task, ok := item.(*tb.Task); ok {
			if task.InProgress {
				pausedTasks = append(pausedTasks, arg)
			} else {
				if task.IsComplete {
					taskBook.UpdateTask(id, func(task *tb.Task) *tb.Task {
						task.IsComplete = !task.IsComplete
						return task
					})
				}
				startedTasks = append(startedTasks, arg)
			}
			taskBook.UpdateTask(id, func(task *tb.Task) *tb.Task {
				task.InProgress = !task.InProgress
				return task
			})
		} else {
			fmt.Println(tb.ItemIsNote(id))
			return
		}
	}

	fmt.Println(tb.MarkOrUnmarkAttribute(startedTasks, pausedTasks, "Started task", "Paused task"))
}

func init() {
	rootCmd.AddCommand(beginCmd)
}
