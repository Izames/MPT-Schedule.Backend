package utils

import "MPT-Schedule/Models"

func FindGroupByName(name string, rData *Models.RequestData) int {
	for i, group := range rData.Groups {
		if group.Name == name {
			return i
		}
	}
	return -1
}
