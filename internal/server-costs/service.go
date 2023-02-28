package servercosts

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/hetznercloud/hcloud-go/hcloud"
)

const (
	apiTokenHeader  = "API_TOKEN"
	errorMessageKey = "error_message"
)

type ServerCostsResponse struct {
	Monthly   float64
	Hourly    float64
	errorCode int
}

func ServerCosts(c *gin.Context) {

	apiToken := c.GetHeader(apiTokenHeader)

	response, err := serverCosts(apiToken)
	if err != nil {
		c.JSON(response.errorCode, gin.H{
			errorMessageKey: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, structs.Map(&response))
}

func serverCosts(apiToken string) (*ServerCostsResponse, error) {
	response := &ServerCostsResponse{}
	client := hcloud.NewClient(hcloud.WithToken(apiToken))
	server, err := client.Server.All(context.TODO())
	if err != nil {
		response.errorCode = http.StatusBadRequest
		return response, err
	}
	for _, instance := range server {
		datacenter := instance.Datacenter.Location.Name
		for _, pricing := range instance.ServerType.Pricings {
			if pricing.Location.Name == datacenter {
				pricePerMonth, _ := strconv.ParseFloat(pricing.Monthly.Gross, 64)
				pricePerHour, _ := strconv.ParseFloat(pricing.Hourly.Gross, 64)
				response.Monthly += pricePerMonth
				response.Hourly += pricePerHour
			}
		}
	}
	return response, nil
}
