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
	answerA := gui.NewAnswer(gui.BoxA, gui.BoxA, "answer string")
	answerB := gui.NewAnswer(gui.BoxB, gui.BoxB, "ppppppppppppppppppppppppppppppp")
	answerC := gui.NewAnswer(gui.BoxC, gui.BoxC, "C-Town")
	answerD := gui.NewAnswer(gui.BoxD, gui.BoxD, "last but not least")
	infoBar := gui.NewInfoBar(gui.InfoBarName, gui.InfoBarABCD)

	g.SetManager(questionFrame, question, answerA, answerB, answerC, answerD, infoBar)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, gui.ABCDNextView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
