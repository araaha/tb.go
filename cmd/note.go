package cmd

import (
	"fmt"
	tb "github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
	"strings"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Create note",
	Run: func(_ *cobra.Command, args []string) {
		note(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return taskBook.GetAllBoard(true), cobra.ShellCompDirectiveNoFileComp
	},
}

// note adds a Note
func note(args []string) {
	var boards []string
	var notes []string
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
			notes = append(notes, arg)
		}
	}

	if len(notes) == 0 || len(notes) == 1 && len(notes[0]) == 0 {
		fmt.Println(tb.MissingDesc())
		return
	}

	if len(boards) == 0 {
		boards = append(boards, "My Board")
	}

	desc := strings.Join(notes, " ")

	taskBook.AddNote(desc, str, boards, level[prio])
}

func init() {
	rootCmd.AddCommand(noteCmd)
	noteCmd.Flags().StringVarP(&prio, "priority", "p", "low", "Set priorty for note")
	_ = noteCmd.RegisterFlagCompletionFunc("priority", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"low", "medium", "high"}, cobra.ShellCompDirectiveDefault
	})

	noteCmd.Flags().BoolVarP(&str, "star", "s", false, "Star note")
}
