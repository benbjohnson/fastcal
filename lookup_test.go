package fastcal

import (
	"testing"
)

func assert_equal(t *testing.T, x interface{}, y interface{}) {
	if x != y {
		t.Errorf("%v is not equal to %v", x, y)
	}
}

func TestLookup(t *testing.T) {
	var time Lookup
	time.Set(1382374809)
	assert_equal(t, time.Hour(), 18)
	assert_equal(t, time.Day(), 21)
	assert_equal(t, time.Month(), 10)
	assert_equal(t, time.Year(), 2013)
}

func TestYearSearch(t *testing.T) {
	assert_equal(t, yearSearch(1382374809), 43)
}

func TestBinarySearch(t *testing.T) {
	var buckets = []uint{0, 10, 20}
	assert_equal(t, binarySearch(buckets, 0), 0)
	assert_equal(t, binarySearch(buckets, 9), 0)
	assert_equal(t, binarySearch(buckets, 10), 1)
	assert_equal(t, binarySearch(buckets, 19), 1)
	assert_equal(t, binarySearch(buckets, 20), 2)
	assert_equal(t, binarySearch(buckets, 29), 2)
}

func TestBinarySearch2(t *testing.T) {
	var buckets = []uint{0, 10}
	assert_equal(t, binarySearch(buckets, 0), 0)
	assert_equal(t, binarySearch(buckets, 9), 0)
	assert_equal(t, binarySearch(buckets, 10), 1)
	assert_equal(t, binarySearch(buckets, 19), 1)
}
