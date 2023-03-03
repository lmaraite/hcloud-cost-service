package controller_test

import (
	"context"
	"testing"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/lmaraite/hcloud-cost-service/internal/controller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockHcloudServerClient struct {
	mock.Mock
}

func (m mockHcloudServerClient) All(ctx context.Context) ([]*hcloud.Server, error) {
	args := m.Called(ctx)
	return []*hcloud.Server{}, args.Error(1)
}

func TestServerCosts(t *testing.T) {
	mock := &mockHcloudServerClient{}

	server := make([]*hcloud.Server, 0)
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
				},
			},
		},
	}
	server = append(server, instance)

	mock.On("All", context.TODO()).Return(server, nil)
	response, err := controller.CalculateCosts(mock)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 13.13, response.Monthly)
}
