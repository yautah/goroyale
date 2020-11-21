package json

type GameMode struct {
	Name                string   `json:"name"`
	Tid                 string   `json:"tid"`
	CardLevelAdjustment string   `json:"card_level_adjustment"`
	DeckSelection       string   `json:"deck_selection"`
	BattleTimeline      string   `json:"battle_timeline"`
	PredefinedDecks     []string `json:"predefined_decks"`
	SeparateTeamDecks   bool     `json:"separate_team_decks"`
	Players             string   `json:"players"`
	Icon                string   `json:"icon"`
	ClanWarDescription  string   `json:"clan_war_description"`
	DescriptionTid      string   `json:"description_tid"`
	BattleStartCooldown int      `json:"battle_start_cooldown"`
	ValidChallengeMode  bool     `json:"valid_challenge_mode"`
	ValidFriendlyMode   bool     `json:"valid_friendly_mode"`
	NoDraws             bool     `json:"no_draws"`
	ID                  int      `json:"id"`
	NameEn              string   `json:"name_en"`
	Lang                struct {
		Tid struct {
			En    string `json:"en"`
			Fr    string `json:"fr"`
			De    string `json:"de"`
			Es    string `json:"es"`
			It    string `json:"it"`
			Nl    string `json:"nl"`
			No    string `json:"no"`
			Tr    string `json:"tr"`
			Jp    string `json:"jp"`
			Kr    string `json:"kr"`
			Ru    string `json:"ru"`
			Ar    string `json:"ar"`
			Pt    string `json:"pt"`
			Cn    string `json:"cn"`
			Cnt   string `json:"cnt"`
			Fa    string `json:"fa"`
			ID    string `json:"id"`
			Ms    string `json:"ms"`
			Th    string `json:"th"`
			Fi    string `json:"fi"`
			Vi    string `json:"vi"`
			ScKey string `json:"sc_key"`
		} `json:"tid"`
		ClanWarDescription struct {
			En    string `json:"en"`
			Fr    string `json:"fr"`
			De    string `json:"de"`
			Es    string `json:"es"`
			It    string `json:"it"`
			Nl    string `json:"nl"`
			No    string `json:"no"`
			Tr    string `json:"tr"`
			Jp    string `json:"jp"`
			Kr    string `json:"kr"`
			Ru    string `json:"ru"`
			Ar    string `json:"ar"`
			Pt    string `json:"pt"`
			Cn    string `json:"cn"`
			Cnt   string `json:"cnt"`
			Fa    string `json:"fa"`
			ID    string `json:"id"`
			Ms    string `json:"ms"`
			Th    string `json:"th"`
			Fi    string `json:"fi"`
			Vi    string `json:"vi"`
			ScKey string `json:"sc_key"`
		} `json:"clan_war_description"`
		DescriptionTid struct {
			En    string `json:"en"`
			Fr    string `json:"fr"`
			De    string `json:"de"`
			Es    string `json:"es"`
			It    string `json:"it"`
			Nl    string `json:"nl"`
			No    string `json:"no"`
			Tr    string `json:"tr"`
			Jp    string `json:"jp"`
			Kr    string `json:"kr"`
			Ru    string `json:"ru"`
			Ar    string `json:"ar"`
			Pt    string `json:"pt"`
			Cn    string `json:"cn"`
			Cnt   string `json:"cnt"`
			Fa    string `json:"fa"`
			ID    string `json:"id"`
			Ms    string `json:"ms"`
			Th    string `json:"th"`
			Fi    string `json:"fi"`
			Vi    string `json:"vi"`
			ScKey string `json:"sc_key"`
		} `json:"description_tid"`
	} `json:"_lang"`
}