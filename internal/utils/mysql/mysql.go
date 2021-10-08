package mysql

import (
	"gorm.io/gorm"
	"log"
	"qx/internal/config"
	"qx/pkg/common/client"
)

var BatchSize = 20000

var UserDatapackDB *gorm.DB

func InitUserDatapackDB() {
	if UserDatapackDB != nil {
		return
	}
	db, err := client.InitGormV2(config.CONFIG.UserDatapackConfig)
	if err != nil {
		panic(err)
	}
	log.Println("InitDealDatapackDB success")
	UserDatapackDB = db
}
