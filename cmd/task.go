package cmd

import (
	"fmt"
	"strings"

	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
)

var (
	prio string
	str  bool
	comp bool
	prog bool
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Create task",
	Run: func(_ *cobra.Command, args []string) {
		task(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllBoard(true), cobra.ShellCompDirectiveNoFileComp
	},
}

// task adds a Task
func task(args []string) {
	var boards []string
	var tasks []string
	level := map[string]int{"low": 1, "medium": 2, "high": 3}

	if len(args) == 0 {
		fmt.Println(tb.MissingDesc())
		return
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "@") {
			if len(arg) == 1 {
				fmt.Println(tb.MissingBoards())
				return
			}
			boards = append(boards, arg)
		} else {
			tasks = append(tasks, arg)
		}
	}

	if len(tasks) == 0 || len(tasks) == 1 && len(tasks[0]) == 0 {
		fmt.Println(tb.MissingDesc())
		return
	}

	if len(boards) == 0 {
		//TODO add board.default in viper
		boards = append(boards, "My Board")
	}

	desc := strings.Join(tasks, " ")

	taskBook.AddTask(desc, str, boards, comp, prog, level[prio])
}

func init() {
	rootCmd.AddCommand(taskCmd)
	taskCmd.Flags().StringVarP(&prio, "priority", "p", "low", "Set priorty for task")
	_ = taskCmd.RegisterFlagCompletionFunc("priority", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"low", "medium", "high"}, cobra.ShellCompDirectiveDefault
	})

	taskCmd.Flags().BoolVarP(&str, "star", "s", false, "Star task")
	taskCmd.Flags().BoolVarP(&comp, "check", "c", false, "Check task")
	taskCmd.Flags().BoolVarP(&prog, "begin", "b", false, "Start task")
	taskCmd.MarkFlagsMutuallyExclusive("check", "begin")
}
