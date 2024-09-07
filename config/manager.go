package config

import (
	"fmt"
	"sync"
)

var manager *Manager

type Manager struct {
	mutex          sync.RWMutex
	mapKeyToConfig map[KeyType]Config
}

func init() {
	manager = &Manager{
		mapKeyToConfig: make(map[KeyType]Config),
	}
}

type KeyType string

const (
	Common KeyType = "common"
	Login  KeyType = "login"
)

func Register(key KeyType, mapEnvToConfig map[EnvType]Param, env EnvType) error {
	param, ok := mapEnvToConfig[env]
	if !ok {
		return fmt.Errorf("invalid env:%v", env)
	}

	var cfg Config
	var err error
	cfg, err = NewApolloClient(
		WithAppID(param.APPID),
		WithCluster(param.Cluster),
		WithIP(param.IP),
		WithNamespace(param.Namespace),
		WithIsBackupConfig(param.IsBackupConfig),
		WithSecret(param.Secret),
	)

	if err != nil {
		return err
	}

	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	manager.mapKeyToConfig[key] = cfg
	return nil
}

func Get(key KeyType) (Config, error) {
	cfg, ok := manager.mapKeyToConfig[key]
	if !ok {
		return nil, fmt.Errorf("key:%v not exist", key)
	}

	return cfg, nil
}
