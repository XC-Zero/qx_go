package association

import (
	"qx/pkg/common/essentials"
)

type ClassTeacherAssociation struct {
	essentials.BasicModel
	ClassId             string  `gorm:"type:VARCHAR(38);not null;" `
	ClassName           string  `gorm:"type:VARCHAR(50);not null" `
	TeacherId           string  `gorm:"type:VARCHAR(38);not null;" `
	TeacherName         string  `gorm:"type:VARCHAR(50);not null" `
	TeacherPositionId   *string `gorm:"type:VARCHAR(38);" `
	TeacherPositionName *string `gorm:"type:VARCHAR(50);" `
}

// TeacherPosition todo 同 class_student_association
type TeacherPosition struct {
	essentials.BasicModel
	TeacherId         string  `gorm:"type:VARCHAR(38);index;not null;" `
	PositionName      string  `gorm:"type:VARCHAR(50);not null;" `
	PositionIntroduce *string `gorm:"type:VARCHAR(200);" `
}
