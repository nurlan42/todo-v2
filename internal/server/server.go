package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nurlan42/todo/config"

	log "github.com/sirupsen/logrus"
)

func Run(httpCfg config.HTTP, router *gin.Engine) error {
	httpSrv := http.Server{
		Addr:         ":" + httpCfg.Port,
		Handler:      router,
		ReadTimeout:  httpCfg.ReadTimeout * time.Second,
		WriteTimeout: httpCfg.WriteTimeout * time.Second,
	}

	ch := make(chan os.Signal, 1)

	go func() {
		err := httpSrv.ListenAndServe()

		fmt.Printf("\nserver is being shutdown\n")

		if !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("failed to stop server: %s\n", err.Error())
			ch <- os.Kill
		}
	}()

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to stop srv.Shutdown(): %s\n", err.Error())
	}

	return nil
}
