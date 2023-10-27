package clients

import (
	"io"
	"net/http"
)

const MaxBodySize = 2_000_000

type Response struct {
	HTTPResponseCode int
	ResponseMessage  interface{}
	Data             []byte
}

func NewResponse(resp http.Response) (*Response, error) {
	r := Response{
		HTTPResponseCode: resp.StatusCode,
	}

	data := io.LimitReader(resp.Body, MaxBodySize)
	body, err := io.ReadAll(data)

	if err != nil {
		return nil, err
	}

	r.Data = body

	return &r, nil
}
