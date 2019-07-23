package structs

import (
	"encoding/json"
	"fmt"
	"github.com/yautah/goroyale/helper"
)

const Endpoint = "http://www.madn.xyz/assets/images"

type Card struct {
	ID           int    `json:"id"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Elixir       int    `json:"elixir"`
	Level        int    `json:"level"`
	MaxLevel     int    `json:"maxLevel"`
	OLevel       int    `json:"oLevel"`
	OMaxLevel    int    `json:"oMaxLevel"`
	StarLevel    int    `json:"starLevel"`
	Count        int    `json:"count"`
	UpgradeCount int    `json:"upgradeCount"`
	MaxCount     int    `json:"maxCount"`
	MaxGold      int    `json:"maxGold"`
	CanUpgradeTo int    `json:"canUpgradeTo"`
	Rarity       string `json:"rarity"`
	IconUrl      string `json:"iconUrl"`
}

func (u *Card) UnmarshalJSON(data []byte) error {
	// type Alias Card
	// aux := &struct {
	// *Alias
	// }{
	// Alias: (*Alias)(u),
	// }

	ori := &struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Level     int    `json:"level"`
		MaxLevel  int    `json:"maxLevel"`
		StarLevel int    `json:"starLevel"`
		Count     int    `json:"count"`
	}{}

	if err := json.Unmarshal(data, &ori); err != nil {
		return err
	}

	c := helper.GetCardById(ori.ID)

	u.ID = c.ID
	u.Name = c.Name
	u.Elixir = c.Elixir
	u.Level = calcLevel(c.Rarity, ori.Level)
	u.MaxLevel = 13
	u.StarLevel = ori.StarLevel
	u.Count = ori.Count

	u.Rarity = c.Rarity
	u.OLevel = ori.Level
	u.OMaxLevel = ori.MaxLevel
	u.IconUrl = fmt.Sprintf("%s/cards/%s.png", Endpoint, c.Key)
	u.Key = c.Key
	// u.MaxLevel = 13

	u.UpgradeCount = helper.GetCardUpgradeCount(u.Rarity, ori.Level)
	u.MaxCount = helper.GetCardMaxCount(u.Rarity, ori.Level, ori.Count)
	u.MaxGold = helper.GetCardMaxGold(u.Rarity, ori.Level)
	u.CanUpgradeTo = helper.GetCardCanUpgradeTo(u.Rarity, ori.Level, ori.Count)
	return nil
}

func calcLevel(rarity string, level int) int {
	switch rarity {
	case "Common":
		return level
	case "Rare":
		return level + 2
	case "Epic":
		return level + 5
	case "Legendary":
		return level + 8
	}
	return level
}

// func (card *Card) UnmarshalJSON(bytes []byte) error {
// err := json.Unmarshal(bytes, card)
// if err != nil {
// return err
// }

// c := helper.GetCardByName(card.Name)

// card.ID = c.ID
// card.Elixir = c.Elixir

// return nil
// }
