package main

import (
	"fmt"
	// "github.com/yautah/goroyale/client"
	"github.com/yautah/goroyale/helper"
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
	s := helper.GetGameModeByID(72000266)
	fmt.Println(s)

	// fmt.Println(helper.TOTAL_CARD_GOLD)
	// fmt.Println(len(helper.CardsIdMap))

}
