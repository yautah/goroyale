package helper

import (
	data "github.com/yautah/goroyale/structs/json"
)

const Endpoint = "http://www.madn.xyz/assets/statics"

var ArenasIdMap map[int]data.Arena = make(map[int]data.Arena)
var ArenasNameMap map[string]data.Arena = make(map[string]data.Arena)
var ArenasNumberMap map[int]data.Arena = make(map[int]data.Arena)

func GetArenaById(id int) data.Arena {
	return ArenasIdMap[id]
}

func GetArenaByName(name string) data.Arena {
	return ArenasNameMap[name]
}

func GetArenaByNumber(number int) data.Arena {
	return ArenasNumberMap[number]
}
