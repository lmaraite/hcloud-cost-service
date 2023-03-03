package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

type HcloudServerClient interface {
	All(ctx context.Context) ([]*hcloud.Server, error)
}

type ServerCostsResponse struct {
	Monthly   float64
	Hourly    float64
	ErrorCode int
}

func CalculateCosts(client HcloudServerClient) (*ServerCostsResponse, error) {
	response := &ServerCostsResponse{}
	server, err := client.All(context.TODO())
	fmt.Println("DEBUG:", server)
	if err != nil {
		response.ErrorCode = http.StatusBadRequest
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
