package client

import (
	"encoding/json"
	// "fmt"
	. "github.com/yautah/goroyale/structs"
	"net/url"
	"strings"
)

// Player retrieves a player by their tag.
func (c *Client) Player(tag string) (player Player, err error) {
	var b []byte
	path := "/players/%23" + strings.ToUpper(tag)
	if b, err = c.get(path, map[string][]string{}); err == nil {
		err = json.Unmarshal(b, &player)
	}
	return
}

// PlayerBattles s battles a player participated in.
func (c *Client) PlayerBattles(tag string) (battles []BattleLog, err error) {
	var b []byte
	path := "/players/%23" + strings.ToUpper(tag) + "/battlelog"
	if b, err = c.get(path, map[string][]string{}); err == nil {
		err = json.Unmarshal(b, &battles)
	}
	return
}

// PlayerChests s a player's upcoming chests.
func (c *Client) PlayerChests(tag string) (chests UpcomingChestList, err error) {
	var b []byte
	path := "/players/%23" + strings.ToUpper(tag) + "/upcomingchests"
	if b, err = c.get(path, map[string][]string{}); err == nil {
		err = json.Unmarshal(b, &chests)
	}
	return
}

// ClanSearch searches for a clan using the provided parmameters.
func (c *Client) ClanSearch(params url.Values) (clans ClanSearchResult, err error) {
	var b []byte
	path := "/clans"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &clans)
	}
	return
}

// Clan returns info about a specific clan.
func (c *Client) Clan(tag string) (clan Clan, err error) {
	var b []byte
	path := "/clans/%23" + strings.ToUpper(tag)
	if b, err = c.get(path, map[string][]string{}); err == nil {
		err = json.Unmarshal(b, &clan)
	}
	return
}

// Clans works like Clan but can return multiple clans.
func (c *Client) ClanMembers(tag string, params url.Values) (members ClanMemberList, err error) {
	var b []byte
	path := "/clans/%23" + strings.ToUpper(tag) + "/members"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &members)
	}
	return
}

// ClanWar returns data about the current clan war.
func (c *Client) ClanWar(tag string) (war CurrentWar, err error) {
	var b []byte
	path := "/clans/%23" + strings.ToUpper(tag) + "/currentwar"
	if b, err = c.get(path, map[string][]string{}); err == nil {
		err = json.Unmarshal(b, &war)
	}
	return
}

// ClanWarLog returns data about past clan wars.
func (c *Client) ClanWarLog(tag string, params url.Values) (warlog WarLogList, err error) {
	var b []byte
	path := "/clans/%23" + strings.ToUpper(tag) + "/warlog"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &warlog)
	}
	return
}

// ClanWar returns data about the current clan war.
func (c *Client) CurrentRiverRace(tag string) (race RiverRace, err error) {
	var b []byte
	path := "/clans/%23" + strings.ToUpper(tag) + "/currentriverrace"
	if b, err = c.get(path, map[string][]string{}); err == nil {
		err = json.Unmarshal(b, &race)
	}
	return
}

// ClanWarLog returns data about past clan wars.
func (c *Client) RiverRaceLog(tag string, params url.Values) (warlog RiverRaceLogList, err error) {
	var b []byte
	path := "/clans/%23" + strings.ToUpper(tag) + "/riverracelog"
	if b, err = c.get(path, params); err == nil {
		err = json.Unmarshal(b, &warlog)
	}
	return
}
