package marathon

import (
	"github.com/MustWin/gomarathon"
)

type mockClient struct{}

func NewMockClient() Client {
	return &mockClient{}
}

func (c *mockClient) GetApp(appID string) (*gomarathon.Application, error) {
	return &gomarathon.Application{
		ID: appID,
		Tasks: []*gomarathon.Task{
			&gomarathon.Task{
				AppID: appID,
				Host:  "localhost",
				Ports: []int{8888},
			},
			&gomarathon.Task{
				AppID: appID,
				Host:  "localhost",
				Ports: []int{8888},
			},
		},
		Mem: 512.0,
	}, nil
}
