package main

import (
	"github.com/AggroSec/Go-Blog-Aggregator/internal/config"
	"github.com/AggroSec/Go-Blog-Aggregator/internal/database"
)

type state struct {
	conf *config.Config
	db   *database.Queries
}
