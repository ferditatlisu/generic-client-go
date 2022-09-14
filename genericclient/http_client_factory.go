package genericclient

type HttpClientFactory interface {
	GetClient(string) *GenericClient
}

type httpClientFactory struct {
	httpClients map[string]*GenericClient
}

var _httpClientFactory HttpClientFactory

func getClientFactory() HttpClientFactory {
	return _httpClientFactory
}

func NewHttpClientFactory() *httpClientFactory {
	clients := make(map[string]*GenericClient)
	factory := &httpClientFactory{httpClients: clients}
	_httpClientFactory = factory

	return factory
}

func (factory *httpClientFactory) GetClient(host string) *GenericClient {
	return factory.httpClients[host]
}

func (factory *httpClientFactory) AppendClient(config ClientConfig) {
	clientGeneric := NewGenericClient(config)
	factory.httpClients[config.GetHost()] = clientGeneric
}
