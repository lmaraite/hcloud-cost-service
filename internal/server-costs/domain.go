package servercosts

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

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
