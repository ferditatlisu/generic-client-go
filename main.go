package genericclient

import (
	"fmt"
	"github.com/ferditatlisu/generic-client-go/example"
	"github.com/ferditatlisu/generic-client-go/genericclient"
)

func main() {
	applicationConfig := getConfig()
	CreateClients(applicationConfig)
	get(applicationConfig)
	put(applicationConfig)
	post(applicationConfig)
	delete(applicationConfig)

}

func post(applicationConfig *example.ApplicationConfig) {
	rp := example.NewRequestPayload(1)

	httpMethod := genericclient.NewGenericHttpMethod[example.RequestPayload, example.ResponsePayload](
		applicationConfig.LocalApi.Api.Host,
		applicationConfig.LocalApi.PostDefault)

	responsePayload, failedResponse := httpMethod.Post(*rp)
	if failedResponse != nil && failedResponse.Err != nil {
		fmt.Println(failedResponse.Err)
	}

	fmt.Println(responsePayload.IsSuccess)
}

func put(applicationConfig *example.ApplicationConfig) {
	rp := example.NewRequestPayload(1)
	defaultEndPoint := applicationConfig.LocalApi.PutDefault
	defaultEndPoint.Url = fmt.Sprintf(defaultEndPoint.Url, "2222")

	httpMethod := genericclient.NewGenericHttpMethod[example.RequestPayload, example.ResponsePayload](
		applicationConfig.LocalApi.Api.Host,
		defaultEndPoint)

	responsePayload, failedResponse := httpMethod.Put(*rp)
	fmt.Println(responsePayload, failedResponse)
}

func get(applicationConfig *example.ApplicationConfig) {
	contentId := 1
	listingsIds := []string{"1", "2", "3"}
	queryParams := make([]map[string]any, 0)

	queryParams = append(queryParams, map[string]any{"contentId": contentId})
	for v := range listingsIds {
		myMap := make(map[string]any)
		myMap["listingIds"] = v
		queryParams = append(queryParams, myMap)
	}

	httpMethod := genericclient.NewGenericHttpMethod[any, example.ResponsePayload](
		applicationConfig.LocalApi.Api.Host,
		applicationConfig.LocalApi.Default)

	responsePayload, failedResponse := httpMethod.Get(queryParams)
	fmt.Println(responsePayload, failedResponse)
}

func delete(applicationConfig *example.ApplicationConfig) {
	defaultEndPoint := applicationConfig.LocalApi.DeleteDefault
	defaultEndPoint.Url = fmt.Sprintf(defaultEndPoint.Url, "1111")
	httpMethod := genericclient.NewGenericHttpMethod[any, example.ResponsePayload](
		applicationConfig.LocalApi.Api.Host,
		defaultEndPoint)

	responsePayload, failedResponse := httpMethod.Delete()
	fmt.Println(responsePayload, failedResponse)
}

func getConfig() *example.ApplicationConfig {
	configInstance := example.CreateConfigInstance()
	applicationConfig, _ := configInstance.GetConfig()
	return applicationConfig
}

func CreateClients(applicationConfig *example.ApplicationConfig) {
	factory := genericclient.NewHttpClientFactory()

	localhostConfig := applicationConfig.LocalApi.Api
	localApiConfig := genericclient.NewClientConfigBuilder(localhostConfig.Host).
		SetTimeout(localhostConfig.Timeout).
		SetMaxIdleConnection(localhostConfig.MaxIdleConnection).
		Build()
	factory.AppendClient(*localApiConfig)
}
