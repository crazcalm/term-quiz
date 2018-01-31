package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

//InfoBar -- Gui component that holds one line of information
type InfoBar struct {
	name    string
	content string
}

//NewInfoBar -- Creates a new InfoBar gui component
func NewInfoBar(name, content string) *InfoBar {
	return &InfoBar{name: name, content: content}
}

func (i *InfoBar) location(g *gocui.Gui) (x, y, w, h int) {
	maxX, maxY := g.Size()
	x = -1
	y = maxY - 3
	w = maxX
	h = y + 2
	return
}

//Layout -- tells gocui.Gui how to display this component
func (i *InfoBar) Layout(g *gocui.Gui) error {
	x, y, w, h := i.location(g)
	v, err := g.SetView(i.name, x, y, w, h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		fmt.Fprint(v, i.content)
	}
	return nil
}
