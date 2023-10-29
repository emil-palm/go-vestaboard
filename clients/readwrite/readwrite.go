package readwrite

import v2 "github.com/mikehelmick/go-vestaboard/v2/clients/api/v2"

const (
	rwApiKey = "X-Vestaboard-Read-Write-Key"
)

func NewBoard(name, token string) *v2.Board {
	return v2.NewBoard(name, token, rwApiKey)
}
