/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dkartachov/govm/pkg/semver"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall [version]",
	Short: "uninstall a locally installed version",
	Long:  `Uninstall a locally installed version of Go.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filteredArgs, err := validateManyArgs(args)

		if err != nil {
			return err
		}

		for _, version := range filteredArgs {
			if !versionExists(version) {
				log.Printf("skipping %s: not installed", version)
				continue
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
		}

		return nil
	},
}

func validateManyArgs(args []string) ([]string, error) {
	if len(args) == 0 {
		return []string{}, fmt.Errorf("must provide at least 1 version to uninstall")
	}

	var filteredArgs []string

	for i := 0; i < len(args); i++ {
		if semver.Valid(args[i]) {
			filteredArgs = append(filteredArgs, args[i])
		} else {
			log.Printf("skipping %s: invalid version", args[i])
		}
	}

	return filteredArgs, nil
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
