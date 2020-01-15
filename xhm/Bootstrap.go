package xhm

import (
	"flag"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/xhminc/xhm-framework/component/database"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/config"
	"github.com/xhminc/xhm-framework/util/tripledes"
	"go.uber.org/zap"
	"os"
	"reflect"
	"regexp"
)

var (
	log                *zap.Logger
	globalConfig       *config.YAMLConfig
	v                  *viper.Viper
	applicationProfile string
	systemEnvRegex, _  = regexp.Compile("\\$\\{([^\\$\\{\\}]+)\\}")
	tripleDESRegex, _  = regexp.Compile("^ENC\\(.*\\)$")
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
		panic(fmt.Errorf("Loading config file fail, exception: %s\n", err))
	}

	err = v.Unmarshal(globalConfig, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			systemEnvHookFunc(),
			tripleDESHookFunc(),
		)
	})

	if err != nil {
		panic(fmt.Errorf("Extraing config file fail, exception: %s\n", err))
	}

}

func systemEnvHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Kind, t reflect.Kind, data interface{}) (interface{}, error) {

		if f != reflect.String {
			return data, nil
		}

		d := data.(string)
		env := systemEnvRegex.ReplaceAllStringFunc(d, func(s string) string {
			return os.Getenv(s[2 : len(s)-1])
		})

		return env, nil
	}
}

func tripleDESHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Kind, t reflect.Kind, data interface{}) (interface{}, error) {

		if f != reflect.String {
			return data, nil
		}

		d := data.(string)
		env := tripleDESRegex.ReplaceAllStringFunc(d, func(s string) string {

			key := os.Getenv(config.DES_KEY)
			iv := os.Getenv(config.DES_IV)
			decryptString, err := tripledes.DecryptToString(d[4:len(d)-1], key, []byte(iv)...)

			if err != nil {
				panic(fmt.Errorf("3DES decrypt ENC() config fail, exception: %s\n", err))
			}

			return decryptString
		})

		return env, nil
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
