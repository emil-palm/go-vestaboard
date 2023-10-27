package readwrite

import (
	"fmt"
	"net/http"
)

const (
	rwapikeyheader = "X-Vestaboard-Read-Write-Key"
)

type Board struct {
	token string
	name  string
}

func NewBoard(name, token string) *Board {
	return &Board{
		token: token,
		name:  name,
	}
}

func (rw *Board) Apply(req *http.Request) error {
	req.Header.Set(rwapikeyheader, rw.token)
	return nil
}

func (rw *Board) String() string {
	return fmt.Sprintf("[ReadWrite] %s", rw.name)
}
