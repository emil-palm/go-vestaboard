package clients

import (
	"context"
	"net/http"
	"os"
	"time"
	"log"

	"github.com/motemen/go-loghttp"

	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

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
	HTTPClient *http.Client
	impl       ClientImplementation
}

func NewClient(impl ClientImplementation) (*Client, error) {
	client := Client{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		impl: impl,
	}

	if os.Getenv("VBHTTPDEBUG") == "1" {
		client.HTTPClient.Transport = &loghttp.Transport{}
	}

	return &client, nil
}

func (c *Client) do(req *http.Request) (*Response, error) {

	httpResponse, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("%+v", httpResponse)
	resp, err := c.impl.ParseResponse(httpResponse)

	return resp, err
}

func (c *Client) SendMessage(ctx context.Context, b Board, m layout.Layout) (*Response, error) {
	req, err := c.impl.SendMessage(ctx, b, m)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *Client) SendText(ctx context.Context, b Board, t string) (*Response, error) {
	req, err := c.impl.SendText(ctx, b, t)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}
