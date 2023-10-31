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
// limitations under the License.package installables

package subscription

import (
	"fmt"
	"net/http"

	"github.com/mikehelmick/go-vestaboard/v2/client"
)

type Board struct {
	ID string `json:"_id"`
}

type Installable struct {
	ID string `json:"_id"`
}

type Installation struct {
	ID          string `json:"_id"`
	Installable `json:"installable"`
}

type Subscription struct {
	ID           string `json:"_id"`
	Created      string `json:"_created"`
	Installation `json:"installation"`
	Boards       []Board `json:"boards"`
}

func (s *Subscription) Apply(rt client.RequestType, req *http.Request) error {
	req.URL.Path = fmt.Sprintf("%s/%s", req.URL.Path, s.ID)
	return nil
}
