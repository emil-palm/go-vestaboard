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

	"github.com/mikehelmick/go-vestaboard/internal/config"
	client "github.com/mikehelmick/go-vestaboard/v2/clients/installables"
	"github.com/mikehelmick/go-vestaboard/v2/layout"
)

// Writes out a test-pattern of characters and colors.
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

	l := layout.NewLayout()
	l.Print(0, 0, " ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$()-+&=;:'\"%,./?°")
	l.SetColor(2, 14, layout.PoppyRed)
	l.SetColor(2, 15, layout.Orange)
	l.SetColor(2, 16, layout.Yellow)
	l.SetColor(2, 17, layout.Green)
	l.SetColor(2, 18, layout.ParisBlue)
	l.SetColor(2, 19, layout.Violet)
	l.SetColor(2, 20, layout.White)

	l.SetColor(3, 1, layout.PoppyRed)
	l.SetColor(3, 2, layout.Orange)
	l.SetColor(3, 3, layout.Yellow)
	l.SetColor(3, 4, layout.Green)
	l.SetColor(3, 5, layout.ParisBlue)
	l.SetColor(3, 6, layout.Violet)
	l.SetColor(3, 7, layout.White)
	l.Print(3, 9, " ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$()-+&=;:'\"%,./?°")

	msg, err := client.SendMessage(ctx, subs[0], l)
	if err != nil {
		log.Fatalf("error sending message: %v", err)
	}

	log.Printf("result: %+v", msg)
}
