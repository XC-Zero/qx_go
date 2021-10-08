package user

/*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *
*
*   @Title:
*       Role
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
*       2021/10/8
*
*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  */
import (
	"qx/pkg/common/essentials"
	"strconv"
)

// Role 预定义角色表
type Role struct {
	essentials.BasicModel
	RoleName   string `gorm:"type:VARCHAR(20);not null"`
	privileges string `gorm:"type:VARCHAR(500);not null"`
}

func (r *Role) GrantPrivileges(privilege []Privilege) {
	privilegeStr := ""
	for i := range privilege {
		privilegeStr += privilege[i].ModuleID + ":" + strconv.Itoa(int(privilege[i].PrivilegeLevel)) + ";"
	}
	r.privileges = privilegeStr
}
