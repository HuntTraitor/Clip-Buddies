package config

import (
	"flag"
	"os"
	"time"
)

type Config struct {
	Port int
	Env  string
	DB   DBConfig

	Limiter LimiterConfig
}

type DBConfig struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}

type LimiterConfig struct {
	RPS        float64
	Burst      int
	Enabled    bool
	Expiration time.Duration
}

func Load() Config {
	var cfg Config

	databaseURL := os.Getenv("DATABASE_URL")

	flag.IntVar(&cfg.Port, "port", 3000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.DB.DSN, "db-dsn", databaseURL, "PostgreSQL DSN")
	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")

	flag.Float64Var(&cfg.Limiter.RPS, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.Limiter.Burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.Limiter.Enabled, "limiter-enabled", true, "Enable rate limiter")
	flag.DurationVar(&cfg.Limiter.Expiration, "limiter-expiration", 3*time.Minute, "Set limiter expiration")

	flag.Parse()

	return cfg
}
