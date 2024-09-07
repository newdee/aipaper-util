package common

import "github.com/newdee/aipaper-util/config"

var MapEnvToConfig = map[config.EnvType]config.Param{
	config.DevEnv: {
		APPID:          "common",
		Cluster:        "DEV",
		IP:             "http://124.222.11.142:8080",
		Namespace:      "mix-paper.dev",
		IsBackupConfig: true,
	},
	config.ProdEnv: {
		APPID:          "common",
		Cluster:        "DEV",
		IP:             "http://124.222.11.142:8080",
		Namespace:      "mix-paper.prod",
		IsBackupConfig: true,
	},
}
