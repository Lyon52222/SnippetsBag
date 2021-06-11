package gui

import (
	"fmt"
	"path"

	"github.com/jroimartin/gocui"
)

type SnipeetsPanel struct {
	v        *gocui.View
	snippets []string
}

func NewSnippetsPanel(v *gocui.View) (*SnipeetsPanel, error) {
	snipeetsPanel := &SnipeetsPanel{
		v: v,
	}

	return snipeetsPanel, nil
}

func (s *SnipeetsPanel) ShowSnippets() {
	s.v.Clear()
	for _, snippet := range s.snippets {
		_, name := path.Split(snippet)
		fmt.Fprintln(s.v, name)
	}
}

func (s *SnipeetsPanel) GetCurrentSnippetPath() string {
	_, cy := s.v.Cursor()
	return s.snippets[cy]
}

func (s *SnipeetsPanel) AddSnippets(snippets []string) {
	s.snippets = append(s.snippets, snippets...)
}

func (s *SnipeetsPanel) SetSnippets(snippets []string) {
	s.snippets = snippets
}

func (s *SnipeetsPanel) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy+1 < len(s.snippets) {
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

func (s *SnipeetsPanel) cursorUp(g *gocui.Gui, v *gocui.View) error {
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
