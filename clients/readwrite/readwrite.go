package readwrite

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mikehelmick/go-vestaboard/v2/clients"
	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

type ReadWriteBoard struct {
	token string
	name  string
}

func NewReadWriteBoard(name, token string) *ReadWriteBoard {
	return &ReadWriteBoard{
		token: token,
		name:  name,
	}
}

func (rw *ReadWriteBoard) Apply(req *http.Request) error {
	return nil
}

func (rw *ReadWriteBoard) String() string {
	return fmt.Sprintf("[ReadWrite] %s", rw.name)
}

func (c *Client) SendMessage(ctx context.Context, board clients.Board, layout layout.Layout) (*http.Request, error) {
	return nil, nil
}

func (c *Client) SendText(ctx context.Context, board clients.Board, text string) (*http.Request, error) {
	return nil, nil
}
func (c *Client) ParseResponse(*http.Response) (*clients.Response, error) {
	return nil, nil
}
