/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use [version]",
	Short: "Change the default version of Go on your machine",
	Long: `Change the default version of Go on your machine to another locally 
installed version. This version will be used when running 'go'. 

This change persists until you switch to a different version. 
If you simply want to test your app on a different version you should 
use the locally installed binary directly. For example: 

go1.20.7 run main.go`,
	Args: func(cmd *cobra.Command, args []string) error {
		return validateVersion(args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return localVersions(), cobra.ShellCompDirectiveDefault
	},
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		previousVersion, _ := currentVersion()

		if version == previousVersion {
			log.Print("already using version")
			return
		}

		if !versionExists(version) {
			log.Printf("version %v is not installed", version)
			return
		}

		govmBin := viper.GetString("GOVM_BIN")
		govmVersions := viper.GetString("GOVM_VERSIONS")

		// links go to new version x
		if err := os.RemoveAll(filepath.Join(govmBin, "go")); err != nil {
			log.Fatalf("error removing default version symlink: %v", err)
		}

		if err := os.Symlink(filepath.Join(govmVersions, version, "bin/go"), filepath.Join(govmBin, "go")); err != nil {
			log.Fatalf("error linking default version: %v", err)
		}

		// remove x symlink and replace with previous version
		if err := os.RemoveAll(filepath.Join(govmBin, "go"+version)); err != nil {
			log.Fatalf("error removing versioned symlink: %v", err)
		}

		if previousVersion != "" {
			if err := os.Symlink(filepath.Join(govmVersions, previousVersion, "bin/go"), filepath.Join(govmBin, "go"+previousVersion)); err != nil {
				log.Fatalf("error linking previous version: %v", err)
			}
		}

		log.Printf("go ==> %s", version)
	},
}

func versionExists(version string) bool {
	govmVersions := viper.GetString("GOVM_VERSIONS")
	if err := os.MkdirAll(govmVersions, os.ModePerm); err != nil {
		panic(err)
	}

	dir, err := os.ReadDir(govmVersions)

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

func currentVersion() (string, error) {
	govmBin := viper.GetString("GOVM_BIN")
	regex := regexp.MustCompile(`\d+(\.\d+)?(\.\d+)?`)
	goLink, err := filepath.EvalSymlinks(filepath.Join(govmBin, "go"))

	if err != nil {
		return "", fmt.Errorf("error evaluating symlink: %v", err)
	}

	return regex.FindString(goLink), nil
}

func init() {
	rootCmd.AddCommand(useCmd)
}
