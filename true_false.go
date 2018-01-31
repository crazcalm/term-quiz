package quiz

import (
	"github.com/crazcalm/term-quiz/interface"
	"github.com/crazcalm/term-quiz/questions"
	"github.com/jroimartin/gocui"
	"log"
)

//TFInit -- Initializes the ABCD gui interface
func TFInit(g *gocui.Gui, q *questions.Question, count string) (err error) {
	//The Answers
	as := q.Answers

	//Highlight the selected view and make it green
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//Add content to gui
	questionFrame := gui.NewQuestionFrame("questionFrame")
	question := gui.NewQuestion("question", count, q.Question)
	answerTrue := gui.NewAnswer(gui.BoxTrue, as.Answers[0].Answer, "")
	AnswersToBoxViews[gui.BoxTrue] = as.Answers[0]
	answerFalse := gui.NewAnswer(gui.BoxFalse, as.Answers[1].Answer, "")
	AnswersToBoxViews[gui.BoxFalse] = as.Answers[1]
	infoBar := gui.NewInfoBar(gui.InfoBarName, gui.InfoBarTrueFalse)

	g.SetManager(questionFrame, question, answerTrue, answerFalse, infoBar)

	//Add keybindings
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, gui.TFNextView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, SelectAnswer); err != nil {
		log.Panicln(err)
	}

	return nil
}
