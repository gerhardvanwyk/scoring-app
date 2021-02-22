package app

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Teams interface {
	add(name string, score int) int
	Score(lines []string) error
	Print()
}

type teams struct {
	teams []team
}

type team struct {
	name  string
	score int
}

func NewTeams() Teams {
	return &teams{
		teams: make([]team, 0),
	}
}

//Adds the team and score
func (ts *teams) add(name string, score int) int {

	//found
	for i, v := range ts.teams {
		if name == v.name {
			ts.teams[i].score = v.score + score
			return v.score
		}
	}

	//Not found
	t := team{
		name:  name,
		score: 0,
	}

	ts.teams = append(ts.teams, t)

	for i, v := range ts.teams {
		if name == v.name {
			ts.teams[i].score = score
		}
	}
	return score
}

//Takes in the lines in the file as a slice
func (ts *teams) Score(lines []string) error {

	for _, line := range lines {

		//Last line in the file
		if line == "" {
			return nil
		}

		input := strings.Split(line, ",")

		in1 := input[0]
		in2 := input[1]

		err, n1, t1  := ts.teamPoints(in1)
		if err != nil {
			return err
		}

		err, n2, t2  := ts.teamPoints(in2)
		if err != nil {
			return err
		}
		var score1 int
		var score2 int
		if t1 < t2{
			score2 = 3
		}else if t1 == t2{
			score1 = 1
			score2 = 1
		}else{
			score1 = 3
		}

		ts.add(n1, score1)
		ts.add(n2, score2)

	}
	return nil
}

//
func (ts *teams) teamPoints(in string) (error, string, int) {
	//Removing leading whitespace
	if strings.HasPrefix(in, " ") {
		in = strings.TrimPrefix(in, " ")
	}

	name := strings.TrimRightFunc(in, func(r rune) bool {
		if unicode.IsDigit(r) || unicode.IsSpace(r) {
			return true
		}
		return false
	})

	s := strings.TrimLeftFunc(in, func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return true
		}
		return false
	})

	score, err := strconv.Atoi(s)
	if err != nil {
		return err, "", 0
	}

	return nil, name, score
}

func (ts *teams) Print() {

	//Sorts the points
	sort.Slice(ts.teams, func(i, j int) bool {
		return ts.teams[i].score > ts.teams[j].score
	})

	//Sorts the names where pts == pts
	sort.Slice(ts.teams, func(i, j int) bool {
		if ts.teams[i].score == ts.teams[j].score{
			return ts.teams[i].name < ts.teams[j].name
		}
		return false
	})

	for i, v := range ts.teams {
		i = i + 1
		fmt.Printf("%d. %s pts %d \n", i, v.name, v.score )
	}
}
