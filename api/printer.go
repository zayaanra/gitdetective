package api

import "fmt"

func PrintBasic(repoName string, creationDate string, numFiles int, numLines int, numCommits int) {
	fmt.Printf("Repo: %s\nCreation Date: %s\nFiles: %d\nLines of Code: %d\nCommits: %d\n", repoName, creationDate, numFiles, numLines, numCommits)
}
