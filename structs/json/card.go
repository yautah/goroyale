package json

import (
	// "fmt"
	"encoding/json"
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
	IconUrl     string `json:"icon_url"`
	IconGoldUrl string `json:"icon_gold_url"`
}

func slugName(name string) string {
	// fmt.Println(name)
	s := utils.CamelSplit(replacer.Replace(name))
	return strings.ToLower(strings.Join(s, "-"))
}

func (self *Card) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID          int    `json:"id"`
		Key         string `json:"key"`
		Name        string `json:"name"`
		Elixir      int    `json:"elixir"`
		Type        string `json:"type"`
		Rarity      string `json:"rarity"`
		Arena       int    `json:"arena"`
		Description string `json:"description"`
		IconUrl     string `json:"iconUrl"`
		IconGoldUrl string `json:"iconGoldUrl"`
	}{
		Key:         self.Key,
		Name:        self.Name,
		Elixir:      self.Elixir,
		Type:        self.Type,
		Rarity:      self.Rarity,
		Arena:       self.Arena,
		Description: self.Description,
		ID:          self.ID,
		IconUrl:     self.IconUrl,
		IconGoldUrl: self.IconGoldUrl,
	})
}
