package structs

type ClanSearchResult struct {
	Items  []Clan `json:"items"`
	Paging Paging `json:"paging"`
}

type Clan struct {
	Tag               string       `json:"tag"`
	Name              string       `json:"name"`
	BadgeId           int          `json:"badgeId"`
	Type              string       `json:"type"`
	ClanScore         int          `json:"clanScore"`
	ClanWarTrophies   int          `json:"clanWarTrophies"`
	RequiredTrophies  int          `json:"requiredTrophies"`
	DonationsPerWeek  int          `json:"donationsPerWeek"`
	ClanChestLevel    int          `json:"clanChestLevel"`
	ClanChestMaxLevel int          `json:"clanChestMaxLevel"`
	Members           int          `json:"members"`
	Location          Location     `json:"location"`
	Description       string       `json:"description"`
	ClanChestStatus   string       `json:"clanChestStatus"`
	ClanChestPoints   int          `json:"clanChestPoints"`
	MemberList        []ClanMember `json:"memberList"`
	Participants      int          `json:"participants"`
	BattlesPlayed     int          `json:"battlesPlayed"`
	Wins              int          `json:"wins"`
	Crowns            int          `json:"crowns"`

	Fame         int `json:"fame"`
	RepairPoints int `json:"repairPoints"`
}

type ClanMemberList struct {
	Items  []ClanMember `json:"items"`
	Paging Paging       `json:"paging"`
}

type ClanMember struct {
	Tag               string `json:"tag"`
	Name              string `json:"name"`
	ExpLevel          int    `json:"expLevel"`
	Trophies          int    `json:"trophies"`
	Arena             Arena  `json:"arena"`
	Role              string `json:"role"`
	ClanRank          int    `json:"clanRank"`
	PreviousClanRank  int    `json:"previousClanRank"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
	ClanChestPoints   int    `json:"clanChestPoints"`
	Fame              int    `json:"fame"`
	RepairPoints      int    `json:"repairPoints"`
}

type CurrentWar struct {
	State             string           `json:"state"`
	WarEndTime        string           `json:"warEndTime"`
	CollectionEndTime string           `json:"collectionEndTime"`
	Clan              Clan             `json:"clan"`
	Participants      []WarParticipant `json:"participants"`
	Clans             []Clan           `json:"clans"`
}

type WarLogList struct {
	Items  []WarLog `json:"items"`
	Paging Paging   `json:"paging"`
}

type WarLog struct {
	SeasonId     int              `json:"seasonId"`
	CreatedDate  string           `json:"createdDate"`
	Participants []WarParticipant `json:"participants"`
	Standings    []WarStanding    `json:"standings"`
}

type WarParticipant struct {
	Tag                        string `json:"tag"`
	Name                       string `json:"name"`
	CardsEarned                int    `json:"cardsEarned"`
	BattlesPlayed              int    `json:"battlesPlayed"`
	Wins                       int    `json:"wins"`
	CollectionDayBattlesPlayed int    `json:"collectionDayBattlesPlayed"`
}

type WarStanding struct {
	Clan         Clan `json:"clan"`
	TrophyChange int  `json:"trophyChange"`
}

type RiverClan struct {
	Tag          string       `json:"tag"`
	Name         string       `json:"name"`
	BadgeId      int          `json:"badgeId"`
	ClanScore    int          `json:"clanScore"`
	Participants []ClanMember `json:"participants"`
	Fame         int          `json:"fame"`
	RepairPoints int          `json:"repairPoints"`
	FinishTime   string       `json:"finishTime"`
}

type RiverRaceStanding struct {
	Rank         int       `json:"rank"`
	TrophyChange int       `json:"trophyChange"`
	Clan         RiverClan `json:"clan"`
}

type RiverRace struct {
	State        string      `json:"state"`
	Clan         RiverClan   `json:"clan"`
	Clans        []RiverClan `json:"clans"`
	SectionIndex int         `json: sectionIndex`
}

type RiverRaceLog struct {
	SeasonId     int                 `json:"seasonId"`
	CreatedDate  string              `json:"createdDate"`
	SectionIndex int                 `json:"sectionIndex"`
	Standings    []RiverRaceStanding `json:"standings"`
}

type RiverRaceLogList struct {
	Items  []RiverRaceLog `json:"items"`
	Paging Paging         `json:"paging"`
}
