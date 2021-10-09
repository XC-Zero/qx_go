package main

import (
	"flag"
	"fmt"
	"gorm.io/gorm"
	cfg "qx/internal/config"
	"qx/internal/utils/mysql"
	"qx/pkg/common/config"
	"qx/pkg/models/record"
	"qx/pkg/models/user"

	"strings"
)

var configFileName = flag.String("cfn", "config", "name of config file")
var configFilePath = flag.String("cfp", "./configs/", "path of config file")

func main() {
	flag.Parse()
	err := config.InitConfiguration(*configFileName, strings.Split(*configFilePath, ","), &cfg.CONFIG)
	if err != nil {
		panic(err)
	}
	mysql.InitUserDatapackDB()
	users := []interface{}{
		&user.Role{},
		&user.User{},
		&user.Privilege{},
		&user.Module{},
		&user.Account{},
	}
	records := []interface{}{
		&record.AccountRecord{},
		&record.LoginRecord{},
	}
	var tables []interface{}
	tables = append(tables, records...)
	tables = append(tables, users...)

	createDeltaTables(mysql.UserDatapackDB, tables)
}

func createDeltaTables(db *gorm.DB, tables []interface{}) {
	err := db.AutoMigrate(tables...)
	if err != nil {
		fmt.Println(err)
	}
}
