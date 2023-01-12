package util

import "strconv"

func ToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func ToInt64Or(str string, defaultValue int64) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}
	return i
}
