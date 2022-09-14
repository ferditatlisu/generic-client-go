package genericclient

import (
	"net/http"
	"time"
)

type ClientConfig interface {
	GetHost() string
	GetTimeout() time.Duration
	GetMaxIdleConnection() int
}

type clientConfig struct {
	builder *clientConfigBuilder
}

const (
	defaultTimeOut time.Duration = 30000
)

func NewClientConfig(builder *clientConfigBuilder) ClientConfig {
	return &clientConfig{builder: builder}
}

func (c *clientConfig) GetHost() string {
	if c.builder.host != "" {
		return c.builder.host
	}

	panic("Host field required")
}

func (c *clientConfig) GetTimeout() time.Duration {
	if c.builder.timeout != nil {
		return *c.builder.timeout
	}

	return defaultTimeOut
}

func (c *clientConfig) GetMaxIdleConnection() int {
	if c.builder.maxIdleConnection != nil {
		return *c.builder.maxIdleConnection
	}

	return http.DefaultMaxIdleConnsPerHost
}
