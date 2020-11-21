package structs

import (
// "encoding/json"
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
