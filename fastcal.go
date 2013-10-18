package fastcal

type Fastcal interface {
  Set(int64)
  Hour() int
  Day() int
  Month() int
  Year() int
}


