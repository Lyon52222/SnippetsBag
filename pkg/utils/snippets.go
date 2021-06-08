package utils

import "path"

func GetAllSnippets(configDir string) []string {
	var allSnippets []string
	for _, f := range GetChildFolders(configDir) {
		for _, file := range GetChildFiles(f) {
			_, snippet := path.Split(file)
			allSnippets = append(allSnippets, snippet)
		}
	}
	return allSnippets
}

func GetAllFolders(configDir string) []string {
	var allFolders []string
	for _, f := range GetChildFolders(configDir) {
		_, file := path.Split(f)
		allFolders = append(allFolders, file)
	}
	return allFolders
}
