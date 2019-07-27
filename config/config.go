package config

import (
	"github.com/spf13/viper"
	"time"
	"os"
	"log"
)

// LogInfo 初始化日志
func LogInfo()  {
	file := "./" + time.Now().Format("2019-07-26") + ".log"
	logFile, _ := os.OpenFile(file, os.O_RDWR| os.O_CREATE| os.O_APPEND, 0766)
	log.SetFlags(log.Ldata| log.Ltime | log.Lshortfile)
	log.SetOutPut(logFile)
}

// Init 读取初始化配置
func Init() error(
	// 初始化配置文件
	if err:= Config(); err != nil{
		return err
	}
	// 初始化日志包
	LogInfo()
	return nil
)

// Config viper 解析配置文件
func Config() error{
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil{
		return err
	}
}


