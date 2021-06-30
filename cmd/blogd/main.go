package main

import (
	"os"

	"github.com/slandymani/blog/cmd/blogd/cmd"
    svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/slandymani/blog/app"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
    if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
