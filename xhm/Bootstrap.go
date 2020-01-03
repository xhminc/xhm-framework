package xhm

import (
	"flag"
	"fmt"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/config"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var (
	log          *zap.Logger
	globalConfig *config.YAMLConfig
)

func Bootstrap() {
	bootstrap()
	//initDataSource()
	log.Info("Bootstrap init finished !!!")
}

func GetGlobalConfig() *config.YAMLConfig {
	if globalConfig != nil {
		return globalConfig
	}
	globalConfig = &config.YAMLConfig{}
	return globalConfig
}

func bootstrap() {

	var applicationProfile string
	flag.StringVar(&applicationProfile, "profile", "dev",
		"Please enter application profile name for loading configure.")
	flag.Parse()

	if applicationProfile != "dev" && applicationProfile != "test" &&
		applicationProfile != "prev" && applicationProfile != "prod" {
		panic(fmt.Errorf("profile incorrect, usage: dev | test | prev | prod"))
	}

	loadYAMLConfig("application.yml")
	loadYAMLConfig("application-" + applicationProfile + ".yml")
	globalConfig.Application.Profile = applicationProfile

	log = logger.InitLogger(globalConfig)
	log.Info("Load YAML configure finished, profiles: [application.yml, application-" + applicationProfile + ".yml]")
}

func loadYAMLConfig(filename string) {

	realPath, pathErr := filepath.Abs("resource/" + filename)

	if pathErr != nil {
		panic(fmt.Errorf("generate absolute path fail, exception: %s", pathErr))
	}

	content, ioErr := ioutil.ReadFile(realPath)

	if ioErr != nil {
		panic(fmt.Errorf("loading config file fail, exception: %s", ioErr))
	}

	e := yaml.Unmarshal(content, globalConfig)

	if e != nil {
		panic(fmt.Errorf("parsing yaml config fail, exception: %s", e))
	}

}
