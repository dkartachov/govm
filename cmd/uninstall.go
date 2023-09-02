/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/dkartachov/govm/pkg/semver"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall [version]",
	Short: "uninstall a locally installed version",
	Long:  `Uninstall a locally installed version of Go.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return uninstall(args)
	},
}

// TODO add "all" option to facilate uninstalling all versions? It will most likely be used
// to remove govm from machine. As a fail safe user will need to confirm this action.
func validateVersionsToUninstall(args []string) ([]string, error) {
	if len(args) == 0 {
		return []string{}, fmt.Errorf("must provide at least 1 argument")
	}

	filteredArgs := []string{}

	for i := 0; i < len(args); i++ {
		version := args[i]
		if semver.Valid(version) {
			if !slices.Contains[[]string](filteredArgs, version) {
				filteredArgs = append(filteredArgs, version)
			}
		} else {
			log.Printf("skipping %s: invalid version", version)
		}
	}

	return filteredArgs, nil
}

func uninstall(args []string) error {
	current, _ := currentVersion()
	filteredArgs, err := validateVersionsToUninstall(args)

	if err != nil {
		return err
	}

	for _, version := range filteredArgs {
		if !versionExists(version) {
			log.Printf("skipping %s: not installed", version)
			continue
		}

		home, _ := os.UserHomeDir()
		var err error

		if version == current {
			err = os.Remove(filepath.Join(home, ".govm/go"))
		} else {
			err = os.Remove(filepath.Join(home, ".govm", "go"+version))
		}

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
