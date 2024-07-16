package cmd

import (
	"github.com/araaha/tb/taskbook"
	"github.com/spf13/cobra"
	"os"
)

var (
	taskBook tb.Book
	rootCmd  = &cobra.Command{
		Use:   "tb",
		Short: "A Taskbook",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			taskBook.Read("/home/araaha/taskbook.json")
		},
		Run: func(cmd *cobra.Command, args []string) {
			taskBook.DisplayByBoard()
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			taskBook.Store("/home/araaha/taskbook.json")
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	tb.Create()
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.taskbook.go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
