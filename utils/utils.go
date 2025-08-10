package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Contains[T any](arr []T, needle func(T) bool) int {
	for i := range arr {
		if needle(arr[i]) {
			return i
		}
	}
	return -1
}

func Filter[T any](ss []T, callback func(T) bool) (ret []T) {
	for _, s := range ss {
		if callback(s) {
			ret = append(ret, s)
		}
	}
	return
}

func IsInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func ValidateName(module string) (string, error) {

	module = strings.TrimSpace(module)
	module = strings.ToLower(module)

	pattern := `^[a-z][a-z0-9]*(?:_[a-z0-9]+)*$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(module) {
		return "", os.ErrInvalid
	}

	return module, nil
}

func Normalize(name string) (string, string, string, error) {

	name, err := ValidateName(name)
	if err != nil {
		fmt.Println("Filename is invalid")
		return "", "", "", err
	}

	splitName := strings.Split(name, "_")
	for i, s := range splitName {
		splitName[i] = cases.Title(language.English).String(s)
	}
	data := map[string]string{
		"Name":    strings.Join(splitName, ""),
		"NameVar": strings.ToLower(splitName[0]) + strings.Join(splitName[1:], ""),
	}

	return data["Name"], data["NameVar"], name, nil
}
