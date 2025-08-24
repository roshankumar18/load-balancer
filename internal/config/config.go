package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server       ServerConfig       `yaml:"server"`
	Backends     []BackendConfig    `yaml:"backends"`
	LoadBalancer LoadBalancerConfig `yaml:"load_balancer"`
	Health       HealthCheckConfig  `yaml:"health_check"`
}

type ServerConfig struct {
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type BackendConfig struct {
	URL string `yaml:"url"`
}

type LoadBalancerConfig struct {
	Algorithm string `yaml:"algorithm"`
}

type HealthCheckConfig struct {
	Interval time.Duration `yaml:"interval"`
	Timeout  time.Duration `yaml:"timeout"`
}

func Load(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	setDefaults(&config)

	return &config, nil
}

func setDefaults(cfg *Config) {
	if cfg.Server.Port == 0 {
		cfg.Server.Port = 8084
	}

	if cfg.Server.ReadTimeout == 0 {
		cfg.Server.ReadTimeout = 30 * time.Second
	}

	if cfg.Server.WriteTimeout == 0 {
		cfg.Server.WriteTimeout = 30 * time.Second
	}

	if cfg.LoadBalancer.Algorithm == "" {
		cfg.LoadBalancer.Algorithm = "round_robin"
	}

}
