package structs

import (
	"encoding/json"
	"github.com/yautah/goroyale/helper"
)

type GameMode struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
	Locale struct {
		TID map[string]string `json:"tid"`
	} `json:"_lang"`
}

func (u *GameMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		NameCn string `json:"nameCn"`
	}{
		ID:     u.ID,
		Name:   u.Name,
		NameCn: helper.GetGameModeByID(u.ID).Lang.Tid.Cn,
	})
}
