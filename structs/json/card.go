package json

import (
	// "fmt"
	"github.com/yautah/goroyale/utils"
	"strings"
)

var replacer = strings.NewReplacer(
	".", "",
	" ", "",
	"-", "",
)

type Card struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Elixir      int    `json:"elixir"`
	Type        string `json:"type"`
	Rarity      string `json:"rarity"`
	Arena       int    `json:"arena"`
	Description string `json:"description"`
	ID          int    `json:"id"`
}

func slugName(name string) string {
	// fmt.Println(name)
	s := utils.CamelSplit(replacer.Replace(name))
	return strings.ToLower(strings.Join(s, "-"))
}
