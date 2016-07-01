package marathon

import (
	"crypto/tls"
	"time"

	"github.com/MustWin/gomarathon"
)

const CACHE_TIMEOUT int64 = 20000

type Client interface {
	GetApp(appID string) (*gomarathon.Application, error)
}

type marathonClient struct {
	*gomarathon.Client
	lastResult *gomarathon.Application
	lastLookup int64
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
		Client:     client,
		lastResult: nil,
		lastLookup: 0,
	}, nil
}

func (c *marathonClient) GetApp(appID string) (*gomarathon.Application, error) {
	if c.lastResult != nil && time.Now().Unix() <= (CACHE_TIMEOUT+c.lastLookup) {
		return c.lastResult, nil
	}

	c.lastLookup = time.Now().Unix()
	r, err := c.Client.GetApp(appID)
	if err != nil {
		return nil, err
	}

	c.lastResult = r.App
	return r.App, nil
}
