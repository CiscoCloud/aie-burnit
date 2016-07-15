package marathon

type mockClient struct{}

func NewMockClient() Client {
	return &mockClient{}
}

func (c *mockClient) GetApp(appID string) (*App, error) {
	return &App{
		ID:     appID,
		Memory: 512.0,
		Tasks: []*Task{
			&Task{
				Alive:       true,
				HostAddress: "localhost:8887",
			},
			&Task{
				Alive:       true,
				HostAddress: "localhost:8888",
			},
		},
	}, nil
}
