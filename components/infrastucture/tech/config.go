package tech

import (
	_ "embed"

	"github.com/ihatiko/olymp/temple/infrastucture/http"
	"github.com/ihatiko/olymp/temple/infrastucture/logger"
	"github.com/ihatiko/olymp/temple/infrastucture/tracer"
)

//go:embed config/tech.config.toml
var defaultConfig []byte

type Service struct {
	Name string
}

type Config struct {
	Tech struct {
		Service Service
		Log     logger.Config
		Tracer  tracer.Config
		Http    http.Config
	}
}