package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

//Score -- Gui component that holds the question
type Score struct {
	name           string
	score          string
	questionNumber string
}

//NewScore -- creates new question gui component
func NewScore(name, score, questionNumber string) *Score {
	return &Score{name: name, score: score, questionNumber: questionNumber}
}

//location -- holds the location logic
func (s *Score) location(g *gocui.Gui) (x, y, w, h int) {
	maxX, maxY := g.Size()
	x = int(0.35 * float32(maxX))
	y = int(0.1 * float32(maxY))
	w = int(0.65 * float32(maxX))
	h = int(0.2 * float32(maxY))
	return
}

//Layout -- Tells gocui.Gui how to display this component
func (s *Score) Layout(g *gocui.Gui) error {
	x, y, w, h := s.location(g)
	v, err := g.SetView(s.name, x, y, w, h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Title = fmt.Sprintf("Total Score -- %s", s.score)
		fmt.Fprint(v, fmt.Sprintf("Question Number %s", s.questionNumber))
	}
	return nil
}
