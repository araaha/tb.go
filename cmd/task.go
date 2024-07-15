/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var (
	prio string
	star bool
	comp bool
	prog bool

	taskCmd = &cobra.Command{
		Use:   "task",
		Short: "Create task",
		Run:   task,
		PostRun: func(cmd *cobra.Command, args []string) {
			taskBook.Store("/home/araaha/taskbook.json")
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return taskBook.AllBoard(), cobra.ShellCompDirectiveNoFileComp
		},
	}
)

func task(cmd *cobra.Command, args []string) {
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

	taskBook.AddTask(desc, star, b, comp, prog, level[prio])
}

func init() {
	rootCmd.AddCommand(taskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	taskCmd.Flags().StringVarP(&prio, "priority", "p", "low", "Set priorty for task")
	taskCmd.RegisterFlagCompletionFunc("priority", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"low", "medium", "high"}, cobra.ShellCompDirectiveDefault
	})

	taskCmd.Flags().BoolVarP(&star, "star", "s", false, "Star task")
	taskCmd.Flags().BoolVarP(&comp, "check", "c", false, "Check task")
	taskCmd.Flags().BoolVarP(&prog, "begin", "b", false, "Start task")
	taskCmd.MarkFlagsMutuallyExclusive("check", "begin")
}
