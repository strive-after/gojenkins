package api

import (
	"io"
	"net/http"
)

type APIRequest struct {
	Method   string
	Endpoint string
	Payload  io.Reader
	Headers  http.Header
	Suffix   string
}


func NewAPIRequest(method string, endpoint string, payload io.Reader) *APIRequest {
	var headers = http.Header{}
	var suffix string
	ar := &APIRequest{method, endpoint, payload, headers, suffix}
	return ar
}

func (ar *APIRequest) SetHeader(key string, value string) *APIRequest {
	ar.Headers.Set(key, value)
	return ar
}
