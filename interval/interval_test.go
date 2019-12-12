package interval

import "testing"

func equalIntervals(actual, expected []Interval) bool {
	if len(actual) != len(expected) {
		return false
	}
	for index := range actual {
		if actual[index] != expected[index] {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	tables := []struct {
		input  []Interval
		output []Interval
	}{
		{
			[]Interval{{15, 20}, {2, 6}, {17, 24}},
			[]Interval{{2, 6}, {15, 24}},
		},
		{
			[]Interval{},
			[]Interval{},
		},
		{
			[]Interval{{0, 1}},
			[]Interval{{0, 1}},
		},
		{
				[]Interval{{1, 15}, {2, 10}},
				[]Interval{{1, 15}},
		},
		{
			[]Interval{{1, 15}, {0, 1}, {15, 16}},
			[]Interval{{0, 16}},
		},
		{
				[]Interval{{0, -1},{2,16}},
				[]Interval{{-1, 0},{2,16}},
		},
	}

	for _, table := range tables {
		actual := Merge(table.input)
		if !equalIntervals(actual, table.output) {
			t.Errorf("Merge of %v was incorrect. expected: %v; actual: %v", table.input, table.output, actual)
		}
	}
}
