package common

import (
	"github.com/newdee/aipaper-util/config"
	"testing"
)

func TestGetWordsPackageList(t *testing.T) {
	err := config.Register(config.Common, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	wordsPackageConfig, err := GetWordsPackageList()
	if err != nil {
		t.Errorf("get words_package config failed, err:%v", err)
		return
	}
	t.Logf("words_package config:+%v", wordsPackageConfig)
}

func TestGetChatallPackageList(t *testing.T) {
	err := config.Register(config.Common, MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	chatallPackageConfig, err := GetChatallPackageList()
	if err != nil {
		t.Errorf("get chatall_package config failed, err:%v", err)
		return
	}
	t.Logf("chatall_packageconfig:+%v", chatallPackageConfig)
}
