package main

import (
	"os"

	"github.com/cmwaters/rook/cmd/rookd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
