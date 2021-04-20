package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/zj0395/golib/conf"
)

func InitDB(dbConf *conf.DBConf) (*gorm.DB, error) {
	rawDsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(rawDsn, dbConf.Info.User, dbConf.Info.Pwd, dbConf.Info.Host, dbConf.Info.Port, dbConf.Info.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
