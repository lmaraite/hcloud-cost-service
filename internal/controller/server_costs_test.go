package controller_test

import (
	"context"
	"testing"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/lmaraite/hcloud-cost-service/internal/controller"
	"github.com/stretchr/testify/assert"
)

type hcloudServerClientMock struct {
	server []*hcloud.Server
	err    error
}

func (mock hcloudServerClientMock) All(ctx context.Context) ([]*hcloud.Server, error) {
	return mock.server, mock.err
}

func TestServerCosts(t *testing.T) {
	testCases := []struct {
		name     string
		server   []*hcloud.Server
		err      error
		expected *controller.ServerCostsResponse
	}{
		{
			name: "TestSingleServer",
			server: []*hcloud.Server{
				createServer("test-server", "ffm-1", createHcloudPricings([]serverTypeLocationPricing{
					{
						locationName: "ffm-1",
						monthlyPrice: "15.13",
						hourlyPrice:  "0.012",
					},
					{
						locationName: "ber-1",
						monthlyPrice: "17.13",
						hourlyPrice:  "0.015",
					},
				},
				)),
			},
			err: nil,
			expected: &controller.ServerCostsResponse{
				Monthly: 13.13,
				Hourly:  0.018,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mock := hcloudServerClientMock{
				server: tc.server,
				err:    tc.err,
			}
			response, err := controller.CalculateServerCosts(mock)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, response)
		})
	}

}

type serverTypeLocationPricing struct {
	locationName string
	monthlyPrice string
	hourlyPrice  string
}

func createHcloudPricings(pricings []serverTypeLocationPricing) []hcloud.ServerTypeLocationPricing {
	result := make([]hcloud.ServerTypeLocationPricing, 0)
	for _, pricing := range pricings {
		result = append(result, hcloud.ServerTypeLocationPricing{
			Location: &hcloud.Location{
				Name: pricing.locationName,
			},
			Monthly: hcloud.Price{
				Gross: pricing.monthlyPrice,
			},
			Hourly: hcloud.Price{
				Gross: pricing.hourlyPrice,
			},
		})
	}
	return nil
}

func createServer(name string, locationName string, pricings []hcloud.ServerTypeLocationPricing) *hcloud.Server {
	return &hcloud.Server{
		Name: name,
		Datacenter: &hcloud.Datacenter{
			Location: &hcloud.Location{
				Name: locationName,
			},
		},
		ServerType: &hcloud.ServerType{
			Pricings: pricings,
		},
	}
}
