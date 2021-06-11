package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type CollectionsPanel struct {
	v          *gocui.View
	colletions []string
}

func NewColletionsPanel(v *gocui.View) (*CollectionsPanel, error) {
	collectionsPanel := &CollectionsPanel{
		v:          v,
		colletions: []string{"\uf719 All Snippets", "\ue7c5 Vim", "\ue795 Shell"},
	}
	return collectionsPanel, nil

}

func (c *CollectionsPanel) ShowCollections() {
	for _, colletion := range c.colletions {
		fmt.Fprintln(c.v, colletion)
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
	return nil
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
	return nil
}
