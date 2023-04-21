package main

import (
	"github.com/nurlan42/todo/config"
	"github.com/nurlan42/todo/internal/composites"
	log "github.com/sirupsen/logrus"
)

const cfgPath = "config.json"

func main() {
	cfg, err := config.New(cfgPath)
	if err != nil {
		log.Fatalf("config.New(): %v", err)
	}

	if err := composites.New(cfg); err != nil {
		log.Fatalf("NewApp(): %v", err)
	}
}
