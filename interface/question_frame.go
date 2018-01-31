package gui

import (
	"github.com/jroimartin/gocui"
)

//QuestionFrame -- Gui component that holds the question frame
type QuestionFrame struct {
	name string
}

//NewQuestionFrame -- creates new question frame gui component
func NewQuestionFrame(name string) *QuestionFrame {
	return &QuestionFrame{name: name}
}

//Layout -- Tells gocui.Gui how to display this compenent
func (qf *QuestionFrame) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	_, err := g.SetView(qf.name, -1, -1, maxX, int(0.4*float32(maxY)))
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}
