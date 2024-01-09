package config

import (
	commonConfig "github.com/arhamj/go-commons/pkg/config"
	sqlite "github.com/shardeum/service-validator/common/db/sqlite"
)

type Config struct {
	AppSqliteConfig sqlite.DatabaseConfig `mapstructure:"app_sqlite_config"`
	HelperFlags     HelperFlags           `mapstructure:"helper_flags"`
	FeatureFlags    FeatureFlags          `mapstructure:"feature_flags"`
	ServerConfig    ServerConfig          `mapstructure:"server_config"`
	AuthConfig      AuthConfig            `mapstructure:"auth_config"`
}

type ServerConfig struct {
}

type SqliteConfig struct {
	Path string `mapstructure:"path"`
}

type HelperFlags struct {
}

type FeatureFlags struct {
}

type AuthConfig struct {
}

func NewConfig(configFile string) (Config, error) {
	var cfg Config
	err := commonConfig.LoadConfig(configFile, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
