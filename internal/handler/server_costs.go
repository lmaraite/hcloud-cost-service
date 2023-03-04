package handler

import (
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/lmaraite/hcloud-cost-service/internal/controller"
)

const (
	apiTokenHeader  = "API_TOKEN"
	errorMessageKey = "error_message"
)

func ServerCosts(c *gin.Context) {

	apiToken := c.GetHeader(apiTokenHeader)

	client := &hcloud.NewClient(hcloud.WithToken(apiToken)).Server

    controller := controller.NewServerCostsController(client)

	response, err := controller.CalculateServerCosts()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			errorMessageKey: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, structs.Map(&response))
}
