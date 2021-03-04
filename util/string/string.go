package string

import (
	"strings"
)

func TrimQuotes(s string) string {
	return strings.Replace(strings.Replace(s, "\"", "", -1), "'", "", -1)
}

func InArray(val string, array []string) bool{
	var exists = false
	for _, v := range array {
		if val == v {
			exists = true
		}
	}
	return exists
}

