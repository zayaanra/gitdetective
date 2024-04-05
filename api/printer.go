package api

import "fmt"

func PrintBasic(repoName string, creationDate string, numFiles int, numLines int, numCommits int) {
	fmt.Printf("Repository Name: %s\nRepository Creation Date: %s\n# of Tracked Files: %d\n# of Lines of Code: %d\n# of Commits: %d\n",
		repoName,
		creationDate,
		numFiles,
		numLines,
		numCommits)
}
