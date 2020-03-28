package config

import (
	"time"
)

type Server struct {
	Port         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}
