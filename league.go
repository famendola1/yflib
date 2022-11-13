package yflib

import (
	"fmt"
	"net/http"

	"github.com/famendola1/yfquery"
	"github.com/famendola1/yfquery/schema"
)

// MakeLeagueKey creates a league key from the gameKey and leagueID.
func MakeLeagueKey(gameKey string, leagueID int) string {
	return fmt.Sprintf("%s.l.%d", gameKey, leagueID)
}

// GetLeague queries the Yahoo Fantasy API for a League.
func GetLeague(client *http.Client, leagueKey string) (*schema.League, error) {
	fc, err := yfquery.League().Key(leagueKey).Standings().Get(client)
	if err != nil {
		return nil, err
	}
	return &fc.League, nil
}

// GetLeagueStandings queries the Yahoo Fantasy API for a leagues Standings.
func GetLeagueStandings(client *http.Client, leagueKey string) (*schema.Standings, error) {
	fc, err := yfquery.League().Key(leagueKey).Standings().Get(client)
	if err != nil {
		return nil, err
	}

	return &fc.League.Standings, nil
}

// GetCurrentScoreboard queries the Yahoo Fantasy API for a league's current scoreboard.
func GetCurrentScoreboard(client *http.Client, leagueKey string) (*schema.Scoreboard, error) {
	fc, err := yfquery.League().Key(leagueKey).CurrentScoreboard().Get(client)
	if err != nil {
		return nil, err
	}

	return &fc.League.Scoreboard, nil
}

// GetScoreboard queries the Yahoo Fantasy API for the scoreboard of a given week.
func GetScoreboard(client *http.Client, leagueKey string, week int) (*schema.Scoreboard, error) {
	fc, err := yfquery.League().Key(leagueKey).Scoreboard(week).Get(client)
	if err != nil {
		return nil, err
	}

	return &fc.League.Scoreboard, nil
}

// GetLeagueRosters queries the Yahoo Fantasy API for all the team rosters in a league.
func GetLeagueRosters(client *http.Client, leagueKey string) (*schema.Teams, error) {
	fc, err := yfquery.League().Key(leagueKey).Teams().Roster().Get(client)
	if err != nil {
		return nil, err
	}

	return &fc.Teams, nil
}
