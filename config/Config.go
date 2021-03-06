package config

import "time"

type YAMLConfig struct {
	Server struct {
		Addr string `yaml:"addr"`
	}

	Application struct {
		Name    string `yaml:"name"`
		Profile string
		Session struct {
			IgnoreUrls []string `yaml:"ignoreUrls"`
		}
		Jwt struct {
			Timeout *time.Duration `yaml:"timeout"`
			Method  string         `yaml:"method"`
			Key     string         `yaml:"key"`
			DataKey string         `yaml:"dataKey"`
		}
		Cors struct {
			AccessControlAllowOrigin      []string       `yaml:"accessControlAllowOrigin"`
			AccessControlAllowMethods     []string       `yaml:"accessControlAllowMethods"`
			AccessControlAllowHeaders     []string       `yaml:"accessControlAllowHeaders"`
			AccessControlExposeHeaders    []string       `yaml:"accessControlExposeHeaders"`
			AccessControlAllowCredentials *bool          `yaml:"accessControlAllowCredentials"`
			AccessControlMaxAge           *time.Duration `yaml:"accessControlMaxAge"`
		}
	}

	Logging struct {
		Encoding string `yaml:"encoding"`
		FileName string `yaml:"filename"`
		FilePath string `yaml:"filepath"`
	}

	Website struct {
		Host  string                 `yaml:"host"`
		Title string                 `yaml:"title"`
		Props map[string]interface{} `yaml:"props"`
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
