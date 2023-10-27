package clients

import (
	"net/http"
	"os"
	"time"

	"github.com/motemen/go-loghttp"
)

type Board interface {
	Apply(*http.Request) error
	String() string
}

func NewHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	if os.Getenv("VBHTTPDEBUG") == "1" {
		client.Transport = &loghttp.Transport{}
	}

	return client
}
