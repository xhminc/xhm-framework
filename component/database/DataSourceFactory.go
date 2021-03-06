package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/config"
	"go.uber.org/zap"
	t "time"
)

var (
	log          *zap.Logger
	dbMap        map[string]*gorm.DB
	globalConfig *config.YAMLConfig
)

func InitDataSource() {

	var myLogger MyLogger
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

			db.SetLogger(&myLogger)
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

func CloseDB() {
	for k, v := range dbMap {
		err := v.Close()
		if err != nil {
			log.Error(err.Error())
		} else {
			log.Info("Closing db success", zap.String("name", k))
		}
	}
}

type MyLogger struct {
}

func (logger *MyLogger) Print(values ...interface{}) {
	var (
		level  = values[0]
		source = values[1]
	)
	if level == "sql" {
		sql := values[3].(string)
		cost := values[2].(t.Duration)
		params := values[4]
		log.Info(
			sql,
			zap.Any("level", level),
			zap.Any("source", source),
			zap.Duration("cost", cost),
			zap.Any("params", params))
	} else if level == "log" {
		log.Error(
			"",
			zap.Any("level", level),
			zap.Any("source", source),
			zap.Any("message", values[2]))
	} else {
		log.Info("", zap.Any("values", values))
	}
}
