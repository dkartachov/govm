/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List versions of Go",
	Long:    `List all locally installed versions of Go.`,
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func list() {
	home, _ := os.UserHomeDir()
	d, err := os.ReadDir(filepath.Join(home, ".govm/versions"))

	if err != nil {
		log.Fatal(err)
	}

	for _, e := range d {
		fmt.Println(e.Name())
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
