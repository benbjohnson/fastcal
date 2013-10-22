package fastcal

import (
	"testing"
)

// 10MM events per second, so the entire query has 100ns.
var epoch = uint(1287429458)

// 300ns to convert
func BenchmarkStdlibEpochToCalendar(b *testing.B) {
	t := Lookup{}

	for i := 0; i < b.N; i++ {
		t.Set(epoch)
		_ = t.Hour()
		_ = t.Day()
		_ = t.Month()
		_ = t.Year()
	}
}

func BenchmarkYearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, time := range year_starts[1:] {
			yearSearch(time - 2*30*24*3600)
			yearSearch(time - 8*30*24*3600)
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	slice := year_starts[:]
	for i := 0; i < b.N; i++ {
		for _, time := range year_starts[1:] {
			binarySearch(slice, time-2*30*24*3600)
			binarySearch(slice, time-8*30*24*3600)
		}
	}
}
