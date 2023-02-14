package config

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type DBConfig struct {
	User     string
	Password string
	Port     int
	Host     string
	Name     string
}

type RestConfig struct {
	Port string
}

func getConfigPath() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	return path.Join(filepath.Dir(basePath), "..", "config", "config.yaml")
}

func NewDBConfig(logger *zap.Logger) *DBConfig {
	cfgPath := getConfigPath()

	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("read db config fail", zap.Error(err))
	}
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	port := viper.GetInt("db.port")
	host := viper.GetString("db.host")

	return &DBConfig{
		User:     user,
		Password: password,
		Port:     port,
		Host:     host,
	}
}

func NewRestConfig(logger *zap.Logger) *RestConfig {
	cfgPath := getConfigPath()

	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("read rest config fail", zap.Error(err))
	}

	port := viper.GetString("rest.port")

	return &RestConfig{
		Port: port,
	}
}

func (cfg *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name)
}
