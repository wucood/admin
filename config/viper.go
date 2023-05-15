package config

import (
	"github.com/spf13/viper"
	"log"
)

// NewSetting 读取配置文件
func newSetting() (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName("config") // 配置文件名
	vp.AddConfigPath("config") // 配置文件路径
	vp.SetConfigType("yaml")   // 配置文件格式
	err := vp.ReadInConfig()   // 读取配置文件
	if err != nil {
		return nil, err
	}
	return vp, nil
}

func readConfig(k string, v interface{}) {
	set, err := newSetting()
	if err != nil {
		log.Fatalln("读取配置文件失败: ", err)
	}
	err = set.UnmarshalKey(k, v)
	if err != nil {
		log.Fatalln("解析配置失败: ", err)
	}
}

func SetupConfig() {
	readConfig("App", &AppSetting)
	readConfig("Database", &DBSetting)
	log.Println("初始化配置完成")
}
