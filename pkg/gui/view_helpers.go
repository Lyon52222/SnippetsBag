package gui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Lyon52222/snippetsbag/pkg/utils"
	"github.com/jroimartin/gocui"
	"github.com/spkg/bom"
)

func (gui *Gui) isPopupPanel(viewName string) bool {
	return viewName == "confirmation" || viewName == "menu"
}

func (gui *Gui) focusPointInView(view *gocui.View) error {
	return nil
}

func (gui *Gui) popPreviousView() string {
	if gui.PreviewViews.Len() > 0 {
		return gui.PreviewViews.Pop().(string)
	}
	return ""
}

func (gui *Gui) peekPreviousView() string {
	if gui.PreviewViews.Len() > 0 {
		return gui.PreviewViews.Peek().(string)
	}
	return ""
}

func (gui *Gui) pushPreviousView(name string) {
	gui.PreviewViews.Push(name)
}

func (gui *Gui) returnFocus(g *gocui.Gui, v *gocui.View) error {
	previousViewName := gui.popPreviousView()
	previousView, err := g.View(previousViewName)
	if err != nil {
		previousView, err = g.View(gui.initiallyFocusedView())
		if err != nil {
			gui.Log.Error(err)
		}
	}
	return gui.switchFocus(g, v, previousView, true)
}

func (gui *Gui) switchFocus(g *gocui.Gui, oldView, newView *gocui.View, returning bool) error {
	if oldView != nil && !gui.isPopupPanel(oldView.Name()) && !returning {
		gui.pushPreviousView(oldView.Name())
	}

	gui.Log.Info("setting highlight to true for view" + newView.Name())
	gui.Log.Info("new focused view is " + newView.Name())

	if _, err := g.SetCurrentView(newView.Name()); err != nil {
		return err
	}

	if _, err := g.SetViewOnTop(newView.Name()); err != nil {
		return err
	}
	g.Cursor = newView.Editable

	return nil

}

func (gui *Gui) cleanString(s string) string {
	output := string(bom.Clean([]byte(s)))
	return utils.NormalizeLinefeeds(output)
}

func (gui *Gui) setViewContet(g *gocui.Gui, v *gocui.View, s string) error {
	v.Clear()
	fmt.Fprint(v, gui.cleanString(s))
	return nil
}

func (gui *Gui) renderString(g *gocui.Gui, viewName, s string) error {
	g.Update(func(*gocui.Gui) error {
		v, err := g.View(viewName)

		if err != nil {
			return nil
		}
		if err := v.SetOrigin(0, 0); err != nil {
			return err
		}
		if err := v.SetOrigin(0, 0); err != nil {
			return err
		}
		return gui.setViewContet(gui.g, v, s)
	})
	return nil
}

func (gui *Gui) optionsMapToString(optionsMap map[string]string) string {
	optionsArray := make([]string, 0)
	for key, description := range optionsMap {
		optionsArray = append(optionsArray, key+": "+description)

	}
	sort.Strings(optionsArray)
	return strings.Join(optionsArray, ", ")
}

func (gui *Gui) renderOptionsMap(optionMap map[string]string) error {
	return gui.renderString(gui.g, "options", gui.optionsMapToString(optionMap))
}
