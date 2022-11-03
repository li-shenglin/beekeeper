package socket

import (
	"backend/common"
	"time"
)

type Client struct {
	*Application
	address  []string
	clientID []byte
}

func (client *Client) Run() {
	client.instances = make([]*Instance, len(client.address))
	for i := range client.address {
		client.instances[i] = NewClientInstance(client.address[i], client.dispatch)
		client.instances[i].Start()
	}
	for {
		client.UpStatus()
	}
}

func (application *Application) UpStatus() {
	for i := range application.instances {
		if application.instances[i].Status == InstanceOnline {
			invoke, err := application.instances[i].Invoke(&Parameter{
				Opt:  0,
				Data: []byte("Hello Server"),
			})
			log.Infof("err=%s, return=%s", err, invoke)
		}
	}
	time.Sleep(time.Second * 5)

}

func NewClient(address ...string) *Client {
	if len(address) == 0 {
		panic("server address not found")
	}
	return &Client{
		Application: GetApplication(),
		address:     address,
		clientID:    common.UUID(),
	}
}
