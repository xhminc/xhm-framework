package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/config"
	"go.uber.org/zap"
)

var (
	log          *zap.Logger
	dbMap        map[string]*gorm.DB
	globalConfig *config.YAMLConfig
)

func InitDataSource() {

	globalConfig = config.GetGlobalConfig()
	log = logger.GetLogger()

	if globalConfig.DB == nil || len(globalConfig.DB) == 0 {
		panic(fmt.Errorf("Data source config not exists\n"))
	}

	dbMap = map[string]*gorm.DB{}
	for k, v := range globalConfig.DB {

		url := fmt.Sprintf("%s:%s@(%s:%d)/%s?"+
			"charset=%s&parseTime=%s&loc=%s&timeout=%s&readTimeout=%s&writeTimeout=%s&rejectReadOnly=%s",
			v.Username,
			v.Password,
			v.Host,
			v.Port,
			v.DbName,
			v.Charset,
			v.ParseTime,
			v.Loc,
			v.Timeout,
			v.ReadTimeout,
			v.WriteTimeout,
			v.RejectReadOnly,
		)

		if db, err := gorm.Open(v.DriverName, url); db != nil && err == nil {

			if globalConfig.Application.Profile == "dev" || globalConfig.Application.Profile == "test" {
				db.LogMode(true)
			} else {
				db.LogMode(false)
			}

			db.DB().SetMaxIdleConns(v.MaxIdleConnections)
			db.DB().SetMaxOpenConns(v.MaxOpenConnections)
			db.DB().SetConnMaxLifetime(v.ConnectionMaxLifetime)
			dbMap[k] = db

		} else {
			panic(fmt.Errorf("loading data source fail, exception: %s", err))
		}
	}

	log.Info("Loading data source config success")
}

func GetDB(dbname string) *gorm.DB {
	return dbMap[dbname]
}
