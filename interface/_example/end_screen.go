package main

import (
	"github.com/crazcalm/term-quiz/interface"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Highlight the selected view and make it green
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	score := gui.NewScore(gui.ScoreName, "11/12", "4")
	explain := gui.NewExplanation(gui.Explain, gui.Right, "question string", "Answer string", "explaination string")

	g.SetManager(score, explain)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
