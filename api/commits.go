package api

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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
// The -t flag produces repo. statistics for only the past 24 hours.
// If an argument is present, then it will show depending on the number of past hours provided.
// Specifically, it will show a table consisting of each hour of the day and the number of commits per hour.
func (c *Commits) T() {
	now := time.Now()
	pastDay := now.Add(-24 * time.Hour)
	since := pastDay.Format("2006-01-02T15:04:05")

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

	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	table := widgets.NewTable()
	table.Rows = [][]string{
		{"Hour", "Commits"},
	}
	for hour, count := range commits {
		table.Rows = append(table.Rows, []string{hour, strconv.Itoa(count)})
	}

	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.SetRect(0, 0, 50, 10)
	table.BorderStyle = ui.NewStyle(ui.ColorYellow)
	table.RowSeparator = true

	ui.Render(table)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
