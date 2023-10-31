package client

import (
	"context"
	"log"
	"net/http"
	"net/url"

	"github.com/mikehelmick/go-vestaboard/v2/client/api"
	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

type RequestType int

const (
	SendMessageRequest RequestType = iota
	SendTextRequest
	ReadMessageRequest
)

func (rt RequestType) String() string {
	return [...]string{"SendMessage", "SendTextRequest", "ReadMessageRequest"}[rt]
}

func ErrorNotImplemented(rt RequestType) {
	log.Fatalf("%s is not implemented in this board type", rt)
}

type Board interface {
	Apply(RequestType, *http.Request) error
}

type BaseBoard struct {
	token   string
	baseurl string
	header  string
}

func NewBoard(baseurl, token, header string) *BaseBoard {
	return &BaseBoard{
		token:   token,
		baseurl: baseurl,
		header:  header,
	}
}

func (bb *BaseBoard) Apply(rt RequestType, req *http.Request) error {
	if req.URL.Host == "" {
		var err error
		req.URL, err = url.Parse(bb.baseurl)
		if err != nil {
			return err
		}
	}
	req.Header.Set(bb.header, bb.token)

	return nil
}

func (bb *BaseBoard) SendText(ctx context.Context, text string) (*api.MessageResponse, error) {
	return NewClient().SendText(ctx, bb, text)
}

func (bb *BaseBoard) SendMessage(ctx context.Context, layout layout.Layout) (*api.MessageResponse, error) {
	return NewClient().SendMessage(ctx, bb, layout)
}
