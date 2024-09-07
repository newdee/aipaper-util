package config

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// EnvType 环境类型
type EnvType string

const (
	// DevEnv dev环境
	DevEnv EnvType = "dev"
	// ProdEnv 正式环境
	ProdEnv EnvType = "prod"
)

type Config interface {
	Get(key string) (interface{}, error)

	GetStringValue(key string) string
	GetIntValue(key string) int

	GetIntWithDefault(key string, defaultValue int) int
	GetStringWithDefault(key string, defaultValue string) string

	GetWithUnmarshal(key string, entity interface{}, unmarshaler Unmarshaler) error
}

type Unmarshaler interface {
	// Unmarshal deserializes the data bytes into value parameter.
	Unmarshal(data []byte, value interface{}) error
}

// JSONUnmarshaler is json unmarshaler.
type JSONUnmarshaler struct{}

// Unmarshal deserializes the data bytes into parameter val in json protocol.
func (ju *JSONUnmarshaler) Unmarshal(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}

// YamlUnmarshaler is yaml unmarshaler.
type YamlUnmarshaler struct{}

// Unmarshal deserializes the data bytes into parameter val in yaml protocol.
func (yu *YamlUnmarshaler) Unmarshal(data []byte, val interface{}) error {
	return yaml.Unmarshal(data, val)
}
