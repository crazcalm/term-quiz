package main

import (
	"github.com/crazcalm/term-quiz"
	"github.com/crazcalm/term-quiz/questions"
	"github.com/jroimartin/gocui"
	"log"
	"path/filepath"
)

var (
	q = questions.NewQuestions()
)

func main() {
	//Get gui driver
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Need to create questions
	q, err := questions.CreateQuestions(q, filepath.Join("test_data", "fill_in_the_blank.csv"))
	if err != nil {
		log.Fatal(err)
	}

	//Need to initialize screen
	quiz.FBInit(g, q.Questions[0], "1")

	//Run main loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
