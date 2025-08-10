package utils

import "strconv"

func IsInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
