package main

import (
	"fmt"
	"runtime"

	"github.com/Lyon52222/snippetsbag/pkg/app"
	"github.com/Lyon52222/snippetsbag/pkg/config"
	"github.com/integrii/flaggy"
)

var (
	snippetsDir = "/Users/admin/.snippets"
	version     = "unversioned"
	date        string
)

func main() {
	info := fmt.Sprintf(
		"%s\nData: %s\n SnippetsDir: %s\nOS: %s\n",
		version,
		date,
		snippetsDir,
		runtime.GOOS,
	)
	flaggy.SetName("SnippetsBag")
	flaggy.SetDescription("Manage your code snippets")
	flaggy.SetVersion(info)

	flaggy.Parse()

	appConfig, err := config.NewAppConfig("SnippetsBag", version, snippetsDir, date)

	app, err := app.NewApp(appConfig)
	if err == nil {
		err = app.Run()
	}
}
