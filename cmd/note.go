/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Create note",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: note,
	PostRun: func(cmd *cobra.Command, args []string) {
		taskBook.Store("/home/araaha/taskbook.json")
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return taskBook.AllBoard(), cobra.ShellCompDirectiveNoFileComp
	},
}

func note(cmd *cobra.Command, args []string) {
	var b []string
	var t []string
	level := map[string]int{"low": 1, "medium": 2, "high": 3}

	if len(args) == 0 {
		fmt.Printf("No description was given as input")
		os.Exit(1)
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "@") {
			b = append(b, arg)
		} else {
			t = append(t, arg)
		}
	}

	if len(b) == 0 {
		b = append(b, "My Board")
	}

	desc := strings.Join(t, " ")

	taskBook.AddNote(desc, star, b, level[prio])
}

func init() {
	rootCmd.AddCommand(noteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// noteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// noteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	noteCmd.Flags().StringVarP(&prio, "priority", "p", "low", "Set priorty for note")
	noteCmd.RegisterFlagCompletionFunc("priority", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"low", "medium", "high"}, cobra.ShellCompDirectiveDefault
	})

	noteCmd.Flags().BoolVarP(&star, "star", "s", false, "Star note")
}
