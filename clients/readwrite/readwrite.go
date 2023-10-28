package readwrite

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mikehelmick/go-vestaboard/v2/clients"
	"github.com/mikehelmick/go-vestaboard/v2/clients/errors"
	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

const MaxBodySize = 2_000_000

var (
	NonOKStatusError = fmt.Errorf("Non-OK status code")
)

type Client struct {
	baseURL string
	client  *http.Client
}

func New() *Client {
	return NewWithHTTPClient(clients.NewHTTPClient())
}

func NewWithHTTPClient(client *http.Client) *Client {
	return &Client{
		baseURL: "https://rw.vestaboard.com",
		client:  client,
	}
}

func (c *Client) do(req *http.Request) (*Response, error) {
	httpResponse, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	resp, err := c.parseResponse(httpResponse)
	return resp, err
}

func (c *Client) SendMessage(ctx context.Context, board clients.Board, layout layout.Layout) (*Response, error) {
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

func (c *Client) SendText(ctx context.Context, board clients.Board, text string) (*Response, error) {
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

func (c *Client) parseResponse(resp *http.Response) (*Response, error) {

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewStatusError(resp.StatusCode, NonOKStatusError)
	}

	var parsedResponse Response

	data := io.LimitReader(resp.Body, MaxBodySize)
	body, err := io.ReadAll(data)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		return &parsedResponse, err
	}

	return &parsedResponse, nil
}
