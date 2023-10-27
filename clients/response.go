package clients

import (
	"net/http"
)

type Response struct {
	HTTPResponseStatusCode int
}

func NewResponse(resp http.Response) *Response {
	r := Response{
		HTTPResponseStatusCode: resp.StatusCode,
	}

	return &r
}
