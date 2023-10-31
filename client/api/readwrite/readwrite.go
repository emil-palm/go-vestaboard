package readwrite

import (
	"net/http"

	"github.com/mikehelmick/go-vestaboard/v2/client"
)

const (
	rwApiKey = "X-Vestaboard-Read-Write-Key"
)

type Board struct {
	*client.BaseBoard
}

func NewBoard(token string) *Board {
	return &Board{
		client.NewBoard("https://rw.vestaboard.com/", token, rwApiKey),
	}
}

func (b *Board) Apply(rt client.RequestType, req *http.Request) error {
	b.BaseBoard.Apply(rt, req)

	switch rt {
	case client.SendMessageRequest:
	case client.SendTextRequest:
		req.Method = http.MethodPost
		req.URL.Path = "/message"
		break
	case client.ReadMessageRequest:
		req.Method = http.MethodGet
		req.URL.Path = "/message"
		break

	default:
		client.ErrorNotImplemented(rt)
		return nil
	}

	return nil
}
