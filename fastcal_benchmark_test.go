package fastcal

import (
  "testing"
)

// 10MM events per second, so the entire query has 100ns.
var epoch = 1287429458

// 300ns to convert
func BenchmarkStdlibEpochToCalendar(b *testing.B) {
  t := Lookup{}

  for i := 0; i < b.N; i++ {
    t.Set((int64)(epoch))
    _ = t.Hour()
    _ = t.Day()
    _ = t.Month()
    _ = t.Year()
  }
}


