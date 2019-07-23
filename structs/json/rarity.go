package json

type Rarity struct {
	Name                 string `json:"name"`
	LevelCount           int    `json:"level_count"`
	RelativeLevel        int    `json:"relative_level"`
	MirrorRelativeLevel  int    `json:"mirror_relative_level"`
	CloneRelativeLevel   int    `json:"clone_relative_level"`
	DonateCapacity       int    `json:"donate_capacity"`
	SortCapacity         int    `json:"sort_capacity"`
	DonateReward         int    `json:"donate_reward"`
	DonateXp             int    `json:"donate_xp"`
	GoldConversionValue  int    `json:"gold_conversion_value"`
	ChanceWeight         int    `json:"chance_weight"`
	BalanceMultiplier    int    `json:"balance_multiplier"`
	UpgradeExp           []int  `json:"upgrade_exp"`
	UpgradeMaterialCount []int  `json:"upgrade_material_count"`
	UpgradeCost          []int  `json:"upgrade_cost"`
	PowerLevelMultiplier []int  `json:"power_level_multiplier"`
	RefundGems           int    `json:"refund_gems"`
}
