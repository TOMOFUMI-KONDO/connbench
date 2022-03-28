package client

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/go-echarts/go-echarts/v2/charts"
)

func Render(durations []time.Duration, network string) error {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    fmt.Sprintf("Connection Establishment Time[μs] for %s", strings.ToUpper(network)),
			Subtitle: fmt.Sprintf("Average: %s[μs]", strconv.FormatInt(average(durations), 10)),
		}),
		charts.WithYAxisOpts(opts.YAxis{Max: 18000, Type: "value"}),
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

func average(durations []time.Duration) int64 {
	var sum int64 = 0
	for _, d := range durations {
		sum += d.Microseconds()
	}
	return sum / int64(len(durations))
}

func zeroToN(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return slice
}
