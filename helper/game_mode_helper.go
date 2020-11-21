package helper

import (
	data "github.com/yautah/goroyale/structs/json"
)

var GameModeMap map[int]data.GameMode = make(map[int]data.GameMode)

func GetGameModeByID(id int) data.GameMode {
	return GameModeMap[id]
}
