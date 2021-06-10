package config

type AppConfig struct {
	Name        string
	Version     string
	SnippetsDir string
}

func NewAppConfig(name, version, snippetsdir, date string) (*AppConfig, error) {
	appConfig := &AppConfig{
		Name:        name,
		Version:     version,
		SnippetsDir: snippetsdir,
	}
	return appConfig, nil
}
