package quiz

import (
	"fmt"
	"github.com/crazcalm/term-quiz/interface"
	"github.com/crazcalm/term-quiz/user"
	"github.com/jroimartin/gocui"
	"log"
	"os"
)

//ESInit -- End Screen Initialization. Presents the results
func ESInit(g *gocui.Gui, u user.Answers) (err error) {
	//End quiz when you run out of answers
	if CurrentUserAnswer > len(UserAnswers) {
		g.Close()
		log.Println("Game Over")
		os.Exit(0)
	}

	//Place holder for 'correct' string that gets placed
	//in the gui
	var correct string

	numCorrect, err := u.TotalCorrect()
	if err != nil {
		log.Panicln(err)
	}
	//Score = num of correct answers over total
	score := fmt.Sprintf("%d/%d", numCorrect, u.Total())

	//Currently need a global counter...
	questionCount := fmt.Sprintf("%d", CurrentUserAnswer)

	currentUserAnswer := u[questionCount]
	answerCorrect, err := currentUserAnswer.Correct()
	if err != nil {
		log.Panicln(err)
	}
	if answerCorrect {
		correct = gui.Right
	} else {
		correct = gui.Wrong
	}

	//Highlight the selected view and make it green
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//create widgets
	scoreWidget := gui.NewScore(gui.ScoreName, score, questionCount)
	explainWidget := gui.NewExplaination(gui.Explain, correct, currentUserAnswer.Question.Question, currentUserAnswer.Answer.Answer, currentUserAnswer.Question.Explaination)

	g.SetManager(scoreWidget, explainWidget)

	//Keybindings
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, NextUserAnswer); err != nil {
		log.Panicln(err)
	}

	return nil
}
