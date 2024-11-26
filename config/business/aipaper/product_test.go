package aipaper

import (
	"testing"

	"github.com/newdee/aipaper-util/config"
)

func TestGetProductList(t *testing.T) {
	err := config.Register(config.AIPaper, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	cfg, err := GetProductList()
	if err != nil {
		t.Errorf("get product config failed, err:%v", err)
		return
	}
	t.Logf("product config:+%v", cfg)
}
