package fastcal

type Fastcal interface {
	Set(uint)
	Hour() int
	Day() int
	Month() int
	Year() int
}
