package gui

import (
	"github.com/Lyon52222/snippetsbag/pkg/config"
	"github.com/Lyon52222/snippetsbag/pkg/data"
	"github.com/golang-collections/collections/stack"
	"github.com/jroimartin/gocui"
	"github.com/sirupsen/logrus"
)

type Gui struct {
	g            *gocui.Gui
	Log          *logrus.Entry
	Config       *config.AppConfig
	Data         *data.DataLoader
	Collections  *CollectionsPanel
	Folders      *FoldersPanel
	Snippets     *SnipeetsPanel
	Preview      *PreviewPanel
	PreviewViews *stack.Stack
}

func NewGui(config *config.AppConfig) (*Gui, error) {
	gui := &Gui{
		Config:       config,
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
func (gui *Gui) handleCollectionsNextLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Collections.cursorDown(g, v)
}

func (gui *Gui) handleCollectionsPreLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Collections.cursorUp(g, v)
}

func (gui *Gui) handleFoldersNextLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Folders.cursorDown(g, v)
}

func (gui *Gui) handleFoldersPreLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Folders.cursorUp(g, v)
}

func (gui *Gui) handleSnippetsNextLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Snippets.cursorDown(g, v)
}

func (gui *Gui) handleSnippetsPreLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Snippets.cursorUp(g, v)
}

func (gui *Gui) focusCollectionsPanel(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(COLLECTIONS_PANEL)
	gui.Collections.v.Highlight = true
	gui.Folders.v.Highlight = false
	return err
}

func (gui *Gui) focusFoldersPanel(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(FOLDERS_PANEL)
	gui.Collections.v.Highlight = false
	gui.Folders.v.Highlight = true
	return err
}

func (gui *Gui) focusSnippetsPanel(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(SNIPPETS_PANEL)
	return err
}

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
