package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/zj0395/golib/conf"
	"github.com/zj0395/golib/golog"
)

func InitDB(dbConf *conf.DBConf) (*gorm.DB, error) {
	rawDsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(rawDsn, dbConf.Info.User, dbConf.Info.Pwd, dbConf.Info.Host, dbConf.Info.Port, dbConf.Info.Database)

	dbLogger := logger.New(
		golog.LogForwarder(),
		logger.Config{
			SlowThreshold:             time.Millisecond * 500, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: false,
			Colorful:                  false, // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
