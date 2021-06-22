package data

import (
	"io/ioutil"
	"os"

	"github.com/Lyon52222/snippetsbag/pkg/config"
	"github.com/Lyon52222/snippetsbag/pkg/utils"
)

type DataLoader struct {
	Config *config.AppConfig
}

func NewData(config *config.AppConfig) (*DataLoader, error) {
	data := &DataLoader{
		Config: config,
	}
	return data, nil
}

func (data *DataLoader) GetAllSnippets() []string {
	var allSnippets []string
	for _, f := range utils.GetChildFolders(data.Config.SnippetsDir) {
		for _, file := range utils.GetChildFiles(f) {
			//_, snippet := path.Split(file)
			allSnippets = append(allSnippets, file)
		}
	}
	return allSnippets
}

func (data *DataLoader) GetSnippetsFromPath(path string) []string {
	var snippets []string
	for _, file := range utils.GetChildFiles(path) {
		snippets = append(snippets, file)
	}
	return snippets
}

func (data *DataLoader) GetAllFolders() []string {
	var allFolders []string
	for _, f := range utils.GetChildFolders(data.Config.SnippetsDir) {
		//_, file := path.Split(f)
		allFolders = append(allFolders, f)
	}
	return allFolders
}

func (data *DataLoader) ReadSnippet(snippetPath string) ([]byte, error) {
	f, err := os.Open(snippetPath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func (data *DataLoader) CreateNewFolder(name string) error {
	return utils.NewDir(name)
}

func (data *DataLoader) CreateNewSnippet(name string) error {
	return utils.NewFile(name)
}
