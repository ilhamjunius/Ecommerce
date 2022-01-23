package configs

import (
	"os"
	"sync"
)

type AppConfig struct {
	Driver   string `yaml:"driver"`
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	DB_Port  int    `yaml:"portdb"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Name = getEnv("NAME", "ecommerce")
	defaultConfig.Address = getEnv("ADDRESS", "dbbe5.cozcj7dbvsdr.ap-southeast-1.rds.amazonaws.com")
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = getEnv("USERNAME", "admin")
	defaultConfig.Password = getEnv("PASSWORD", "admin123")

	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback

}
