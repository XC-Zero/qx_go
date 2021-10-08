package migrate

import (
	"fmt"
	"gorm.io/gorm"
)

func Migrate() {

}

func createDeltaTables(db *gorm.DB, tables []interface{}) {
	err := db.AutoMigrate(tables...)
	if err != nil {
		fmt.Println(err)
	}
}
