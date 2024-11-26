package common

import (
	"github.com/newdee/aipaper-util/config"
	"testing"
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

func TestGetCategoryList(t *testing.T) {
	err := config.Register(config.Common, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	cfg, err := GetCategoryList()
	if err != nil {
		t.Errorf("get category config failed, err:%v", err)
		return
	}
	t.Logf("category config:+%v", cfg)
}

func TestGetFeatureList(t *testing.T) {
	err := config.Register(config.Common, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	cfg, err := GetFeatureList()
	if err != nil {
		t.Errorf("get feature config failed, err:%v", err)
		return
	}
	t.Logf("feature config:+%v", cfg)
}

func TestGetSubjectList(t *testing.T) {
	err := config.Register(config.Common, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	cfg, err := GetSubjectList()
	if err != nil {
		t.Errorf("get subject config failed, err:%v", err)
		return
	}
	t.Logf("subject config:+%v", cfg)
}

func TestGetMsgQueueConfig(t *testing.T) {
	err := config.Register(config.Common, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	cfg, err := GetMsgQueueConfig()
	if err != nil {
		t.Errorf("get redis config failed, err:%v", err)
		return
	}
	t.Logf("msg queue config:+%v", cfg)
}
