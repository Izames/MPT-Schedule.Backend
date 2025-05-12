package utils

import "MPT-Schedule/Models"

func TeacherWindowCheck(teachers []Models.TeacherModel) bool {
	fail := false
	for _, teacher := range teachers {
		if teacher.Window {
			continue
		} else {
			for _, tWeek := range teacher.Week {
				space := 0
				switcher := false
				for _, tDay := range tWeek.Lessons {
					if tDay && !switcher {
						space++
						switcher = true
					} else {
						switcher = false
					}
				}
				if space >= 2 {
					fail = true
				}
			}
		}
	}
	return fail
}
