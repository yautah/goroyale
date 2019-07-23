package json

type Arena struct {
	Name                       string      `json:"name"`
	Arena                      int         `json:"arena"`
	ArenaName                  string      `json:"arena_name"`
	ChestArena                 string      `json:"chest_arena"`
	TvArena                    string      `json:"tv_arena"`
	IsInUse                    bool        `json:"is_in_use"`
	TrainingCamp               bool        `json:"training_camp"`
	TrophyLimit                int         `json:"trophy_limit"`
	DemoteTrophyLimit          int         `json:"demote_trophy_limit"`
	ChestRewardMultiplier      int         `json:"chest_reward_multiplier"`
	ShopChestRewardMultiplier  int         `json:"shop_chest_reward_multiplier"`
	RequestSize                int         `json:"request_size"`
	MaxDonationCountCommon     int         `json:"max_donation_count_common"`
	MaxDonationCountRare       int         `json:"max_donation_count_rare"`
	MaxDonationCountEpic       int         `json:"max_donation_count_epic"`
	MatchmakingMinTrophyDelta  int         `json:"matchmaking_min_trophy_delta"`
	MatchmakingMaxTrophyDelta  int         `json:"matchmaking_max_trophy_delta"`
	MatchmakingMaxSeconds      int         `json:"matchmaking_max_seconds"`
	DailyDonationCapacityLimit int         `json:"daily_donation_capacity_limit"`
	BattleRewardGold           int         `json:"battle_reward_gold"`
	SeasonRewardChest          interface{} `json:"season_reward_chest"`
	QuestCycle                 string      `json:"quest_cycle"`
	Key                        string      `json:"key"`
	Title                      string      `json:"title"`
	Subtitle                   string      `json:"subtitle"`
	ArenaID                    int         `json:"arena_id"`
	LeagueID                   interface{} `json:"league_id"`
	IconUrl                    string      `json:"icon_url"`
	ID                         int         `json:"id"`
}
