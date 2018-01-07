package quiz

import (
	"fmt"
	"github.com/crazcalm/term-quiz/answers"
	"github.com/crazcalm/term-quiz/user"
	"github.com/jroimartin/gocui"
	"strconv"
)

var (
	//AnswersToBoxViews -- used to map the user selected box view to the actual answer
	AnswersToBoxViews = map[string]*answers.Answer{}
)

//SelectAnswer -- Callback used to select and answer in the ABCDLayout
func SelectAnswer(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintln(v, "Selected")

	cQuestion, err := Questions.Current()
	if err != nil {
		return err
	}

	selectedAnswer := AnswersToBoxViews[v.Name()]

	a := user.Answer{
		cQuestion,
		selectedAnswer,
	}

	//User answers
	UserAnswers[strconv.Itoa(len(UserAnswers))] = &a

	//After all that work, I still have to manually increment
	//the questions...
	Questions.Index++

	//Next Screen? I am not sure if this is going to work...
	err = Init(g)
	if err != nil {
		return err
	}

	return nil
}
