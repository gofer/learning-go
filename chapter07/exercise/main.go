package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1Name string, team1Score int, team2Name string, team2Score int) {
	if team1Score > team2Score {
		l.Wins[team1Name]++
	} else if team2Score > team1Score {
		l.Wins[team2Name]++
	}
}

func (l League) Ranking() []string {
	var result = make([]string, len(l.Teams))
	for i, team := range l.Teams {
		result[i] = team.Name
	}
	sort.Slice(result, func(i, j int) bool {
		if _, ok := l.Wins[result[i]]; !ok {
			return false
		}
		if _, ok := l.Wins[result[j]]; !ok {
			return false
		}
		return l.Wins[result[i]] > l.Wins[result[j]]
	})
	return result
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, writer io.Writer) {
	ranking := r.Ranking()
	for _, team := range ranking {
		_, err := io.WriteString(writer, team+"\n")
		if err != nil {
			fmt.Printf("Error writing to writer: %v\n", err)
		}
	}
}

func main() {
	l := League{
		Teams: []Team{
			Team{
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			Team{
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			Team{
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			Team{
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)

	results := l.Ranking()
	fmt.Println(results)

	RankPrinter(l, os.Stdout)
}
