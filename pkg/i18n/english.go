package i18n

type TranslationSet struct {
	No                        string
	Yes                       string
	ErrorTitle                string
	CreateNewFolderPanelTitle string
}

func englishSet() TranslationSet {
	return TranslationSet{
		No:                        "No",
		Yes:                       "Yes",
		ErrorTitle:                "Error",
		CreateNewFolderPanelTitle: "Folder's Name",
	}
}
