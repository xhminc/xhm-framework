package generator

import (
	"github.com/jinzhu/gorm"
	"github.com/xhminc/xhm-framework/component/logger"
	"go.uber.org/zap"
	"sync"
)

type IdGenerator struct {
	lock sync.Mutex
	tags map[string]*tag
}

func InitGenerator(tags []string, database *gorm.DB) {
	db = database
	log = logger.GetLogger()
	Generator = &IdGenerator{}
	Generator.initTags(tags)
}

func (g *IdGenerator) initTags(tags []string) {
	for _, t := range tags {
		realTag := tag{}
		err := realTag.InitTag(t)
		if err != nil {
			log.Error("Init biz tag fail", zap.String("error", err.Error()))
		} else {
			log.Info("Init biz tag success", zap.String("tag", t))
		}
	}
}

func (g *IdGenerator) NextId(tag string) (int64, error) {
	return g.tags[tag].nextId()
}
