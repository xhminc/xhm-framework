package config

type YAMLConfig struct {
	Server struct {
		Port string `yaml:"port"`
	}

	Application struct {
		Name    string `yaml:"name"`
		Profile string
	}

	Website struct {
		Host  string `yaml:"host"`
		Title string `yaml:"title"`
	}
}
