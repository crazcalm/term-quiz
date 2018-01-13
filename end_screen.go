package quiz

import (
	"github.com/crazcalm/term-quiz/interface"
	"github.com/crazcalm/term-quiz/user"
	"github.com/jroimartin/gocui"
	"log"
	"fmt"
)

//ESInit -- End Screen Initialization. Presents the results
func ESInit(g *gocui.Gui, u user.Answers) (err error){
	score = fmt.Sprintf("%d/%d", u.TotalCorrect(), u.Total())

	//TODO: need a good way of getting the current userAnswer
	//and interating to the next userAnswer

	//Highlight the selected view and make it green
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//TODO: replace the score and question number...
	scoreWidget := gui.NewScore(gui.ScoreName, score, questionCount)
	explainWidget := gui.NewExplaination(gui.Explain, correct, u.Question.Question, u.Answer.Answer, u.Question.Explaination)

	g.SetManager(scoreWidget, explainWidget)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	//TODO: add a keybinding so that enter goes to the next result

	return nil
}