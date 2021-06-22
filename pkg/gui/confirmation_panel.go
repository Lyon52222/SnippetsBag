package gui

import (
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
)

type createPopupPanelOpts struct {
	hasLoader           bool
	editable            bool
	title               string
	prompt              string
	handleConfirm       func() error
	handleConfirmPrompt func(string) error
	handleClose         func() error

	// when handlersManageFocus is true, do not return from the confirmation context automatically. It's expected that the handlers will manage focus, whether that means switching to another context, or manually returning the context.
	handlersManageFocus bool
}

type askOpts struct {
	title               string
	prompt              string
	handleConfirm       func() error
	handleClose         func() error
	handlersManageFocus bool
}

type promptOpts struct {
	title          string
	initialContent string
	handleConfirm  func(string) error
}

//func (gui *Gui) ask(opts askOpts) error {
//return gui.createPopupPanel(createPopupPanelOpts{
//title:               opts.title,
//prompt:              opts.prompt,
//handleConfirm:       opts.handleConfirm,
//handleClose:         opts.handleClose,
//handlersManageFocus: opts.handlersManageFocus,
//findSuggestionsFunc: opts.findSuggestionsFunc,
//})
//}

//func (gui *Gui) prompt(opts promptOpts) error {
//return gui.createPopupPanel(createPopupPanelOpts{
//title:               opts.title,
//prompt:              opts.initialContent,
//editable:            true,
//handleConfirmPrompt: opts.handleConfirm,
//findSuggestionsFunc: opts.findSuggestionsFunc,
//})
//}

func (gui *Gui) wrappedConfirmationFunction(function func(*gocui.Gui, *gocui.View) error) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if function != nil {
			if err := function(g, v); err != nil {
				return err
			}
		}
		return gui.closeConfirmationPrompt(g)
	}
}

func (gui *Gui) closeConfirmationPrompt(g *gocui.Gui) error {
	view, err := g.View(CONFIRMATION_PANEL)
	if err != nil {
		return nil // if it's already been closed we can just return
	}
	if err := gui.returnFocus(g, view); err != nil {
		panic(err)
	}
	g.DeleteKeybindings(CONFIRMATION_PANEL)
	return g.DeleteView(CONFIRMATION_PANEL)
}

func (gui *Gui) getMessageHeight(wrap bool, message string, width int) int {
	lines := strings.Split(message, "\n")
	lineCount := 0
	// if we need to wrap, calculate height to fit content within view's width
	if wrap {
		for _, line := range lines {
			lineCount += len(line)/width + 1
		}
	} else {
		lineCount = len(lines)
	}
	return lineCount
}

func (gui *Gui) getConfirmationPanelDimensions(g *gocui.Gui, wrap bool, prompt string) (int, int, int, int) {
	width, height := g.Size()
	panelWidth := width / 2
	panelHeight := gui.getMessageHeight(wrap, prompt, panelWidth)
	return width/2 - panelWidth/2,
		height/2 - panelHeight/2 - panelHeight%2 - 1,
		width/2 + panelWidth/2,
		height/2 + panelHeight/2
}

func (gui *Gui) createPromptPanel(g *gocui.Gui, currentView *gocui.View, title string, handleConfirm func(*gocui.Gui, *gocui.View) error) error {
	confirmationView, err := gui.prepareConfirmationPanel(currentView, title, "")
	if err != nil {
		return err
	}
	confirmationView.Editable = true
	return gui.setKeyBindings(g, handleConfirm, nil)
}

func (gui *Gui) prepareConfirmationPanel(currentView *gocui.View, title, prompt string) (*gocui.View, error) {
	x0, y0, x1, y1 := gui.getConfirmationPanelDimensions(gui.g, true, prompt)
	confirmationView, err := gui.g.SetView(CONFIRMATION_PANEL, x0, y0, x1, y1)
	if err != nil {
		if err.Error() != "unknown view" {
			return nil, err
		}
		confirmationView.Title = title
		confirmationView.Wrap = true
		confirmationView.FgColor = gocui.ColorWhite
	}
	gui.g.Update(func(g *gocui.Gui) error {
		return gui.switchFocus(gui.g, currentView, confirmationView, false)
	})
	return confirmationView, nil
}

// it is very important that within this function we never include the original prompt in any error messages, because it may contain e.g. a user password
func (gui *Gui) createConfirmationPanel(g *gocui.Gui, currentView *gocui.View, title, prompt string, handleConfirm, handleClose func(*gocui.Gui, *gocui.View) error) error {
	return gui.createPopupPanel(g, currentView, title, prompt, handleConfirm, handleClose)
}

func (gui *Gui) createPopupPanel(g *gocui.Gui, currentView *gocui.View, title, prompt string, handleConfirm, handleClose func(*gocui.Gui, *gocui.View) error) error {
	g.Update(func(g *gocui.Gui) error {
		// delete the existing confirmation panel if it exists
		if view, _ := g.View(CONFIRMATION_PANEL); view != nil {
			if err := gui.closeConfirmationPrompt(g); err != nil {
				gui.Log.Error(err.Error())
			}
		}
		confirmationView, err := gui.prepareConfirmationPanel(currentView, title, prompt)
		if err != nil {
			return err
		}
		confirmationView.Editable = false
		if err := gui.renderString(g, CONFIRMATION_PANEL, prompt); err != nil {
			return err
		}
		return gui.setKeyBindings(g, handleConfirm, handleClose)
	})
	return nil
}

func (gui *Gui) setKeyBindings(g *gocui.Gui, handleConfirm, handleClose func(*gocui.Gui, *gocui.View) error) error {
	// would use a loop here but because the function takes an interface{} and slices of interfaces require even more boilerplate
	if err := g.SetKeybinding(CONFIRMATION_PANEL, gocui.KeyEnter, gocui.ModNone, gui.wrappedConfirmationFunction(handleConfirm)); err != nil {
		return err
	}

	if err := g.SetKeybinding(CONFIRMATION_PANEL, gocui.KeyCtrlC, gocui.ModNone, gui.wrappedConfirmationFunction(handleClose)); err != nil {
		return err
	}

	return nil
}

// createSpecificErrorPanel allows you to create an error popup, specifying the
//  view to be focused when the user closes the popup, and a boolean specifying
// whether we will log the error. If the message may include a user password,
// this function is to be used over the more generic createErrorPanel, with
// willLog set to false
func (gui *Gui) createSpecificErrorPanel(message string, nextView *gocui.View, willLog bool) error {
	if willLog {
		go func() {
			// when reporting is switched on this log call sometimes introduces
			// a delay on the error panel popping up. Here I'm adding a second wait
			// so that the error is logged while the user is reading the error message
			time.Sleep(time.Second)
			gui.Log.Error(message)
		}()
	}

	colorFunction := color.New(color.FgRed).SprintFunc()
	coloredMessage := colorFunction(strings.TrimSpace(message))
	return gui.createConfirmationPanel(gui.g, nextView, gui.Tr.ErrorTitle, coloredMessage, nil, nil)
}

func (gui *Gui) createErrorPanel(g *gocui.Gui, message string) error {
	return gui.createSpecificErrorPanel(message, g.CurrentView(), true)
}

func (gui *Gui) renderConfirmationOptions() error {
	optionsMap := map[string]string{
		"n/esc":   gui.Tr.No,
		"y/enter": gui.Tr.Yes,
	}
	return gui.renderOptionsMap(optionsMap)
}
