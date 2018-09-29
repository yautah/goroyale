package helper

import (
	"encoding/json"
	"fmt"
	. "github.com/yautah/goroyale/assets"
	data "github.com/yautah/goroyale/structs/json"
)

const Endpoint = "http://www.madn.xyz/assets/statics"

var ArenasIdMap map[int]data.Arena = make(map[int]data.Arena)
var ArenasNameMap map[string]data.Arena = make(map[string]data.Arena)
var ArenasNumberMap map[int]data.Arena = make(map[int]data.Arena)

func init() {
	var arenas []data.Arena

	data, err := Asset("assets/arenas.json")
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &arenas)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, arena := range arenas {
		arena.IconUrl = fmt.Sprintf("%s/arenas/%d.png", Endpoint, arena.Scid)
		ArenasNameMap[arena.Name] = arena
		ArenasIdMap[arena.Scid] = arena
		ArenasNumberMap[arena.Arena] = arena
	}
}

func GetArenaById(id int) data.Arena {
	return ArenasIdMap[id]
}

func GetArenaByName(name string) data.Arena {
	return ArenasNameMap[name]
}

func GetArenaByNumber(number int) data.Arena {
	return ArenasNumberMap[number]
}
