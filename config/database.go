package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// InitDB 初始化数据库，返回数据库对象
func initDB() (*gorm.DB, error) {
	dsn := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	dsn = fmt.Sprintf(dsn,
		DBSetting.Username,
		DBSetting.Password,
		DBSetting.Host,
		DBSetting.DBName,
		DBSetting.Charset,
		DBSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	// 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(DBSetting.MaxIdleConns)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(DBSetting.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, nil
}

func SetupDB() {
	var err error
	DB, err = initDB()
	if err != nil {
		log.Fatalln("初始化数据库失败: ", err)
	}
	log.Println("初始化数据库完成")
}
