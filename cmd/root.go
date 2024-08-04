package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/araaha/tb.go/taskbook"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Version represents the version of tb.go
	Version string
	// Revision represents the revision of tb.go
	Revision string
)

var (
	// taskBook represents the taskbook
	taskBook tb.Book
	// cfgFile represents the config file
	cfgFile string
	// ver is the version information
	ver bool
)

// rootCmd displays board or version. It reads the data and stores it.
var rootCmd = &cobra.Command{
	Use:   "tb",
	Short: "A Taskbook",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := taskBook.Read(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	Run: func(_ *cobra.Command, _ []string) {
		if ver {
			fmt.Printf("%s (%s)\n", Version, Revision)
			return
		}
		taskBook.DisplayByBoard()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if err := taskBook.Store(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/taskbook/taskbook.toml)")
	rootCmd.PersistentFlags().BoolVarP(&ver, "version", "v", false, "Display current version")
	_ = rootCmd.RegisterFlagCompletionFunc("config", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"toml"}, cobra.ShellCompDirectiveFilterFileExt
	})

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cobra.CheckErr(err)

		var cfgPath string

		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")

		if xdgConfigHome != "" {
			cfgPath = filepath.Join(xdgConfigHome, "taskbook")
		} else {
			cfgPath = filepath.Join(home, ".config", "taskbook")
		}

		// Search config in XDG first, followed by .config/taskbook
		viper.AddConfigPath(cfgPath)
		viper.SetConfigName("taskbook")
		viper.SetConfigType("toml")

	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return
		}
	}
}
