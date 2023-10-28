// Copyright 2021 Mike Helmick
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package installables

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mikehelmick/go-vestaboard/v2/clients"
	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

var (
	ErrMessageTruncated  = fmt.Errorf("message truncated")
	ErrInvalidCoordinate = fmt.Errorf("invalid coordinate")
)

type TextMessage struct {
	Text string `json:"text"`
}

type LayoutMessage struct {
	Layout layout.Layout `json:"characters"`
}

type Message struct {
	ID      string `json:"id"`
	Created int    `json:"created"`
	Text    string `json:"text,omitempty"`
}

type MessageResponse struct {
	Message `json:"message"`
}

func (c *Client) SendMessage(ctx context.Context, subscription clients.Board, l layout.Layout) (*MessageResponse, error) {
	var b bytes.Buffer
	body := &LayoutMessage{
		Layout: l,
	}
	if err := json.NewEncoder(&b).Encode(body); err != nil {
		return nil, fmt.Errorf("failed to encode JSON: %w", err)
	}

	// We create a URL without the subscriptionid
	// pass that request to Subscription.Apply() which will alter that path
	// Then we append the action we want /message

	url := fmt.Sprintf("%s%s", c.baseURL, subscriptionsPath)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &b)
	if err != nil {
		return nil, err
	}

	req.URL.Path = fmt.Sprintf("%s/message", req.URL.Path)

	var response MessageResponse
	_, err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SendText(ctx context.Context, subscription clients.Board, text string) (*MessageResponse, error) {
	text = strings.ToUpper(text)
	if err := layout.ValidText(text, true); err != nil {
		return nil, fmt.Errorf("invalid message: %w", err)
	}

	var b bytes.Buffer
	body := &TextMessage{
		Text: text,
	}
	if err := json.NewEncoder(&b).Encode(body); err != nil {
		return nil, fmt.Errorf("failed to encode JSON: %w", err)
	}

	// We create a URL without the subscriptionid
	// pass that request to Subscription.Apply() which will alter that path
	// Then we append the action we want /message

	url := fmt.Sprintf("%s%s", c.baseURL, subscriptionsPath)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &b)
	if err != nil {
		return nil, err
	}
	err = subscription.Apply(req)
	if err != nil {
		return nil, err
	}

	req.URL.Path = fmt.Sprintf("%s/message", req.URL.Path)

	var response MessageResponse
	_, err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
