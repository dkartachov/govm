/*
Copyright © 2023 Denis Kartachov <kartachovd@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "govm [command]",
	Short: "Manage different versions of Go on your machine",
	Long: `Govm is a CLI application that allows developers to manage
different versions of Go on their machine. Add, remove and switch
between locally installed versions with ease.

You can also simply use this tool to easily upgrade to the latest version
of Go. We all know how tedious it is to navigate to the website, download
the latest release, extract the archive, remove the previous installation...`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.govm.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
