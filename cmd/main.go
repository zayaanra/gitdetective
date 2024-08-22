package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/spf13/cobra"

	"github.com/zayaanra/gitdetective/api"
	"github.com/zayaanra/gitdetective/visuals"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var rootCmd = &cobra.Command{
		Use: "gd",
		Annotations: map[string]string{
			cobra.CommandDisplayNameAnnotation: "gd",
		},
	}
	rootCmd.PersistentFlags().StringVar(&cwd, "output", "", "Path to where output files are generated. Default is directory from which this command is run.")

	commitsCmd := &cobra.Command{
		Use:   "commits",
		Short: "Generates several statistics based on the number of commits for some time period",
		Long: `Generates several statistics based on the number of commits for some time period.
		Specifically, it will generate charts showcasing:
		Commits per hour for the past 24 hours
		Commits per day of the week for the past week
		Commits per hour of the week by day for the past week
		Commits by month of the year for the past year
		No. of lines committed (added/removed) by day for the past month
		`,
		Run: func(cmd *cobra.Command, args []string) {
			page := components.NewPage()
			data := api.PerformCommits()

			bar_byHour := visuals.GenerateBar("Commits By Hour in the Past 24 Hours", "", "Hour of Day", "# of Commits", "No. of Commits", data.ByHour)
			page.AddCharts(
				bar_byHour,
			)

			f, err := os.Create(filepath.Join(cwd, "commits_report.html"))
			if err != nil {
				panic(err)
			}
			page.Render(io.MultiWriter(f))
		},
	}

	authorsCmd := &cobra.Command{
		Use:   "authors",
		Short: "Generates several statistics based on the authors of the repository",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("gd authors")
		},
	}
	rootCmd.AddCommand(commitsCmd)
	rootCmd.AddCommand(authorsCmd)

	rootCmd.Execute()

}
