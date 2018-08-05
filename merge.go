package main

import "sort"

// Definition for an interval.
type Interval struct {
	Start int
	End   int
}
type Intervals []Interval

func (intervals Intervals) Less(i, j int) bool {
	if intervals[i].Start < intervals[j].Start {
		return true
	} else {
		return false
	}
}
func (intervals Intervals) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]

}
func (intervals Intervals) Len() int {
	return len(intervals)
}

func mergeRange(intervals []Interval) []Interval {
	var result []Interval
	wrap := Intervals(intervals)
	sort.Sort(wrap)
	temp := wrap[0]
	for i := 1; i < wrap.Len(); i++ {
		if temp.End > wrap[i].Start {
			if temp.End < wrap[i].End {
				temp.End = wrap[i].End
			}
		} else {
			result = append(result, temp)
			temp = wrap[i]
		}
	}
	result = append(result, temp)
	return result

}
