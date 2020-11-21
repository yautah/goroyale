package helper

import (
	"encoding/json"
	"fmt"
	. "github.com/yautah/goroyale/assets"
	CardModel "github.com/yautah/goroyale/structs/json"
	data "github.com/yautah/goroyale/structs/json"
)

var jsons = "assets/cards.json"

var CardsIdMap = make(map[int]CardModel.Card)
var CardsNameMap = make(map[string]CardModel.Card)
var TOTAL_CARD_GOLD int

func init() {
	var rarities []data.Rarity

	data, err := Asset("assets/rarities.json")
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &rarities)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, rarity := range rarities {
		raritiesMap[rarity.Name] = rarity
	}
}

func init() {
	for i := MinBadgeId; i <= MaxBadgeId; i++ {
		BadgesMap[i] = fmt.Sprintf("%s/badges/%d.png", Endpoint, i)
	}
}

func init() {
	var arenas []data.Arena

	data, err := Asset("assets/arenas.json")
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &arenas)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, arena := range arenas {
		arena.IconUrl = fmt.Sprintf("%s/arenas/arena%d.png", Endpoint, arena.Arena)
		ArenasNameMap[arena.ArenaName] = arena
		ArenasIdMap[arena.ID] = arena
		ArenasNumberMap[arena.Arena] = arena
	}
}

func init() {
	var gameModes []data.GameMode

	data, err := Asset("assets/game_modes.json")
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &gameModes)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, gameMode := range gameModes {
		GameModeMap[gameMode.ID] = gameMode
	}
}

func init() {
	var cards []CardModel.Card
	data, err := Asset(jsons)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, &cards)
	if err != nil {
		fmt.Println(err)
		return
	}

	gold := 0
	for _, c := range cards {
		CardsIdMap[c.ID] = c
		CardsNameMap[c.Key] = c
		gold = gold + GetCardMaxGold(c.Rarity, 1)
	}
	TOTAL_CARD_GOLD = gold

}
