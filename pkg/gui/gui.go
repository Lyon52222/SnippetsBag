package gui

import (
	"github.com/Lyon52222/snippetsbag/pkg/config"
	"github.com/Lyon52222/snippetsbag/pkg/data"
	"github.com/Lyon52222/snippetsbag/pkg/i18n"
	"github.com/golang-collections/collections/stack"
	"github.com/jroimartin/gocui"
	"github.com/sirupsen/logrus"
)

type Gui struct {
	g            *gocui.Gui
	Log          *logrus.Entry
	Config       *config.AppConfig
	Tr           *i18n.TranslationSet
	Data         *data.DataLoader
	Collections  *CollectionsPanel
	Folders      *FoldersPanel
	Snippets     *SnipeetsPanel
	Preview      *PreviewPanel
	PreviewViews *stack.Stack
}

func NewGui(log *logrus.Entry, config *config.AppConfig, tr *i18n.TranslationSet) (*Gui, error) {
	gui := &Gui{
		Log:          log,
		Config:       config,
		Tr:           tr,
		PreviewViews: stack.New(),
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
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	gui.g = g

	//g.SetManager(gocui.ManagerFunc(gui.layout), gocui.ManagerFunc(gui.getFocusLayout()))
	g.SetManagerFunc(gui.layout)

	if err = gui.keybindings(g); err != nil {
		return err
	}

	err = g.MainLoop()
	return err
}

//----------

func (gui *Gui) focusPreviewPanel(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(PREVIEW_PANEL)
	return err
}

func (gui *Gui) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

//----------

func (gui *Gui) initiallyFocusedView() string {
	return SNIPPETS_PANEL
}
