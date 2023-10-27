package clients

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/motemen/go-loghttp"

	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

type Response struct {
}

type Board interface {
	Apply(*http.Request) error
	String() string
}

type ClientImplementation interface {
	SendMessage(context.Context, Board, layout.Layout) (*http.Request, error)
	SendText(context.Context, Board, string) (*http.Request, error)

	ParseResponse(*http.Response) (*Response, error)
}

type Client struct {
	httpClient *http.Client
	impl       ClientImplementation
}

func NewClient(impl ClientImplementation) (*Client, error) {
	client := Client{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		impl: impl,
	}

	if os.Getenv("VBHTTPDEBUG") == "1" {
		client.httpClient.Transport = &loghttp.Transport{}
	}

	return &client, nil
}

func (c *Client) SendMessage(ctx context.Context, b Board, m layout.Layout) (*Response, error) {
	req, err := c.impl.SendMessage(ctx, b, m)
	if err != nil {
		return nil, err
	}

	httpResponse, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.impl.ParseResponse(httpResponse)

	return resp, nil
}

func (c *Client) SendText(ctx context.Context, b Board, t string) (Response, error) {
	return Response{}, nil
}
