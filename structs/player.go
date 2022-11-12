package structs

import (
	// "encoding/json"

	"fmt"
	//log "github.com/sirupsen/logrus"
	"github.com/yautah/goroyale/helper"
	"sort"
)

type Player struct {
	Tag                   string `json:"tag"`
	Name                  string `json:"name"`
	ExpLevel              int    `json:"expLevel"`
	Trophies              int    `json:"trophies"`
	BestTrophies          int    `json:"bestTrophies"`
	Wins                  int    `json:"wins"`
	Losses                int    `json:"losses"`
	BattleCount           int    `json:"battleCount"`
	ThreeCrownWins        int    `json:"threeCrownWins"`
	ChallengeCardsWon     int    `json:"challengeCardsWon"`
	ChallengeMaxWins      int    `json:"challengeMaxWins"`
	TournamentCardsWon    int    `json:"tournamentCardsWon"`
	TournamentBattleCount int    `json:"tournamentBattleCount"`
	Role                  string `json:"role"`
	Donations             int    `json:"donations"`
	DonationsReceived     int    `json:"donationsReceived"`
	TotalDonations        int    `json:"totalDonations"`
	WarDayWins            int    `json:"warDayWins"`
	ClanCardsCollected    int    `json:"clanCardsCollected"`
	StarPoints            int    `json:"starPoints"`

	Clan                 ClanBase         `json:"clan"`
	Arena                Arena            `json:"arena"`
	LeagueStatistics     LeagueStatistics `json:"leagueStatistics"`
	Badges               []Badge          `json:"badges"`
	Achievements         []Achievement    `json:"-"`
	Cards                []Card           `json:"cards"`
	CurrentDeck          []Card           `json:"currentDeck"`
	CurrentFavouriteCard Card             `json:"currentFavouriteCard"`
}

type Badge struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
	Progress int    `json:"progress"`
	Target   int    `json:"target"`
}

type LeagueStatistics struct {
	CurrentSeason  SeasonStatistics `json:"currentSeason"`
	PreviousSeason SeasonStatistics `json:"previousSeason"`
	BestSeason     SeasonStatistics `json:"bestSeason"`
}

type Achievement struct {
	Name   string `json:"name"`
	Stars  int    `json:"stars"`
	Value  int    `json:"value"`
	Target int    `json:"target"`
	Info   string `json:"info"`
}

type SeasonStatistics struct {
	ID           string `json:"id"`
	Rank         int    `json:Rank`
	Trophies     int    `json:"trophies"`
	BestTrophies int    `json:"bestTrophies"`
}

type IconUrl struct {
	Medium string `json:"medium"`
}

type UpcomingChestList struct {
	Items []UpcomingChest `json:"items"`
}

func (uc UpcomingChestList) Len() int {
	return len(uc.Items)
}

func (uc UpcomingChestList) Swap(i, j int) {
	uc.Items[i], uc.Items[j] = uc.Items[j], uc.Items[i]
}

func (uc UpcomingChestList) Less(i, j int) bool {
	return uc.Items[i].Index < uc.Items[j].Index
}

func (uc UpcomingChestList) SpecialChestsCount() int {
	count := 0
	for i := 0; i < 9; i++ {
		item := uc.Items[i]
		//log.Infoln(item.Name)
		if item.Name == "Epic Chest" || item.Name == "Legendary Chest" || item.Name == "Mega Lightning Chest" {
			count = count + 1
		}
	}
	return count
}

func (uc UpcomingChestList) HasWild() bool {
	for _, item := range uc.Items {
		if item.Name == "Royal Wild Chest" {
			return true
		}
	}
	return false
}

func (uc UpcomingChestList) HasCrate() bool {
	for _, item := range uc.Items {
		if item.Name == "Overflowing Gold Crate" {
			return true
		}
	}
	return false
}

func (uc UpcomingChestList) GiantPosition() int {
	for _, item := range uc.Items {
		if item.Name == "Giant Chest" {
			return item.Index
		}
	}
	return -1
}

