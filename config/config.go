package config

import "gorm.io/gorm"

// 定义application参数
type application struct {
	Port        uint
	LogSavePath string
	LogFileName string
}

// 定义数据库参数
type database struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

var AppSetting *application
var DBSetting *database
var DB *gorm.DB
