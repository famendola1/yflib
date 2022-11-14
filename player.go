package yflib

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/famendola1/yfquery"
	"github.com/famendola1/yfquery/schema"
)

func findPlayer(players *schema.Players, name string) (*schema.Player, error) {
	for _, p := range players.Player {
		if strings.ToLower(p.Name.Full) == strings.ToLower(name) {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("player %q not found", name)
}

// GetPlayerKey returns the key of the player, if the player is found.
func GetPlayerKey(client *http.Client, leagueKey, name string) (string, error) {
	if len(name) < 3 {
		return "", fmt.Errorf("name (%q) must contain at least 3 letters", name)
	}

	fc, err := yfquery.League().Key(leagueKey).Players().Search(name).Get(client)
	if err != nil {
		return "", err
	}

	player, err := findPlayer(&fc.League.Players, name)
	if err != nil {
		return "", err
	}
	return player.PlayerKey, nil
}

// GetPlayer searches the given league for a player with the provided player name.
// If the player is not found, an error is returned. name should contain at
// least 3 letters.
func GetPlayer(client *http.Client, leagueKey, name string) (*schema.Player, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("name (%q) must contain at least 3 letters", name)
	}

	fc, err := yfquery.League().Key(leagueKey).Players().Search(name).Get(client)
	if err != nil {
		return nil, err
	}

	return findPlayer(&fc.League.Players, name)
}

// SearchPlayers searches the given league for players with the provided player
// name. name should contain at least 3 letters.
func SearchPlayers(client *http.Client, leagueKey, name string) ([]*schema.Player, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("name (%q) must contain at least 3 letters", name)
	}

	fc, err := yfquery.League().Key(leagueKey).Players().Search(name).Get(client)
	if err != nil {
		return nil, err
	}

	var players []*schema.Player
	for _, p := range fc.League.Players.Player {
		players = append(players, &p)
	}

	return players, nil
}

// SearchMultiPlayers searches the given league for players with the provided player
// names. Each name should contain at least 3 letters.
func SearchMultiPlayers(client *http.Client, leagueKey string, names []string) ([]*schema.Player, error) {
	players := []*schema.Player{}

	for _, name := range names {
		found, err := SearchPlayers(client, leagueKey, name)
		if err != nil {
			return nil, err
		}
		players = append(players, found...)
	}

	return players, nil
}

// GetPlayerStats searches the given league for a player with the provided player name.
// and returns their average stats for the current season. If the player is not
// found, an error is returned. name should contain at least 3 letters.
func GetPlayerStats(client *http.Client, leagueKey, name string, statsType int) (*schema.Player, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("name (%q) must contain at least 3 letters", name)
	}

	q := yfquery.League().Key(leagueKey).Players().Search(name).Stats()
	switch statsType {
	case StatsTypeUnknown:
		return nil, fmt.Errorf("unknown stats type requested")
	case StatsTypeSeason:
		q = q.CurrentSeason()
		break
	case StatsTypeAverageSeason:
		q = q.CurrentSeasonAverage()
		break
	case StatsTypeDate:
		q = q.Today()
		break
	case StatsTypeLastWeek:
		q = q.LastWeek()
		break
	case StatsTypeLastWeekAverage:
		q = q.LastWeekAverage()
		break
	case StatsTypeLastMonth:
		q = q.LastMonth()
		break
	case StatsTypeLastMonthAverage:
		q = q.LastMonthAverage()
		break
	}

	fc, err := q.Get(client)
	if err != nil {
		return nil, err
	}

	return findPlayer(&fc.League.Players, name)
}

// GetPlayersStats searches the given league for players with the provided player names.
// and returns their requested stats. If the player is not found, an error is
// returned. Each name should contain at least 3 letters.
func GetPlayersStats(client *http.Client, leagueKey string, names []string, statsType int) ([]*schema.Player, error) {
	players := []*schema.Player{}

	for _, name := range names {
		player, err := GetPlayerStats(client, leagueKey, name, statsType)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

// GetPlayerAdvancedStats searches the given league for a player with the provided
// player name and returns their advanced stats. If the player is not found, an
// error is returned. name should contain at least 3 letters.
func GetPlayerAdvancedStats(client *http.Client, leagueKey, name string) (*schema.Player, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("name (%q) must contain at least 3 letters", name)
	}
	fc, err := yfquery.League().Key(leagueKey).Players().Search(name).Stats().Get(client)
	if err != nil {
		return nil, err
	}

	return findPlayer(&fc.League.Players, name)
}

// GetPlayerOwnership searches the league for a player with the provided named and
// returns their ownership status.
func GetPlayerOwnership(client *http.Client, leagueKey, name string) (*schema.Player, error) {
	fc, err := yfquery.League().Key(leagueKey).Players().Search(name).Ownership().Get(client)
	if err != nil {
		return nil, err
	}

	return findPlayer(&fc.League.Players, name)
}

// SortFreeAgentsByStat searches the league for the top free agents by the provided stat.
func SortFreeAgentsByStat(client *http.Client, leagueKey string, statID, count int) ([]*schema.Player, error) {
	fc, err := yfquery.
		League().
		Key(leagueKey).
		Players().
		Status(yfquery.PlayerStatusFreeAgent).
		SortByStat(statID).
		Count(count).
		Get(client)
	if err != nil {
		return nil, err
	}

	var players []*schema.Player
	for _, p := range fc.League.Players.Player {
		players = append(players, &p)
	}

	return players, nil
}

// ComparePlayersNBA9CAT computes the diff in stats between two players in standard NBA 9 category leagues.
func ComparePlayersNBA9CAT(client *http.Client, leagueKey, playerA, playerB string, statsType int) (*StatsDiff, error) {
	players, err := GetPlayersStats(client, leagueKey, []string{playerA, playerB}, statsType)
	if err != nil {
		return nil, err
	}

	if len(players) != 2 {
		return nil, fmt.Errorf("encountered problem fetching stats for %q and %q", playerA, playerB)
	}

	diff := &StatsDiff{
		PlayerA: players[0].Name.Full,
		PlayerB: players[1].Name.Full,
		Diffs:   make(map[string]float64),
	}

	for i, stat := range players[0].PlayerStats.Stats.Stat {
		if !nba9CATIDs[stat.StatID] {
			continue
		}

		if stat.Value == "-" {
			return nil, fmt.Errorf("stats unavailable for %q", players[0].Name.Full)
		}

		statA, err := strconv.ParseFloat(stat.Value, 64)
		if err != nil {
			return nil, err
		}

		if players[1].PlayerStats.Stats.Stat[i].Value == "-" {
			return nil, fmt.Errorf("stats unavailable for %q", players[1].Name.Full)
		}

		statB, err := strconv.ParseFloat(players[1].PlayerStats.Stats.Stat[i].Value, 64)
		if err != nil {
			return nil, err
		}

		diff.Diffs[statIDToName[stat.StatID]] = statA - statB
	}
	return diff, nil
}
