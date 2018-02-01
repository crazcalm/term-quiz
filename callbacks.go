package quiz

import (
	"github.com/crazcalm/term-quiz/answers"
	"github.com/crazcalm/term-quiz/interface"
	"github.com/crazcalm/term-quiz/user"
	"github.com/jroimartin/gocui"
	"strconv"
)

var (
	//AnswersToBoxViews -- used to map the user selected box view to the actual answer
	AnswersToBoxViews = map[string]*answers.Answer{}
)

//FillInAnswer -- Callback used for the fill in the blank answers
func FillInAnswer(g *gocui.Gui, v *gocui.View) error {
	cQuestion, err := Questions.Current()
	if err != nil {
		return err
	}

	filledInAnswer := &answers.Answer{v.Buffer(), true}

	a := user.Answer{
		cQuestion,
		filledInAnswer,
	}

	//User answers -- The plus one is so the count starts at 1
	UserAnswers[strconv.Itoa(len(UserAnswers)+1)] = &a

	//Increment the questions index!
	Questions.Index++

	//Next Screen
	err = Init(g)
	if err != nil {
		return err
	}

	return nil

}

//SelectAnswer -- Callback used to select an answer in quiz layouts that have 
//multiple answers to select from
func SelectAnswer(g *gocui.Gui, v *gocui.View) error {
	//Reset variable used for tabbing through solutions
	gui.ActiveView = 0

	cQuestion, err := Questions.Current()
	if err != nil {
		return err
	}

	selectedAnswer := AnswersToBoxViews[v.Name()]

	a := user.Answer{
		cQuestion,
		selectedAnswer,
	}

	//User answers -- The plus one is so the count starts at 1
	UserAnswers[strconv.Itoa(len(UserAnswers)+1)] = &a

	//Increment the questions index!
	Questions.Index++

	//Next Screen
	err = Init(g)
	if err != nil {
		return err
	}

	return nil
}

//NextUserAnswer -- View next user answer
func NextUserAnswer(g *gocui.Gui, v *gocui.View) error {
	//Increment count
	CurrentUserAnswer++

	//Next Screen
	err := Init(g)
	if err != nil {
		return err
	}
	return nil
}
