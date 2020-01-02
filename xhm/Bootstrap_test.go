package xhm

import (
	"github.com/xhminc/xhm-framework/config"
	"testing"
)

func TestBootstrap(t *testing.T) {
	var CommonConfig config.YAMLConfig
	Bootstrap(&CommonConfig)
}
