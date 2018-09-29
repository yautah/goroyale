package main

import (
	"fmt"
	// "github.com/yautah/goroyale/client"
	"github.com/yautah/goroyale/helper"
	// "net/url"
)

func main() {
	// arena := helper.GetArenaById(54000013)
	// fmt.Println(arena)
	// fmt.Println(helper.GetTextById(arena.ID))

	// card := helper.GetCardById(27000001)
	// fmt.Println(card)

	// card = helper.GetCardById(26000062)
	// fmt.Println(card)

	// cc := helper.GetCardByName(card.NameEn)
	// fmt.Println(cc)

	// r := helper.GetRarityByName("Common")
	// fmt.Println(r)

	// fmt.Println(helper.GetCardUpgradeCount("Rare", 10))
	// fmt.Println(helper.GetCardUpgradeGold("Rare", 10))
	// fmt.Println(helper.GetCardMaxCount("Rare", 9, 1100))
	// fmt.Println(helper.GetCardMaxGold("Rare", 9))
	// fmt.Println(helper.GetCardLevelGold("Legendary", 5))

	var gold = 0
	for _, card := range helper.CardsIdMap {
		switch card.Rarity {
		case "Common":
			gold = gold + helper.GetCardLevelGold(card.Rarity, 13)
		case "Rare":
			gold = gold + helper.GetCardLevelGold(card.Rarity, 11)
		case "Epic":
			gold = gold + helper.GetCardLevelGold(card.Rarity, 8)
		case "Legendary":
			gold = gold + helper.GetCardLevelGold(card.Rarity, 5)
		}
	}
	fmt.Println(gold)

	// fmt.Println(helper.GetCardCanUpgradeTo("Rare", 10, 1000))
	// fmt.Println(helper.GetCardCanUpgradeTo("Common", 10, 50))
	// fmt.Println(helper.GetCardCanUpgradeTo("Common", 9, 50))
	// fmt.Println(helper.GetCardCanUpgradeTo("Common", 9, 500))

}
