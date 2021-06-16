package gui

import (
	"github.com/Lyon52222/snippetsbag/pkg/data"
	"github.com/jroimartin/gocui"
)

type PreviewPanel struct {
	v           *gocui.View
	snippetPath string
	snippet     []byte
	dataloader  *data.DataLoader
}

func NewPreviewPanel(v *gocui.View, dataloader *data.DataLoader) (*PreviewPanel, error) {
	previewPanel := &PreviewPanel{
		v:          v,
		dataloader: dataloader,
	}
	return previewPanel, nil
}

func (p *PreviewPanel) Refresh(file string) error {
	p.snippetPath = file
	var err error
	p.snippet, err = p.dataloader.ReadSnippet(file)
	if err != nil {
		return err
	}
	p.ShowSnippet()
	return nil
}

func (p *PreviewPanel) SetSnippetPath(path string) {
	p.snippetPath = path
}

func (p *PreviewPanel) SetSnippet(snippet []byte) {
	p.snippet = snippet
}

func (p *PreviewPanel) ShowSnippet() {
	p.v.Clear()
	p.v.Write(p.snippet)
}

func (p *PreviewPanel) ResetSnippet(path string, snippet []byte) {
	p.SetSnippetPath(path)
	p.SetSnippet(snippet)
	p.ShowSnippet()
}
