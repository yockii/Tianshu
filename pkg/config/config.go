package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const DefaultConfigPath = "config.toml"

// Config 全局配置结构体
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	Tenant   TenantConfig
	MQTT     MQTTConfig
	Dji      DjiConfig
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type JWTConfig struct {
	Secret      string `mapstructure:"secret"`
	ExpireHours int    `mapstructure:"expire_hours"`
}

type TenantConfig struct {
	DefaultLogo  string `mapstructure:"default_logo"`
	DefaultTheme string `mapstructure:"default_theme"`
}

// MQTTConfig holds settings for embedded broker and DJI Cloud MQTT connection
type MQTTConfig struct {
	UseEmbedded bool   `mapstructure:"use_embedded"`
	ListenAddr  string `mapstructure:"listen_addr"`
	BrokerURL   string `mapstructure:"broker_url"`
	ClientID    string `mapstructure:"client_id"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
}

type DjiConfig struct {
	AppId      string `mapstructure:"app_id"`
	AppKey     string `mapstructure:"app_key"`
	AppLicense string `mapstructure:"app_license"`
}

var Cfg *Config

// InitConfig 读取配置文件
func InitConfig(configPath string) error {
	if configPath == "" {
		configPath = DefaultConfigPath
	}
	viper.SetConfigFile(configPath)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config failed: %w", err)
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return fmt.Errorf("unmarshal config failed: %w", err)
	}
	Cfg = &c
	return nil
}
