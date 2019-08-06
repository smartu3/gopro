package model

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB // 全局变量DB,sql.DB的指针

func Init() error {
	var err, error
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/devops?parseTime=true&charset=utf8&loc=Local") // 仅仅初始化sql.DB对象，在第一次查询才会建立Sql链接
	if nil != err {
		return err
	}
	// 设置最大闲置链接数
	DB.SetMaxIdleConns(50)
	
	// 建立链接
	err = DB.Ping() // Ping方法验证
	if nil != err{
		return err
	}else{
		log.Println("Mysql startup successfully.")
	}
	return DB

}