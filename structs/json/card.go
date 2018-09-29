package json

import (
	// "fmt"
	"github.com/yautah/goroyale/locale"
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
	NameEn      string `json:"nameEn"`
	Name        string `json:"name"`
	Elixir      int    `json:"elixir"`
	Type        string `json:"type"`
	Rarity      string `json:"rarity"`
	Arena       int    `json:"arena"`
	Description string `json:"description"`
	ID          int    `json:"id"`
}

type CardTroop struct {
	Name                       string `json:"name"`
	IconFile                   string `json:"iconFile"`
	UnlockArena                string `json:"unlockArena"`
	UnlockLevel                int    `json:"unlockLevel"`
	Rarity                     string `json:"rarity"`
	ManaCost                   int    `json:"manaCost"`
	ManaCostFromSummonerMana   bool   `json:"manaCostFromSummonerMana"`
	NotInUse                   bool   `json:"notInUse"`
	Mirror                     bool   `json:"mirror"`
	CustomDeployTime           int    `json:"customDeployTime"`
	SummonCharacter            string `json:"summonCharacter"`
	SummonNumber               int    `json:"summonNumber"`
	SummonCharacterLevelIndex  int    `json:"summonCharacterLevelIndex"`
	SummonCharacterSecondCount int    `json:"summonCharacterSecondCount"`
	SummonRadius               int    `json:"summonRadius"`
	SummonWidth                int    `json:"summonWidth"`
	SummonDeployDelay          int    `json:"summonDeployDelay"`
	SummonDeployDelaySecond    int    `json:"summonDeployDelaySecond"`
	Radius                     int    `json:"radius"`
	Height                     int    `json:"height"`
	SpellAsDeploy              bool   `json:"spellAsDeploy"`
	CanPlaceOnBuildings        bool   `json:"canPlaceOnBuildings"`
	CanPlaceOnWater            bool   `json:"canPlaceOnWater"`
	InstantDamage              int    `json:"instantDamage"`
	DurationSeconds            int    `json:"durationSeconds"`
	InstantHeal                int    `json:"instantHeal"`
	HealPerSecond              int    `json:"healPerSecond"`
	Effect                     string `json:"effect"`
	Pushback                   int    `json:"pushback"`
	MultipleProjectiles        int    `json:"multipleProjectiles"`
	BuffTime                   int    `json:"buffTime"`
	BuffTimeIncreasePerLevel   int    `json:"buffTimeIncreasePerLevel"`
	BuffNumber                 int    `json:"buffNumber"`
	OnlyOwnTroops              bool   `json:"onlyOwnTroops"`
	OnlyEnemies                bool   `json:"onlyEnemies"`
	CanDeployOnEnemySide       bool   `json:"canDeployOnEnemySide"`
	TouchdownLimitedDeploy     bool   `json:"touchdownLimitedDeploy"`
	TID                        string `json:"TID"`
	TIDINFO                    string `json:"TID_INFO"`
	HideRadiusIndicator        bool   `json:"hideRadiusIndicator"`
	ElixirProductionStopTime   int    `json:"elixirProductionStopTime"`
	DarkMirror                 bool   `json:"darkMirror"`
	StatsUnderInfo             bool   `json:"statsUnderInfo"`
	HpModifier                 int    `json:"hpModifier"`
	DamageModifier             int    `json:"damageModifier"`
	DeployWTileMargin          int    `json:"deployWTileMargin"`
	FullLaneDeploy             bool   `json:"fullLaneDeploy"`
	Scid                       int    `json:"scid"`
}

type CardBuilding struct {
	Name                       string `json:"name"`
	IconFile                   string `json:"iconFile"`
	UnlockArena                string `json:"unlockArena"`
	UnlockLevel                int    `json:"unlockLevel"`
	Rarity                     string `json:"rarity"`
	ManaCost                   int    `json:"manaCost"`
	ManaCostFromSummonerMana   bool   `json:"manaCostFromSummonerMana"`
	NotInUse                   bool   `json:"notInUse"`
	Mirror                     bool   `json:"mirror"`
	CustomDeployTime           int    `json:"customDeployTime"`
	SummonCharacter            string `json:"summonCharacter"`
	SummonNumber               int    `json:"summonNumber"`
	SummonCharacterLevelIndex  int    `json:"summonCharacterLevelIndex"`
	SummonCharacterSecondCount int    `json:"summonCharacterSecondCount"`
	SummonRadius               int    `json:"summonRadius"`
	SummonWidth                int    `json:"summonWidth"`
	SummonDeployDelay          int    `json:"summonDeployDelay"`
	SummonDeployDelaySecond    int    `json:"summonDeployDelaySecond"`
	Radius                     int    `json:"radius"`
	Height                     int    `json:"height"`
	SpellAsDeploy              bool   `json:"spellAsDeploy"`
	CanPlaceOnBuildings        bool   `json:"canPlaceOnBuildings"`
	CanPlaceOnWater            bool   `json:"canPlaceOnWater"`
	InstantDamage              int    `json:"instantDamage"`
	DurationSeconds            int    `json:"durationSeconds"`
	InstantHeal                int    `json:"instantHeal"`
	HealPerSecond              int    `json:"healPerSecond"`
	Pushback                   int    `json:"pushback"`
	MultipleProjectiles        int    `json:"multipleProjectiles"`
	BuffTime                   int    `json:"buffTime"`
	BuffTimeIncreasePerLevel   int    `json:"buffTimeIncreasePerLevel"`
	BuffNumber                 int    `json:"buffNumber"`
	OnlyOwnTroops              bool   `json:"onlyOwnTroops"`
	OnlyEnemies                bool   `json:"onlyEnemies"`
	CanDeployOnEnemySide       bool   `json:"canDeployOnEnemySide"`
	TouchdownLimitedDeploy     bool   `json:"touchdownLimitedDeploy"`
	CastSound                  string `json:"castSound"`
	TID                        string `json:"TID"`
	TIDINFO                    string `json:"TID_INFO"`
	HideRadiusIndicator        bool   `json:"hideRadiusIndicator"`
	ElixirProductionStopTime   int    `json:"elixirProductionStopTime"`
	DarkMirror                 bool   `json:"darkMirror"`
	StatsUnderInfo             bool   `json:"statsUnderInfo"`
	HpModifier                 int    `json:"hpModifier"`
	DamageModifier             int    `json:"damageModifier"`
	DeployWTileMargin          int    `json:"deployWTileMargin"`
	FullLaneDeploy             bool   `json:"fullLaneDeploy"`
	Scid                       int    `json:"scid"`
}

