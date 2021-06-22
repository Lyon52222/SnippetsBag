package gui

import (
	"fmt"
	"path"

	"github.com/Lyon52222/snippetsbag/pkg/data"
	"github.com/jroimartin/gocui"
)

type FoldersPanel struct {
	v            *gocui.View
	folders      []string
	dataloader   *data.DataLoader
	snippetPanel *SnipeetsPanel
}

func NewFoldersPanel(v *gocui.View, dataloader *data.DataLoader, snippetPanel *SnipeetsPanel) (*FoldersPanel, error) {
	foldersPanel := &FoldersPanel{
		v:            v,
		dataloader:   dataloader,
		snippetPanel: snippetPanel,
	}
	foldersPanel.folders = append(foldersPanel.folders, foldersPanel.dataloader.GetAllFolders()...)
	return foldersPanel, nil
}

func (f *FoldersPanel) ShowFolders() {
	for _, folder := range f.folders {
		_, name := path.Split(folder)
		fmt.Fprintln(f.v, name)
	}
	f.setCursorY(0)
}

func (f *FoldersPanel) setCursorY(y int) error {
	cx, _ := f.v.Cursor()
	ox, _ := f.v.Origin()
	if err := f.v.SetCursor(cx, 0); err != nil {
		if err := f.v.SetOrigin(ox, 0); err != nil {
			return err
		}
	}
	if err := f.snippetPanel.Refresh(f.GetCurrentFolder()); err != nil {
		return err
	}
	return nil
}

func (f *FoldersPanel) AddFolders(folders []string) {
	f.folders = append(f.folders, folders...)
	for _, folder := range folders {
		_, name := path.Split(folder)
		fmt.Fprintln(f.v, name)
	}
}

func (f *FoldersPanel) AddFolder(folder string) {
	f.folders = append(f.folders, folder)
	_, name := path.Split(folder)
	fmt.Fprintln(f.v, name)
}

func (f *FoldersPanel) GetCurrentFolder() string {
	_, cy := f.v.Cursor()
	return f.folders[cy]
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
	if err := f.snippetPanel.Refresh(f.GetCurrentFolder()); err != nil {
		return err
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
	if err := f.snippetPanel.Refresh(f.GetCurrentFolder()); err != nil {
		return err
	}
	return nil
}
