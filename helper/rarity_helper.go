package helper

import (
	// "fmt"
	data "github.com/yautah/goroyale/structs/json"
)

var raritiesMap map[string]data.Rarity = make(map[string]data.Rarity)

func GetRarityByName(name string) data.Rarity {
	return raritiesMap[name]
}

func GetCardUpgradeCount(rarityName string, currentLevel int) int {
	if currentLevel == 0 || rarityName == "" {
		return 0
	}
	return GetRarityByName(rarityName).UpgradeMaterialCount[currentLevel-1]
}

func GetCardUpgradeGold(rarityName string, currentLevel int) int {
	if currentLevel == 0 || rarityName == "" {
		return 0
	}
	return GetRarityByName(rarityName).UpgradeCost[currentLevel-1]
}

func GetCardMaxCount(rarityName string, currentLevel int, currentCount int) int {
	if currentLevel == 0 || rarityName == "" {
		return 0
	}
	rarity := GetRarityByName(rarityName)
	var count = 0
	for i := currentLevel; i < rarity.LevelCount; i++ {
		count = count + GetCardUpgradeCount(rarity.Name, i)
	}
	return count - currentCount
}

func GetCardMaxGold(rarityName string, currentLevel int) int {
	if currentLevel == 0 || rarityName == "" {
		return 0
	}
	rarity := GetRarityByName(rarityName)
	var gold = 0
	for i := currentLevel; i < rarity.LevelCount; i++ {
		gold = gold + GetCardUpgradeGold(rarity.Name, i)
	}
	return gold
}

func GetCardLevelGold(rarityName string, currentLevel int) int {
	if currentLevel == 0 || rarityName == "" {
		return 0
	}
	rarity := GetRarityByName(rarityName)
	var gold = 0
	for i := 1; i < currentLevel; i++ {
		gold = gold + GetCardUpgradeGold(rarity.Name, i)
	}
	return gold

}

func GetCardCanUpgradeTo(rarityName string, currentLevel int, currentCount int) (level int) {
	if currentLevel == 0 || rarityName == "" {
		return 0
	}
	rarity := GetRarityByName(rarityName)
	for i := currentLevel - 1; i < len(rarity.UpgradeMaterialCount); i++ {
		currentCount = currentCount - rarity.UpgradeMaterialCount[i]
		level = i + 1
		if currentCount < 0 {
			break
		}
	}

	switch rarityName {
	case "Rare":
		level = level + 2
	case "Epic":
		level = level + 5
	case "Legendary":
		level = level + 8
	}

	return

}
