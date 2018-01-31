package gui

import (
	"github.com/jroimartin/gocui"
	"log"
)

var (
	//ActiveView -- Index counter used when tabbing through answers
	ActiveView = 0
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
	nextIndex := (ActiveView + 1) % len(ABCDBoxes)
	name := ABCDBoxes[nextIndex]

	_, err := setCurrentViewOnTop(g, name)
	if err != nil {
		log.Panicln(err)
	}

	ActiveView = nextIndex
	return nil
}

//TFNextView -- Callback used to interate through the True and False choices
func TFNextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (ActiveView + 1) % len(TFBoxes)
	name := TFBoxes[nextIndex]

	_, err := setCurrentViewOnTop(g, name)
	if err != nil {
		log.Panicln(err)
	}

	ActiveView = nextIndex
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

//CursorUp -- Callback used to scroll up
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
*/
