package common

import (
	"testing"

	"github.com/newdee/aipaper-util/config"
)

func TestGetMysqlConfig(t *testing.T) {
	err := config.Register(config.Common, MapEnvToConfig, config.ProdEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	mysqlConfig, err := GetMysqlConfig()
	if err != nil {
		t.Errorf("get mysql config failed, err:%v", err)
		return
	}
	t.Logf("mysql config:+%v", mysqlConfig)
}

func TestGetRedidConfig(t *testing.T) {
	err := config.Register(config.Common, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	cfg, err := GetRedisConfig()
	if err != nil {
		t.Errorf("get redis config failed, err:%v", err)
		return
	}
	t.Logf("redis config:+%v", cfg)
}

func TestGetSMSConfig(t *testing.T) {
	err := config.Register(config.Common, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	cfg, err := GetSMSConfig()
	if err != nil {
		t.Errorf("get sms config failed, err:%v", err)
		return
	}
	t.Logf("sms config:+%v", cfg)
}

func TestGetAudioCloneRouteConfig(t *testing.T) {
	err := config.Register(config.Common, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	cfg, err := GetAudioCloneRouteConfig()
	if err != nil {
		t.Errorf("get sms config failed, err:%v", err)
		return
	}
	t.Logf("audio config config:+%v", cfg)
}
