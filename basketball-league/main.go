package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// Team - Keeps a record of the team details including the names of the players
type Team struct {
	PlayerNames []string
	TeamName    string
}

// League - Maps the number of wins to the Team
type League struct {
	Teams []Team
	Wins  map[string]int
}

// MatchResult - Takes the name of the first team (ft), it's score (fts),
// and the name of the second team (st) and it's score (stc).
func (l *League) MatchResult(ft string, fts int, st string, stc int) {
	if fts > stc {
		l.Wins[ft]++
		return
	} else if stc > fts {
		l.Wins[st]++
		return
	}
}

// Ranking - Returns a slice of the teams in order of wins
func (l League) Ranking() []string {
	teams := make([]string, 0, len(l.Wins))

	for team := range l.Wins {
		teams = append(teams, team)
	}

	fmt.Println("Before sort: ", teams)

	sort.Slice(teams, func(i, j int) bool {
		return l.Wins[teams[i]] > l.Wins[teams[j]]
	})

	fmt.Println("After sort: ", teams)

	return teams
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) error {
	output := strings.Join(r.Ranking(), "\n")
	_, err := io.WriteString(w, output)

	return err
}

func main() {
	league := League{
		Wins: make(map[string]int),
	}

	league.MatchResult("Lakers", 100, "Bulls", 90)
	league.MatchResult("Bulls", 199, "Chicago Fire", 90)
	league.MatchResult("Bulls", 110, "Heat", 105)
	league.MatchResult("Lakers", 95, "Heat", 100)

	err := RankPrinter(league, os.Stdout)

	if err != nil {
		log.Fatal("error has occured: ", err)
	}

}
