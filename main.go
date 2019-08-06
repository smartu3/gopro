package main

import (
	"./config"
	"./model"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/spf13/viper"
	"./router"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	if err := model.Init();  err != nil {
		panic(err)
	}
	gin.SetMode(viper.GetString("runmode"))

	gin := gin.New()
	router.InitRouter(g)
	log.Printf("Start to listening the incoming requests on http address: %s \n", viper.GetString("addr"))

	if err := g.Run(viper.GetString("addr"); err != nil {
		log.Fatal("ListenAndServer", err)
	}
}