func (uc UpcomingChestList) MagicalPosition() int {
	for _, item := range uc.Items {
		if item.Name == "Magical Chest" {
			return item.Index
		}
	}
	return -1
}

type BattleLog struct {
	Type               string          `json:"type"`
	BattleTime         string          `json:"battleTime"`
	IsLadderTournament bool            `json:"isLadderTournament"`
	Arena              Arena           `json:"arena"`
	GameMode           GameMode        `json:"gameMode"`
	DeckSelection      string          `json:"deckSelection"`
	Team               []BattleLogTeam `json:"team"`
	Opponent           []BattleLogTeam `json:"opponent"`
}

type BattleLogTeam struct {
	Tag              string   `json:"tag"`
	Name             string   `json:"name"`
	StartingTrophies int      `json:"startingTrophies"`
	TrophyChange     int      `json:"trophyChange"`
	Crowns           int      `json:"crowns"`
	Clan             ClanBase `json:"clan"`
	Cards            []Card   `json:"cards"`
}

func (chestList *UpcomingChestList) GetChestOrder(level int) UpcomingChestList {

	var cycle = helper.GetChestCycle(level)
	position := getChestPosition(cycle, *chestList)
	fmt.Println(position)

	if position == -1 {
		return *chestList
	}

	wildPosition, goldPosition := 0, 0

	//log.Infoln(chestList)
	//specialCount := chestList.SpecialChestsCount()

	for i := position; i < len(cycle); i++ {
		if cycle[i] == "WildChampion" {
			wildPosition = i - position
			break
		}
	}

	for i := position; i < len(cycle); i++ {
		if cycle[i] == "GoldCage_Large" {
			goldPosition = i - position
			break
		}
	}

	//log.Infoln(position, specialCount, wildPosition, goldPosition)

	if !chestList.HasWild() && wildPosition > 0 {
		up := UpcomingChest{Index: wildPosition, Name: "Royal Wild Chest"}
		//chestList.Items = append(chestList.Items[:9], append([]UpcomingChest{up}, chestList.Items[9:]...)...)
		chestList.Items = append(chestList.Items, up)
	}
	if !chestList.HasCrate() && goldPosition > 0 {
		up := UpcomingChest{Index: goldPosition, Name: "Overflowing Gold Crate"}
		//chestList.Items = append(chestList.Items[:9], append([]UpcomingChest{up}, chestList.Items[9:]...)...)
		chestList.Items = append(chestList.Items, up)
	}

	sort.Sort(*chestList)
	return *chestList
}

func getChestPosition(cycle []string, chestList UpcomingChestList) int {

	list := make([]UpcomingChest, 0)
	magicalPosition := chestList.MagicalPosition()
	giantPosition := chestList.GiantPosition()

	for i := 0; i < 9; i++ {
		item := chestList.Items[i]
		//if item.Name != "Epic Chest" && item.Name != "Legendary Chest" && item.Name != "Mega Lightning Chest" {
		//list = append(list, item)
		//}
		list = append(list, item)
	}

	//log.Infoln(list, len(cycle))
	//log.Infoln(cycle)

	for i := 0; i < len(cycle); i++ {
		var equal = 0
		for j := 0; j < len(list); j++ {

			fmt.Println(helper.ChestNameMap[cycle[i+j]]+" | "+list[j].Name, helper.ChestNameMap[cycle[i+j]] == list[j].Name)
			if list[j].Name == "Epic Chest" ||
				list[j].Name == "Legendary Chest" ||
				list[j].Name == "Mega Lightning Chest" ||
				helper.ChestNameMap[cycle[i+j]] == list[j].Name {
				equal = equal + 1
			} else {
				break
			}
		}
		if equal == len(list) {
			//log.Infoln(i, magicalPosition, giantPosition)
			//return i
			if cycle[i+magicalPosition] == "Magic" && cycle[i+giantPosition] == "Giant" {
				return i
			} else {
				continue
			}
		}
	}
	return -1
}
