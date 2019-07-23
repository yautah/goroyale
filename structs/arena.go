package structs

import (
	"encoding/json"
	"github.com/yautah/goroyale/helper"
)

type Arena struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IconUrl string `json:"iconUrl"`
}

func (u *Arena) MarshalJSON() ([]byte, error) {
	c := helper.GetArenaByName(u.Name)
	return json.Marshal(&struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		IconUrl string `json:"iconUrl"`
	}{
		ID:      c.ID,
		Name:    c.Title,
		IconUrl: c.IconUrl,
	})
}
