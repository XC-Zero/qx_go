package business_logic

import (
	"qx/pkg/common/essentials"
)

type Course struct {
	essentials.BasicModel
	CourseName  string `gorm:"type:VARCHAR(50);not null;" `
	TeacherId   string `gorm:"type:VARCHAR(38);not null;" `
	TeacherName string `gorm:"type:VARCHAR(50);not null;" `
}
