package api

import (
	"log"
	"os"
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

	for i := 0; i < 24; i++ {
		_, ok := commits[i]
		if !ok {
			commits[i] = 0
		}
	}

	return commits
}

func commitsByDay() map[int]int {
	cmd := exec.Command("git", "log", "--since=last.week", "--pretty=format:'%cd'", "--date=format:%A")
	cmd2 := exec.Command("sort")
	cmd3 := exec.Command("uniq -c")
	cmd4 := exec.Command("sort -nr")

	cmd2.Stdin, _ = cmd.StdoutPipe()
	cmd3.Stdin, _ = cmd2.StdoutPipe()
	cmd4.Stdin, _ = cmd3.StdoutPipe()
	cmd4.Stdout = os.Stdout

	cmd4.Start()
	cmd3.Start()
	cmd2.Start()
	cmd.Run()
	cmd2.Wait()
	cmd3.Wait()
	cmd4.Wait()


	//log.Println(cmd3.Stdout)


	// output, err := cmd.Output()
	// if err != nil {
	// 	log.Fatalf("Error in fetching number of commits per hour: %v", err)
	// }

	// lines := strings.Split(string(output), "\n")
	// log.Println(lines)
	commits := make(map[int]int)

	// for _, line := range lines {
	// 	if len(line) >= 10 {
	// 		hour, _ := strconv.Atoi(line[11:13])
	// 		commits[hour]++
	// 	}
	// }

	// log.Println(commits)

	// for i := 0; i < 24; i++ {
	// 	_, ok := commits[i]
	// 	if !ok {
	// 		commits[i] = 0
	// 	}
	// }

	return commits
}

func PerformCommits() *CommitsData {
	data := &CommitsData{ByHour: commitsByHour(), ByDay: commitsByDay()}

	return data
}
