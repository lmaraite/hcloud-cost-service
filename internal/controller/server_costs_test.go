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
	mock := hcloudServerClientMock{}
	instance := &hcloud.Server{
		Name: "instance-01",
		Datacenter: &hcloud.Datacenter{
			Location: &hcloud.Location{
				Name: "ffm-1",
			},
		},
		ServerType: &hcloud.ServerType{
			Pricings: []hcloud.ServerTypeLocationPricing{
				{
					Location: &hcloud.Location{
						Name: "ffm-1",
					},
					Monthly: hcloud.Price{
						Gross: "13.13",
					},
                    Hourly: hcloud.Price{
                        Gross: "0.018",
                    },
				},
			},
		},
	}
    mock.server = []*hcloud.Server{instance}
    mock.err = nil

    response, err := controller.CalculateCosts(mock)
    assert.NoError(t, err)
    assert.Equal(t, 13.13, response.Monthly)
    assert.Equal(t, 0.018, response.Hourly)

}
