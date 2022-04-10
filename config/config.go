package config

import (
	"strconv"

	"example.com/m/v2/utils/env"
)

type AppConfig struct {
	Port         int    `json:"port"`
	Timeout      int    `json:"timeout"`
	Dialect      string `json:"dialect"`
	dbConnString string `json:"db_conn_string"`
}

func GetAppConfig() AppConfig {
	return AppConfig{
		Port:         parseEnvToInt("PORT", "8080"),
		Timeout:      parseEnvToInt("TIMEOUT", "60"),
		Dialect:      env.GetEnv("DIALECT", "sqlite3"),
		dbConnString: env.GetEnv("DB_CONN_STRING", ":memory:"),
	}
}

func parseEnvToInt(envName string, defaultValue string) int {
	envValue := env.GetEnv(envName, defaultValue)
	intValue, err := strconv.Atoi(envValue)
	if err != nil {
		return 0
	}
	return intValue
}
