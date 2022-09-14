package genericclient

import "time"

type ClientConfigBuilder interface {
	Build() *ClientConfig
	SetMaxIdleConnection(m int) ClientConfigBuilder
	SetTimeout(millisecond time.Duration) ClientConfigBuilder
}

type clientConfigBuilder struct {
	host              string
	maxIdleConnection *int
	timeout           *time.Duration
}

func NewClientConfigBuilder(host string) ClientConfigBuilder {
	return &clientConfigBuilder{host: host}
}

func (c *clientConfigBuilder) SetMaxIdleConnection(millisecond int) ClientConfigBuilder {
	c.maxIdleConnection = &millisecond

	return c
}

func (c *clientConfigBuilder) SetTimeout(millisecond time.Duration) ClientConfigBuilder {
	c.timeout = &millisecond

	return c
}

func (c *clientConfigBuilder) Build() *ClientConfig {
	config := NewClientConfig(c)
	return &config
}
