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
}

type ServerConfig struct {
	Port                 int `mapstructure:"port"`
	UserPortalPort       int `mapstructure:"user_portal_port"`
	AdminPortalPort      int `mapstructure:"admin_portal_port"`
	CockpitDashboardPort int `mapstructure:"cockpit_dashboard_port"`
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
