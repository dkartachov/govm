/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall [version]",
	Short: "uninstall a locally installed version",
	Long:  `Uninstall a locally installed version of Go.`,
	Args: func(cmd *cobra.Command, args []string) error {
		return validateArgs(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]

		if !versionExists(version) {
			log.Printf("version %s is not installed", version)
			return
		}

		home, _ := os.UserHomeDir()
		err := os.Remove(filepath.Join(home, ".govm", "go"+version))

		if err != nil {
			log.Fatalf("error removing version symlink: %v", err)
		}

		err = os.RemoveAll(filepath.Join(home, ".govm/versions", version))

		if err != nil {
			log.Fatalf("error removing version: %v", err)
		}

		log.Printf("version %s has been uninstalled", version)
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
