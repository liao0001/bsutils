package bsutils

import (
	"fmt"
	"testing"
)

func TestStringToIntWithDefault(b *testing.T) {
	fmt.Println("right int : ", StringToIntWithDefault("3301", 0))
	fmt.Println("right int : ", StringToIntWithDefault("03301", 0))
	fmt.Println("wrong int : ", StringToIntWithDefault("3301a", 1))
}

//func StringToInt64WithDefault(s string, def int64) int64 {
//	i, err := strconv.ParseInt(s, 10, 64)
//	if err != nil {
//		return def
//	}
//	return i
//}
//
//func StringToFloat64WithDefault(s string, def float64) float64 {
//	i, err := strconv.ParseFloat(s, 64)
//	if err != nil {
//		return def
//	}
//	return i
//}
//
//func StringToDateWithDefault(s string, def time.Time, formats ...string) time.Time {
//	if len(formats) == 0 {
//		formats = []string{}
//	}
//	formats = append(formats, time.RFC3339, time.RFC3339Nano, time.RFC1123Z, time.RFC1123)
//	for _, f := range formats {
//		if t, err := time.Parse(f, s); err == nil {
//			return t
//		}
//	}
//	return def
//}
