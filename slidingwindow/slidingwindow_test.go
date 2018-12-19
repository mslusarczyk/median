package slidingwindow

import (
	"math/rand"
	"testing"
	"time"
)

func TestGetMedianSingleElem(t *testing.T) {
	// given
	window := NewSlidingWindow(3)
	window.AddDelay(1)

	// when
	median := window.GetMedian()

	//then
	excepted := -1
	if median != excepted {
		t.Fatalf("%v != %v", median, excepted)
	}
}

func TestGetMedianOddShort(t *testing.T) {
	// given
	window := NewSlidingWindow(3)
	window.AddDelay(1)
	window.AddDelay(3)
	window.AddDelay(2)
	window.AddDelay(4)
	window.AddDelay(7)
	window.AddDelay(6)
	window.AddDelay(9)

	// when
	median := window.GetMedian()

	//then
	excepted := 7
	if median != excepted {
		t.Fatalf("%v != %v", median, excepted)
	}
}

func TestGetMedianEvenShort(t *testing.T) {
	// given
	window := NewSlidingWindow(4)
	window.AddDelay(1)
	window.AddDelay(3)
	window.AddDelay(2)
	window.AddDelay(4)
	window.AddDelay(6)
	window.AddDelay(9)

	// when
	median := window.GetMedian()

	//then
	excepted := 5
	if median != excepted {
		t.Fatalf("%v != %v", median, excepted)
	}
}

func TestGetMedianOddLong(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	// given
	rand.Seed(time.Now().Unix())
	window := NewSlidingWindow(9999)
	for i := 0; i < 200000; i++ {
		window.AddDelay(rand.Intn(1000))
		window.AddDelay(rand.Intn(1000) + 1001)
	}

	window.AddDelay(1000)

	// when
	median := window.GetMedian()

	//then
	expected := 1000
	if median != expected {
		t.Fatalf("%v != %v", median, expected)
	}
}

func TestGetMedianEvenLong(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	// given
	rand.Seed(time.Now().Unix())
	window := NewSlidingWindow(10000)
	for i := 0; i < 200000; i++ {
		window.AddDelay(rand.Intn(1000))
		window.AddDelay(rand.Intn(1000) + 1003)
	}

	window.AddDelay(1000)
	window.AddDelay(1002)

	// when
	median := window.GetMedian()

	//then
	expected := 1001
	if median != expected {
		t.Fatalf("%v != %v", median, expected)
	}
}

func BenchmarkAddDelaySize100(b *testing.B) {
	window := NewSlidingWindow(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		window.AddDelay(i)
	}
}

func BenchmarkAddDelaySize1000(b *testing.B) {
	window := NewSlidingWindow(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		window.AddDelay(i)
	}
}

func BenchmarkAddDelaySize10000(b *testing.B) {
	window := NewSlidingWindow(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		window.AddDelay(i)
	}
}
