package utils

import "strings"

func Contains(value string, array []string) bool {
	for _, v := range array {
		if strings.ToLower(v) == strings.ToLower(value) {
			return true
		}
	}
	return false
}
