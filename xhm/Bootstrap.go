package xhm

import (
	"flag"
	"fmt"
	"github.com/xhminc/xhm-framework/component/database"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/config"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var (
	log                *zap.Logger
	globalConfig       *config.YAMLConfig
	applicationProfile string
)

func init() {

	flag.StringVar(&applicationProfile, "profile", "dev",
		"Please enter application profile name for loading config.")
	flag.Parse()

	bootstrap()
}

func GetGlobalConfig() *config.YAMLConfig {
	if globalConfig != nil {
		return globalConfig
	}
	globalConfig = &config.YAMLConfig{}
	return globalConfig
}

func bootstrap() {

	if applicationProfile != "dev" && applicationProfile != "test" &&
		applicationProfile != "prev" && applicationProfile != "prod" {
		panic(fmt.Errorf("profile incorrect, usage: dev | test | prev | prod"))
	}

	GetGlobalConfig()

	loadYAMLConfig("application.yml")
	loadYAMLConfig("application-" + applicationProfile + ".yml")
	globalConfig.Application.Profile = applicationProfile

	log = logger.InitLogger(globalConfig)
	log.Info("Loading yaml config finished, profiles: [application.yml, application-" + applicationProfile + ".yml]")

	database.InitDataSource(globalConfig)
	log.Info("Bootstrap framework finished !!!")
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
