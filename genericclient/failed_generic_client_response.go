package genericclient

import "net/http"

type FailedGenericClientResponse struct {
	HttpResponse *http.Response
	Err          error
}

func NewFailedGenericClientResponse(response *http.Response, err error) *FailedGenericClientResponse {
	return &FailedGenericClientResponse{
		response, err,
	}
}
