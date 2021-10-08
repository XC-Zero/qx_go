package models

import (
	"qx/pkg/common/essentials"
)

type User struct {
	essentials.BasicModel
	UserName      string  `gorm:"type:VARCHAR(50);not null"`
	UserAlias     *string `gorm:"type:VARCHAR(50);"`
	UserLegalName *string `gorm:"type:VARCHAR(50);"`
	UserEmail     string  `gorm:"type:VARCHAR(30);not null;index;"`
	UserPhone     *string `gorm:"type:VARCHAR(20);index;"`
	RoleId        string  `gorm:"type:VARCHAR(38);not null"`
	RoleName      string  `gorm:"type:VARCHAR(20);not null"`
	AccountID     string  `gorm:"type:VARCHAR(38);not null"`
	Balances      float64 `gorm:"type:decimal(20,4);not null;default:0.0"`
	Wechat
}

type Wechat struct {
	UserOpenId    *string `gorm:"type:VARCHAR(50);"`
	UserNickName  *string `gorm:"type:VARCHAR(50);"`
	UserGender    *int    `gorm:"type:VARCHAR(2);"`
	UserPhone     *string `gorm:"type:VARCHAR(20);"`
	UserCountry   *string `gorm:"type:VARCHAR(50);"`
	UserProvince  *string `gorm:"type:VARCHAR(50);"`
	UserCity      *string `gorm:"type:VARCHAR(50);"`
	UserAvatarUrl *string `gorm:"type:VARCHAR(300);"`
}
