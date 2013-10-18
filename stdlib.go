package fastcal

import (
  "time"
)

type GoStdlib struct {
  timeobj time.Time
}

func (t *GoStdlib) Set(epoch int64) {
  t.timeobj = time.Unix(epoch, 0)
}

func (t *GoStdlib) Hour() int {
  return t.timeobj.Hour()
}

func (t *GoStdlib) Day() int {
  return t.timeobj.Day()
}

func (t *GoStdlib) Month() int {
  return (int)(t.timeobj.Month())
}

func (t *GoStdlib) Year() int {
  return t.timeobj.Year()
}

