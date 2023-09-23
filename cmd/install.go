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
	"slices"

	"github.com/dkartachov/govm/pkg/semver"
	"github.com/dkartachov/govm/pkg/targz"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	Run: func(cmd *cobra.Command, args []string) {
		versions, err := validateVersionsToInstall(args)

		if err != nil {
			log.Fatalf("error validating versions: %v", err)
		}

		install(versions)
	},
}

func validateVersion(args []string) error {
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

func validateVersionsToInstall(args []string) ([]string, error) {
	if len(args) == 0 {
		return []string{}, fmt.Errorf("must provide at least 1 argument")
	}

	filteredArgs := []string{}

	for i := 0; i < len(args); i++ {
		version := args[i]
		if semver.Valid(version) || version == "go" {
			if !slices.Contains[[]string](filteredArgs, version) {
				filteredArgs = append(filteredArgs, version)
			}
		} else {
			log.Printf("skipping %s: invalid version", version)
		}
	}

	return filteredArgs, nil
}

func latestVersion() string {
	versions := availableVersions()
	err := semver.Sort(versions, semver.Desc)

	if err != nil {
		log.Fatalf("error sorting versions: %v", err)
	}

	return versions[0]
}

func install(versions []string) {
	for _, v := range versions {
		if v == "go" {
			v = latestVersion()
			log.Printf("found latest version %s", v)
		}

		if versionExists(v) {
			log.Printf("version %s is already installed", v)
			continue
		}

		log.Printf("downloading version %s from remote", v)
		url := fmt.Sprintf("https://go.dev/dl/go%s.linux-amd64.tar.gz", v)
		resp, err := http.Get(url)

		if err != nil {
			log.Fatal("error connecting to Go release archive", err)
		}

		if resp.StatusCode == http.StatusNotFound {
			log.Printf("skipping %s: not found", v)
			continue
		}

		// TODO validate config file earlier to stop execution if empty values are found
		govmBin := viper.GetString("GOVM_BIN")
		govmVersions := viper.GetString("GOVM_VERSIONS")

		err = os.MkdirAll(govmVersions, os.ModePerm)
		if err != nil {
			panic(err)
		}

		log.Printf("extracting archive for %s", v)
		if err = targz.Extract(resp.Body, govmVersions); err != nil {
			log.Fatalf("error extracting archive for %s: %v", v, err)
		}

		if err = os.Rename(filepath.Join(govmVersions, "go"), filepath.Join(govmVersions, v)); err != nil {
			log.Fatal(err)
		}

		// CHECKME Although the new version can be used immediately shell completion doesn't seem to work until
		// the terminal is refreshed or rc file is resourced. Is there a way to fix this?
		log.Printf("linking files for %s", v)
		if err = os.Symlink(filepath.Join(govmVersions, v, "bin/go"), filepath.Join(govmBin, "go"+v)); err != nil {
			log.Fatal(err)
		}

		log.Printf("version %s installed", v)
		resp.Body.Close()
	}
}

func init() {
	rootCmd.AddCommand(installCmd)
	// TODO add flag to switch versions during install. Maybe --use? Then the command
	// to install a new version and automatically switch to it becomes
	// govm install 1.20.7 --use
}
