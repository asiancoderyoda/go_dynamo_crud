package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

type Config struct {
	timeout time.Duration
}

func getRouteConfig() *Config {
	return &Config{}
}

func (c *Config) Cors(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           86400, // Maximum value not ignored by any of major browsers
	}).Handler(next)
}

func (c *Config) SetTimeout(timeInSec int) *Config {
	c.timeout = time.Duration(timeInSec) * time.Second
	return c
}

func (c *Config) GetTimeout() time.Duration {
	return c.timeout
}
