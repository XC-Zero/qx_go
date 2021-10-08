package client

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	cfg "qx/pkg/common/config"
)

// InitGormV2 GormV2 returns a MySQL DB engine from config
func InitGormV2(config cfg.MySQLConfiguration) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
	var log logger.Interface
	if config.LogMode == cfg.None {
		log = logger.Default.LogMode(logger.Silent)
	} else {
		log = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: log,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)

	return db, nil
}
