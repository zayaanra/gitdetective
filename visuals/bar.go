package visuals

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GenerateBar(title string, subtitle string, category string, data map[string]int) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title, Subtitle: subtitle}),
	)

	items := make([]opts.BarData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.BarData{Value: data[fmt.Sprint(i)]})
	}

	keySet := make([]string, 0)
	for k := range data {
		keySet = append(keySet, k)
	}

	bar.SetXAxis(keySet).
		AddSeries(category, items)

	return bar
}
