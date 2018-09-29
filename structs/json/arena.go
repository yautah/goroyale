package json

type Arena struct {
	Name                       string `json:"name"`
	TID                        string `json:"TID"`
	SubtitleTID                string `json:"subtitleTID"`
	Arena                      int    `json:"arena"`
	ChestArena                 string `json:"chestArena"`
	TvArena                    string `json:"tvArena"`
	IsInUse                    bool   `json:"isInUse"`
	TrainingCamp               bool   `json:"trainingCamp"`
	PVEArena                   bool   `json:"pVEArena"`
	TrophyLimit                int    `json:"trophyLimit"`
	DemoteTrophyLimit          int    `json:"demoteTrophyLimit"`
	ChestRewardMultiplier      int    `json:"chestRewardMultiplier"`
	ShopChestRewardMultiplier  int    `json:"shopChestRewardMultiplier"`
	RequestSize                int    `json:"requestSize"`
	MaxDonationCountCommon     int    `json:"maxDonationCountCommon"`
	MaxDonationCountRare       int    `json:"maxDonationCountRare"`
	MaxDonationCountEpic       int    `json:"maxDonationCountEpic"`
	IconSWF                    string `json:"iconSWF"`
	IconExportName             string `json:"iconExportName"`
	MainMenuIconExportName     string `json:"mainMenuIconExportName"`
	MatchmakingMinTrophyDelta  int    `json:"matchmakingMinTrophyDelta"`
	MatchmakingMaxTrophyDelta  int    `json:"matchmakingMaxTrophyDelta"`
	MatchmakingMaxSeconds      int    `json:"matchmakingMaxSeconds"`
	PvpLocation                string `json:"pvpLocation"`
	TeamVsTeamLocation         string `json:"teamVsTeamLocation"`
	DailyDonationCapacityLimit int    `json:"dailyDonationCapacityLimit"`
	BattleRewardGold           int    `json:"battleRewardGold"`
	QuestCycle                 string `json:"questCycle"`
	DisableIn2V2               bool   `json:"disableIn2v2"`
	Scid                       int    `json:"scid"`
	IconUrl                    string `json:"iconUrl"`
}
