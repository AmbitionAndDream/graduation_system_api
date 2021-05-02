package global

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var sqlDB *sql.DB
var db *gorm.DB

func initMysql() (err error) {
	config := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBConfig.User,
		config.DBConfig.PassWord,
		config.DBConfig.Server,
		config.DBConfig.Ports,
		config.DBConfig.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{SlowThreshold: time.Nanosecond,
			LogLevel: logger.Info,
			Colorful: true})})
	if err != nil {
		logrus.Errorf("customize driver failed, error is : %s", err.Error())
		return err
	}
	sqlDB, err = db.DB()
	if err != nil {
		logrus.Errorf("get generic database sqlDB failed, error is : %s", err.Error())
	}
	sqlDB.SetMaxIdleConns(config.DBConfig.Connection.MaxIdleConn)
	sqlDB.SetMaxOpenConns(config.DBConfig.Connection.MaxOpenConn)
	sqlDB.SetConnMaxIdleTime(time.Millisecond * time.Duration(config.DBConfig.Connection.ConnMaxIdleTime))

	return nil
}
func GetDb() *gorm.DB {
	if db == nil {
		panic("mysqlDB not initialize")
	}
	return db
}
func closeMysql() {
	if sqlDB != nil {
		sqlDB.Close()
	}
}
