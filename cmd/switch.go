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

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch [version]",
	Short: "Change the preferred version of Go on your machine",
	Long: `Change the preferred version of Go on your machine to another locally 
installed version. If the specified version does not exist it will be downloaded. 

This version will be used when running 'go'. This change persists until you switch 
to a different version. If you simply want to test your app on a different version 
you should use the locally installed binary directly. For example: 

go1.20.7 run main.go`,
	Aliases: []string{"use"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateArgs(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]

		if !versionExists(version) {
			log.Printf("version %v is not installed", version)
			return
		}

		home, _ := os.UserHomeDir()
		targetDir := filepath.Join(home, ".govm/versions")

		if err := os.RemoveAll(filepath.Join(home, ".govm/go")); err != nil {
			log.Fatalf("error removing symlink: %v", err)
		}

		if err := os.Symlink(filepath.Join(targetDir, version, "bin/go"), filepath.Join(home, ".govm/go")); err != nil {
			log.Fatal(err)
		}
	},
}

func versionExists(version string) bool {
	home, _ := os.UserHomeDir()
	dir, err := os.ReadDir(filepath.Join(home, ".govm/versions"))

	if err != nil {
		log.Fatalf("error reading directory: %v", err)
	}

	for _, e := range dir {
		if version == e.Name() {
			return true
		}
	}

	return false
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
