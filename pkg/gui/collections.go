package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type CollectionsPanel struct {
	v          *gocui.View
	title      string
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

func (c *CollectionsPanel) currentCursorY() int {
	_, y := c.v.Cursor()
	return y
}

func (c *CollectionsPanel) Moveup() error {
	y := c.currentCursorY() - 1
	if y < 0 {
		return nil
	}
	return c.v.SetCursor(0, y)
}

func (c *CollectionsPanel) Movedown() error {
	y := c.currentCursorY() + 1
	if y >= len(c.colletions) {
		return nil
	}
	return c.v.SetCursor(0, y)
}
