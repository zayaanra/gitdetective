package api

import (
	"log"

	"fmt"
	"os/exec"
	

)

type AuthorDF struct {
	username string
	numOfInsertions int
	numOfDeletions int
	filename string
}

func AuthorStats() []AuthorEntry {
	cmd := exec.Command("git", "log", "--format=\"author: %ae\"", "--numstat")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching author data: %v", err)
	}

	p := NewParser(string(output))
	return p.ParseAuthorOutput()

}

func GeneratePieChart() {

}