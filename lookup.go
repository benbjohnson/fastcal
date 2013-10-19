package fastcal

import (
  "sort"
)

type Lookup struct {
  epoch int64
  hour int
  day int
  month int
  year int
}

func (t *Lookup) Set(epoch int64) {
  t.epoch = epoch

  i := sort.Search(len(year_starts), func(i int) bool { return year_starts[i] >= epoch })
  t.year = 1999 + i
  offset := epoch - year_starts[i-1]

  if dst[i-1] {
    i = sort.Search(len(dst_months), func(i int) bool { return dst_months[i] >= offset })
    offset = offset - dst_months[i-1]
  } else {
    i = sort.Search(len(months), func(i int) bool { return months[i] >= offset })
    offset = offset - months[i-1]
  }
  t.month = i

  i = sort.Search(len(days), func(i int) bool { return days[i] >= offset })
  t.day = i+1
  offset -= days[i]

  i = sort.Search(len(hours), func(i int) bool { return hours[i] >= offset })
  t.hour = i+1

}

func (t *Lookup) Hour() int {
  return t.hour
}

func (t *Lookup) Day() int {
  return t.day
}

func (t *Lookup) Month() int {
  return t.month
}

func (t *Lookup) Year() int {
  return t.year
}

var year_starts = []int64{946684800,
 978307200,
 1009843200,
 1041379200,
 1072915200,
 1104537600,
 1136073600,
 1167609600,
 1199145600,
 1230768000,
 1262304000,
 1293840000,
 1325376000,
 1356998400,
 1388534400,
 1420070400,
 1451606400,
 1483228800,
 1514764800,
 1546300800,
 1577836800,
 1609459200,
 1640995200,
 1672531200,
 1704067200,
 1735689600,
 1767225600,
 1798761600,
 1830297600,
 1861920000,
 1893456000,
 1924992000,
 1956528000,
 1988150400,
 2019686400,
 2051222400,
 2082758400,
 2114380800,
 2145916800,
 2177452800,
 2208988800,
 2240611200,
 2272147200,
 2303683200,
 2335219200,
 2366841600,
 2398377600,
 2429913600,
 2461449600,
 2493072000,
 2524608000,
 2556144000,
 2587680000,
 2619302400,
 2650838400,
 2682374400,
 2713910400,
 2745532800,
 2777068800,
 2808604800,
 2840140800,
 2871763200,
 2903299200,
 2934835200,
 2966371200,
 2997993600,
 3029529600,
 3061065600,
 3092601600,
 3124224000,
 3155760000,
 3187296000,
 3218832000,
 3250454400,
 3281990400,
 3313526400,
 3345062400,
 3376684800,
 3408220800,
 3439756800,
 3471292800,
 3502915200,
 3534451200,
 3565987200}

var dst = []bool{true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false,
 true,
 false,
 false,
 false}

var months= []int64 {0,
 2678400,
 5097600,
 7776000,
 10368000,
 13046400,
 15638400,
 18316800,
 20995200,
 23587200,
 26265600,
 28857600}

var dst_months = []int64{0,
 2678400,
 5184000,
 7862400,
 10454400,
 13132800,
 15724800,
 18403200,
 21081600,
 23673600,
 26352000,
 28944000}

var days = []int64 {0,
 86400,
 172800,
 259200,
 345600,
 432000,
 518400,
 604800,
 691200,
 777600,
 864000,
 950400,
 1036800,
 1123200,
 1209600,
 1296000,
 1382400,
 1468800,
 1555200,
 1641600,
 1728000,
 1814400,
 1900800,
 1987200,
 2073600,
 2160000,
 2246400,
 2332800,
 2419200,
 2505600,
 2592000 }

var hours = []int64{0,
 3600,
 7200,
 10800,
 14400,
 18000,
 21600,
 25200,
 28800,
 32400,
 36000,
 39600,
 43200,
 46800,
 50400,
 54000,
 57600,
 61200,
 64800,
 68400,
 72000,
 75600,
 79200}
