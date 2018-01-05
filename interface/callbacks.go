package gui

import (
	"github.com/jroimartin/gocui"
	"log"
)

var (
	activeView = 0
	//ABCDBoxes -- slice of the A, B, C, and D answer box names
	ABCDBoxes = []string{BoxA, BoxB, BoxC, BoxD}
	//TFBoxes -- slice of the True and False answer box names
	TFBoxes = []string{BoxTrue, BoxFalse}
)

//Quit -- Callback used to quit application
func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

//setCurrentViewOnTop -- Sets the passed in view on top
func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	_, err := g.SetCurrentView(name)
	if err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)

}

//ABCDNextView -- Callback used to interate through the A, B, C, D choices
func ABCDNextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (activeView + 1) % len(ABCDBoxes)
	name := ABCDBoxes[nextIndex]

	_, err := setCurrentViewOnTop(g, name)
	if err != nil {
		log.Panicln(err)
	}

	activeView = nextIndex
	return nil
}

//TFNextView -- Callback used to interate through the True and False choices
func TFNextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (activeView + 1) % len(TFBoxes)
	name := TFBoxes[nextIndex]

	_, err := setCurrentViewOnTop(g, name)
	if err != nil {
		log.Panicln(err)
	}

	activeView = nextIndex
	return nil
}

/*
//CursorDown -- Callback used to scroll down
func CursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

//CursorUp -- Callback used to scoll up
func CursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

//SelectAnswer -- Callback used to select and answer in the ABCDLayout
func SelectAnswer(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintln(v, "Selected")

	cQuestion := currentQuestion()
	selectedAnswer := answersToBoxViews[v.Name()]

	a := UserAnswer{
		v.Name(),
		&cQuestion,
		&selectedAnswer,
	}

	//User answers
	userAnswers = append(userAnswers, a)

	if !nextQuestionExist() || len(userAnswers) >= QuestionsLimit {
		g.SetManagerFunc(endScreenLayout)
		err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit)
		if err != nil {
			log.Panicln(err)
			return err
		}
		err = g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, CursorDown)
		if err != nil {
			log.Panicln(err)
			return err
		}

		err = g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, CursorUp)
		if err != nil {
			log.Panicln(err)
			return err
		}

		return nil
	}
	question, err := nextQuestion()
	if err != nil {
		log.Fatal(err)
	}

	//Write Question and Answers to layout
	writeInfoToLayout(g, question)

	return nil
}

func writeInfoToLayout(g *gocui.Gui, q Question) {
	//Write question
	questionBox := getQuestionBoxView(g)
	questionBox.Clear()
	questionBox.Title = fmt.Sprintf("Question %d", QuestionCount+1)
	fmt.Fprintln(questionBox, q.Question)

	//Write answers
	answerBoxViews := getAnswerBoxViews(g)
	for i, answer := range q.Answers {
		answerBoxViews[i].Clear()

		//Adding it to the map
		answersToBoxViews[answerBoxViews[i].Name()] = answer

		//Write the answer to the layout
		fmt.Fprintln(answerBoxViews[i], answer.Answer)
	}

}

func getQuestionBoxView(g *gocui.Gui) *gocui.View {
	var result *gocui.View
	views := g.Views()
	for _, view := range views {
		if strings.EqualFold(QuestionBox, view.Name()) {
			result = view
		}
	}
	return result
}

func getAnswerBoxViews(g *gocui.Gui) []*gocui.View {
	var questionViews []*gocui.View
	views := g.Views()
	for _, view := range views {
		if isViewInSlice(BoxesView, view) {
			questionViews = append(questionViews, view)
		}
	}
	return questionViews
}

func isViewInSlice(viewNames []string, v *gocui.View) bool {
	result := false
	for _, name := range viewNames {
		if strings.EqualFold(name, v.Name()) {
			result = true
		}
	}
	return result
}
*/
