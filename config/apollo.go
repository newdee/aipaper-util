package config

import (
	"fmt"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

type Param struct {
	APPID          string `json:"appid"`
	Cluster        string `json:"cluster"`
	IP             string `json:"ip"`
	Namespace      string `json:"namespace"`
	IsBackupConfig bool   `json:"is_backup_config"`
	Secret         string `json:"secret"`
}

func NewApolloClient(opts ...Option) (*ApolloClient, error) {
	conf := ApolloConfig{AppConfig: config.AppConfig{}}
	for _, opt := range opts {
		opt(&conf)
	}

	cli, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return &conf.AppConfig, nil
	})
	cli.AddChangeListener(&CustomChangeListener{})

	apolloCache, err := getApolloCache(conf.NamespaceName, cli)
	if err != nil {
		return nil, err
	}

	return &ApolloClient{Cache: apolloCache}, err
}

func getApolloCache(namespace string, cli agollo.Client) (agcache.CacheInterface, error) {
	configCache := cli.GetConfigCache(namespace)
	if configCache == nil {
		return nil, fmt.Errorf("get config cache failed, namespace:%v", namespace)
	}
	return configCache, nil
}

type ApolloClient struct {
	Cache agcache.CacheInterface
}

func (c *ApolloClient) GetIntWithDefault(key string, defaultValue int) int {
	val, err := c.Cache.Get(key)
	if err != nil {
		return defaultValue
	}

	if intVal, ok := val.(int); ok {
		return intVal
	}
	return defaultValue
}

func (c *ApolloClient) GetStringWithDefault(key string, defaultValue string) string {
	val, err := c.Cache.Get(key)
	if err != nil {
		return defaultValue
	}

	if strVal, ok := val.(string); ok {
		return strVal
	}
	return defaultValue
}

func (c *ApolloClient) Get(key string) (interface{}, error) {
	val, err := c.Cache.Get(key)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (c *ApolloClient) GetWithUnmarshal(key string, entity interface{}, unmarshaler Unmarshaler) error {
	val, err := c.Cache.Get(key)
	if err != nil {
		return err
	}

	return unmarshaler.Unmarshal([]byte(val.(string)), entity)
}

func (c *ApolloClient) GetStringValue(key string) string {
	value, err := c.Cache.Get(key)
	if err != nil {
		return ""
	}

	return value.(string)
}

func (c *ApolloClient) GetIntValue(key string) int {
	value, err := c.Cache.Get(key)
	if err != nil {
		return 0
	}

	return value.(int)
}

func (c *ApolloClient) GetFloatValue(key string) float64 {
	return 0
}

func (c *ApolloClient) GetBoolValue(key string) bool {
	return false
}

type CustomChangeListener struct {
}

// OnChange
// 触发时机:	配置变动时，客户端启动时不会触发
// 行为:		只获取变动配置的键值对
func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
}

// OnNewestChange
// 触发时机:	客户端启动时以及配置变动时
// 行为:		获取最新的该namespace下的全部配置
func (c *CustomChangeListener) OnNewestChange(newestChangeEvent *storage.FullChangeEvent) {

}
