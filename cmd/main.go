package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nurlan42/todo/config"
	v1 "github.com/nurlan42/todo/internal/adapter/delivery/http/v1"
	"github.com/nurlan42/todo/internal/server"
	"github.com/nurlan42/todo/pkg/db"
	log "github.com/sirupsen/logrus"
)

const cfgPath = "config.json"

func main() {
	cfg, err := config.New(cfgPath)
	if err != nil {
		log.Fatalf("config.New(): %v", err)
	}

	sqlDB, err := db.Connect(cfg.DB.Psql)
	defer sqlDB.Close()

	if err != nil {
		//return err
	}

	router := gin.New()

	v1.NewHTTPRoutes(sqlDB, router)

	if err := server.Run(cfg.HTTP, router); err != nil {
		//return err
	}

}
