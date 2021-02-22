package app

import "testing"

func Test_parsing(t *testing.T) {

	teams := NewTeams()

	lines := []string{
		"Lions 3, Snakes 3455",
		"Tarantulas 1, FC Awesome 0",
		"Lions 1, FC Awesome 1",
		"Tarantulas 3, Snakes 1",
		"Lions 4, Grouches 0",
	}

	err := teams.Score(lines)

	if err != nil {
		t.Fatal(err)
	}

	teams.Print()
}

func Test_Game(t *testing.T){

	teams := NewTeams()

	lines := []string{
		"Lions 3, Snakes 7",
	}

	err := teams.Score(lines)

	if err != nil {
		t.Fatal(err)
	}
}
