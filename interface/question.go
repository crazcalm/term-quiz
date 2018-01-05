package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

//Question -- Gui component that holds the question
type Question struct {
	name  string
	title string
	body  string // will replace with questions.Question
}

//NewQuestion -- creates new question gui component
func NewQuestion(name, title string, body string) *Question {
	return &Question{name: name, title: title, body: body}
}

func (q *Question) location(g *gocui.Gui) (x, y, w, h int) {
	maxX, maxY := g.Size()
	x = int(0.2 * float32(maxX))
	y = int(0.1 * float32(maxY))
	w = int(0.8 * float32(maxX))
	h = int(0.4 * float32(maxY))
	return
}

//Layout -- Tells gocui.Gui how to display this compenent
func (q *Question) Layout(g *gocui.Gui) error {
	x, y, w, h := q.location(g)
	v, err := g.SetView(q.name, x, y, w, h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		fmt.Fprint(v, q.body)
	}
	return nil
}
