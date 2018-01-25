package main

import (
	"github.com/crazcalm/term-quiz"
	"github.com/crazcalm/term-quiz/questions"
	"github.com/jroimartin/gocui"
	"log"
	"os"
)

/*
Run Examples:

- go run terminal_args.go test_data/abcd.csv
- go run terminal_args.go test_data/abcd.csv test_data/true_false.csv
- go run terminal_args.go test_data/abcd.csv test_data/true_false.csv test_data/fill_in_the_blank.csv

Run Exmaples with errors:
- go run terminal_args.go
- go run terminal_args.go test_data
- go run terminal_args.go test_data/DoesNotExist
*/

func main() {
	//Check passed in files
	err := quiz.FileArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	//Get gui driver
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Need to create questions
	quiz.Questions, err = questions.CreateQuestions(quiz.Questions, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	//Shuffle Questions
	err = quiz.Questions.Shuffle()
	if err != nil {
		log.Fatal(err)
	}

	//Need to initialize screen
	err = quiz.Init(g)
	if err != nil {
		log.Fatal(err)
	}

	//Run main loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
