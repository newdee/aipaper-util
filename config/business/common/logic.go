package common

import (
	"github.com/newdee/aipaper-util/config"
	"github.com/newdee/aipaper-util/database"
	oaconfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

func getDBConfig(key string) (*database.DBConfig, error) {
	var dbConfig database.DBConfig
	err := getConfig(key, &dbConfig)
	if err != nil {
		return nil, err
	}
	return &dbConfig, nil
}

func GetMysqlConfig() (*database.DBConfig, error) {
	return getDBConfig("mysql")
}

func GetRedisConfig() (*database.DBConfig, error) {
	return getDBConfig("redis")
}

type MsgQueueConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Vhost    string `json:"vhost"`
}

func getConfig(key string, conf interface{}) error {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return err
	}

	err = cfg.GetWithUnmarshal(key, conf, &config.JSONUnmarshaler{})
	if err != nil {
		return err
	}
	return nil
}

func GetMsgQueueConfig() (*MsgQueueConfig, error) {
	var conf MsgQueueConfig
	err := getConfig("msg_queue", &conf)
	if err != nil {
		return nil, err
	}
	return &conf, err
}

// Param 短信参数
type Param struct {
	Signature  string `json:"signature"`
	TemplateID string `json:"template_id"`
	TTL        string `json:"ttl"`
}

func GetSMSConfig() (*Param, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var param Param
	err = cfg.GetWithUnmarshal("sms", &param, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return &param, nil
}

func GetOfficialAccountConfig() (*oaconfig.Config, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}

	oaCfg := &oaconfig.Config{}
	err = cfg.GetWithUnmarshal("wechat_official_account", oaCfg, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return oaCfg, nil
}
