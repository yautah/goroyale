package helper

import (
	chest "github.com/yautah/goroyale/structs/json"
)

var chestOrder chest.ChestCycle = chest.ChestCycle{}
var ChestNameMap = map[string]string{
	"Silver":          "Silver Chest",
	"Gold":            "Golden Chest",
	"Magic":           "Magical Chest",
	"Giant":           "Giant Chest",
	"GoldCage_Small":  "Gold Crate",
	"GoldCage_Medium": "Plentiful Gold Crate",
	"GoldCage_Large":  "Overflowing Gold Crate",
	"WildChampion":    "Royal Wild Chest",
}

//GetCGetChestOrder  get chest order

func GetChestCycle(level int) []string {
	var cycle []string
	if level == 14 {
		cycle = chestOrder.MainCycleKL14
	} else if level == 13 {
		cycle = chestOrder.MainCycleKL13
	} else if level == 12 {
		cycle = chestOrder.MainCycleKL12
	} else if level == 11 {
		cycle = chestOrder.MainCycleKL11
	} else if level == 10 {
		cycle = chestOrder.MainCycleKL10
	} else {
		cycle = chestOrder.MainCycle
	}
	return append(cycle, append(cycle, cycle...)...)
}
