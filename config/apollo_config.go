package config

import "github.com/apolloconfig/agollo/v4/env/config"

type ApolloConfig struct {
	config.AppConfig
}

type Option func(config *ApolloConfig)

func WithAppID(appID string) Option {
	return func(config *ApolloConfig) {
		config.AppID = appID
	}
}

func WithCluster(cluster string) Option {
	return func(config *ApolloConfig) {
		config.Cluster = cluster
	}
}

func WithIP(ip string) Option {
	return func(config *ApolloConfig) {
		config.IP = ip
	}
}

func WithNamespace(namespace string) Option {
	return func(config *ApolloConfig) {
		config.NamespaceName = namespace
	}
}

func WithIsBackupConfig(isBackupConfig bool) Option {
	return func(config *ApolloConfig) {
		config.IsBackupConfig = isBackupConfig
	}
}

func WithSecret(secret string) Option {
	return func(config *ApolloConfig) {
		config.Secret = secret
	}
}
