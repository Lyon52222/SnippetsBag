package app

import (
	"github.com/Lyon52222/snippetsbag/pkg/config"
	"github.com/Lyon52222/snippetsbag/pkg/gui"
	"github.com/Lyon52222/snippetsbag/pkg/i18n"
	"github.com/Lyon52222/snippetsbag/pkg/log"
	"github.com/Lyon52222/snippetsbag/pkg/utils"
	"github.com/sirupsen/logrus"
)

type App struct {
	Gui    *gui.Gui
	Config *config.AppConfig
	Log    *logrus.Entry
	Tr     *i18n.TranslationSet
}

func NewApp(config *config.AppConfig) (*App, error) {
	utils.SetupApp(config.SnippetsDir)
	app := &App{
		Config: config,
	}
	app.Log = log.NewLogger(config)
	app.Tr = i18n.NewTranslationSet(app.Log)
	var err error
	app.Gui, err = gui.NewGui(app.Log, config, app.Tr)
	if err != nil {
		return app, err
	}
	return app, nil
}

func (app *App) Run() error {
	err := app.Gui.Run()
	return err
}
