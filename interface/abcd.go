package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strings"
)

const (
	//BoxA -- title for question box A
	BoxA = "A"
	//BoxB -- title for question box B
	BoxB = "B"
	//BoxC -- title for question box C
	BoxC = "C"
	//BoxD -- title for question box D
	BoxD = "D"
	//BoxTrue -- title for question box true
	BoxTrue = "True"
	//BoxFalse -- title for question box false
	BoxFalse = "False"
	//BoxBlank -- title for question box fill in the blank
	BoxBlank = "Write"
)

//QuestionFrame -- Gui component that holds the question frame
type QuestionFrame struct {
	name string
	title string
}

//NewQuestionFrame -- creates new question frame gui component
func NewQuestionFrame(name, title string) *QuestionFrame {
	return &QuestionFrame{name: name, title: title}
}

//Layout -- Tells gocui.Gui how to display this compenent
func (qf *QuestionFrame) Layout (g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(qf.name, -1, -1, maxX, int(0.5 * float32(maxY)))
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = qf.title
	}
	return nil
}

//Question -- Gui component that holds the question
type Question struct {
	name string
	title string
	body string // will replace with questions.Question
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

//Answer -- Gui component that holds a Answer
type Answer struct {
	name string
	title string
	body string // will replace with answers.Answer
}

//NewAnswer -- creates a new Answer gui component
func NewAnswer(name, title string, body string) *Answer {
	return &Answer{name: name, title: title, body: body}
}

func (a *Answer) location(title string, g *gocui.Gui) (x, y, w, h int){
	maxX, maxY := g.Size()
	if strings.EqualFold(title, BoxA){
		x = -1
		y = int(0.5 * float32(maxY))
		w = int(0.5 * float32(maxX))
		h = int(0.73 * float32(maxX))
		return

	} else if strings.EqualFold(title, BoxB) {
		x = int(0.5 * float32(maxX))
		y = int(0.5 * float32(maxY))
		w = maxX
		h = int(0.73 * float32(maxY))
		return
		
	} else if strings.EqualFold(title, BoxC) {
		x = -1
		y = int(0.77 * float32(maxY))
		w = int(0.5 * float32(maxX))
		h = maxY
		return
	} else if strings.EqualFold(title, BoxD) {
		x = int(0.5 * float32(maxX))
		y = int(0.77 * float32(maxY))
		w = maxX
		h = maxY
		return 
	}

	fmt.Printf("%d, %d, %d, %d, %d, %d\n", x, y, w, h, maxX, maxY)
	return -1, -1, maxX, maxY
}

//Layout -- Tells gocui.Gui how to display this compenent
func (a *Answer) Layout(g *gocui.Gui) error {
	x, y, w, h := a.location(a.title, g)
	v, err := g.SetView(a.name, x, y, w, h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = a.title
		v.Highlight = true
		v.Wrap = true
		fmt.Fprint(v, a.body)
	}
	return nil
}