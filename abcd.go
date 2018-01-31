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
	//QuestionLimit -- Limits the number of questions that are shown to the user
	QuestionLimit = 10
	//CurrentUserAnswer -- index for user answers
	CurrentUserAnswer = 1
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
	AnswersToBoxViews[gui.BoxA] = as.Answers[0]
	answerB := gui.NewAnswer(gui.BoxB, gui.BoxB, as.Answers[1].Answer)
	AnswersToBoxViews[gui.BoxB] = as.Answers[1]
	answerC := gui.NewAnswer(gui.BoxC, gui.BoxC, as.Answers[2].Answer)
	AnswersToBoxViews[gui.BoxC] = as.Answers[2]
	answerD := gui.NewAnswer(gui.BoxD, gui.BoxD, as.Answers[3].Answer)
	AnswersToBoxViews[gui.BoxD] = as.Answers[3]
	infoBar := gui.NewInfoBar(gui.InfoBarName, gui.InfoBarABCD)

	g.SetManager(questionFrame, question, answerA, answerB, answerC, answerD, infoBar)

	//Add keybindings
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, gui.ABCDNextView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, SelectAnswer); err != nil {
		log.Panicln(err)
	}
	return nil
}
