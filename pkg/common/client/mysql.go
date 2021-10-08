package client

import (
	"fmt"

	"github.com/jinzhu/gorm"
	cfg "qx/pkg/common/config"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitMySQL returns a MySQL DB engine from config
func InitMySQL(config cfg.MySQLConfiguration) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)

	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(20)
	if config.LogMode == cfg.None {
		db.LogMode(false)
	} else {
		db.LogMode(true)
	}

	return db, nil
}
