package utils

import (
	"MPT-Schedule/Models"
	"strings"
)

func FindTeacherByName(name string) *Models.TeacherModel {
	name = strings.ReplaceAll(name, " ", "")
	parts := strings.Split(name, ".")
	reversedName := parts[len(parts)-1] + " " + strings.Join(parts[:len(parts)-1], ".") + "."

	for _, Teacher := range Models.Teachers {

		if Teacher.FIO == reversedName {
			return &Teacher
		}
	}
	return nil
}
