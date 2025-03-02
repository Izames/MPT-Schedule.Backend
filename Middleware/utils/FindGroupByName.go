package utils

import "MPT-Schedule/Models"

func FindGroupByName(name string) int {
	for i, group := range Models.Groups {
		if group.Name == name {
			return i
		}
	}
	return -1
}
