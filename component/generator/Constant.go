package generator

import (
	"github.com/jinzhu/gorm"
	"github.com/xhminc/xhm-framework/component/logger"
)

var (
	db        *gorm.DB
	log       = logger.GetLogger()
	tableName = "xhm_segment"
	Generator *IdGenerator
)
