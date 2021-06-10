package gui

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

// getFocusLayout returns a manager function for when view gain and lose focus
func (gui *Gui) getFocusLayout() func(g *gocui.Gui) error {
	var previousView *gocui.View
	return func(g *gocui.Gui) error {
		newView := gui.g.CurrentView()
		if err := gui.onFocusChange(); err != nil {
			return err
		}
		// for now we don't consider losing focus to a popup panel as actually losing focus
		if newView != previousView && !gui.isPopupPanel(newView.Name()) {
			if err := gui.onFocusLost(previousView, newView); err != nil {
				return err
			}
			if err := gui.onFocus(newView); err != nil {
				return err
			}
			previousView = newView
		}
		return nil
	}
}

func (gui *Gui) onFocusChange() error {
	currentView := gui.g.CurrentView()
	for _, view := range gui.g.Views() {
		view.Highlight = view == currentView && view.Name() != "main"
	}
	return nil
}

func (gui *Gui) onFocusLost(v *gocui.View, newView *gocui.View) error {
	if v == nil {
		return nil
	}

	//if !gui.isPopupPanel(newView.Name()) {
	//v.ParentView = nil
	//}

	// refocusing because in responsive mode (when the window is very short) we want to ensure that after the view size changes we can still see the last selected item
	if err := gui.focusPointInView(v); err != nil {
		return err
	}

	gui.Log.Info(v.Name() + " focus lost")
	return nil
}

func (gui *Gui) onFocus(v *gocui.View) error {
	if v == nil {
		return nil
	}

	if err := gui.focusPointInView(v); err != nil {
		return err
	}

	gui.Log.Info(v.Name() + " focus gained")
	return nil
}

func (gui *Gui) layout(g *gocui.Gui) error {
	g.Highlight = true
	width, height := g.Size()

	if v, err := g.SetView(COLLECTIONS_PANEL, 0, 0, width/6, height/4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Personal Collections"
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorWhite | gocui.AttrBold

		v.Highlight = true
		gui.Collections, err = NewColletionsPanel(v)
		if err != nil {
			return err
		}
		gui.Collections.ShowCollections()
	}

	if v, err := g.SetView(FOLDERS_PANEL, 0, height/4, width/6, height-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Folders"
		childFolders := gui.Data.GetAllFolders()
		for _, f := range childFolders {
			fmt.Fprintln(v, f)
		}
	}

	if v, err := g.SetView(SNIPPETS_PANEL, width/6, 0, width/5*2, height-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Snippets"
		snippets := gui.Data.GetAllSnippets()
		for _, s := range snippets {
			fmt.Fprintln(v, s)
		}
	}

	if v, err := g.SetView(PREVIEW_PANEL, width/5*2, 0, width-1, height-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Preview"
		snippet, err := gui.Data.ReadSnippet("/Users/admin/.snippets/Python/test.py")
		if err == nil {
			v.Write(snippet)
		} else {
			log.Panicln(err)
		}
	}

	return nil
}
