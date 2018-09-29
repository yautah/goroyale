package helper

import (
	"fmt"
)

const (
	MinBadgeId = 16000000
	MaxBadgeId = 16000179
)

var BadgesMap map[int]string = make(map[int]string)

func init() {
	for i := MinBadgeId; i <= MaxBadgeId; i++ {
		BadgesMap[i] = fmt.Sprintf("%s/badges/%d.png", Endpoint, i)
	}
}

func GetBadgeById(id int) string {
	return BadgesMap[id]
}
