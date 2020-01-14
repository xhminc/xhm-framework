package xhm

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"github.com/xhminc/xhm-framework/component/database"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/config"
	"go.uber.org/zap"
)

var (
	log                *zap.Logger
	globalConfig       *config.YAMLConfig
	v                  *viper.Viper
	applicationProfile string
)

func init() {

	flag.StringVar(&applicationProfile, "profile", "dev",
		"Please enter application profile name for loading config.")
	flag.Parse()

	bootstrap()
}

func bootstrap() {

	if applicationProfile != "dev" && applicationProfile != "test" &&
		applicationProfile != "prev" && applicationProfile != "prod" {
		panic(fmt.Errorf("profile incorrect, usage: dev | test | prev | prod"))
	}

	globalConfig = config.GetGlobalConfig()

	initViper()
	loadYAMLConfig("application.yml")
	loadYAMLConfig("application-" + applicationProfile + ".yml")
	globalConfig.Application.Profile = applicationProfile

	log = logger.InitLogger()
	database.InitDataSource()

	log.Info("Bootstrap framework finished !!!")
}

func initViper() {
	v = viper.New()
	v.AddConfigPath("./resource/")
	v.SetConfigType("yaml")
}

func loadYAMLConfig(filename string) {

	v.SetConfigName(filename)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Loading config file fail, exception: %s \n", err))
	}

	err = v.Unmarshal(globalConfig)
	if err != nil {
		panic(fmt.Errorf("Extraing config file fail, exception: %s \n", err))
	}

}

//func loadYAMLConfig(filename string) {
//
//	realPath, pathErr := filepath.Abs("resource/" + filename)
//
//	if pathErr != nil {
//		panic(fmt.Errorf("generate absolute path fail, exception: %s", pathErr))
//	}
//
//	content, ioErr := ioutil.ReadFile(realPath)
//
//	if ioErr != nil {
//		panic(fmt.Errorf("loading config file fail, exception: %s", ioErr))
//	}
//
//	e := yaml.Unmarshal(content, globalConfig)
//
//	if e != nil {
//		panic(fmt.Errorf("parsing yaml config fail, exception: %s", e))
//	}
//
//}
