package example

type RequestPayload struct {
	ClientId int `json:"clientId"`
}

func NewRequestPayload(clientId int) *RequestPayload {
	return &RequestPayload{clientId}
}
