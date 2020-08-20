package storage

import (
	"encoding/json"
	"time"
)

type storage struct {
	events []event
}

type event struct {
	Time    time.Time `json:"Time"`
	Action  action    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Output  string    `json:"Output,omitempty"`
	Elapsed float32   `json:"Elapsed,omitempty"`
}

type action string

const (
	run    = action("run")
	output = action("output")
	pass   = action("pass")
)

func (s *storage) Put(b []byte) error {
	var e event
	if err := json.Unmarshal(b, &e); err != nil {
		return err
	}
	s.events = append(s.events, e)
	return nil
}
