package api

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Commits struct {
	count int // Keeps track of the # of commits
}

func NewCommits() *Commits {
	return &Commits{}
}

func (c *Commits) GetNumCommits() int {
	cmd := exec.Command("git", "rev-list", "--count", "--all")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching number of commits: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	count, _ := strconv.Atoi(lines[0])

	c.count = count

	return count
}

/* gd commits -t */
// The -t flag produces repo. statistics for only today.
// Specifically, it will show a table consisting of each hour of the day and the number of commits per hour.
func (c *Commits) T() {
	now := time.Now()
	pastDay := now.Add(-24 * time.Hour)
	since := pastDay.Format("2006-01-02T:15:04:05")

	cmd := exec.Command("git", "log", "--since="+since, "--format=%ci")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching number of commits per hour: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	commits := make(map[string]int)

	for _, line := range lines {
		if len(line) >= 10 {
			hour := line[11:13]
			commits[hour]++
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Hour", "Commits"})

	for hour, count := range commits {
		var color text.Color
		switch {
		case count >= 10:
			color = text.FgHiRed
		case count >= 5:
			color = text.FgHiYellow
		default:
			color = text.FgHiGreen
		}
		t.AppendRow([]interface{}{color.Sprint(hour), color.Sprint(count)})
	}

	t.Render()
}
