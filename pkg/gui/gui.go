package gui

import (
	"path"
	"strings"

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
func (gui *Gui) handleCollectionsNextLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Collections.cursorDown(g, v)
}

func (gui *Gui) handleCollectionsPreLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Collections.cursorUp(g, v)
}

func (gui *Gui) createNewFolder(g *gocui.Gui, v *gocui.View) error {
	confirmationPanel, _ := g.View(CONFIRMATION_PANEL)
	dirName := confirmationPanel.Buffer()
	dirName = strings.Replace(dirName, "\n", "", -1)
	dirName = path.Join(gui.Config.SnippetsDir, dirName)
	if err := gui.Data.CreateNewFolder(dirName); err != nil {
		return err
	}
	gui.Folders.AddFolder(dirName)
	return nil
}

func (gui *Gui) handleCreateNewFolder(g *gocui.Gui, v *gocui.View) error {
	//return gui.createConfirmationPanel(g, v, "title", "prompt", nil, nil)
	return gui.createPromptPanel(g, v, gui.Tr.CreateNewFolderPanelTitle, gui.createNewFolder)
}

func (gui *Gui) createNewSnippet(g *gocui.Gui, v *gocui.View) error {
	confirmationPanel, _ := g.View(CONFIRMATION_PANEL)
	snippetName := confirmationPanel.Buffer()
	snippetName = strings.Replace(snippetName, "\n", "", -1)
	snippetName = path.Join(gui.Folders.GetCurrentFolder(), snippetName)
	if err := gui.Data.CreateNewSnippet(snippetName); err != nil {
		return err
	}
	gui.Snippets.AddSnippet(snippetName)
	return nil
}
func (gui *Gui) handleCreateNewSnippet(g *gocui.Gui, v *gocui.View) error {
	return gui.createPromptPanel(g, v, gui.Tr.CreateNewFolderPanelTitle, gui.createNewSnippet)
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
