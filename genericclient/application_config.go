package genericclient

import "time"

type EndPoint struct {
	Url string
}

type ApiConfig struct {
	Host              string
	Timeout           time.Duration
	MaxIdleConnection int
}
