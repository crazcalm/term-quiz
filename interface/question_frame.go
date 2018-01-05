package gui

import (
	"github.com/jroimartin/gocui"
)

//QuestionFrame -- Gui component that holds the question frame
type QuestionFrame struct {
	name  string
	title string
}

//NewQuestionFrame -- creates new question frame gui component
func NewQuestionFrame(name, title string) *QuestionFrame {
	return &QuestionFrame{name: name, title: title}
}

//Layout -- Tells gocui.Gui how to display this compenent
func (qf *QuestionFrame) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(qf.name, -1, -1, maxX, int(0.5*float32(maxY)))
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = qf.title
	}
	return nil
}
