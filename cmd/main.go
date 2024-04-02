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
		/*
			TODO: Pretty print the following
			- repo name
			- creation date
			- # of files
			- total lines of code
			- total # of commits
			- # of authors
		*/
	}
}
