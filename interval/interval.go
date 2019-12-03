// Package interval provides a structure for intervals of integers and merging
// of overlapping intervals.
package interval

import (
	"fmt"
	"github.com/schneereich/merge-intervals/intmath"
	"log"
	"sort"
)

type Interval struct {
	Start, End int
}

// Stringer to define native format (square brackets)
func (i Interval) String() string {
	return fmt.Sprintf("[%v,%v]", i.Start, i.End)
}

// Merge returns a sorted slice of merged overlapping intervals
func Merge(intervals []Interval) []Interval {
	switch len(intervals) {
	case 0:
		log.Println("Please provide at least one interval. expected: len(intervals) > 0; actual:",
			len(intervals))
		return intervals
	case 1:
		log.Println("Minimum of two intervals required for merging, skipping execution")
		return intervals
	}

	// intervals sorted by Start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	// merged overlapping intervals
	merged := []Interval{intervals[0]}
	for _, interval := range intervals[1:] {
		if merged[len(merged)-1].End < interval.Start {
			merged = append(merged, interval)
			continue
		}
		log.Println("Merging the intervals", merged[len(merged)-1], interval)
		merged[len(merged)-1].End = intmath.Max(merged[len(merged)-1].End, interval.End)
	}

	return merged
}
