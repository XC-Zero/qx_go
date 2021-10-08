package record

/*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *
*
*   @Title:
*       Account record
*
*   @Description:
*		账户交易记录
*
*   @Remarks:
*		AccountID		账户ID 逻辑外键
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

type AccountChangeType int

const (
	RECHARGE AccountChangeType = iota + 1
	CONSUME
	REFUND
	OTHER
)

type AccountRecord struct {
	essentials.BasicModel
	AccountID         string
	ChangeAmount      float64
	AccountChangeType AccountChangeType
}
