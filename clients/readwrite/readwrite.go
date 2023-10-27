package readwrite

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mikehelmick/go-vestaboard/v2/clients"
	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

const MaxBodySize = 2_000_000

type messageResponse struct {
	Message string `json:"status"`
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
	return nil, nil
}
func (c *Client) ParseResponse(resp *http.Response) (*clients.Response, error) {

	var response messageResponse

	r := io.LimitReader(resp.Body, MaxBodySize)
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	/*ct := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "text/plain") && !strings.HasPrefix(ct, "application/json") {
		return nil, fmt.Errorf("%s: response content-type is not text/plain or application/json (got %s): body: %s",
			errPrefix, ct, body)
	}
	*/
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	log.Print(response)

	return nil, nil
}
