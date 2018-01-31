package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

//Explanation -- Gui component that holds the question
type Explanation struct {
	name     string
	result   string
	question string
	answer   string
	explain  string
}

//NewExplanation -- creates new question gui component
func NewExplanation(name, result, question, answer, explain string) *Explanation {
	return &Explanation{name: name, result: result, question: question, answer: answer, explain: explain}
}

//location -- holds the location logic
func (e *Explanation) location(g *gocui.Gui) (x, y, w, h int) {
	maxX, maxY := g.Size()
	x = int(0.1 * float32(maxX))
	y = int(0.3 * float32(maxY))
	w = int(0.9 * float32(maxX))
	h = int(0.8 * float32(maxY))
	return
}

//Layout -- Tells gocui.Gui how to display this compenent
func (e *Explanation) Layout(g *gocui.Gui) error {
	x, y, w, h := e.location(g)
	v, err := g.SetView(e.name, x, y, w, h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true

		//result as title -- you were right or wrong
		v.Title = e.result

		//Display question
		fmt.Fprint(v, fmt.Sprintf("(Question) -- %s\n\n", e.question))

		//Display answer
		fmt.Fprint(v, fmt.Sprintf("(Answer) -- %s\n\n", e.answer))

		//Display explanation
		fmt.Fprint(v, fmt.Sprintf("(Explanation) -- %s", e.explain))
	}
	return nil
}