type CardOther struct {
	Name                       string `json:"name"`
	IconFile                   string `json:"iconFile"`
	UnlockArena                string `json:"unlockArena"`
	UnlockLevel                int    `json:"unlockLevel"`
	Rarity                     string `json:"rarity"`
	ManaCost                   int    `json:"manaCost"`
	ManaCostFromSummonerMana   bool   `json:"manaCostFromSummonerMana"`
	NotInUse                   bool   `json:"notInUse"`
	Mirror                     bool   `json:"mirror"`
	CustomDeployTime           int    `json:"customDeployTime"`
	SummonNumber               int    `json:"summonNumber"`
	SummonCharacterLevelIndex  int    `json:"summonCharacterLevelIndex"`
	SummonCharacterSecondCount int    `json:"summonCharacterSecondCount"`
	SummonRadius               int    `json:"summonRadius"`
	SummonWidth                int    `json:"summonWidth"`
	SummonDeployDelay          int    `json:"summonDeployDelay"`
	SummonDeployDelaySecond    int    `json:"summonDeployDelaySecond"`
	Radius                     int    `json:"radius"`
	Height                     int    `json:"height"`
	Projectile                 string `json:"projectile"`
	SpellAsDeploy              bool   `json:"spellAsDeploy"`
	CanPlaceOnBuildings        bool   `json:"canPlaceOnBuildings"`
	CanPlaceOnWater            bool   `json:"canPlaceOnWater"`
	InstantDamage              int    `json:"instantDamage"`
	DurationSeconds            int    `json:"durationSeconds"`
	InstantHeal                int    `json:"instantHeal"`
	HealPerSecond              int    `json:"healPerSecond"`
	Pushback                   int    `json:"pushback"`
	MultipleProjectiles        int    `json:"multipleProjectiles"`
	BuffTime                   int    `json:"buffTime"`
	BuffTimeIncreasePerLevel   int    `json:"buffTimeIncreasePerLevel"`
	BuffNumber                 int    `json:"buffNumber"`
	OnlyOwnTroops              bool   `json:"onlyOwnTroops"`
	OnlyEnemies                bool   `json:"onlyEnemies"`
	CanDeployOnEnemySide       bool   `json:"canDeployOnEnemySide"`
	TouchdownLimitedDeploy     bool   `json:"touchdownLimitedDeploy"`
	CastSound                  string `json:"castSound"`
	TID                        string `json:"TID"`
	TIDINFO                    string `json:"TID_INFO"`
	HideRadiusIndicator        bool   `json:"hideRadiusIndicator"`
	ElixirProductionStopTime   int    `json:"elixirProductionStopTime"`
	DarkMirror                 bool   `json:"darkMirror"`
	StatsUnderInfo             bool   `json:"statsUnderInfo"`
	HpModifier                 int    `json:"hpModifier"`
	DamageModifier             int    `json:"damageModifier"`
	DeployWTileMargin          int    `json:"deployWTileMargin"`
	FullLaneDeploy             bool   `json:"fullLaneDeploy"`
	Scid                       int    `json:"scid"`
}

func slugName(name string) string {
	// fmt.Println(name)
	s := utils.CamelSplit(replacer.Replace(name))
	return strings.ToLower(strings.Join(s, "-"))
}

func (c *CardTroop) ConvertToCard() (card *Card) {
	if c.TID == "" || c.NotInUse {
		return nil
	}
	return &Card{
		Key:         slugName(locale.GetTranslationById(c.TID, "En")),
		NameEn:      locale.GetTranslationById(c.TID, "En"),
		Name:        locale.GetLocaleById(c.TID),
		Elixir:      c.ManaCost,
		Type:        "Troop",
		Rarity:      c.Rarity,
		Description: locale.GetLocaleById(c.TIDINFO),
		ID:          c.Scid,
	}

}

func (c *CardBuilding) ConvertToCard() (card *Card) {
	if c.TID == "" || c.NotInUse {
		return nil
	}
	return &Card{
		Key:         slugName(locale.GetTranslationById(c.TID, "En")),
		NameEn:      locale.GetTranslationById(c.TID, "En"),
		Name:        locale.GetLocaleById(c.TID),
		Elixir:      c.ManaCost,
		Type:        "Troop",
		Rarity:      c.Rarity,
		Description: locale.GetLocaleById(c.TIDINFO),
		ID:          c.Scid,
	}
}

func (c *CardOther) ConvertToCard() (card *Card) {
	if c.TID == "" || c.NotInUse {
		return nil
	}
	return &Card{
		Key:         slugName(locale.GetTranslationById(c.TID, "En")),
		NameEn:      locale.GetTranslationById(c.TID, "En"),
		Name:        locale.GetLocaleById(c.TID),
		Elixir:      c.ManaCost,
		Type:        "Troop",
		Rarity:      c.Rarity,
		Description: locale.GetLocaleById(c.TIDINFO),
		ID:          c.Scid,
	}
}
