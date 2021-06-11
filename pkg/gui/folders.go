package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type FoldersPanel struct {
	v       *gocui.View
	folders []string
}

func NewFoldersPanel(v *gocui.View) (*FoldersPanel, error) {
	foldersPanel := &FoldersPanel{
		v: v,
	}
	return foldersPanel, nil
}

func (f *FoldersPanel) ShowFolders() {
	for _, folder := range f.folders {
		fmt.Fprintln(f.v, folder)
	}
}

func (f *FoldersPanel) AddFolders(folders []string) {
	f.folders = append(f.folders, folders...)
}
func (f *FoldersPanel) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy+1 < len(f.folders) {
			if err := v.SetCursor(cx, cy+1); err != nil {
				ox, oy := v.Origin()
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}

		}
	}
	return nil
}

func (f *FoldersPanel) cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}
