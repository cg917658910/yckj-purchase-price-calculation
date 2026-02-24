package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env         string `yaml:"env"`
	Addr        string `yaml:"addr"`
	DBDSN       string `yaml:"db_dsn"`
	AllowOrigin string `yaml:"allow_origin"`
}

// Load 会优先尝试从 config.yaml 读取配置,找不到或解析失败时回退到环境变量/默认值
func Load() *Config {
	// 默认值
	cfg := &Config{
		Env:         "dev",
		Addr:        ":8080",
		AllowOrigin: "*",
	}
	configFile := getEnv("CONFIG_FILE", "config.yaml")
	b, err := os.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("failed to read config file: %v", err))
	}
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		panic(fmt.Sprintf("failed to parse config file: %v", err))
	}
	return cfg
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getEnvInt64(key string, def int64) int64 {
	if v := os.Getenv(key); v != "" {
		var out int64
		_, err := fmt.Sscanf(v, "%d", &out)
		if err == nil {
			return out
		}
	}
	return def
}
