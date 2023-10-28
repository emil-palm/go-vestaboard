package readwrite

import (
	"encoding/json"
	"time"
)

type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0)
	return nil
}

type Response struct {
	Status  string   `json:"status"`
	ID      string   `json:"id"`
	Created UnixTime `json:"created"`
}
