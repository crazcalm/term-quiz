package quiz

import (
	"github.com/jroimartin/gocui"
	"strconv"
)

//Init -- The Init function decides which sub Init funtion should be called
//ABCDInint, TFInit or FBInit
func Init(g *gocui.Gui) (err error) {
	//Need to check if we have reached the question limit
	//or if we have run out of question.
	//TODO: add end screen

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
