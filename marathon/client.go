package marathon

import (
	"crypto/tls"
	"fmt"

	"github.com/MustWin/gomarathon"
)

type Client interface {
	GetApp(appID string) (*App, error)
}

type Task struct {
	Alive       bool
	HostAddress string
}

type App struct {
	ID     string
	Tasks  []*Task
	Memory float32
}

func makeApp(app *gomarathon.Application) *App {
	tasks := make([]*Task, 0)
	a := &App{
		ID:     app.ID,
		Memory: app.Mem,
	}

	for _, t := range app.Tasks {
		tasks = append(tasks, makeTask(t))
	}

	a.Tasks = tasks
	return a
}

func makeTask(task *gomarathon.Task) *Task {
	t := &Task{
		Alive: true,
	}

	if len(task.HealthCheckResults) > 0 {
		t.Alive = task.HealthCheckResults[0].Alive
	}

	port := 80
	if len(task.Ports) > 0 {
		port = task.Ports[0]
	}

	t.HostAddress = fmt.Sprintf("%s:%d", task.Host, port)
	return t
}

type marathonClient struct {
	*gomarathon.Client
}

func NewClient() (Client, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	client, err := gomarathon.NewClient("http://marathon.service.consul:18080", nil, tlsConfig)
	if err != nil {
		return nil, err
	}

	return &marathonClient{
		Client: client,
	}, nil
}

func (c *marathonClient) GetApp(appID string) (*App, error) {
	r, err := c.Client.GetApp(appID)
	if err != nil {
		return nil, err
	}

	return makeApp(r.App), nil
}
