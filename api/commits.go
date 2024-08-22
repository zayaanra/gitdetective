package api

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

type CommitsData struct {
	ByMonth map[string]int
	ByHour  map[string]int
	ByDay   map[string]int
}

func commitsByMonth() {
}

func commitsByHour() map[string]int {
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

	return commits
}

func commitsByDay() {
}

func PerformCommits() *CommitsData {
	data := &CommitsData{ByHour: commitsByHour()}

	return data
}
