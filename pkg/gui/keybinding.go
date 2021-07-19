package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type Binding struct {
	ViewName    string
	Handler     func(*gocui.Gui, *gocui.View) error
	Key         interface{}
	Modifier    gocui.Modifier
	Description string
}

func (b *Binding) GetDisplayStrings(isFocused bool) []string {
	return []string{b.GetKey(), b.Description}
}

func (b *Binding) GetKey() string {
	key := 0

	switch b.Key.(type) {
	case rune:
		key = int(b.Key.(rune))
	case gocui.Key:
		key = int(b.Key.(gocui.Key))
	}

	// special keys
	switch key {
	case 27:
		return "esc"
	case 13:
		return "enter"
	case 32:
		return "space"
	case 65514:
		return "►"
	case 65515:
		return "◄"
	case 65517:
		return "▲"
	case 65516:
		return "▼"
	case 65508:
		return "PgUp"
	case 65507:
		return "PgDn"
	}

	return fmt.Sprintf("%c", key)

}

func (gui *Gui) GetInitialKeybindings() []*Binding {
	bindings := []*Binding{
		//{
		//ViewName: "",
		//Key:      'q',
		//Modifier: gocui.ModNone,
		//Handler:  gui.quit,
		//},
		{
			ViewName: "",
			Key:      gocui.KeyCtrlQ,
			Modifier: gocui.ModNone,
			Handler:  gui.quit,
		},
		{
			ViewName: COLLECTIONS_PANEL,
			Key:      gocui.KeyArrowDown,
			Modifier: gocui.ModNone,
			Handler:  gui.focusFoldersPanel,
		},
		{
			ViewName: COLLECTIONS_PANEL,
			Key:      gocui.KeyArrowRight,
			Modifier: gocui.ModNone,
			Handler:  gui.focusSnippetsPanel,
		},
		{
			ViewName: FOLDERS_PANEL,
			Key:      'a',
			Modifier: gocui.ModNone,
			Handler:  gui.handleCreateNewFolder,
		},
		{
			ViewName: FOLDERS_PANEL,
			Key:      gocui.KeyArrowUp,
			Modifier: gocui.ModNone,
			Handler:  gui.focusCollectionsPanel,
		},
		{
			ViewName: FOLDERS_PANEL,
			Key:      gocui.KeyArrowRight,
			Modifier: gocui.ModNone,
			Handler:  gui.focusSnippetsPanel,
		},
		{
			ViewName: SNIPPETS_PANEL,
			Key:      'a',
			Modifier: gocui.ModNone,
			Handler:  gui.handleCreateNewSnippet,
		},
		{
			ViewName: SNIPPETS_PANEL,
			Key:      gocui.KeyArrowLeft,
			Modifier: gocui.ModNone,
			Handler:  gui.focusCollectionsPanel,
		},
		{
			ViewName: SNIPPETS_PANEL,
			Key:      gocui.KeyArrowRight,
			Modifier: gocui.ModNone,
			Handler:  gui.focusPreviewPanel,
		},
		{
			ViewName: PREVIEW_PANEL,
			Key:      gocui.KeyArrowLeft,
			Modifier: gocui.ModNone,
			Handler:  gui.focusSnippetsPanel,
		},
		{
			ViewName: PREVIEW_PANEL,
			Key:      'e',
			Modifier: gocui.ModNone,
			Handler:  gui.OpenSnippetWithEditor,
		},
	}

	panelMap := map[string]struct {
		onKeyUpPress   func(*gocui.Gui, *gocui.View) error
		onKeyDownPress func(*gocui.Gui, *gocui.View) error
	}{
		COLLECTIONS_PANEL: {onKeyUpPress: gui.handleCollectionsPreLine, onKeyDownPress: gui.handleCollectionsNextLine},
		FOLDERS_PANEL:     {onKeyUpPress: gui.handleFoldersPreLine, onKeyDownPress: gui.handleFoldersNextLine},
		SNIPPETS_PANEL:    {onKeyUpPress: gui.handleSnippetsPreLine, onKeyDownPress: gui.handleSnippetsNextLine},
	}

	for viewName, functions := range panelMap {
		bindings = append(bindings, []*Binding{
			{ViewName: viewName, Key: 'k', Modifier: gocui.ModNone, Handler: functions.onKeyUpPress},
			{ViewName: viewName, Key: 'j', Modifier: gocui.ModNone, Handler: functions.onKeyDownPress},
		}...)
	}
	return bindings
}

func (gui *Gui) keybindings(g *gocui.Gui) error {
	bindings := gui.GetInitialKeybindings()
	for _, binding := range bindings {
		if err := g.SetKeybinding(binding.ViewName, binding.Key, binding.Modifier, binding.Handler); err != nil {
			return err
		}
	}

	return nil
}
