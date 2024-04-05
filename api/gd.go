package api

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	LINUX   = 0
	WINDOWS = 1
	DARWIN  = 2
)

type GitDetective struct {
	cmd   string
	flags []string

	os int
}

func NewGitDetective(cmd string, flags []string) *GitDetective {
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

	return &GitDetective{cmd: cmd, flags: flags, os: t}
}

// Basic prints basic statistics such as:
// - repo name
// - creation date
// - # of files
// - total lines of code
// - total # of commits
// - # of authors
func (g *GitDetective) Basic() {
	name := g.GetRepoName()
	date := g.GetRepoCreationDate()
	numTracked := g.GetNumTrackedFiles()
	numLines := g.GetNumTotalLinesOfCode()
	numCommits := g.GetNumCommits()
	PrintBasic(name, date, numTracked, numLines, numCommits)

}

func (g *GitDetective) GetRepoName() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error in fetching repo name: %v", err)
	}

	return path[strings.LastIndex(path, "/")+1:]
}

func (g *GitDetective) GetRepoCreationDate() string {
	cmd := exec.Command("git", "log", "--reverse", "--format=%aI")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching repo creation date: %v", err)
	}

	lines := strings.Split(string(output), "\n")

	return lines[0]
}

func (g *GitDetective) GetNumTrackedFiles() int {
	cmd := exec.Command("/bin/sh", "-c", "git ls-files | wc -l")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching number of tracked files: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	count, _ := strconv.Atoi(lines[0])

	return count
}

func (g *GitDetective) GetNumTotalLinesOfCode() int {
	cmd := exec.Command("/bin/sh", "-c", "git ls-files | xargs wc -l")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching number of total lines of code: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	lastLine := lines[len(lines)-2]
	parsed := strings.Split(lastLine, " ")
	total, _ := strconv.Atoi(parsed[1])

	return total
}

func (g *GitDetective) GetNumCommits() int {
	cmd := exec.Command("git", "rev-list", "--count", "--all")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching number of commits: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	count, _ := strconv.Atoi(lines[0])

	return count
}
