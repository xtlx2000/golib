package util

import (
	"strconv"
)

func IntToString(num int) string {
	return strconv.Itoa(num)
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringToInt(numStr string) (int, error) {
	num, err := strconv.Atoi(numStr)
	return num, err
}

func StringToInt64(numStr string) (int64, error) {
	num, err := strconv.ParseInt(numStr, 10, 64)
	return num, err
}
