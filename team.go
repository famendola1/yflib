package yflib

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/famendola1/yfquery"
	"github.com/famendola1/yfquery/schema"
)

func findTeam(teams *schema.Teams, teamName string) (*schema.Team, error) {
	for _, tm := range teams.Team {
		if strings.ToLower(tm.Name) == strings.ToLower(teamName) {
			return &tm, nil
		}
	}

	return nil, fmt.Errorf("team %q not found", teamName)
}

// GetTeamKey returns the key of the team, if the team is found.
func GetTeamKey(client *http.Client, leagueKey, teamName string) (string, error) {
	fc, err := yfquery.League().Key(leagueKey).Teams().Get(client)
	if err != nil {
		return "", err
	}

	tm, err := findTeam(&fc.League.Teams, teamName)
	if err != nil {
		return "", err
	}

	return tm.TeamKey, nil
}

// GetTeam searches the given league for a team with the provided team name.
// If the team is not found an error is returned.
func GetTeam(client *http.Client, leagueKey, teamName string) (*schema.Team, error) {
	fc, err := yfquery.League().Key(leagueKey).Teams().Get(client)
	if err != nil {
		return nil, err
	}

	return findTeam(&fc.League.Teams, teamName)
}

// GetTeamRoster searches the given league for a team with the provided team name
// and return's its roster. If the team is not found an error is returned.
func GetTeamRoster(client *http.Client, leagueKey, teamName string) (*schema.Team, error) {
	fc, err := yfquery.League().Key(leagueKey).Teams().Roster().Get(client)
	if err != nil {
		return nil, err
	}

	return findTeam(&fc.League.Teams, teamName)
}

// GetTeamStats searches the given league for a team with the provided team name
// and return's its stats. If the team is not found an error is returned.
func GetTeamStats(client *http.Client, leagueKey, teamName string, statsType int) (*schema.Team, error) {
	q, err := addStatsTypeToQuery(yfquery.League().Key(leagueKey).Teams().Stats(), statsType)
	if err != nil {
		return nil, err
	}

	fc, err := q.Get(client)
	if err != nil {
		return nil, err
	}

	return findTeam(&fc.League.Teams, teamName)
}
