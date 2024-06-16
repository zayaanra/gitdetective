package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/zayaanra/gitdetective/gd"
)

const (
	BASIC   = 0
	COMMITS = 1
	AUTHORS = 2
)

// Represents the command provided by the user in struct form
type Command struct {
	cmdtype int // Specifies what type of output the user wants

	save_opts *SaveOptions // Used if 'save' flag is enabled

}

type SaveOptions struct {
	path     string
	filename string
}

func main() {
	cmd := parseFlags()

	gd := gd.NewGitDetective()

	var df dataframe.DataFrame

	// If user entered -b, perform basic stats.
	if cmd.cmdtype == BASIC {
		df = gd.Basic()
	} else if cmd.cmdtype == COMMITS {
		// TODO
	}

	// If save option is enabled, save DF as CSV to the given path as the given filename
	if cmd.save_opts != nil {
		fullPath := fmt.Sprintf("%s%s", cmd.save_opts.path, cmd.save_opts.filename)
		file, err := os.Create(fullPath)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)

		}
		defer file.Close()

		if err := df.WriteCSV(file); err != nil {
			log.Fatalf("Failed to write DF to CSV: %v", err)
		}

		fmt.Printf("DataFrame was saved to %s\n", fullPath)
	}

}

func parseFlags() *Command {
	// Main command flags
	basicCmd := flag.Bool("b", false, "Show basic stats")
	commitsCmd := flag.Bool("c", false, "Show commit stats")
	authorsCmd := flag.Bool("a", false, "Show author stats")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}

	var saveOpts *SaveOptions

	// TODO: Will need to change args indexing as commands grow
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

	// fmt.Println("commits:", *commitsCmd)
	// fmt.Println("authors:", *authorsCmd)
	// fmt.Println("basic:", *basicCmd)

	var cmdtype int
	if *commitsCmd {
		cmdtype = COMMITS
	} else if *authorsCmd {
		cmdtype = AUTHORS
	} else if *basicCmd {
		cmdtype = BASIC
	}

	return &Command{cmdtype: cmdtype, save_opts: saveOpts}
}
