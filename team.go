package yflib

import (
	"fmt"
	"net/http"
	"strconv"
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

// GetTeamMatchups searches the given league for a team with the provided team name
// and returns all of their matchups. If the team is not found an error is returned.
func GetTeamMatchups(client *http.Client, leagueKey, teamName string) (*schema.Team, error) {
	fc, err := yfquery.League().Key(leagueKey).Teams().AllMatchups().Get(client)
	if err != nil {
		return nil, err
	}

	return findTeam(&fc.League.Teams, teamName)
}

// CategoryMatchupResult represents the results from a matchup in a category
// from the perspective of the HomeTeam.
type CategoryMatchupResult struct {
	HomeTeam       string
	AwayTeam       string
	CategoriesWon  []int
	CategoriesLost []int
	CategoriesTied []int
}

// CalculateCategoryMathchupResultsVsLeague computes the results of the provided
// team against every other team in the league for the given week based on the
// provided stats to use.
func CalculateCategoryMathchupResultsVsLeague(client *http.Client, leagueKey, teamName string, eligibleStats StatIDSet, week int) ([]CategoryMatchupResult, error) {
	fc, err := yfquery.League().Key(leagueKey).Teams().Stats().Week(week).Get(client)
	if err != nil {
		return nil, err
	}

	return computeCategoryMatchupResultsVsLeague(teamName, &fc.League.Teams, eligibleStats)
}

func computeCategoryMatchupResultsVsLeague(teamName string, teams *schema.Teams, eligibleStats StatIDSet) ([]CategoryMatchupResult, error) {
	_, err := findTeam(teams, teamName)
	if err != nil {
		return nil, err
	}

	teamStats := make(map[string]map[int]float64)
	for _, tm := range teams.Team {
		teamStats[tm.Name] = make(map[int]float64)
		for _, stat := range tm.TeamStats.Stats.Stat {
			if !eligibleStats[stat.StatID] {
				continue
			}
			val, _ := strconv.ParseFloat(stat.Value, 64)
			teamStats[tm.Name][stat.StatID] = val
		}
	}

	results := []CategoryMatchupResult{}
	for _, tm := range teams.Team {
		if tm.Name == teamName {
			continue
		}

		res := CategoryMatchupResult{HomeTeam: teamName, AwayTeam: tm.Name}
		for stat := range eligibleStats {
			if inverseStats[stat] {
				if teamStats[teamName][stat] < teamStats[tm.Name][stat] {
					res.CategoriesWon = append(res.CategoriesWon, stat)
					continue
				}

				if teamStats[teamName][stat] > teamStats[tm.Name][stat] {
					res.CategoriesLost = append(res.CategoriesLost, stat)
					continue
				}

				res.CategoriesTied = append(res.CategoriesTied, stat)
				continue
			}

			if teamStats[teamName][stat] > teamStats[tm.Name][stat] {
				res.CategoriesWon = append(res.CategoriesWon, stat)
				continue
			}

			if teamStats[teamName][stat] < teamStats[tm.Name][stat] {
				res.CategoriesLost = append(res.CategoriesLost, stat)
				continue
			}

			res.CategoriesTied = append(res.CategoriesTied, stat)
			continue
		}
		results = append(results, res)
	}

	return results, nil
}
