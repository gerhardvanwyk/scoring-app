package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"scoring-app/app"
	"strings"
)

var filename string

func main() {

	const (
		short       = " (shorthand)"
		defaultFile = "input.txt"
		usageFile   = "path to the score file"
	)

	flag.StringVar(&filename, "file", defaultFile, usageFile)
	flag.StringVar(&filename, "f", defaultFile, usageFile+short)

	flag.Parse()

	bCont, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Could not read file at %s ", filename)
		os.Exit(1)
	}
	fileCont := string(bCont)

	lines := strings.Split(fileCont, "\n")
	//Get init with the amount of lines (Will be less teams)
	teams := app.NewTeams()
	err = teams.Score(lines)
	if err != nil{
		fmt.Errorf("Error occured while adding scores, panic %s ", filename)
		os.Exit(1)
	}
	teams.Print()
}
