package xhm

import (
	"github.com/xhminc/xhm-framework/component/config"
	"testing"
)

func TestBootstrap(t *testing.T) {

	var GlobalConfig config.YAMLConfig

	Bootstrap(&GlobalConfig)
}
