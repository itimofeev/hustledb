package util

import (
	"github.com/kelseyhightower/envconfig"
)

type configEnv struct {
	DbEnv
	AppEnv
}

// Config interface
type Config interface {
	Db() *DbEnv
	App() *AppEnv
}

// DbEnv db settings
type DbEnv struct {
	URL          string `envconfig:"HUSTLESA_DB_URL"`
	MaxIdleConns int    `envconfig:"HUSTLESA_DB_MAX_IDLE_CONNS" default:"4"`
	MaxOpenConns int    `envconfig:"HUSTLESA_DB_MAX_OPEN_CONNS" default:"16"`
	StrictMode   bool   `envconfig:"HUSTLESA_DB_STRICT_MODE" default:"false"`
}
type AppEnv struct {
	JsonFilesPath string `envconfig:"HUSTLESA_APP_JSON_FILES_PATH"`
}

// Db settings
func (env configEnv) Db() *DbEnv {
	return &env.DbEnv
}
func (env configEnv) App() *AppEnv {
	return &env.AppEnv
}

// ReadConfig read config from environment
func ReadConfig() Config {
	var c configEnv
	err := envconfig.Process("hustlesa", &c)

	CheckErr(err, "read envconfig")

	return &c
}
