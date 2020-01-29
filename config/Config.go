package config

import "time"

type YAMLConfig struct {
	Server struct {
		Addr string `yaml:"addr"`
	}

	Application struct {
		Name    string `yaml:"name"`
		Profile string
		Cors    struct {
			AccessControlAllowOrigin      []string       `yaml:"access-control-allow-origin"`
			AccessControlAllowMethods     []string       `yaml:"access-control-allow-methods"`
			AccessControlAllowHeaders     []string       `yaml:"access-control-allow-headers"`
			AccessControlExposeHeaders    []string       `yaml:"access-control-expose-headers"`
			AccessControlAllowCredentials *bool          `yaml:"access-control-allow-credentials"`
			AccessControlMaxAge           *time.Duration `yaml:"access-control-max-age"`
		}
	}

	Logging struct {
		Encoding string `yaml:"encoding"`
		FileName string `yaml:"filename"`
		FilePath string `yaml:"filepath"`
	}

	Website struct {
		Host  string `yaml:"host"`
		Title string `yaml:"title"`
	}

	DB map[string]struct {
		DriverName            string        `yaml:"driverName"`
		Host                  string        `yaml:"host"`
		Port                  uint16        `yaml:"port"`
		Username              string        `yaml:"username"`
		Password              string        `yaml:"password"`
		DbName                string        `yaml:"dbname"`
		Charset               string        `yaml:"charset"`
		ParseTime             string        `yaml:"parseTime"`
		Loc                   string        `yaml:"loc"`
		Timeout               string        `yaml:"timeout"`
		ReadTimeout           string        `yaml:"readTimeout"`
		WriteTimeout          string        `yaml:"writeTimeout"`
		RejectReadOnly        string        `yaml:"rejectReadOnly"`
		MaxIdleConnections    int           `yaml:"maxIdleConnections"`
		MaxOpenConnections    int           `yaml:"maxOpenConnections"`
		ConnectionMaxLifetime time.Duration `yaml:"connectionMaxLifetime"`
	}
}

var (
	globalConfig *YAMLConfig
)

func GetGlobalConfig() *YAMLConfig {
	if globalConfig != nil {
		return globalConfig
	}
	globalConfig = &YAMLConfig{}
	return globalConfig
}

func (config *YAMLConfig) IsDevelop() bool {
	return !config.IsProduct()
}

func (config *YAMLConfig) IsProduct() bool {
	if config.Application.Profile == "prev" || config.Application.Profile == "prod" {
		return true
	} else {
		return false
	}
}
