package gui

import (
	"fmt"

	"github.com/crazcalm/text-to-width"

	"github.com/jroimartin/gocui"
)

//Question -- Gui component that holds the question
type Question struct {
	name  string
	title string
	body  string
}

//NewQuestion -- creates new question gui component
func NewQuestion(name, title string, body string) *Question {
	return &Question{name: name, title: title, body: body}
}

func (q *Question) location(g *gocui.Gui) (x, y, w, h int) {
	maxX, maxY := g.Size()
	x = int(0.2 * float32(maxX))
	y = int(0.05 * float32(maxY))
	w = int(0.8 * float32(maxX))
	h = int(0.35 * float32(maxY))
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
		//Set to false because text-to-width will word wrap for us
		v.Wrap = false

		v.Title = q.title

		//The allowed amount of space that the text can take
		length := w - x - 1
		fmt.Fprint(v, texttowidth.Format(q.body, length))
	}
	return nil
}
