package yflib

import (
	"fmt"
	"net/http"
	"strconv"
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

var (
	statIDToName = map[int]string{
		5:       "FG%",
		8:       "FT%",
		10:      "3PM",
		12:      "PTS",
		15:      "REB",
		16:      "AST",
		17:      "STL",
		18:      "BLK",
		19:      "TOV",
		9004003: "FG",
		9007006: "FT",
	}

	statNameToID = map[string]int{
		"FG%": 5,
		"FT%": 8,
		"3PM": 10,
		"PTS": 12,
		"REB": 15,
		"AST": 16,
		"STL": 17,
		"BLK": 18,
		"TOV": 19,
		"FG":  9004003,
		"FT":  9007006,
	}

	nba9CATIDs = map[int]bool{
		5:  true,
		8:  true,
		10: true,
		12: true,
		15: true,
		16: true,
		17: true,
		18: true,
		19: true,
	}
)

// StatsDiff represents the difference of stats between PlayerA and PlayerB.
type StatsDiff struct {
	PlayerA string
	PlayerB string
	Diffs   map[string]float64
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
