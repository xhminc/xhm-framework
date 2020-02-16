package security

import (
	"github.com/xhminc/xhm-framework/config"
	"go.uber.org/zap"
)

var (
	log          *zap.Logger
	globalConfig *config.YAMLConfig
)
