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

		home, _ := os.UserHomeDir()
		targetDir := filepath.Join(home, ".govm/versions")

		// links go to new version x
		if err := os.RemoveAll(filepath.Join(home, ".govm/go")); err != nil {
			log.Fatalf("error removing default version symlink: %v", err)
		}

		if err := os.Symlink(filepath.Join(targetDir, version, "bin/go"), filepath.Join(home, ".govm/go")); err != nil {
			log.Fatalf("error linking default version: %v", err)
		}

		// remove x symlink and replace with previous version
		if err := os.RemoveAll(filepath.Join(home, ".govm/go"+version)); err != nil {
			log.Fatalf("error removing versioned symlink: %v", err)
		}

		if previousVersion != "" {
			if err := os.Symlink(filepath.Join(targetDir, previousVersion, "bin/go"), filepath.Join(home, ".govm/go"+previousVersion)); err != nil {
				log.Fatalf("error linking previous version: %v", err)
			}
		}

		log.Printf("go ==> %s", version)
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

func currentVersion() (string, error) {
	home, _ := os.UserHomeDir()
	regex := regexp.MustCompile(`\d+(\.\d+)?(\.\d+)?`)
	goLink, err := filepath.EvalSymlinks(filepath.Join(home, ".govm/go"))

	if err != nil {
		return "", fmt.Errorf("error evaluating symlink: %v", err)
	}

	return regex.FindString(goLink), nil
}

func init() {
	rootCmd.AddCommand(useCmd)
}
