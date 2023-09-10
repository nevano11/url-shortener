package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
)

// @title           Url shortener api
// @version         1.0
// @description     implementation of a test task for Altcraft company

// @host      localhost:8080
// @BasePath  /
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Info("Starting url-shortener app")

	// Read config
	if err := initConfig(); err != nil {
		logrus.Fatalf("Failed to read config: %s", err.Error())
	}

	// Configure logger
	if err := configureLogger(viper.GetString("logger.log-level")); err != nil {
		logrus.Fatalf("Failed to configure logger: %s", err.Error())
	}

	// Create context
	ctx := context.Background()

	// Database connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     ":" + viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.entryTime"),
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		logrus.Fatalf("Failed to ping database: %s", err.Error())
	}

	// Hasher
	hashProvider := service.NewMd5Hasher()

	// Repository, service, handler
	rep := repository.NewRedisRepository(rdb, &ctx)
	//rep := repository.NewFakeKVRepository()
	ser := service.NewService(service.NewUrlShortenerService(rep, hashProvider))
	han := handler.NewHandler(ser)

	// Routes
	routes := han.InitRoutes()

	// Server
	server := createServer(viper.GetString("server.port"), routes)
	logrus.Infof("Server running on http://localhost%s", server.Addr)
	logrus.Infof("Swagger: http://localhost%s/swagger/index.html", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		logrus.Fatalf("Failed to start server: %s", err.Error())
	}
}

func createServer(port string, routes *gin.Engine) *http.Server {
	return &http.Server{
		Addr:              ":" + port,
		Handler:           routes,
		ReadHeaderTimeout: 2 << 20,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
}

// Configuration
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func configureLogger(logLevel string) error {
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	return nil
}
