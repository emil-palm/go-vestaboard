package readwrite

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mikehelmick/go-vestaboard/v2/clients"
	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0)
	return nil
}

type messageResponse struct {
	Status  string `json:"status"`
	ID      string `json:"id"`
	Created UnixTime
}

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient() *Client {
	return NewClientWithHTTPClient(clients.NewHTTPClient())
}

func NewClientWithHTTPClient(client *http.Client) *Client {
	return &Client{
		baseURL: "https://rw.vestaboard.com",
		client:  client,
	}
}

func (c *Client) do(req *http.Request) (*clients.Response, error) {
	httpResponse, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	resp, err := c.parseResponse(httpResponse)
	return resp, err
}

func (c *Client) SendMessage(ctx context.Context, board clients.Board, layout layout.Layout) (*clients.Response, error) {
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(layout); err != nil {
		return nil, fmt.Errorf("failed to encode JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Content-Type", "application/json")

	err = board.Apply(req)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *Client) SendText(ctx context.Context, board clients.Board, text string) (*clients.Response, error) {
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Content-Type", "application/json")

	err = board.Apply(req)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *Client) parseResponse(httpResp *http.Response) (*clients.Response, error) {
	resp, err := clients.NewResponse(*httpResp)
	if err != nil {
		return resp, err
	}

	var parsedResponse messageResponse

	if err := json.Unmarshal(resp.Data, &parsedResponse); err != nil {
		return resp, err
	}

	resp.ResponseMessage = parsedResponse

	if resp.HTTPResponseCode != http.StatusOK {
		return resp, fmt.Errorf("unexpected status code: %d", resp.HTTPResponseCode)
	}

	return resp, nil
}
