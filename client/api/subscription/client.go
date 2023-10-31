package subscription

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mikehelmick/go-vestaboard/v2/client"
)

const (
	keyHeader    = "X-Vestaboard-Api-Key"
	secretHeader = "X-Vestaboard-Api-Secret"
)

type Client struct {
	*client.Client

	baseurl   string
	apiKey    string
	apiSecret string
}

func NewClient(key, secret string) *Client {
	return NewWithHTTPClient(client.NewHTTPClient(), key, secret)
}

func NewWithHTTPClient(c *http.Client, key, secret string) *Client {
	return &Client{
		baseurl:   "https://platform.vestaboard.com",
		apiKey:    key,
		apiSecret: secret,
		Client:    client.NewWithHTTPClient(c),
	}
}

func (c *Client) Do(rt RequestType, req *http.Request) (interface{}, error) {
	req.Header.Set(keyHeader, c.apiKey)
	req.Header.Set(secretHeader, c.apiSecret)

	httpResponse, err := c.Client.Client.Do(req)

	if err != nil {
		return nil, err
	}

	resp, err := c.ParseResponse(rt, httpResponse)
	return resp, err
}

func (c *Client) ParseResponse(rt RequestType, resp *http.Response) (interface{}, error) {
	log.Print(rt == SubscriptionsRequest)
	switch rt {
	case SubscriptionsRequest, ViewerRequest:
		log.Print("BLAH")
		body, err := c.ReadResponse(resp)
		if err != nil {
			return nil, err
		}

		if rt == SubscriptionsRequest {
			var response struct{ Subscriptions []*Subscription }

			if err := json.Unmarshal(body, &response); err != nil {
				return nil, err
			}

			return response.Subscriptions, nil

		} else if rt == ViewerRequest {
			var response Viewer
			if err := json.Unmarshal(body, &response); err != nil {
				return nil, err
			}
		}

		return nil, fmt.Errorf("Missing response parser implementation")

	default:
		return c.Client.ParseResponse(client.RequestType(rt), resp)
	}

}

const subscriptionsPath = "/subscriptions"

func (c *Client) Subscriptions(ctx context.Context) ([]*Subscription, error) {
	ctx.
		url := c.baseurl + subscriptionsPath

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	response, err := c.Do(SubscriptionsRequest, req)

	if err != nil {
		return nil, err
	}

	log.Printf("%T", response)

	return response.([]*Subscription), nil
}

const viewerPath = "/viewer"

func (c *Client) Viewer(ctx context.Context) (*Viewer, error) {
	url := c.baseurl + viewerPath
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.Do(ViewerRequest, req)
	if err != nil {
		return nil, err
	}

	return response.(*Viewer), nil
}

type RequestType client.RequestType

const (
	ViewerRequest RequestType = iota + 3
	SubscriptionsRequest
)
