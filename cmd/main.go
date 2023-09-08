package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})
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

	// Repository, server, handler

	// Routes

	// Server
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
