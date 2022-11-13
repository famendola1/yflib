package yflib

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/famendola1/yfquery"
	"github.com/famendola1/yfquery/schema"
)

// Enum of types when requesting for stats.
const (
	StatsTypeUnknown = iota
	StatsTypeSeason
	StatsTypeAverageSeason
	StatsTypeDate
	StatsTypeLastWeek
	StatsTypeLastWeekAverage
	StatsTypeLastMonth
	StatsTypeLastMonthAverage
)

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

	for _, p := range fc.League.Players.Player {
		if strings.ToLower(p.Name.Full) == strings.ToLower(name) {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("player %q not found", name)
}

// SearchPlayers searches the given league for a players with the provided player
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

	for _, p := range fc.League.Players.Player {
		if strings.ToLower(p.Name.Full) == strings.ToLower(name) {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("player %q not found", name)
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

	for _, p := range fc.League.Players.Player {
		if strings.ToLower(p.Name.Full) == strings.ToLower(name) {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("player %q not found", name)
}

// GetPlayerOwnership searches the league for a player with the provided named and
// returns their ownership status.
func GetPlayerOwnership(client *http.Client, leagueKey, name string) (*schema.Player, error) {
	fc, err := yfquery.League().Key(leagueKey).Players().Search(name).Ownership().Get(client)
	if err != nil {
		return nil, err
	}

	for _, p := range fc.League.Players.Player {
		if strings.ToLower(p.Name.Full) == strings.ToLower(name) {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("player %q not found", name)
}
