package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	servercosts "github.com/lmaraite/hcloud-cost-service/internal/server-costs"
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
	costs.GET("/server", servercosts.ServerCosts)

	var port int
	if viper.InConfig(configPort) {
		port = viper.GetInt(configPort)
	} else {
		port = defaultPort
	}
	route.Run(fmt.Sprintf(":%d", port))
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
