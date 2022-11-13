package yflib

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/famendola1/yfquery"
	"github.com/famendola1/yfquery/schema"
)

// GetTeam searches the given league for a team with the provided team name.
// If the team is not found an error is returned.
func GetTeam(client *http.Client, leagueKey, teamName string) (*schema.Team, error) {
	fc, err := yfquery.League().Key(leagueKey).Teams().Get(client)
	if err != nil {
		return nil, err
	}

	for _, tm := range fc.League.Teams.Team {
		if strings.ToLower(tm.Name) == strings.ToLower(teamName) {
			return &tm, nil
		}
	}

	return nil, fmt.Errorf("team %q not found", teamName)
}

// GetTeamRoster searches the given league for a team with the provided team name
// and return's its roster. If the team is not found an error is returned.
func GetTeamRoster(client *http.Client, leagueKey, teamName string) (*schema.Team, error) {
	fc, err := yfquery.League().Key(leagueKey).Teams().Roster().Get(client)
	if err != nil {
		return nil, err
	}

	for _, tm := range fc.League.Teams.Team {
		if strings.ToLower(tm.Name) == strings.ToLower(teamName) {
			return &tm, nil
		}
	}

	return nil, fmt.Errorf("team %q not found", teamName)
}

// GetTeamStats searches the given league for a team with the provided team name
// and return's its stats. If the team is not found an error is returned.
func GetTeamStats(client *http.Client, leagueKey, teamName string, statsType int) (*schema.Team, error) {
	q := yfquery.League().Key(leagueKey).Teams().Stats()
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

	for _, tm := range fc.League.Teams.Team {
		if strings.ToLower(tm.Name) == strings.ToLower(teamName) {
			return &tm, nil
		}
	}

	return nil, fmt.Errorf("team %q not found", teamName)
}
