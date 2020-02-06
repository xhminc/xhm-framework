package generator

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	db        *gorm.DB
	log       *zap.Logger
	tableName = "xhm_segment"
	Generator *IdGenerator
)
