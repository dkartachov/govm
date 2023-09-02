/*
Copyright © 2023 Denis Kartachov <kartachovd@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "govm [command]",
	Short: "Manage different versions of Go on your machine",
	Long: `Govm is a CLI application that allows developers to manage
different versions of Go on their machine. Add, remove and switch
between locally installed versions with ease. You can also test
your application on a different version without switching by running 
the locally installed binary directly. For example: go1.20.7 run main.go

"Why spend 5 minutes doing something when you can spend
5 days automating it?" - Anonymous
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

func init() {}
