package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lmaraite/hcloud-cost-service/handler"
	"github.com/spf13/viper"
)

const (
	configPort  = "PORT"
	defaultPort = 8080
)

func main() {
	readConfig()

	route := gin.Default()

	costs := route.Group("/costs")
	costs.GET("/server", func(c *gin.Context) {
		handler := handler.DefaultServerCostsHandler(c)
		handler.ServerCosts(c)
	})

	route.Run(getFormattedPort())
}

func getFormattedPort() string {
	var port int
	if viper.InConfig(configPort) {
		port = viper.GetInt(configPort)
	} else {
		port = defaultPort
	}
	return fmt.Sprintf(":%d", port)
}

func readConfig() {
	viper.SetConfigName(".hcloud")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}
}
