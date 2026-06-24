package config

import (
	"fmt"
	"os"

	driver "github.com/go-sql-driver/mysql"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server ServerConfig `toml:"server"`
	MySQL  MySQLConfig  `toml:"mysql"`
	Redis  RedisConfig  `toml:"redis"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	Mode string `toml:"mode"`
}

type MySQLConfig struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	Username     string `toml:"username"`
	Password     string `toml:"password"`
	Database     string `toml:"database"`
	Charset      string `toml:"charset"`
	ParseTime    bool   `toml:"parse_time"`
	Loc          string `toml:"loc"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxOpenConns int    `toml:"max_open_conns"`
}

type RedisConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Database int    `toml:"database"`
}

func Load(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	cfg := defaultConfig()
	if err := toml.Unmarshal(content, cfg); err != nil {
		return nil, fmt.Errorf("parse toml config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if c.Server.Host == "" {
		return fmt.Errorf("server.host is required")
	}
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return fmt.Errorf("server.port must be between 1 and 65535")
	}
	if c.Server.Mode == "" {
		return fmt.Errorf("server.mode is required")
	}
	if c.MySQL.Host == "" {
		return fmt.Errorf("mysql.host is required")
	}
	if c.MySQL.Port <= 0 || c.MySQL.Port > 65535 {
		return fmt.Errorf("mysql.port must be between 1 and 65535")
	}
	if c.MySQL.Username == "" {
		return fmt.Errorf("mysql.username is required")
	}
	if c.MySQL.Database == "" {
		return fmt.Errorf("mysql.database is required")
	}
	if c.MySQL.Charset == "" {
		return fmt.Errorf("mysql.charset is required")
	}
	if c.MySQL.Loc == "" {
		return fmt.Errorf("mysql.loc is required")
	}
	if c.MySQL.MaxIdleConns < 0 {
		return fmt.Errorf("mysql.max_idle_conns cannot be negative")
	}
	if c.MySQL.MaxOpenConns < 0 {
		return fmt.Errorf("mysql.max_open_conns cannot be negative")
	}
	if c.Redis.Host == "" {
		return fmt.Errorf("redis.host is required")
	}
	if c.Redis.Port <= 0 || c.Redis.Port > 65535 {
		return fmt.Errorf("redis.port must be between 1 and 65535")
	}
	if c.Redis.Database < 0 {
		return fmt.Errorf("redis.database cannot be negative")
	}
	return nil
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

func (c *Config) MySQLDSN() string {
	return (&driver.Config{
		User:                 c.MySQL.Username,
		Passwd:               c.MySQL.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", c.MySQL.Host, c.MySQL.Port),
		DBName:               c.MySQL.Database,
		AllowNativePasswords: true,
		ParseTime:            c.MySQL.ParseTime,
		Params: map[string]string{
			"charset": c.MySQL.Charset,
			"loc":     c.MySQL.Loc,
		},
	}).FormatDSN()
}

func (c *Config) RedisAddr() string {
	return fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port)
}

func defaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host: "0.0.0.0",
			Port: 8080,
			Mode: "debug",
		},
		MySQL: MySQLConfig{
			Host:         "127.0.0.1",
			Port:         3306,
			Username:     "root",
			Password:     "root",
			Database:     "gomind",
			Charset:      "utf8mb4",
			ParseTime:    true,
			Loc:          "Local",
			MaxIdleConns: 10,
			MaxOpenConns: 100,
		},
		Redis: RedisConfig{
			Host:     "127.0.0.1",
			Port:     6379,
			Password: "",
			Database: 0,
		},
	}
}
