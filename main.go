/*
Copyright © 2023 Denis Kartachov <kartachovd@gmail.com>
*/
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/dkartachov/govm/cmd"
)

func main() {
	// CHECKME best place to set this?
	log.SetFlags(0)

	o := runtime.GOOS
	a := runtime.GOARCH

	// TODO support other platforms?
	if o != "linux" && a != "amd64" {
		fmt.Printf("platform not supported: %s/%s. Sorry!\n", o, a)
		os.Exit(1)
	}

	cmd.Execute()
}
