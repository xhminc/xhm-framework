package xhm

import (
	"flag"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"github.com/xhminc/xhm-framework/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	GlobalConfig       *config.YAMLConfig
	log                *logrus.Logger
	applicationProfile string
	timestampFormat    = "2006-01-02 15:04:05.000"
)

func Bootstrap(c *config.YAMLConfig) {
	initCommandLineArguments()
	initLogConfig()
	loadConfigFile(c)
	//initDataSource()
	log.Info("Bootstrap init finished !!!")
	GlobalConfig = c
}

func loadConfigFile(c *config.YAMLConfig) {
	loadYAMLConfig("application.yml", c)
	loadYAMLConfig("application-"+applicationProfile+".yml", c)
	c.Application.Profile = applicationProfile
	log.Infof("Load YAML configure finished, profiles: [application.yml, application-%s.yml]", applicationProfile)
}

func loadYAMLConfig(filename string, c *config.YAMLConfig) {

	realPath, pathErr := filepath.Abs("resource/" + filename)

	if pathErr != nil {
		log.Fatal("Generate Absolute Path error")
		return
	}

	content, ioErr := ioutil.ReadFile(realPath)

	if ioErr != nil {
		log.Fatal("Loading config file fail")
		return
	}

	e := yaml.Unmarshal(content, c)

	if e != nil {
		log.Fatal("Parse YAML config fail")
		return
	}

}

func initCommandLineArguments() {

	flag.StringVar(&applicationProfile, "profile", "dev",
		"Please enter application profile name for loading configure.")
	flag.Parse()

	if applicationProfile != "dev" && applicationProfile != "test" &&
		applicationProfile != "prev" && applicationProfile != "prod" {
		log.Fatal("Profile incorrect")
		return
	}
}

func initLogConfig() {

	log = logrus.New()

	if applicationProfile == "dev" || applicationProfile == "test" {
		log.SetFormatter(&prefixed.TextFormatter{
			ForceColors:      true,
			ForceFormatting:  true,
			FullTimestamp:    true,
			DisableUppercase: true,
			TimestampFormat:  timestampFormat,
		})
		log.SetLevel(logrus.DebugLevel)
		log.SetOutput(os.Stdout)
	} else {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timestampFormat,
		})
		log.SetLevel(logrus.InfoLevel)
		log.SetOutput(os.Stdout)
	}

	log.Info("Logger configure init finished")
}
