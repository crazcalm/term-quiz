package gui

import (
	"fmt"
	"log"
	"strings"

	"github.com/crazcalm/text-to-width"

	"github.com/jroimartin/gocui"
)

//Answer -- Gui component that holds a Answer
type Answer struct {
	name  string
	title string
	body  string // will replace with answers.Answer
}

//NewAnswer -- creates a new Answer gui component
func NewAnswer(name, title string, body string) *Answer {
	return &Answer{name: name, title: title, body: body}
}

func (a *Answer) location(title string, g *gocui.Gui) (x, y, w, h int) {
	maxX, maxY := g.Size()
	if strings.EqualFold(title, BoxA) {
		x = -1
		y = int(0.4 * float32(maxY))
		w = int(0.5 * float32(maxX))
		h = int(0.63 * float32(maxY))
		return

	} else if strings.EqualFold(title, BoxB) {
		x = int(0.5 * float32(maxX))
		y = int(0.4 * float32(maxY))
		w = maxX
		h = int(0.63 * float32(maxY))
		return

	} else if strings.EqualFold(title, BoxC) {
		x = -1
		y = int(0.67 * float32(maxY))
		w = int(0.5 * float32(maxX))
		h = int(0.87 * float32(maxY))
		return
	} else if strings.EqualFold(title, BoxD) {
		x = int(0.5 * float32(maxX))
		y = int(0.67 * float32(maxY))
		w = maxX
		h = int(0.87 * float32(maxY))
		return
	} else if strings.EqualFold(title, BoxTrue) {
		x = int(0.05 * float32(maxX))
		y = int(0.5 * float32(maxY))
		w = int(0.45 * float32(maxX))
		h = int(0.8 * float32(maxY))
		return

	} else if strings.EqualFold(title, BoxFalse) {
		x = int(0.55 * float32(maxX))
		y = int(0.5 * float32(maxY))
		w = int(0.95 * float32(maxX))
		h = int(0.8 * float32(maxY))
		return
	} else if strings.EqualFold(title, BoxBlank) {
		x = int(0.05 * float32(maxX))
		y = int(0.5 * float32(maxY))
		w = int(0.95 * float32(maxX))
		h = int(0.8 * float32(maxY))
	}

	return
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

		//Set to false because text-to-width will do the wrapping
		v.Wrap = false

		//Make sure the fill in the blank is editable
		if strings.EqualFold(a.title, BoxBlank) {
			v.Editable = true
		} else {
			//The allowed amount of space the text can take
			length := w - x - 1
			fmt.Fprint(v, texttowidth.Format(a.body, length))
		}

		//Make sure the first answer is on top
		if strings.EqualFold(a.title, BoxA) || strings.EqualFold(a.title, BoxTrue) || strings.EqualFold(a.title, BoxBlank) {
			_, err := setCurrentViewOnTop(g, a.name)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}
