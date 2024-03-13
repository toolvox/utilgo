package mathutil

import (
	"slices"
)

type Report struct {
	Min    float64
	Bot1p  float64
	Bot10p float64
	Avg    float64
	Top10p float64
	Top1p  float64
	Max    float64
}

func GenerateReport(items ...float64) Report {
	if len(items) == 0 {
		return Report{}
	}

	slices.Sort(items)
	percentile := func(sorted []float64, p float64) float64 {
		index := int(p * float64(len(sorted)))
		return sorted[index]
	}

	return Report{
		Min:    items[0],
		Bot1p:  percentile(items, 0.01),
		Bot10p: percentile(items, 0.10),
		Avg:    Average(items...),
		Top10p: percentile(items, 0.90),
		Top1p:  percentile(items, 0.99),
		Max:    items[len(items)-1],
	}
}
