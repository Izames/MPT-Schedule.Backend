package utils

import (
	"MPT-Schedule/Models"
	"strings"
)

func FindTeacherByName(name string, semester int, rData *Models.RequestData) *Models.TeacherModel {
	name = strings.ReplaceAll(name, " ", "")
	parts := strings.Split(name, ".")
	reversedName := parts[len(parts)-1] + " " + strings.Join(parts[:len(parts)-1], ".") + "."
	if semester == 1 {
		for _, Teacher := range rData.TeachersS1 {

			if Teacher.FIO == reversedName {
				return &Teacher
			}
		}
	}
	if semester == 2 {
		for _, Teacher := range rData.TeachersS2 {

			if Teacher.FIO == reversedName {
				return &Teacher
			}
		}
	}

	return nil
}
