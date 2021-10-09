package record

/*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *
*
*   @Title:
*       Login record
*
*   @Description:
*		登录信息表
*
*   @Remarks:
*		UserId		用户ID 逻辑外键
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
	"qx/pkg/common/ip"
)

type LoginRecord struct {
	essentials.BasicModel
	UserId        string  `gorm:"type:VARCHAR(38);not null;index;"`
	IPAddress     *string `gorm:"type:VARCHAR(20);"`
	IPRealAddress *string `gorm:"type:VARCHAR(100);"`
}

func (l *LoginRecord) TransferIPToRealAddress() {
	if l.IPAddress != nil {
		l.IPRealAddress = &ip.GetIpRealAddress(*l.IPAddress).Data.City
	}
}
