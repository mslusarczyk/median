package slidingwindow

import (
	"sort"
)

type SlidingWindow struct {
	size          int
	sortedMetrics []*metric
	oldest        *metric
	latest        *metric
}

type metric struct {
	value int
	index int
	next  *metric
}

// returns new instance of sliding window with given size
func NewSlidingWindow(size int) *SlidingWindow {
	return &SlidingWindow{size: size, sortedMetrics: make([]*metric, 0, size)}
}

// adds delay value to sliding window. Time: O(sw.size), Space: O(sw.size)
func (sw *SlidingWindow) AddDelay(value int) {
	metric := &metric{value: value}

	// init of first elem
	if len(sw.sortedMetrics) == 0 {
		sw.addFirst(metric)
		return
	}

	sw.latest.next = metric
	sw.latest = metric

	// below cap, simply adding new elem
	if len(sw.sortedMetrics) < sw.size {
		sw.addWhenBelowCap(metric)
		return
	}

	// no sorting needed, moving pointers only
	if metric.value == sw.oldest.value {
		sw.sortedMetrics[sw.oldest.index] = metric
		metric.index = sw.oldest.index
		sw.oldest = sw.oldest.next
		return
	}

	//window at capacity, oldest element needs to be removed
	sw.removeOldest()
	// new metric is added
	sw.addWhenBelowCap(metric)
}

// find median for current state of sliding window. Time: O(1), Space: O(1)
func (sw *SlidingWindow) GetMedian() int {
	if len(sw.sortedMetrics) == 1 {
		return -1
	}

	half := len(sw.sortedMetrics) / 2
	m := sw.sortedMetrics[half].value
	if len(sw.sortedMetrics)%2 == 0 {
		m = (m + sw.sortedMetrics[half-1].value) / 2
	}
	return m

}

// removes metric that was added first to current state of sliding window
func (sw *SlidingWindow) removeOldest() {
	toBeDelIdx := sw.oldest.index
	sw.oldest = sw.oldest.next
	copy(sw.sortedMetrics[toBeDelIdx:], sw.sortedMetrics[toBeDelIdx+1:])
	sw.sortedMetrics[sw.size-1] = nil
	sw.sortedMetrics = sw.sortedMetrics[:sw.size-1]
	for i := toBeDelIdx; i < len(sw.sortedMetrics); i = i + 1 {
		sw.sortedMetrics[i].index--
	}
}

// initializes first elem in sliding window
func (sw *SlidingWindow) addFirst(metric *metric) {
	sw.sortedMetrics = append(sw.sortedMetrics, metric)
	sw.oldest = metric
	sw.latest = metric
	metric.index = 0
}

// adds element to sliding window when capacity is less than size
func (sw *SlidingWindow) addWhenBelowCap(m *metric) {
	newElemIdx := sw.searchIndex(m)
	sw.sortedMetrics = append(sw.sortedMetrics, &metric{})
	copy(sw.sortedMetrics[newElemIdx+1:], sw.sortedMetrics[newElemIdx:])
	for i := newElemIdx + 1; i < len(sw.sortedMetrics); i = i + 1 {
		sw.sortedMetrics[i].index++
	}
	m.index = newElemIdx
	sw.sortedMetrics[newElemIdx] = m
}

// finds index in current sliding window state where new metric should be inserted
func (sw *SlidingWindow) searchIndex(metric *metric) int {
	return sort.Search(len(sw.sortedMetrics), func(i int) bool { return sw.sortedMetrics[i].value > metric.value })
}
