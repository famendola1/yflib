package yflib

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
