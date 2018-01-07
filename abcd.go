package quiz

import (
	"github.com/crazcalm/term-quiz/interface"
	"github.com/crazcalm/term-quiz/questions"
	"github.com/crazcalm/term-quiz/user"
	"github.com/jroimartin/gocui"
	"log"
)

var (
	//Questions -- Global container for questions
	Questions = questions.NewQuestions()
	//UserAnswers -- Global container for UserAnswers
	UserAnswers = user.Answers{}
)

//ABCDInit -- Initializes the ABCD gui interface
func ABCDInit(g *gocui.Gui, q *questions.Question, count string) (err error) {
	//Highlight the selected view and make it green
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	as := q.Answers

	//Add content to gui
	questionFrame := gui.NewQuestionFrame("questionFrame")
	question := gui.NewQuestion("question", count, q.Question)
	answerA := gui.NewAnswer(gui.BoxA, gui.BoxA, as.Answers[0].Answer)
	answerB := gui.NewAnswer(gui.BoxB, gui.BoxB, as.Answers[1].Answer)
	answerC := gui.NewAnswer(gui.BoxC, gui.BoxC, as.Answers[2].Answer)
	answerD := gui.NewAnswer(gui.BoxD, gui.BoxD, as.Answers[3].Answer)

	g.SetManager(questionFrame, question, answerA, answerB, answerC, answerD)

	//Add keybindings
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, gui.ABCDNextView); err != nil {
		log.Panicln(err)
	}
	return nil
}
