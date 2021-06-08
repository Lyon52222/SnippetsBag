package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Lyon52222/snippetsbag/pkg/utils"
	"github.com/jroimartin/gocui"
)

var (
	snippetsDir string
)

func main() {
	setupApp()
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func setupApp() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panicln(err)
	} else {
		snippetsDir = filepath.Join(homeDir, ".snippets")
		if !utils.IsExist(snippetsDir) {
			utils.NewDir(snippetsDir)
		}
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("collections", 0, 0, maxX/6, maxY/4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Personal Collections"
		fmt.Fprintln(v, "\uf719 All Snippets")
		fmt.Fprintln(v, "\ue7c5 Vim")
		fmt.Fprintln(v, "\ue795 Shell")
	}

	if v, err := g.SetView("folders", 0, maxY/4, maxX/6, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Folders"
		childFolders := utils.GetAllFolders(snippetsDir)
		for _, f := range childFolders {
			fmt.Fprintln(v, f)
		}
	}

	if v, err := g.SetView("snippets", maxX/6, 0, maxX/5*2, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Snippets"
		snippets := utils.GetAllSnippets(snippetsDir)
		for _, s := range snippets {
			fmt.Fprintln(v, s)
		}
	}

	if v, err := g.SetView("preview", maxX/5*2, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Preview"
		fmt.Fprintln(v, "preview")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
