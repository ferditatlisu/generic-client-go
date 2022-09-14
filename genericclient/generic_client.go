package genericclient

import (
	"net"
	"net/http"
	"time"
)

type GenericClient struct {
	host   string
	client *http.Client
}

func NewGenericClient(builder ClientConfig) *GenericClient {
	to := builder.GetTimeout()
	timeout := to * time.Millisecond

	tr := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		MaxIdleConns: builder.GetMaxIdleConnection(),
	}

	client := http.Client{
		Transport: &tr,
		Timeout:   timeout,
	}

	genericClient := GenericClient{
		host:   builder.GetHost(),
		client: &client,
	}

	return &genericClient
}
