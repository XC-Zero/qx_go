package user

/*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *
*
*   @Title:
*       Account
*
*   @Description:
*		账户表
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
)

type Account struct {
	essentials.BasicModel
	// 余额
	Balances float64 `gorm:"type:decimal(20,2);not null;default:0.00"`
	// 累计充值
	AccumulatedRecharge float64 `gorm:"type:decimal(20,2);not null;default:0.00"`
}
