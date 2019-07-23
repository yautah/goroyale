package main

import (
	"fmt"
	// "github.com/yautah/goroyale/client"
	"github.com/yautah/goroyale/helper"
	// "net/url"
)

func main() {
	s := helper.GetArenaById(54000014)
	fmt.Println(s)

	fmt.Println(helper.TOTAL_CARD_GOLD)
	fmt.Println(len(helper.CardsIdMap))

}
