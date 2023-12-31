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
	"github.com/spf13/viper"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall [version]",
	Short: "uninstall a locally installed version",
	Long:  `Uninstall a locally installed version of Go.`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return localVersions(), cobra.ShellCompDirectiveDefault
	},
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

		govmBin := viper.GetString("GOVM_BIN")
		govmVersions := viper.GetString("GOVM_VERSIONS")
		var err error

		if version == current {
			err = os.Remove(filepath.Join(govmBin, "go"))
		} else {
			err = os.Remove(filepath.Join(govmBin, "go"+version))
		}

		if err != nil {
			log.Fatalf("error removing version symlink: %v", err)
		}

		err = os.RemoveAll(filepath.Join(govmVersions, version))

		if err != nil {
			log.Fatalf("error removing version: %v", err)
		}

		log.Printf("version %s has been uninstalled", version)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
