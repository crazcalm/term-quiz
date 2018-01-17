package main

import (
	"github.com/crazcalm/term-quiz"
	"github.com/crazcalm/term-quiz/questions"
	"github.com/jroimartin/gocui"
	"log"
	"path/filepath"
)

func main() {
	//Get gui driver
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Csv file paths
	//abcd.csv  error.csv  fill_in_the_blank.csv  true_false.csv
	path1 := filepath.Join("test_data", "abcd.csv")
	path2 := filepath.Join("test_data", "true_false.csv")
	path3 := filepath.Join("test_data", "fill_in_the_blank.csv")

	//Need to create questions
	quiz.Questions, err = questions.CreateQuestions(quiz.Questions, path1, path2, path3, path1)
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
