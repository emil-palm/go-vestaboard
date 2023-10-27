package readwrite

import (
	"fmt"
	"net/http"
)

const (
	rwapikeyheader = "X-Vestaboard-Read-Write-Key"
)

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
	req.Header.Set(rwapikeyheader, rw.token)
	return nil
}

func (rw *ReadWriteBoard) String() string {
	return fmt.Sprintf("[ReadWrite] %s", rw.name)
}
