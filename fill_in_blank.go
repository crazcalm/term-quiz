package quiz

import (
	"github.com/crazcalm/term-quiz/interface"
	"github.com/crazcalm/term-quiz/questions"
	"github.com/jroimartin/gocui"
	"log"
)

//FBInit -- Initializes the ABCD gui interface
func FBInit(g *gocui.Gui, q *questions.Question, count string) (err error) {
	//The Answers
	as := q.Answers

	//Highlight the selected view and make it green
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//Add content to gui
	questionFrame := gui.NewQuestionFrame("questionFrame")
	question := gui.NewQuestion("question", count, q.Question)
	answerBlank := gui.NewAnswer(gui.BoxBlank, gui.BoxBlank, as.Answers[0].Answer)
	infoBar := gui.NewInfoBar(gui.InfoBarName, gui.InfoBarFillInBlank)

	g.SetManager(questionFrame, question, answerBlank, infoBar)

	//Add keybindings
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, FillInAnswer); err != nil {
		log.Panicln(err)
	}

	return nil
}
