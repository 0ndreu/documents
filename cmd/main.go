package main

import (
	"context"
	"fmt"
	"github.com/0ndreu/documents/internal/server"
	service2 "github.com/0ndreu/documents/internal/service"
	"net/http"
	"os"
	"time"

	"github.com/0ndreu/documents/internal/storage"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

type configs struct {
	Port          string        `envconfig:"PORT" default:":8081"`
	LogLevel      string        `envconfig:"LOG_LEVEL" default:"debug"`
	AuthSecret    string        `envconfig:"AUTH_SECRET" default:"secret"`
	UserDBConnStr string        `envconfig:"USER_DB_CONN_STR"  required:"false"`
	DBConnStr     string        `envconfig:"DB_CONN_STR"  required:"false"`
	DBTimeout     time.Duration `envconfig:"DB_TIMEOUT" required:"false"`
	RTimeout      time.Duration `envconfig:"READ_TIMEOUT" required:"false"`
	WTimeout      time.Duration `envconfig:"WRITE_TIMEOUT" required:"false"`
}

func main() {
	log := zerolog.New(os.Stdout).With().Logger()

	var cfg configs
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal().Msgf("read config err: %s", err.Error())
		os.Exit(1)
	}

	connection, err := storage.NewConnection(cfg.DBConnStr)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	userConnection, err := storage.NewConnection(cfg.UserDBConnStr)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	storage := storage.NewStorage(userConnection, connection, cfg.DBTimeout)
	service := service2.NewService(storage)

	ctx := context.Background()
	r := server.NewServer(ctx, service, log)

	api := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		ReadTimeout:  cfg.RTimeout,
		Handler:      r,
		WriteTimeout: cfg.WTimeout,
	}

	go func() {
		log.Info().Msgf("Listen to API on %d port", cfg.Port)
		err := api.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("Listen to API: %v", err)
		}
	}()

}
