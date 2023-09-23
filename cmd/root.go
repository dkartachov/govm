/*
Copyright © 2023 Denis Kartachov <kartachovd@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "govm [command]",
	Short: "Manage different versions of Go on your machine",
	Long: `Govm is a CLI application that allows developers to manage
different versions of Go on their machine. Add, remove and switch
between installed versions with ease. 

You can also test your application on a different version without 
switching by running the installed binary directly. For example: 

go1.20.7 run main.go

Yes, I am aware that upgrading Go and managing different versions
is fairly easy to do without a manager. However, please allow me to 
elaborate on the motivation behind creating this tool:

"Why spend 5 minutes doing something when you can spend
5 days automating it?" - Some developer somewhere
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// CHECKME best place to set this?
	log.SetFlags(0)

	o := runtime.GOOS
	a := runtime.GOARCH

	// TODO support other platforms?
	if o != "linux" || a != "amd64" {
		fmt.Printf("platform not supported: %s/%s. Sorry!\n", o, a)
		os.Exit(1)
	}

	home, _ := os.UserHomeDir()
	viper.Set("GOVM_BIN", filepath.Join(home, ".govm/bin"))
	viper.Set("GOVM_VERSIONS", filepath.Join(home, ".govm/versions"))
}
