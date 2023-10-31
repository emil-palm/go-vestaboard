package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/motemen/go-loghttp"

	"github.com/mikehelmick/go-vestaboard/v2/client/api"
	"github.com/mikehelmick/go-vestaboard/v2/client/errors"
	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

func NewHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	if os.Getenv("VBHTTPDEBUG") == "1" {
		client.Transport = &loghttp.Transport{}
	}

	return client
}

const MaxBodySize = 2_000_000

var (
	NonOKStatusError = fmt.Errorf("Non-OK status code")
)

type Client struct {
	*http.Client
}

func NewClient() *Client {
	return NewWithHTTPClient(NewHTTPClient())
}

func NewWithHTTPClient(client *http.Client) *Client {
	return &Client{
		client,
	}
}

func (c *Client) Do(rt RequestType, req *http.Request) (interface{}, error) {
	httpResponse, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}

	resp, err := c.ParseResponse(rt, httpResponse)
	return resp, err
}

func (c *Client) SendMessage(ctx context.Context, board Board, layout layout.Layout) (*api.MessageResponse, error) {
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(layout); err != nil {
		return nil, fmt.Errorf("failed to encode JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "", &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Content-Type", "application/json")

	err = board.Apply(SendMessageRequest, req)
	if err != nil {
		return nil, err
	}

	messageresponse, err := c.Do(SendMessageRequest, req)

	return messageresponse.(*api.MessageResponse), err
}

func (c *Client) SendText(ctx context.Context, board Board, text string) (*api.MessageResponse, error) {
	text = strings.ToUpper(text)
	if err := layout.ValidText(text, true); err != nil {
		return nil, fmt.Errorf("invalid message: %w", err)
	}

	var b bytes.Buffer
	body := struct {
		Text string `json:"text"`
	}{Text: text}

	if err := json.NewEncoder(&b).Encode(body); err != nil {
		return nil, fmt.Errorf("failed to encode JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "", &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Content-Type", "application/json")

	err = board.Apply(SendTextRequest, req)
	if err != nil {
		return nil, err
	}

	messageresponse, err := c.Do(SendTextRequest, req)

	return messageresponse.(*api.MessageResponse), err
}

func (c *Client) ReadResponse(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewStatusError(resp.StatusCode, NonOKStatusError)
	}

	data := io.LimitReader(resp.Body, MaxBodySize)
	body, err := io.ReadAll(data)

	if err != nil {
		return nil, err
	}

	return body, nil

}

func (c *Client) ParseResponse(rt RequestType, resp *http.Response) (interface{}, error) {
	var parsedResponse interface{}
	switch rt {
	case SendMessageRequest:
	case SendTextRequest:
	case ReadMessageRequest:
		parsedResponse = api.MessageResponse{}
	}

	body, err := c.ReadResponse(resp)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		return &parsedResponse, err
	}

	return &parsedResponse, nil
}
