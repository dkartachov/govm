/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/dkartachov/govm/pkg/semver"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all locally installed versions",
	Long:    "List all locally installed versions of Go.",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		remote, _ := cmd.Flags().GetBool("remote")

		if remote {
			listRemote()
		} else {
			list()
		}
	},
}

func list() {
	home, _ := os.UserHomeDir()
	d, err := os.ReadDir(filepath.Join(home, ".govm/versions"))

	if err != nil {
		log.Fatal(err)
	}

	for i := len(d) - 1; i >= 0; i-- {
		fmt.Println(d[i].Name())
	}
}

func listRemote() {
	r := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		URLs: []string{"https://github.com/golang/go"},
	})

	err := r.Fetch(&git.FetchOptions{})

	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Fatalf("error fetching references from remote repo: %v", err)
	}

	rfs, err := r.List(&git.ListOptions{})

	if err != nil {
		log.Fatalf("error retrieving list of remote references: %v", err)
	}

	var versions = []string{}

	for _, ref := range rfs {
		if ref.Name().IsTag() {
			versionP := ref.Name().Short()

			if semver.ValidP(versionP, "go") {
				version := strings.TrimPrefix(versionP, "go")
				versions = append(versions, version)
			}
		}
	}

	sort.SliceStable(versions, func(i, j int) bool {
		// TODO implement version comparison
		return true
	})

	for _, v := range versions {
		fmt.Println(v)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().Bool("remote", false, "list available versions to download")
}
