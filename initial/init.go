package initial

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"my-blog/db"
	"my-blog/router"
	"os"
)

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}

func ginServer() {
	port := viper.GetString("server.port")
	r := gin.Default()
	router.AddRoutes(r)
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err.Error())
	}
}

func Server() {
	db.Server()
	ginServer()
}
