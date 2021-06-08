package gui

import (
	"github.com/jroimartin/gocui"
)

type Collections struct {
	*gocui.View
	title      string
	colletions []string
}

func (c *Collections) currentCursorY() int {
	_, y := c.Cursor()
	return y
}

func (c *Collections) Moveup() error {
	y := c.currentCursorY() - 1
	if y < 0 {
		return nil
	}
	return c.SetCursor(0, y)
}

func (c *Collections) Movedown() error {
	y := c.currentCursorY() + 1
	if y >= len(c.colletions) {
		return nil
	}
	return c.SetCursor(0, y)
}
