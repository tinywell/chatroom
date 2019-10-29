package ui

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

type UI struct {
	ui     *gocui.Gui
	viewO  *gocui.View
	viewI  *gocui.View
	inputC chan string
}

func NewUI() (*UI, error) {
	ui := &UI{
		inputC: make(chan string),
	}
	err := ui.init()
	if err != nil {
		return nil, err

	}
	return ui, nil
}

func (ui *UI) init() error {
	cui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	cui.Cursor = true
	// cui.Mouse = true
	ui.ui = cui
	sizeX, sizeY := cui.Size()
	viewOX0, viewOY0 := 0, 0
	viewOX1, viewOY1 := sizeX-1, sizeY-4
	viewIX0, viewIY0 := 0, sizeY-3
	viewIX1, viewIY1 := sizeX-1, sizeY-1
	viewO, err := cui.SetView("room", viewOX0, viewOY0, viewOX1, viewOY1)
	if err != gocui.ErrUnknownView {
		fmt.Println(err)
		return err
	}
	viewO.Wrap = true
	viewO.Autoscroll = true
	viewO.Frame = true
	viewO.Title = "CHAT ROOM"
	ui.viewO = viewO
	viewI, err := cui.SetView("input", viewIX0, viewIY0, viewIX1, viewIY1)
	if err != gocui.ErrUnknownView {
		return err
	}
	viewI.Editable = true
	viewI.Frame = true
	viewI.Title = "INPUT"
	viewI.SetCursor(0, 0)
	viewI.Highlight = true
	// ui.viewI, err = cui.SetViewOnBottom("input")
	// if err != nil {
	// 	return err
	// }
	_, err = cui.SetCurrentView("input")
	if err != nil {
		return err
	}
	ui.viewI = viewI
	return nil
}

func (ui *UI) GetInputC() <-chan string {
	return ui.inputC
}

func (ui *UI) Open() {
	if err := ui.ui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := ui.ui.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, ui.input); err != nil {
		log.Panicln(err)
	}

	if err := ui.ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (ui *UI) CLose() {
	ui.ui.Close()
}

func (ui *UI) ApppendMsg(msg, name string) {
	m := fmt.Sprintf("\033[34m%s:\033[0m\n%s\n", name, msg)
	ui.viewO.Write([]byte(m))
}

func (ui *UI) input(g *gocui.Gui, v *gocui.View) error {
	if v.Name() != "input" {
		return nil
	}
	ui.inputC <- v.Buffer()
	v.SetCursor(0, 0)
	v.Clear()
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
