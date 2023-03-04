package handler

import (
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/lmaraite/hcloud-cost-service/controller"
)

const (
	apiTokenHeader  = "API_TOKEN"
	errorMessageKey = "error_message"
)

type ServerCostsController interface {
    CalculateServerCosts() (*controller.ServerCostsResponse, error)
}

type ServerCostsHandler struct {
    Controller ServerCostsController
}

func NewServerCostsHandler(controller ServerCostsController) *ServerCostsHandler {
    return &ServerCostsHandler{
        Controller: controller,    
    }
}

func DefaultServerCostsHandler(c *gin.Context) *ServerCostsHandler {
    apiToken := c.GetHeader(apiTokenHeader)
    client := &hcloud.NewClient(hcloud.WithToken(apiToken)).Server
    controller := controller.NewServerCostsController(client)
    return &ServerCostsHandler{
        Controller: controller,
    }
}

func (handler *ServerCostsHandler) ServerCosts(c *gin.Context) {
	response, err := handler.Controller.CalculateServerCosts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			errorMessageKey: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, structs.Map(&response))
}
