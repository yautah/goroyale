package helper

import (
	CardModel "github.com/yautah/goroyale/structs/json"
)

func GetCardById(id int) CardModel.Card {
	return CardsIdMap[id]
}

func GetCardByName(name string) CardModel.Card {
	return CardsNameMap[name]
}
