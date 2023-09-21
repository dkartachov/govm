/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List locally installed versions",
	Long:    `List all locally installed versions of Go.`,
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func localVersions() []string {
	govmVersions := viper.GetString("GOVM_VERSIONS")
	if err := os.MkdirAll(govmVersions, os.ModePerm); err != nil {
		panic(err)
	}

	d, err := os.ReadDir(govmVersions)

	if err != nil {
		log.Fatalf("error retrieving locally installed versions: %v", err)
	}

	versions := []string{}
	for _, e := range d {
		versions = append(versions, e.Name())
	}

	return versions
}

func list() {
	versions := localVersions()

	for _, v := range versions {
		fmt.Println(v)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
