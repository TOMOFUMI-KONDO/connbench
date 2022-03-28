package client

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/go-echarts/go-echarts/v2/charts"
)

func Render(durations []time.Duration, network string) error {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: fmt.Sprintf("Connection Establishment Time Benchmark for %s", strings.ToUpper(network))}),
		charts.WithYAxisOpts(opts.YAxis{Name: "connection establishment time[μs]", Max: 15000}),
	)

	lineItems := make([]opts.LineData, len(durations))
	for i, d := range durations {
		lineItems[i] = opts.LineData{Value: d.Microseconds()}
	}
	line.SetXAxis(zeroToN(len(durations))).AddSeries(network, lineItems)

	f, err := os.Create(fmt.Sprintf("charts/html/%s.html", network))
	if err != nil {
		return err
	}

	if err := line.Render(f); err != nil {
		return err
	}

	return nil
}

func zeroToN(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return slice
}
