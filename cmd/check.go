package cmd

import (
	"fmt"
	"strconv"

	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:     "check",
	Short:   "Check/uncheck task",
	Example: "tb check 7 8 9",
	Run: func(_ *cobra.Command, args []string) {
		check(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllID(false), cobra.ShellCompDirectiveNoFileComp
	},
}

// check either checks or unchecks task(s) by updating the givens task(s)
func check(args []string) {
	var completeTasks []string
	var incompleteTasks []string

	if err := validArgs(args, true); err != nil {
		fmt.Println(err)
		return
	}

	for _, arg := range args {
		id, _ := strconv.Atoi(arg)
		_, item := taskBook.GetIndexAndItemByID(id)
		if task, ok := item.(*tb.Task); ok {
			if task.IsComplete {
				incompleteTasks = append(incompleteTasks, arg)
			} else {
				if task.InProgress {
					taskBook.UpdateTask(id, func(task *tb.Task) *tb.Task {
						task.InProgress = !task.InProgress
						return task
					})
				}
				completeTasks = append(completeTasks, arg)
			}
			taskBook.UpdateTask(id, func(task *tb.Task) *tb.Task {
				task.IsComplete = !task.IsComplete
				return task
			})
		} else {
			fmt.Println(tb.ItemIsNote(id))
			return
		}
	}

	fmt.Println(tb.MarkOrUnmarkAttribute(completeTasks, incompleteTasks, "Checked task", "Unchecked task"))
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
