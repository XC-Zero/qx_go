package models

import (
	"qx/pkg/common/essentials"
)

type PrivilegeLevel int

const (
	READ PrivilegeLevel = iota + 1
	WRITE
	MANAGE
	SPECIAL
)

type Privilege struct {
	essentials.BasicModel
	ModuleID       string         `gorm:"type:VARCHAR(38);not null"`
	PrivilegeLevel PrivilegeLevel `gorm:"type:VARCHAR(2);not null"`
}
