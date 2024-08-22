package api

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type CommitsData struct {
	ByMonth map[int]int
	ByHour  map[int]int
	ByDay   map[int]int
}

func commitsByMonth() {
}

func commitsByHour() map[int]int {
	now := time.Now()
	pastDay := now.Add(-24 * time.Hour)
	since := pastDay.Format("2006-01-02T15:04:05")

	cmd := exec.Command("git", "log", "--since="+since, "--format=%ci")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error in fetching number of commits per hour: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	commits := make(map[int]int)

	for _, line := range lines {
		if len(line) >= 10 {
			hour, _ := strconv.Atoi(line[11:13])
			commits[hour]++
		}
	}

	log.Println(commits)

	for i := 0; i < 24; i++ {
		_, ok := commits[i]
		if !ok {
			commits[i] = 0
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
