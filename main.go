package main

import (
	"fmt"
	"github.com/nicitapa/firstProgect/internal/configs"
	"github.com/nicitapa/firstProgect/internal/controller"
	"github.com/nicitapa/firstProgect/internal/repository"
	"github.com/nicitapa/firstProgect/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"os"
)

// @title OnlineShop API
// @contact.name OnlineShop API Service
// @contact.url http://test.com
// @contact.email test@test.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().Msg("Starting up application...")

	if err := configs.ReadSettings(); err != nil {
		logger.Error().Err(err).Msg("Error during reading settings")
		return
	}
	logger.Info().Msg("Read settings successfully")

	dsn := fmt.Sprintf(`host=%s 
							port=%s 
							user=%s 
							password=%s 
							dbname=%s 
							sslmode=disable`,
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		os.Getenv("POSTGRES_PASSWORD"),
		configs.AppSettings.PostgresParams.Database,
	)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		logger.Error().Err(err).Msg("Error during connection to database")
	}
	logger.Info().Msg("Database connected successfully")

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", configs.AppSettings.RedisParams.Host, configs.AppSettings.RedisParams.Port),
		DB:   configs.AppSettings.RedisParams.Database,
	})

	cache := repository.NewCache(rdb)
	logger.Info().Msg("Redis connected successfully")

	repo := repository.NewRepository(db)
	svc := service.NewService(repo, cache)
	ctrl := controller.NewController(svc)

	if err = ctrl.RunServer(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun)); err != nil {
		logger.Error().Err(err).Msg("Error during running http-server")
	}

	if err = db.Close(); err != nil {
		logger.Error().Err(err).Msg("Error during closing database connection")
	}
}
