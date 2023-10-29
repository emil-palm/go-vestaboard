package localboard

import (
	"fmt"
	"net/http"

	v2 "github.com/mikehelmick/go-vestaboard/v2/clients/api/v2"
)

const (
	localApiKey = "X-Vestaboard-Local-Api-Key"
)

type LocalBoard struct {
	*v2.Board
}

func NewBoard(name, token string) *LocalBoard {
	return &LocalBoard{
		v2.NewBoard(name, token, localApiKey),
	}
}

func (lb *LocalBoard) Apply(req *http.Request) error {
	err := lb.Board.Apply(req)
	if err != nil {
		return err
	}

	req.URL.Path = fmt.Sprintf("%s/local-api/message", req.URL.Path)

	return nil
}
