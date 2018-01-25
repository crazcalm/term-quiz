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

	questionFrame := gui.NewQuestionFrame("questionFrame")
	question := gui.NewQuestion("question", "title - question", "question string")
	answerTrue := gui.NewAnswer(gui.BoxTrue, gui.BoxTrue, "answer string")
	answerFalse := gui.NewAnswer(gui.BoxFalse, gui.BoxFalse, "ppppppppppppppppppppppppppppppp")

	g.SetManager(questionFrame, question, answerTrue, answerFalse)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, gui.TFNextView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
