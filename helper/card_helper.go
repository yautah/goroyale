package helper

import (
	"encoding/json"
	"fmt"
	. "github.com/yautah/goroyale/assets"
	CardModel "github.com/yautah/goroyale/structs/json"
)

var jsons = []string{
	"assets/spells_buildings.json",
	"assets/spells_characters.json",
	"assets/spells_other.json",
}

var CardsIdMap = make(map[int]CardModel.Card)
var CardsNameMap = make(map[string]CardModel.Card)

func init() {
	var cards []CardModel.CardBuilding
	data, err := Asset(jsons[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, &cards)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range cards {
		card := c.ConvertToCard()
		if card != nil {
			CardsIdMap[card.ID] = *card
			CardsNameMap[card.NameEn] = *card
		}
	}

	var cards1 []CardModel.CardTroop
	data, err = Asset(jsons[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, &cards1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range cards1 {
		card := c.ConvertToCard()
		if card != nil {
			CardsIdMap[card.ID] = *card
			CardsNameMap[card.NameEn] = *card
		}
	}
	var cards2 []CardModel.CardOther
	data, err = Asset(jsons[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, &cards2)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range cards2 {
		card := c.ConvertToCard()
		if card != nil {
			CardsIdMap[card.ID] = *card
			CardsNameMap[card.NameEn] = *card
		}
	}
}

func GetCardById(id int) CardModel.Card {
	return CardsIdMap[id]
}

func GetCardByName(name string) CardModel.Card {
	return CardsNameMap[name]
}
