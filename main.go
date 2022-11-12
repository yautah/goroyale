package main

import (
	"fmt"
	// "github.com/yautah/goroyale/client"
	//"github.com/yautah/goroyale/helper"
	"github.com/yautah/goroyale/structs"
	// "net/url"
	// "encoding/json"
	// "github.com/gocolly/colly"
	// "os"
	// "regexp"
	// "strconv"
	// "strings"
)

type Card struct {
	Name          string                   `json:"name"`
	Rarity        string                   `json:"rarity"`
	Arena         string                   `json:"arena"`
	CardType      string                   `json:"type"`
	Description   string                   `json:"description"`
	CardStats     []map[string]interface{} `json:"cardStats"`
	CardUnits     []string                 `json:"cardUnits"`
	CardUnitsData []map[string]interface{} `json:"cardUnitsData"`
}

func main() {
	//s := helper.GetGameModeByID(72000266)
	//fmt.Println(s)

	// fmt.Println(helper.TOTAL_CARD_GOLD)
	// fmt.Println(len(helper.CardsIdMap))

	up := structs.UpcomingChestList{}

	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 0,
		Name:  "Gold Crate",
	})

	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 1,
		Name:  "Silver Chest",
	})
	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 2,
		Name:  "Golden Chest",
	})
	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 3,
		Name:  "Plentiful Gold Crate",
	})
	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 4,
		Name:  "Silver Chest",
	})
	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 5,
		Name:  "Silver Chest",
	})
	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 6,
		Name:  "Golden Chest",
	})
	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 7,
		Name:  "Silver Chest",
	})
	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 8,
		Name:  "Giant Chest",
	})

	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 44,
		Name:  "Magical Chest",
	})

	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 276,
		Name:  "Legendary Chest",
	})
	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 578,
		Name:  "Mega Lightning Chest",
	})
	up.Items = append(up.Items, structs.UpcomingChest{
		Index: 579,
		Name:  "Epic Chest",
	})
	fmt.Println(up.GetChestOrder(14))

}
