package aipaper

import "github.com/newdee/aipaper-util/config"

var MapEnvToConfig = map[config.EnvType]config.Param{
	config.DevEnv: {
		APPID:          "ai-paper",
		Cluster:        "DEV",
		IP:             "http://124.222.11.142:8080",
		Namespace:      "mix-paper.app",
		IsBackupConfig: true,
	},
	config.ProdEnv: {
		APPID:          "ai-paper",
		Cluster:        "DEV",
		IP:             "http://124.222.11.142:8080",
		Namespace:      "mix-paper.app",
		IsBackupConfig: true,
	},
}
