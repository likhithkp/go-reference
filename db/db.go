package db

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/likhithkp/practice/config"
)

var (
	DB   *pgxpool.Pool
	once sync.Once
)

func ConnectDb() {
	once.Do(func() {
		url := config.Config()

		config, err := pgxpool.ParseConfig(url)
		if err != nil {
			log.Fatalf("Error while parsing the url: %v", err)
			return
		}

		config.MaxConns = 10
		config.MinConns = 2
		config.HealthCheckPeriod = 1 * time.Second

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("Error while getting new config instance: %v", err)
			return
		}

		DB = pool
		log.Println("Connected to DB: students")
	})
}

func GetDb() *pgxpool.Pool {
	if DB == nil {
		ConnectDb()
	}
	return DB
}

func CloseDb() {
	if DB != nil {
		DB.Close()
		log.Println("DB: students closed")
	}
}
