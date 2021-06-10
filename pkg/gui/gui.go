package gui

import (
	"github.com/Lyon52222/snippetsbag/pkg/config"
	"github.com/Lyon52222/snippetsbag/pkg/data"
	"github.com/jroimartin/gocui"
	"github.com/sirupsen/logrus"
)

type Gui struct {
	g           *gocui.Gui
	Log         *logrus.Entry
	Config      *config.AppConfig
	Data        *data.DataLoader
	Collections *CollectionsPanel
}

func NewGui(config *config.AppConfig) (*Gui, error) {
	gui := &Gui{
		Config: config,
	}
	data, err := data.NewData(config)
	if err != nil {
		return gui, err
	}
	gui.Data = data

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

	g.SetCurrentView(COLLECTIONS_PANEL)

	if err = gui.keybindings(g); err != nil {
		return err
	}

	err = g.MainLoop()
	return err
}

func (gui *Gui) handleCollectionsNextLine(g *gocui.Gui, v *gocui.View) error {
	gui.Collections.Movedown()
	return nil
}

func (gui *Gui) handleCollectionsPreLine(g *gocui.Gui, v *gocui.View) error {
	gui.Collections.Moveup()
	return nil
}

func (gui *Gui) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
