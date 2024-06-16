package gd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

const (
	LINUX   = 0
	WINDOWS = 1
	DARWIN  = 2
)

type GitDetective struct {
	os int
}

func NewGitDetective() *GitDetective {
	var t int
	os := runtime.GOOS

	switch {
	case os == "windows":
		t = WINDOWS
	case os == "darwin":
		t = DARWIN
	case os == "linux":
		t = LINUX
	}

	return &GitDetective{os: t}
}

// Basic prints basic statistics such as:
// - repo name
// - creation date
// - # of files
// - total lines of code
// - total # of commits
func (g *GitDetective) Basic() dataframe.DataFrame {
	// Get repo name
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error in fetching repo name: %v", err)
	}
	repoName := path[strings.LastIndex(path, "/")+1:]

	// Get repo creation date
	cmd := exec.Command("git", "log", "--reverse", "--format=%aI")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching repo creation date: %v", err)
	}
	creationDate := strings.Split(string(output), "\n")[0]

	// Get # of tracked files
	cmd = exec.Command("/bin/sh", "-c", "git ls-files | wc -l")
	output, err = cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching number of tracked files: %v", err)
	}
	lines := strings.Split(string(output), "\n")
	numTracked, _ := strconv.Atoi(lines[0])

	// Get # of lines of code
	cmd = exec.Command("/bin/sh", "-c", "git ls-files | xargs cat | wc -l")
	output, err = cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching number of total lines of code: %v", err)
	}
	lines = strings.Split(string(output), "\n")
	parsed := strings.Trim(lines[0], " ")
	numLines, _ := strconv.Atoi(parsed)

	// Get # of commits
	cmd = exec.Command("git", "rev-list", "--count", "--all")
	output, err = cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching number of commits: %v", err)
	}
	lines = strings.Split(string(output), "\n")
	numCommits, _ := strconv.Atoi(lines[0])

	// Print out basic info
	fmt.Printf("Repository Name: %s\nRepository Creation Date: %s\n# of Tracked Files: %d\n# of Lines of Code: %d\n# of Commits: %d\n",
		repoName,
		creationDate,
		numTracked,
		numLines,
		numCommits)

	return dataframe.New(
		series.New([]string{repoName}, series.String, "Repository Name"),
		series.New([]string{creationDate}, series.String, "Creation Date"),
		series.New([]int{numTracked}, series.Int, "Num. of Tracked Files"),
		series.New([]int{numLines}, series.Int, "Num. of Lines of Code"),
		series.New([]int{numCommits}, series.Int, "Num. of Commits"),
	)
}

// This outputs information for the -c flag:
// - # of commits since some time
func (g *GitDetective) Commits() {

}

// func (g *GitDetective) DoCommits(flag string) {
// 	switch {
// 	case flag == "-t":
// 		g.commits.T()
// 	}
// }

// func (g *GitDetective) DoAuthors() []AuthorEntry{
// 	return AuthorStats()
// }
