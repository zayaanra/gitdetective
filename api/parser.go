package api

// parser.go parses any output that is complex and transforms it into a dataframe

import (
	"log"
	"bufio"
	"fmt"
	"strings"
)

type Parser struct {
	output string // The original output produced by a git command

	AuthorEntries []AuthorEntry
}

type AuthorEntry struct {
	Author string
	Add    int
	Del    int
	File   string
}

func NewParser(output string) *Parser {
	return &Parser{output, nil}
}

func (p *Parser) ParseAuthorOutput() []AuthorEntry {
	// Parse output for author stats
	scanner := bufio.NewScanner(strings.NewReader(p.output))
	var entries []AuthorEntry
	var currentAuthor string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "\"author:") {
			currentAuthor = strings.TrimPrefix(line, "\"author: ")
			currentAuthor = strings.Trim(currentAuthor, "\"")
		} else if line != "" {
			fields := strings.Fields(line)
			if len(fields) == 3 {
				add := parseToInt(fields[0])
				del := parseToInt(fields[1])
				file := fields[2]
				entry := AuthorEntry{
					Author: currentAuthor,
					Add:    add,
					Del:    del,
					File:   file,
				}
				entries = append(entries, entry)
			}
		}
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading log data: %v", err)
	}

	p.AuthorEntries = entries
	return entries
}

func parseToInt(s string) int {
	if s == "-" {
		return 0
	}
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		return 0
	}
	return result
}

