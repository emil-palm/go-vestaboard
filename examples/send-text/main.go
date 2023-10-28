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

	"github.com/mikehelmick/go-vestaboard/v2/clients/errors"
	client "github.com/mikehelmick/go-vestaboard/v2/clients/installables"
	"github.com/mikehelmick/go-vestaboard/v2/internal/config"
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

	client := client.New(c.APIKey, c.Secret)
	subs, err := client.Subscriptions(ctx)

	if err != nil {
		log.Fatalf("error calling Viewer: %v", err)
	}
	log.Printf("result: %+v", subs)

	msg, err := client.SendText(ctx, subs[0], *textFlag)
	if err != nil {
		switch err.(type) {
		case errors.NoChangeError:
			break
		default:
			log.Fatalf("error sending message: %v", err)
		}
	} else {
		log.Print(msg)
	}
}
