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
	Folders     *FoldersPanel
	Snippets    *SnipeetsPanel
	Preview     *PreviewPanel
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
	defer g.Close()

	gui.g = g

	//g.SetManager(gocui.ManagerFunc(gui.layout), gocui.ManagerFunc(gui.getFocusLayout()))
	g.SetManagerFunc(gui.layout)

	if err = gui.keybindings(g); err != nil {
		return err
	}

	err = g.MainLoop()
	return err
}

func (gui *Gui) handleCollectionsNextLine(g *gocui.Gui, v *gocui.View) error {
	gui.Collections.cursorDown(g, v)
	return nil
}

func (gui *Gui) handleCollectionsPreLine(g *gocui.Gui, v *gocui.View) error {
	gui.Collections.cursorUp(g, v)
	return nil
}

func (gui *Gui) handleFoldersNextLine(g *gocui.Gui, v *gocui.View) error {
	gui.Folders.cursorDown(g, v)
	folder := gui.Folders.GetCurrentFolder()
	gui.Snippets.SetSnippets(gui.Data.GetSnippetsFromPath(folder))
	gui.Snippets.ShowSnippets()
	return nil
}

func (gui *Gui) handleFoldersPreLine(g *gocui.Gui, v *gocui.View) error {
	gui.Folders.cursorUp(g, v)
	folder := gui.Folders.GetCurrentFolder()
	gui.Snippets.SetSnippets(gui.Data.GetSnippetsFromPath(folder))
	gui.Snippets.ShowSnippets()
	return nil
}

func (gui *Gui) handleSnippetsNextLine(g *gocui.Gui, v *gocui.View) error {
	gui.Snippets.cursorDown(g, v)
	snippetPath := gui.Snippets.GetCurrentSnippetPath()
	snippet, err := gui.Data.ReadSnippet(snippetPath)
	if err != nil {
		return err
	}
	gui.Preview.SetSnippetPath(snippetPath)
	gui.Preview.SetSnippet(snippet)
	gui.Preview.ShowSnippet()
	return nil
}

func (gui *Gui) handleSnippetsPreLine(g *gocui.Gui, v *gocui.View) error {
	gui.Snippets.cursorUp(g, v)
	snippetPath := gui.Snippets.GetCurrentSnippetPath()
	snippet, err := gui.Data.ReadSnippet(snippetPath)
	if err != nil {
		return err
	}
	gui.Preview.SetSnippetPath(snippetPath)
	gui.Preview.SetSnippet(snippet)
	gui.Preview.ShowSnippet()
	return nil
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
