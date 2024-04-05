package main

import (
	"os"

	"github.com/zayaanra/gitdetective/api"
)

func main() {
	args := os.Args

	cmd := args[0]
	flags := args[1:]

	gd := api.NewGitDetective(cmd, flags)

	// If only one argument, then it is just 'gd'. Show basic statistics for the repo.
	if len(args) == 1 {
		gd.Basic()
		return
	}

	// If more than one argument, then it is 'gd <flag>'. Show detailed statistics for the repo.
	if args[1] == "commits" {
		gd.DoCommits(flags[1])
	}

}
