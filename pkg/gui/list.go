package gui

import "github.com/jroimartin/gocui"

type List struct {
	*gocui.View
	title string
	items []string
}

func CreateList(v *gocui.View) *List {
	list := &List{}
	list.View = v
	list.SelBgColor = gocui.ColorBlack
	list.SelFgColor = gocui.ColorWhite | gocui.AttrBold
	list.Autoscroll = true
	return list
}

// Focus hightlights the View of the current List
func (l *List) Focus(g *gocui.Gui) error {
	l.Highlight = true
	_, err := g.SetCurrentView(l.Name())

	return err
}

// Unfocus is used to remove highlighting from the current list
func (l *List) Unfocus() {
	l.Highlight = false
}

func (c *List) currentCursorY() int {
	_, y := c.Cursor()
	return y
}

func (c *List) Moveup() error {
	y := c.currentCursorY() - 1
	if y < 0 {
		return nil
	}
	return c.SetCursor(0, y)
}

func (c *List) Movedown() error {
	y := c.currentCursorY() + 1
	if y >= len(c.items) {
		return nil
	}
	return c.SetCursor(0, y)
}
