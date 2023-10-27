package clients

import (
	"net/http"

	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

type Response struct {
}

type Board interface {
	Apply(*http.Request) error
	String() string
}

type ClientImplementation interface {
	SendMessage(Board, layout.Layout) (Response, error)
	SendText(Board, string) (Response, error)
}
