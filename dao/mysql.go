//Author  :     knight
//Date    :     2020/07/11 21:01:12
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :		初始化MySQL数据库的连接

package dao

import (
	"fmt"
	"holl/config"
	"holl/model"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DB 数据库游标
var DB *gorm.DB

//InitMySQL 初始化数据库连接
func InitMySQL() error {
	var err error
	
	mysqlConfig := config.GetMysqlConfig()
	userAndPassword := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.URL, mysqlConfig.Name)
	
	//连接数据库
	DB, err = gorm.Open("mysql", userAndPassword)
	if err != nil {
		return err
	}

	// 设置连接池
	DB.DB().SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	DB.DB().SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	DB.DB().SetConnMaxLifetime(time.Hour * time.Duration(mysqlConfig.ConnMaxLifetime))

	if err = DB.DB().Ping(); err != nil {
		return err
	}

	log.Println("mysql init success")
	return nil
}

//Migrate 模型数据库的迁移
func Migrate() {
	if !DB.HasTable("deals") {
		DB.CreateTable(&model.Deal{})
		log.Println("deals table create success")
	} else {
		DB.Table("deals").AutoMigrate(&model.Deal{})
		log.Println("deals table exists")
	}

	if !DB.HasTable("users") {
		DB.CreateTable(&model.User{})
		log.Println("users table create success")
	} else {
		DB.Table("users").AutoMigrate(&model.User{})
		log.Println("users table exists")
	}

	if !DB.HasTable("orders") {
		DB.CreateTable(&model.Order{})
		log.Println("orders table create sucess")
	} else {
		DB.Table("orders").AutoMigrate(&model.Order{})
		log.Println("orders table exists")
	}

	if !DB.HasTable("images") {
		DB.CreateTable(&model.Image{})
		log.Println("images table create sucess")
	} else {
		DB.Table("images").AutoMigrate(&model.Image{})
		log.Println("images table exists")
	} 
}

//CloseMysql 关闭数据库连接
func CloseMysql() {
	log.Println("close mysql connection")
	DB.Close()
}
