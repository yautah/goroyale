package helper

import (
	CardModel "github.com/yautah/goroyale/structs/json"
)

var jsons = "assets/cards.json"

var CardsIdMap = make(map[int]CardModel.Card)
var CardsNameMap = make(map[string]CardModel.Card)
var TOTAL_CARD_GOLD int

func GetCardById(id int) CardModel.Card {
	return CardsIdMap[id]
}

func GetCardByName(name string) CardModel.Card {
	return CardsNameMap[name]
}
