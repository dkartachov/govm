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

	"github.com/dkartachov/govm/pkg"
	"github.com/dkartachov/govm/pkg/targz"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Download and install a specific version of Go",
	Long: `Download and install a specific version of Go
from https://go.dev/dl/. Automatically switches to the new
version.`,
	Aliases: []string{"i"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validate(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		/* TODO install command
		1. fetch archive from website
		2. extract archive (probably under ~/.govm/versions)
			- Rename folder to version number (e.g. 1.21.0/)
		3. remove previous symlink. If doesn't exist, attempt to remove previous installation
		4. create new symlink
		*/
		version := args[0]
		url := fmt.Sprintf("https://go.dev/dl/go%s.linux-amd64.tar.gz", version)
		// url := fmt.Sprint("https://nodejs.org/dist/v18.17.1/node-v18.17.1-linux-x64.tar.gz")

		resp, err := http.Get(url)

		if err != nil {
			log.Fatal("error connecting to Go release archive", err)
		}

		defer resp.Body.Close()

		home, _ := os.UserHomeDir()
		targetDir := filepath.Join(home, ".govm/versions")

		if err = targz.Extract(resp.Body, targetDir); err != nil {
			log.Fatal("error extracting archive", err)
		}

		if err = os.Rename(filepath.Join(targetDir, "go"), filepath.Join(targetDir, version)); err != nil {
			log.Fatal(err)
		}

		if err = os.RemoveAll(filepath.Join(home, ".govm/go")); err != nil {
			log.Fatal(err)
		}

		if err = os.Symlink(filepath.Join(targetDir, version), filepath.Join(home, ".govm/go")); err != nil {
			log.Fatal(err)
		}
	},
}

func validate(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Expected only one argument")
	}

	version := args[0]

	if !pkg.ValidVersion(version) {
		return fmt.Errorf("Invalid version %s", version)
	}

	return nil
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
}
