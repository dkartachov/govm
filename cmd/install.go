/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dkartachov/govm/pkg/semver"
	"github.com/dkartachov/govm/pkg/targz"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Download and install a specific version of Go",
	Long: `Download and install a specific version of Go
from https://go.dev/dl/.`,
	Aliases: []string{"i"},
	Example: `Install latest version
> govm install go

Install specific version
> govm install 1.21.0`,
	Args: func(cmd *cobra.Command, args []string) error {
		return validateArgs(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		install(args[0])
	},
}

func validateArgs(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected only one argument")
	}

	version := args[0]

	if !semver.Valid(version) {
		if version != "go" {
			return fmt.Errorf("invalid version %s", version)
		}
	}

	return nil
}

func latestVersion() string {
	versions := availableVersions()
	err := semver.Sort(versions, semver.Desc)

	if err != nil {
		log.Fatalf("error sorting versions: %v", err)
	}

	return versions[0]
}

func install(version string) {
	if version == "go" {
		version = latestVersion()
		log.Printf("found latest version %s", version)
	}

	if versionExists(version) {
		log.Print("version is already installed")
		return
	}

	log.Print("downloading version from remote")
	url := fmt.Sprintf("https://go.dev/dl/go%s.linux-amd64.tar.gz", version)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("error connecting to Go release archive", err)
	}

	defer resp.Body.Close()

	home, _ := os.UserHomeDir()
	targetDir := filepath.Join(home, ".govm/versions")

	log.Print("extracting archive")
	if err = targz.Extract(resp.Body, targetDir); err != nil {
		log.Fatal("error extracting archive", err)
	}

	if err = os.Rename(filepath.Join(targetDir, "go"), filepath.Join(targetDir, version)); err != nil {
		log.Fatal(err)
	}

	// CHECKME Although the new version can be used immediately shell completion doesn't seem to work until
	// the terminal is refreshed or rc file is resourced. Is there a way to fix this?
	log.Print("linking files")
	if err = os.Symlink(filepath.Join(targetDir, version, "bin/go"), filepath.Join(home, ".govm/go"+version)); err != nil {
		log.Fatal(err)
	}

	log.Printf("version %s installed", version)
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// TODO add flag to switch versions during install. Maybe --use? Then the command
	// to install a new version and automatically switch to it becomes
	// govm install 1.20.7 --use
}
