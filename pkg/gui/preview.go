package gui

import "github.com/jroimartin/gocui"

type PreviewPanel struct {
	v           *gocui.View
	snippetPath string
	snippet     []byte
}

func NewPreviewPanel(v *gocui.View) (*PreviewPanel, error) {
	previewPanel := &PreviewPanel{
		v: v,
	}
	return previewPanel, nil
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
