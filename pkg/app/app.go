package app

import (
	"github.com/Lyon52222/snippetsbag/pkg/config"
	"github.com/Lyon52222/snippetsbag/pkg/gui"
	"github.com/Lyon52222/snippetsbag/pkg/utils"
)

type App struct {
	Gui    *gui.Gui
	Config *config.AppConfig
}

func NewApp(config *config.AppConfig) (*App, error) {
	utils.SetupApp(config.SnippetsDir)
	app := &App{
		Config: config,
	}
	var err error
	app.Gui, err = gui.NewGui(config)
	if err != nil {
		return app, err
	}
	return app, nil
}

func (app *App) Run() error {
	err := app.Gui.Run()
	return err
}
