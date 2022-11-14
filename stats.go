package yflib

import (
	"fmt"

	"github.com/famendola1/yfquery"
)

// Enum of types when requesting for stats.
const (
	StatsTypeUnknown = iota
	StatsTypeSeason
	StatsTypeAverageSeason
	StatsTypeDate
	StatsTypeWeek
	StatsTypeLastWeek
	StatsTypeAverageLastWeek
	StatsTypeLastMonth
	StatsTypeAverageLastMonth
)

var (
	// StatIDToName is a map of stat ids to their name.
	StatIDToName = map[int]string{
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

	// StatNameToID is a map of stat names to their ID.
	StatNameToID = map[string]int{
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

func addStatsTypeToQuery(q *yfquery.StatsQuery, statsType int) (*yfquery.StatsQuery, error) {
	switch statsType {
	case StatsTypeUnknown:
		return nil, fmt.Errorf("unknown stats type requested")
	case StatsTypeSeason:
		return q.CurrentSeason(), nil
	case StatsTypeAverageSeason:
		return q.CurrentSeasonAverage(), nil
	case StatsTypeDate:
		return q.Today(), nil
	case StatsTypeLastWeek:
		return q.LastWeek(), nil
	case StatsTypeAverageLastWeek:
		return q.LastWeekAverage(), nil
	case StatsTypeLastMonth:
		return q.LastMonth(), nil
	case StatsTypeAverageLastMonth:
		return q.LastMonthAverage(), nil
	default:
		return nil, fmt.Errorf("unknown stats type requested")
	}
}

func convertStatsTypeToSortType(statsType int) yfquery.PlayerSortType {
	switch statsType {
	case StatsTypeSeason:
		return yfquery.PlayerSortTypeSeason
	case StatsTypeAverageSeason:
		return yfquery.PlayerSortTypeAverageSeason
	case StatsTypeDate:
		return yfquery.PlayerSortTypeDate
	case StatsTypeWeek:
		return yfquery.PlayerSortTypeWeek
	case StatsTypeLastWeek:
		return yfquery.PlayerSortTypeLastWeek
	case StatsTypeAverageLastWeek:
		return yfquery.PlayerSortTypeAverageLastWeek
	case StatsTypeLastMonth:
		return yfquery.PlayerSortTypeLastMonth
	case StatsTypeAverageLastMonth:
		return yfquery.PlayerSortTypeAverageLastMonth
	default:
		return yfquery.PlayerSortTypeUnknown
	}
}
