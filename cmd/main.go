package main

import (
	"log"
	"os"

	"github.com/zayaanra/gitdetective/api"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	args := os.Args

	cmd := args[0]
	flags := args[2:]

	gd := api.NewGitDetective(cmd, flags)

	// If only one argument, then it is just 'gd'. Show basic statistics for the repo.
	if len(args) == 1 {
		gd.Basic()
		return
	}

	var cmdtype int

	// If more than one argument, then it is 'gd <flag>'. Show detailed statistics for the repo.
	if args[1] == "commits" {
		gd.DoCommits(flags[0])
		cmdtype = api.COMMITS
	} else if args[1] == "authors" {
		gd.DoAuthors()
		cmdtype = api.AUTHOR
	}

	for i, flag := range flags {
		if flag == "--save" && i == len(flags)-1 {
			log.Fatalf("Failed - no output path specified")
		}

		if flag == "--save" {
			path := flag[i+1]
			save(string(path), cmdtype)
		}
	}

	// TODO - Need to refactor codebase. It's getting too messy.
	// TODO - Work on --save.

}

// If --save option is provided, then we must save the resulting output to the specified path
func save(path string, cmdtype int) {
	var df dataframe.DataFrame
	switch {
	case cmdtype == api.AUTHOR:
		df = dataframe.LoadStructs()
	}
}
