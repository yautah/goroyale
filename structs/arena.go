package structs

import (
	"encoding/json"
	"github.com/yautah/goroyale/helper"
	"github.com/yautah/goroyale/locale"
)

type Arena struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	IconUrl string `json:"iconUrl"`
}

func (u *Arena) MarshalJSON() ([]byte, error) {
	c := helper.GetArenaById(u.Id)
	return json.Marshal(&struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		IconUrl string `json:"iconUrl"`
	}{
		ID:      c.Scid,
		Name:    locale.GetLocaleById(c.SubtitleTID),
		IconUrl: c.IconUrl,
	})
}
