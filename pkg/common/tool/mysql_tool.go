package tool

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
	"qx/pkg/common/connector"
)

func CalcMysqlBatchSize(data interface{}) int {
	count := countStructFields(reflect.ValueOf(data))
	return 60000 / count
}

func countStructFields(v reflect.Value) int {
	var count int
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Type.Kind() == reflect.Struct && connector.IsInnerStruct(field) {
			num := countStructFields(v.Field(i))
			count += num
		} else if !strings.Contains(field.Tag.Get("gorm"), "foreignKey") {
			count += 1
		}
	}
	return count
}

func AddForeignKey(db *gorm.DB, ownTable string, colName string, dependentTable string, dependentColName string, isCascade bool, fkName string) {
	if fkName == "" {
		fkName = "fk_" + colName + "_" + strconv.Itoa(int(time.Now().UnixNano()))
	}
	sqlExec := "ALTER TABLE " + ownTable + " ADD CONSTRAINT " + fkName + " FOREIGN KEY(" + colName + ") REFERENCES " +
		dependentTable + "(" + dependentColName + ")"
	if isCascade {
		sqlExec += " ON UPDATE CASCADE ON DELETE CASCADE"
	}
	err := db.Exec(sqlExec).Error
	if err != nil {
		panic(err)
	}
}

func AddForeignKeyWithSetNull(db *gorm.DB, ownTable string, colName string, dependentTable string, dependentColName string, isCascade bool, fkName string) {
	if fkName == "" {
		fkName = "fk_" + colName + "_" + strconv.Itoa(int(time.Now().UnixNano()))
	}
	sqlExec := "ALTER TABLE " + ownTable + " ADD CONSTRAINT " + fkName + " FOREIGN KEY(" + colName + ") REFERENCES " +
		dependentTable + "(" + dependentColName + ")"
	if isCascade {
		sqlExec += " ON UPDATE CASCADE ON DELETE SET NULL"
	}
	err := db.Exec(sqlExec).Error
	if err != nil {
		panic(err)
	}
}
