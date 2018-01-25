package quiz

import (
	"github.com/jroimartin/gocui"
	"strconv"
)

//Init -- The Init function decides which sub Init function should be called
//ABCDInint, TFInit or FBInit
func Init(g *gocui.Gui) (err error) {
	//Have we reached the question limit?
	if Questions.Index >= QuestionLimit {
		//Need to call End Screen
		ESInit(g, UserAnswers)
		return nil
	}

	//Have we ran out of questions?
	if Questions.Index >= len(Questions.Questions) {
		//Need to call End Screen
		ESInit(g, UserAnswers)
		return nil
	}

	//Get current question
	q, err := Questions.Current()
	if err != nil {
		return err
	}

	//Use number of answers to figure out which Init function to use
	numOfAnswers := len(q.Answers.Answers)

	count := "Question " + strconv.Itoa(len(UserAnswers)+1)

	if numOfAnswers == 4 {
		err = ABCDInit(g, q, count)
		if err != nil {
			return err
		}

	} else if numOfAnswers == 2 {
		err = TFInit(g, q, count)
		if err != nil {
			return err
		}
	} else if numOfAnswers == 1 {
		err = FBInit(g, q, count)
		if err != nil {
			return err
		}
	}

	return nil
}
