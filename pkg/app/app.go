package app

import "github.com/Lyon52222/snippetsbag/pkg/gui"

type App struct {
	Gui *gui.Gui
}

func NewApp() (*App, error) {
	app := &App{}
	var err error
	app.Gui, err = gui.NewGui()
	if err != nil {
		return app, err
	}
	return app, nil
}

func (app *App) Run() error {
	err := app.Gui.Run()
	return err
}
