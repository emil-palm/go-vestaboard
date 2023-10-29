package v2

import (
	"fmt"
	"net/http"
)

type Board struct {
	token  string
	name   string
	header string
}

func NewBoard(name, token, header string) *Board {
	return &Board{
		token:  token,
		name:   name,
		header: header,
	}
}

func (b *Board) Apply(req *http.Request) error {
	req.Header.Set(b.header, b.token)
	return nil
}

func (b *Board) String() string {
	return fmt.Sprintf("[ReadWrite] %s", b.name)
}
