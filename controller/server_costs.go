package controller

import (
	"context"
	"strconv"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

type HcloudServerClient interface {
	All(ctx context.Context) ([]*hcloud.Server, error)
}

type ServerCostsResponse struct {
	Monthly float64
	Hourly  float64
}

type ServerCostsController struct {
	Client HcloudServerClient
}

func NewServerCostsController(client HcloudServerClient) *ServerCostsController {
	return &ServerCostsController{
		Client: client,
	}
}

func (scc *ServerCostsController) CalculateServerCosts() (*ServerCostsResponse, error) {
	response := &ServerCostsResponse{}
	server, err := scc.Client.All(context.TODO())
	if err != nil {
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
