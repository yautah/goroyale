package helper

import ()

const (
	MinBadgeId = 16000000
	MaxBadgeId = 16000179
)

var BadgesMap map[int]string = make(map[int]string)

func GetBadgeById(id int) string {
	return BadgesMap[id]
}
