package cmd

import (
	"fmt"
	"strconv"
	"strings"

	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
)

// priorityCmd represents the priority command
var priorityCmd = &cobra.Command{
	Use:     "priority",
	Short:   "Update priority of task",
	Example: "tb priority @2 3",
	Run: func(_ *cobra.Command, args []string) {
		priority(args)
	},
	ValidArgsFunction: func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllID(true, false), cobra.ShellCompDirectiveNoFileComp
	},
}

// priority updates the priority of a Task(s)
func priority(args []string) {
	if len(args) != 2 {
		fmt.Println(tb.InvalidPriority())
		return
	}

	priorityLevels := map[int]string{1: "low", 2: "medium", 3: "high"}

	tmpID, tmpPrio := args[0], args[1]
	if !strings.HasPrefix(tmpID, "@") {
		if !strings.HasPrefix(tmpPrio, "@") {
			fmt.Println(tb.MissingID())
			return
		}
		tmpID, tmpPrio = tmpPrio, tmpID
	}

	id, err := strconv.Atoi(tmpID[1:])
	if err != nil {
		fmt.Println(tb.InvalidID(tmpID[1:]))
		return
	}

	prio, err := strconv.Atoi(tmpPrio)
	if err != nil || prio < 1 || prio > 3 {
		fmt.Println(tb.InvalidPriority())
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

	modifyTask := func(task *tb.Task) *tb.Task {
		task.Priority = prio
		return task
	}

	taskBook.UpdateTask(id, modifyTask)

	fmt.Println(tb.ItemPriority(id, priorityLevels[prio]))
}

func init() {
	rootCmd.AddCommand(priorityCmd)
}
