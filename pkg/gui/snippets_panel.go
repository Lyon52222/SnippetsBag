package gui

import (
	"fmt"
	"path"
	"strings"

	"github.com/Lyon52222/snippetsbag/pkg/data"
	"github.com/jroimartin/gocui"
)

type SnipeetsPanel struct {
	v            *gocui.View
	snippets     []string
	dataloader   *data.DataLoader
	previewPanel *PreviewPanel
}

func NewSnippetsPanel(v *gocui.View, dataloader *data.DataLoader, previewPanel *PreviewPanel) (*SnipeetsPanel, error) {
	snipeetsPanel := &SnipeetsPanel{
		v:            v,
		dataloader:   dataloader,
		previewPanel: previewPanel,
	}
	snipeetsPanel.snippets = append(snipeetsPanel.snippets, snipeetsPanel.dataloader.GetAllSnippets()...)
	return snipeetsPanel, nil
}

func (s *SnipeetsPanel) ShowSnippets() {
	s.v.Clear()
	for _, snippet := range s.snippets {
		_, name := path.Split(snippet)
		fmt.Fprintln(s.v, name)
	}
	s.SetCursorY(0)
}

func (s *SnipeetsPanel) Refresh(folder string) error {
	if folder == "\uf719 All Snippets" {
		s.snippets = s.dataloader.GetAllSnippets()
	} else if folder == "\ue7c5 Vim" {
		return nil
	} else if folder == "\ue795 Shell" {
		return nil
	} else {
		s.snippets = s.dataloader.GetSnippetsFromPath(folder)
	}
	if len(s.snippets) == 0 {
		s.v.Clear()
		s.previewPanel.v.Clear()
		return nil
	}
	s.SetCursorY(0)
	s.ShowSnippets()
	return nil
}

func (s *SnipeetsPanel) SetCursorY(y int) error {
	cx, _ := s.v.Cursor()
	ox, _ := s.v.Origin()
	if err := s.v.SetCursor(cx, 0); err != nil {
		if err := s.v.SetOrigin(ox, 0); err != nil {
			return err
		}
	}
	if err := s.previewPanel.Refresh(s.GetCurrentSnippetPath()); err != nil {
		return err
	}
	return nil
}

func (s *SnipeetsPanel) GetCurrentSnippetPath() string {
	_, cy := s.v.Cursor()
	return s.snippets[cy]
}

func (s *SnipeetsPanel) AddSnippets(snippets []string) {
	s.snippets = append(s.snippets, snippets...)
	for _, snippet := range snippets {
		_, name := path.Split(snippet)
		fmt.Fprintln(s.v, name)
	}
}

func (s *SnipeetsPanel) AddSnippet(snippet string) {
	s.snippets = append(s.snippets, snippet)
	_, name := path.Split(snippet)
	fmt.Fprintln(s.v, name)

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
	if err := s.previewPanel.Refresh(s.GetCurrentSnippetPath()); err != nil {
		return err
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
	if err := s.previewPanel.Refresh(s.GetCurrentSnippetPath()); err != nil {
		return err
	}
	return nil
}

func (gui *Gui) handleSnippetsNextLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Snippets.cursorDown(g, v)
}

func (gui *Gui) handleSnippetsPreLine(g *gocui.Gui, v *gocui.View) error {
	return gui.Snippets.cursorUp(g, v)
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

func (gui *Gui) focusSnippetsPanel(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(SNIPPETS_PANEL)
	return err
}
