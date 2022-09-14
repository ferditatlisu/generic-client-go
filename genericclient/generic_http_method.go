package genericclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GenericHttMethod[TRequest interface{}, TResponse interface{}] interface {
	Get(requestMap map[string]any) (TRequest, *FailedGenericClientResponse)
}

type genericHttpMethod[TRequest interface{}, TResponse interface{}] struct {
	genericClient  *GenericClient
	endPointConfig EndPoint
	header         *http.Header
}

func NewGenericHttpMethod[TRequest interface{}, TResponse interface{}](
	host string, epc EndPoint) *genericHttpMethod[TRequest, TResponse] {
	factory := getClientFactory()
	genericClient := factory.GetClient(host)

	return &genericHttpMethod[TRequest, TResponse]{genericClient: genericClient, endPointConfig: epc}
}

func (i *genericHttpMethod[TRequest, TResponse]) Get(requestMap []map[string]any) (TResponse, *FailedGenericClientResponse) {
	queryParameters := addQueryParams(requestMap)
	uri := getUri(i.genericClient.host, i.endPointConfig.Url, &queryParameters)
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	i.setHeader(req)
	httpResponse, err := i.genericClient.client.Do(req)
	successResponse, failedResponse := prepareResponse[TResponse](httpResponse, err)

	return successResponse, failedResponse
}

func (i *genericHttpMethod[TRequest, TResponse]) Post(payload TRequest) (TResponse, *FailedGenericClientResponse) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return Default[TResponse](), NewFailedGenericClientResponse(nil, err)
	}

	uri := getUri(i.genericClient.host, i.endPointConfig.Url, nil)
	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(jsonBody))
	i.setHeader(req)
	httpResponse, err := i.genericClient.client.Do(req)
	successResponse, failedResponse := prepareResponse[TResponse](httpResponse, err)

	return successResponse, failedResponse
}

func (i *genericHttpMethod[TRequest, TResponse]) Put(payload TRequest) (TResponse, *FailedGenericClientResponse) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return Default[TResponse](), NewFailedGenericClientResponse(nil, err)
	}

	uri := getUri(i.genericClient.host, i.endPointConfig.Url, nil)
	req, _ := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(jsonBody))
	i.setHeader(req)
	httpResponse, err := i.genericClient.client.Do(req)
	successResponse, failedResponse := prepareResponse[TResponse](httpResponse, err)

	return successResponse, failedResponse
}

func (i *genericHttpMethod[TRequest, TResponse]) Delete() (TResponse, *FailedGenericClientResponse) {
	uri := getUri(i.genericClient.host, i.endPointConfig.Url, nil)
	req, _ := http.NewRequest(http.MethodDelete, uri, nil)
	i.setHeader(req)
	httpResponse, err := i.genericClient.client.Do(req)
	successResponse, failedResponse := prepareResponse[TResponse](httpResponse, err)

	return successResponse, failedResponse
}

func (i *genericHttpMethod[TRequest, TResponse]) setHeader(req *http.Request) {
	if i.header != nil {
		req.Header = *i.header
	}
}

func (i *genericHttpMethod[TRequest, TResponse]) AddHeaders(h map[string]string) *genericHttpMethod[TRequest, TResponse] {
	header := http.Header{}
	for k, v := range h {
		header.Set(k, v)
	}

	i.header = &header

	return i
}

func prepareResponse[TResponse interface{}](httpResponse *http.Response, err error) (TResponse, *FailedGenericClientResponse) {
	if err != nil {
		return Default[TResponse](), NewFailedGenericClientResponse(httpResponse, err)
	}

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return Default[TResponse](), NewFailedGenericClientResponse(httpResponse, err)
	}

	if httpResponse.StatusCode == http.StatusNoContent {
		return Default[TResponse](), NewFailedGenericClientResponse(httpResponse, nil)
	}

	if is2xx(httpResponse) {
		var response TResponse
		err = json.Unmarshal(body, &response)
		var failedResponse *FailedGenericClientResponse
		if err != nil {
			failedResponse = NewFailedGenericClientResponse(httpResponse, err)
		}
		return response, failedResponse
	}

	return Default[TResponse](), NewFailedGenericClientResponse(httpResponse, nil)
}

func is2xx(httpResponse *http.Response) bool {
	return httpResponse.StatusCode == http.StatusOK ||
		httpResponse.StatusCode == http.StatusCreated ||
		httpResponse.StatusCode == http.StatusAccepted
}

func getUri(host string, endPointUrl string, queryParameters *string) string {
	qp := EmptyString
	if queryParameters != nil {
		qp = *queryParameters
	}

	return host + SLASH + endPointUrl + qp
}

func addQueryParams(queryMap []map[string]any) string {
	var b bytes.Buffer
	if queryMap == nil || len(queryMap) == 0 {
		return ""
	}

	b.WriteString(QuestionMark)
	for _, m := range queryMap {
		if m == nil {
			continue
		}

		for k, v := range m {
			if k == "" || v == nil {
				continue
			}
			b.WriteString(k)
			b.WriteString(EQUAL)
			b.WriteString(fmt.Sprint(v))
			b.WriteString(AND)
		}
	}

	queryParameters := b.String()

	lastIndex := len(queryParameters) - 1
	queryParameters = queryParameters[:lastIndex]

	return queryParameters
}
