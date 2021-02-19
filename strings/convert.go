package strings

import (
	"strconv"
	"time"
)

/**
string 转换
*/
func StrToInt64(s string) (i int64, err error) {
	i, err = strconv.ParseInt(s, 10, 64)
	return i, err
}

func StrToInt32(s string) (i int64, err error) {
	i, err = strconv.ParseInt(s, 10, 32)
	return i, err
}

func StrToInt(s string) (i int, err error) {
	i, err = strconv.Atoi(s)
	return i, err
}

func StrToFloat64(s string) (f float64, err error) {
	f, err = strconv.ParseFloat(s, 64)
	return f, err
}

func StrToByte(s string) []byte {
	return []byte(s)
}

func StrToTime(layout, value string, loc ...*time.Location) (time.Time, error) {
	location := time.Local
	if len(loc) > 0 {
		location = loc[0]
	}
	return time.ParseInLocation(layout, value, location)
}

/**
int 转换
*/
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}
