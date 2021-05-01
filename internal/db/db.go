package db 

import (
  "github.com/jinzhu/gorm"
	"fmt"
	"log"
  _ "github.com/go-sql-driver/mysql"
)

func dbConn(User, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Println(err)
		// return err
	}
	db.SingularTable(true)          //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	db.LogMode(true)                //打印sql语句
	//开启连接池 
	db.DB().SetMaxIdleConns(100)        //最大空闲连接
	db.DB().SetMaxOpenConns(10000)      //最大连接数
	db.DB().SetConnMaxLifetime(30)      //最大生存时间(s)

	return db
}

func GetDb() (conn *gorm.DB) {
	conn = dbConn("root", "zhangyinge", "localhost", "project_management", 3306)
	if conn != nil {
		fmt.Println("连接成功")
		return
	}
	fmt.Println("本次未获取到mysql连接")
	return conn
}