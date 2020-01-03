package config

type YAMLConfig struct {
	Server struct {
		Port string `yaml:"port"`
	}

	Application struct {
		Name    string `yaml:"name"`
		Profile string
	}

	Logging struct {
		FileName string `yaml:"filename"`
		FilePath string `yaml:"filepath"`
	}

	Website struct {
		Host  string `yaml:"host"`
		Title string `yaml:"title"`
	}
}
