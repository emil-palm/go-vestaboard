package readwrite

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	Status string `json:"status"`
	ID string `json:"id"`
	Created UnixTime
}

type Client struct {
	baseURL string
}

func NewClient() *Client {
	return &Client{
		baseURL: "https://rw.vestaboard.com",
	}
}

func (c *Client) SendMessage(ctx context.Context, board clients.Board, layout layout.Layout) (*http.Request, error) {
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

	return req, nil
}

func (c *Client) SendText(ctx context.Context, board clients.Board, text string) (*http.Request, error) {
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

	return req, nil
}

func (c *Client) ParseResponse(httpResp *http.Response) (*clients.Response, error) {
	resp, err := clients.NewResponse(*httpResp)

	if err != nil {
		return nil, err
	}

	var parsedResponse messageResponse
	

	if err := json.Unmarshal(resp.Data, &parsedResponse); err != nil {
		return nil, err
	}

	if resp.HTTPResponseCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.HTTPResponseCode)
	}

	return resp, nil
}
