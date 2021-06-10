package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/sirupsen/logrus"
)

type Gui struct {
	g   *gocui.Gui
	Log *logrus.Entry

	// this tells us whether our views have been initially set up
	ViewsSetup bool
}

func NewGui() (*Gui, error) {
	gui := &Gui{}
	return gui, nil
}

func (gui *Gui) Run() error {
	g, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		return err
	}
	gui.g = g
	defer g.Close()
	//userConfig=gui.Config.GetUserConfig()

	g.SetManager(gocui.ManagerFunc(gui.layout), gocui.ManagerFunc(gui.getFocusLayout()))

	if err = gui.keybindings(g); err != nil {
		return err
	}
	gui.Log.Info("starting main loop")

	err = g.MainLoop()
	return err
}

func (gui *Gui) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
