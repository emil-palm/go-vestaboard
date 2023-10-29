// Copyright 2021 Mike Helmick
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	api "github.com/mikehelmick/go-vestaboard/v2/clients/api/v2"
	"github.com/mikehelmick/go-vestaboard/v2/clients/errors"
	"github.com/mikehelmick/go-vestaboard/v2/clients/localboard"
	"github.com/mikehelmick/go-vestaboard/v2/examples/internal/config"
)

var textFlag = flag.String("text", "HELLO, WORLD!", "text to send")

// Just send a quick 'Hello World' mesasge.
func main() {
	flag.Parse()

	ctx := context.Background()
	c, err := config.New(ctx)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	board := localboard.NewBoard("example", c.Secret)

	client := api.NewClient("http://vestaboard.local:8000/")
RESEND:
	resp, err := client.SendText(ctx, board, *textFlag)

	if err != nil {
		switch err.(type) {
		case errors.StatusError:
			switch err.(errors.StatusError).StatusCode {
			case http.StatusServiceUnavailable:
				time.Sleep(time.Second * 10)
				goto RESEND
			case http.StatusNotModified:
				log.Print("Board is already that")
				break

			default:
				log.Fatal(err)
			}

		default:
			log.Fatal(err)
		}

	} else {
		log.Printf("result: %+v", resp)
	}
}
