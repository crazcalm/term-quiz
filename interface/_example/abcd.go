package main

import (
	"github.com/crazcalm/term-quiz/interface"
	"github.com/jroimartin/gocui"
	"log"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	questionFrame := gui.NewQuestionFrame("questionFrame", "count")
	question := gui.NewQuestion("question", "title - question", "question string")
	answerA := gui.NewAnswer("answerA", gui.BoxA, "answer string")
	answerB := gui.NewAnswer("answerB", gui.BoxB, "ppppppppppppppppppppppppppppppp")
	answerC := gui.NewAnswer("answerC", gui.BoxC, "C-Town")
	answerD := gui.NewAnswer("answerD", gui.BoxD, "last but not least")

	g.SetManager(questionFrame, question, answerA, answerB, answerC, answerD)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
