/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/dkartachov/govm/pkg/semver"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/spf13/cobra"
)

// listRemoteCmd represents the listRemote command
var listRemoteCmd = &cobra.Command{
	Use:     "list-remote",
	Short:   "List versions available to download",
	Long:    `List versions available to download`,
	Aliases: []string{"lsr"},
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		listRemote(limit)
	},
}

func listRemote(limit int) {
	versions := availableVersions()
	err := semver.Sort(versions, semver.Asc)

	if limit > 0 {
		numVersions := len(versions)
		if limit > numVersions {
			limit = numVersions
		}
		versions = versions[numVersions-limit:]
	}

	if err != nil {
		log.Fatalf("error sorting versions: %v", err)
	}

	for _, v := range versions {
		fmt.Println(v)
	}
}

func availableVersions() []string {
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

	return versions
}

func init() {
	rootCmd.AddCommand(listRemoteCmd)
	listRemoteCmd.Flags().Int("limit", 10, "limit the amount returned (default 10), 0 for no limit")
}
