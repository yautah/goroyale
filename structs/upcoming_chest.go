package structs

import (
	"encoding/json"
	"strings"
	// "github.com/yautah/goroyale/helper"
)

type UpcomingChest struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
}

func (u *UpcomingChest) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name  string `json:"name"`
		Index int    `json:"index"`
	}{
		Index: u.Index,
		Name:  strings.ToLower(strings.Replace(u.Name, " ", "-", -1)),
	})
}
