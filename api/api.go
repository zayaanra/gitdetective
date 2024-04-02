package api

import (
	"log"
	"os"
	"runtime"
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

func (g *GitDetective) Basic() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Might need to change depending on OS (\\)
	repoName := path[strings.LastIndex(path, "\\")+1:]
	log.Println(repoName)

}
