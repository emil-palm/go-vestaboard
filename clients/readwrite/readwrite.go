package readwrite

import (
	"fmt"
	"net/http"

	"github.com/mikehelmick/go-vestaboard/v2/clients"
	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

type ReadWriteClient struct {
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

func (rw *ReadWriteClient) SendMessage(board clients.Board, layout layout.Layout) {

}

func (rw *ReadWriteClient) SendText(board clients.Board, text string) {

}
