package association

/*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *
*
*   @Title:
*       Class student association
*
*   @Description:
*
*
*   @Remarks:
*
*
*   @Functions:
*
*
*   @Author:
*       XiangChen
*
*   @Date:
*       2021/10/9
*
*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  */
import (
	"qx/pkg/common/essentials"
)

type ClassStudentAssociation struct {
	essentials.BasicModel
	ClassId             string  `gorm:"type:VARCHAR(38);not null;" `
	ClassName           string  `gorm:"type:VARCHAR(50);not null" `
	StudentId           string  `gorm:"type:VARCHAR(38);not null;" `
	StudentName         string  `gorm:"type:VARCHAR(50);not null" `
	StudentPositionId   *string `gorm:"type:VARCHAR(38);" `
	StudentPositionName *string `gorm:"type:VARCHAR(50);" `
}

// StudentPosition todo 内置职务，id设置为前10个，显示时 select * from student_positions where teacher_id = ? union select * from student_positions where rec_id
type StudentPosition struct {
	essentials.BasicModel
	TeacherId         string  `gorm:"type:VARCHAR(38);index;not null;" `
	PositionName      string  `gorm:"type:VARCHAR(50);not null;" `
	PositionIntroduce *string `gorm:"type:VARCHAR(200);" `
}
