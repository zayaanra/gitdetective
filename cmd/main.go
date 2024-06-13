package main

import (
	"flag"
	"fmt"
	"os"
	// "github.com/zayaanra/gitdetective/api"
	// "github.com/go-gota/gota/dataframe"
)

const (
	AUTHORS = 0
	COMMITS = 1
)

// Represents the command provided by the user in struct form
type Command struct {
	cmdtype int // Either -c (commit) or -a (author)

	save_opts *SaveOptions // Used if 'save' flag is enabled

}

type SaveOptions struct {
	path     string
	filename string
}

func main() {
	// cmd := parseFlags()
	parseFlags()
	// args := os.Args

	// cmd := args[0]
	// flags := args[2:]

	// gd := api.NewGitDetective(cmd, flags)

	// // If only one argument, then it is just 'gd'. Show basic statistics for the repo.
	// if len(args) == 1 {
	// 	gd.Basic()
	// 	return
	// }

	// // var cmdtype int

	// // If more than one argument, then it is 'gd <flag>'. Show detailed statistics for the repo.
	// if args[1] == "commits" {
	// 	gd.DoCommits(flags[0])
	// 	// cmdtype = api.COMMITS
	// } else if args[1] == "authors" {
	// 	gd.DoAuthors()
	// 	// cmdtype = api.AUTHOR
	// }

	// for i, flag := range flags {
	// 	if flag == "--save" && i == len(flags)-1 {
	// 		log.Fatalf("Failed - no output path specified")
	// 	}

	// 	if flag == "--save" {
	// 		// path := flag[i+1]
	// 		//save(string(path), cmdtype)
	// 	}
	// }

	// TODO - Need to refactor codebase. It's getting too messy.
	// TODO - Work on --save.

}

func parseFlags() *Command {
	// Main command flags
	commitsCmd := flag.Bool("c", false, "Show commit stats")
	authorsCmd := flag.Bool("a", false, "Show author stats")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}

	var saveOpts *SaveOptions

	if len(os.Args) > 2 && os.Args[2] == "save" {
		// Save subcommand - used to save output data to a CSV file in some location
		saveCmd := flag.NewFlagSet("save", flag.ExitOnError)
		path := saveCmd.String("path", "", "Specify the path for where the output file should be saved")
		filename := saveCmd.String("filename", "", "Specify the filename for the output file")

		switch os.Args[2] {
		case "save":
			// usage: gd <-c/-a> save -path=<filepath> -filename=<filename>
			saveCmd.Parse(os.Args[3:])

			if len(*path) == 0 || len(*filename) == 0 {
				fmt.Println("Error: File path or filename was empty")
				os.Exit(1)
			}
			saveOpts = &SaveOptions{path: *path, filename: *filename}
			fmt.Println("subcommand 'save'")
			fmt.Println("	path:", *path)
			fmt.Println("	filename:", *filename)
		default:
			flag.Usage()
			os.Exit(1)
		}
	}

	fmt.Println("commits:", *commitsCmd)
	fmt.Println("authors:", *authorsCmd)

	var cmdtype int
	if *commitsCmd {
		cmdtype = COMMITS
	} else {
		cmdtype = AUTHORS
	}

	return &Command{cmdtype: cmdtype, save_opts: saveOpts}
}
