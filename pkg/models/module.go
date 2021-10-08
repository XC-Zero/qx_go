package models

import (
	"qx/pkg/common/essentials"
)

type ModuleType string

const (
	TABLE ModuleType = "table"
	CHART ModuleType = "chart"
	SHEET ModuleType = "sheet"
)

type Module struct {
	essentials.BasicModel
	ModuleName *string    `gorm:"type:VARCHAR(50)"`
	ModuleType ModuleType `gorm:"type:VARCHAR(20)"`
	Remark     *string    `gorm:"type:VARCHAR(100)"`
}
