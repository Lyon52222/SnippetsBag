package gui

import (
	"fmt"
	"path"

	"github.com/Lyon52222/snippetsbag/pkg/data"
	"github.com/jroimartin/gocui"
)

type CollectionsPanel struct {
	v            *gocui.View
	colletions   []string
	dataloader   *data.DataLoader
	snippetPanel *SnipeetsPanel
}

func NewColletionsPanel(v *gocui.View, dataloader *data.DataLoader, snippetPanel *SnipeetsPanel) (*CollectionsPanel, error) {
	collectionsPanel := &CollectionsPanel{
		v:            v,
		dataloader:   dataloader,
		colletions:   []string{"\uf719 All Snippets", "\ue7c5 Vim", "\ue795 Shell"},
		snippetPanel: snippetPanel,
	}
	return collectionsPanel, nil

}

func (c *CollectionsPanel) ShowCollections() {
	for _, colletion := range c.colletions {
		_, file := path.Split(colletion)
		fmt.Fprintln(c.v, file)
	}
}

func (c *CollectionsPanel) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy+1 < len(c.colletions) {
			if err := v.SetCursor(cx, cy+1); err != nil {
				ox, oy := v.Origin()
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}

		}
	}
	if err := c.snippetPanel.Refresh(c.getCurentColletion()); err != nil {
		return err
	}
	return nil
}

func (c *CollectionsPanel) getCurentColletion() string {
	_, cy := c.v.Cursor()
	return c.colletions[cy]
}

func (c *CollectionsPanel) cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	if err := c.snippetPanel.Refresh(c.getCurentColletion()); err != nil {
		return err
	}
	return nil
}

func (gui *Gui) handleCollectionsNextLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Collections.cursorDown(g, v)
}

func (gui *Gui) handleCollectionsPreLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Collections.cursorUp(g, v)
}

func (gui *Gui) focusCollectionsPanel(g *gocui.Gui, v *gocui.View) error {
	if _, err := g.SetCurrentView(COLLECTIONS_PANEL); err != nil {
		return nil
	}

	if err := gui.Snippets.Refresh(gui.Collections.getCurentColletion()); err != nil {
		return err
	}

	gui.Collections.v.Highlight = true
	gui.Folders.v.Highlight = false
	return nil
}
