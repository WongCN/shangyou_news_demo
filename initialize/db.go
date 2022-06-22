package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB
var Err error

func init() {
	dsn := "root:password@tcp(127.0.0.1:3306)/shangyou_news?charset=utf8mb4&parseTime=True&loc=Local"
	DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{ //建立连接时指定打印info级别的sql
		Logger: logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
	})
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db, err := DB.DB()
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)
	if Err == nil {
		fmt.Println(Err)
		fmt.Println(err)
	}
}
