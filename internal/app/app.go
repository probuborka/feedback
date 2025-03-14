package app

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/probuborka/feedback/internal/config"
	handlers "github.com/probuborka/feedback/internal/controller/http"
	"github.com/probuborka/feedback/internal/usecase/feedback"
	"github.com/probuborka/feedback/pkg/route"
	"github.com/sirupsen/logrus"
)

func Run() {
	//-------------------------------------- config
	cfg, err := config.New()
	if err != nil {
		logrus.Error(err)
		return
	}

	//-------------------------------------- log
	//logrus
	log := logrus.New()

	//format log json
	log.SetFormatter(&logrus.JSONFormatter{})

	//saving logs to file
	file, err := os.OpenFile(cfg.Log.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("log file error")
		return
	}
	log.SetOutput(io.MultiWriter(os.Stdout, file))

	//-------------------------------------- client

	//-------------------------------------- infrastructure

	//-------------------------------------- usecase
	useCaseFeedback := feedback.NewFeedbackUseCase()

	//-------------------------------------- handlers
	handlers := handlers.New(
		log,
		useCaseFeedback,
	)

	//-------------------------------------- server
	server := route.New(
		cfg.HTTP.Port,
		handlers.Init(),
	)

	//start server
	log.WithFields(logrus.Fields{
		"service": "nutrial",
		"version": "1.0.0",
		"port":    cfg.HTTP.Port,
	}).Info("Server run")
	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("error occurred while running http server")
		}
	}()

	//stop server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.WithFields(logrus.Fields{
		"service": "nutrial",
		"version": "1.0.0",
		"port":    cfg.HTTP.Port,
	}).Info("server stop")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to stop server")
	}
}
