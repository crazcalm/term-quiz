package quiz

import (
	"log"
	"github.com/jroimartin/gocui"
	"github.com/crazcalm/term-quiz/interface"
	"github.com/crazcalm/term-quiz/questions"
	"github.com/crazcalm/term-quiz/answers"
)

//TFInit -- Initializes the ABCD gui interface
func TFInit(g *gocui.Gui, q *questions.Question, as answers.Answers, count string) (err error){
	//Highlight the selected view and make it green
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//Add content to gui
	questionFrame := gui.NewQuestionFrame("questionFrame", count)
	question := gui.NewQuestion("question", "question", q.Question)	
	answerTrue := gui.NewAnswer(gui.BoxTrue, gui.BoxTrue, as.Answers[0].Answer)
	answerFalse := gui.NewAnswer(gui.BoxTrue, gui.BoxTrue, as.Answers[1].Answer)

	g.SetManager(questionFrame, question, answerTrue, answerFalse)

	//Add keybindings
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, gui.TFNextView); err != nil {
		log.Panicln(err)
	}

	return nil
}