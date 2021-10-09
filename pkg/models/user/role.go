package user

/*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *
*
*   @Title:
*       Role
*
*   @Description:
*		对于qx而言 基础 role 分为 5 类，Administrative Staff , Teacher , Student , Parent , Manager
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
	RoleName   string  `gorm:"type:VARCHAR(20);not null"`
	RoleRemark *string `gorm:"type:VARCHAR(200);"`
	privileges string  `gorm:"type:VARCHAR(500);not null"`
}

func (r *Role) GrantPrivileges(privilege []Privilege) {
	privilegeStr := ""
	for i := range privilege {
		privilegeStr += privilege[i].ModuleID + ":" + strconv.Itoa(int(privilege[i].PrivilegeLevel)) + ";"
	}
	r.privileges = privilegeStr
}
