package structs

import (
	"encoding/json"
	"github.com/yautah/goroyale/helper"
)

type ClanBase struct {
	Tag     string `json:"tag"`
	Name    string `json:"name"`
	BadgeId int    `json:"badgeId"`
}

func (u *ClanBase) MarshalJSON() ([]byte, error) {
	badgeUrl := helper.GetBadgeById(u.BadgeId)
	return json.Marshal(&struct {
		Tag      string `json:"tag"`
		Name     string `json:"name"`
		BadgeId  int    `json:"badgeId"`
		BadgeUrl string `json:"badgeUrl"`
	}{
		Tag:      u.Tag,
		Name:     u.Name,
		BadgeId:  u.BadgeId,
		BadgeUrl: badgeUrl,
	})
}
