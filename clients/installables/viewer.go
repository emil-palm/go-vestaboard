package installables

import (
	"context"
	"net/http"
)

const viewerPath = "/viewer"

type ViewerResponse struct {
	Type         string `json:"type"`
	ID           string `json:"_id"`
	Created      string `json:"_created"`
	Installation `json:"installation"`
}

func (c *Client) Viewer(ctx context.Context) (*ViewerResponse, error) {
	url := c.baseURL + viewerPath
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set(APIKeyHeader, c.apiKey)
	req.Header.Set(APIKeySecret, c.apiSecret)

	var response ViewerResponse
	_, err = c.do(req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
