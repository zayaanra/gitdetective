package visuals

import (
	"sort"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GenerateBar(title string, subtitle string, xName string, yName string, category string, data map[int]int) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title, Subtitle: subtitle}),
		charts.WithXAxisOpts(opts.XAxis{Name: xName}),
		charts.WithYAxisOpts(opts.YAxis{Name: yName}),
	)

	items := make([]opts.BarData, len(data))
	for i := 0; i < len(data); i++ {
		items[i] = opts.BarData{Value: data[i]}
	}

	keySet := make([]int, 0)
	for k := range data {
		keySet = append(keySet, k)
	}

	sort.Ints(keySet)

	bar.SetXAxis(keySet).
		AddSeries(category, items)

	return bar
}
