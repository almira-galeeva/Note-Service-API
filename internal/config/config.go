package config

import (
	"encoding/json"
	"net"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbPassEscSeg = "{password}"
	password     = "note-password"
)

// HTTP ...
type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// GRPC ...
type GRPC struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// DB ...
type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max_open_connections"`
}

// Config ...
type Config struct {
	HTTP HTTP `json:"http"`
	GRPC GRPC `json:"grpc"`
	DB   DB   `json:"db"`
}

// NewConfig ...
func NewConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (g *GRPC) GetAddress() string {
	return net.JoinHostPort(g.Host, g.Port)
}

func (g *HTTP) GetAddress() string {
	return net.JoinHostPort(g.Host, g.Port)
}

// GetDBConfig ...
func (c *Config) GetDBConfig() (*pgxpool.Config, error) {
	dbDsn := strings.ReplaceAll(c.DB.DSN, dbPassEscSeg, password)

	poolConfig, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		return nil, err
	}
	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.DB.MaxOpenConnections

	return poolConfig, nil
}
