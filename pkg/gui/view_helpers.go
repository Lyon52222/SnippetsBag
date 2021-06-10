package gui

import "github.com/jroimartin/gocui"

func (gui *Gui) isPopupPanel(viewName string) bool {
	return viewName == "confirmation" || viewName == "menu"
}

func (gui *Gui) focusPointInView(view *gocui.View) error {
	return nil
}
