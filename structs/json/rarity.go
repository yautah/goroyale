package json

type Rarity struct {
	Name                   string `json:"name"`
	LevelCount             int    `json:"levelCount"`
	RelativeLevel          int    `json:"relativeLevel"`
	MirrorRelativeLevel    int    `json:"mirrorRelativeLevel"`
	CloneRelativeLevel     int    `json:"cloneRelativeLevel"`
	DonateCapacity         int    `json:"donateCapacity"`
	SortCapacity           int    `json:"sortCapacity"`
	DonateReward           int    `json:"donateReward"`
	DonateXP               int    `json:"donateXP"`
	GoldConversionValue    int    `json:"goldConversionValue"`
	ChanceWeight           int    `json:"chanceWeight"`
	BalanceMultiplier      int    `json:"balanceMultiplier"`
	UpgradeExp             []int  `json:"upgradeExp"`
	UpgradeMaterialCount   []int  `json:"upgradeMaterialCount"`
	UpgradeCost            []int  `json:"upgradeCost"`
	PowerLevelMultiplier   []int  `json:"powerLevelMultiplier"`
	RefundGems             int    `json:"refundGems"`
	TID                    string `json:"TID"`
	CardBaseFileName       string `json:"cardBaseFileName"`
	BigFrameExportName     string `json:"bigFrameExportName"`
	CardBaseExportName     string `json:"cardBaseExportName"`
	StackedCardExportName  string `json:"stackedCardExportName"`
	CardRewardExportName   string `json:"cardRewardExportName"`
	CastEffect             string `json:"castEffect"`
	InfoTitleExportName    string `json:"infoTitleExportName"`
	CardRarityBGExportName string `json:"cardRarityBGExportName"`
	SortOrder              int    `json:"sortOrder"`
	Red                    []int  `json:"Red"`
	Green                  []int  `json:"green"`
	Blue                   []int  `json:"blue"`
	AppearEffect           string `json:"appearEffect"`
	BuySound               string `json:"buySound"`
	LoopEffect             string `json:"loopEffect"`
	CardTxtBgFrameIdx      int    `json:"cardTxtBgFrameIdx"`
	RotateExportName       string `json:"rotateExportName,omitempty"`
	IconSWF                string `json:"iconSWF"`
	IconExportName         string `json:"iconExportName"`
	CardGlowInstanceName   string `json:"cardGlowInstanceName,omitempty"`
	SpellSelectedSound     string `json:"spellSelectedSound,omitempty"`
	SpellAvailableSound    string `json:"spellAvailableSound,omitempty"`
}
