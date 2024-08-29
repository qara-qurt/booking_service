package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qara-qurt/booking_service/config"
)

const (
	defaultMaxConn           = int32(90)
	defaultMinConn           = int32(0)
	defaultMaxConnLifetime   = time.Hour
	defaultMaxConnIdleTime   = time.Minute * 30
	defaultHealthCheckPeriod = time.Minute
	defaultConnectTimeout    = time.Second * 5
)

type postgresDB struct {
	DB *pgxpool.Pool
}

func Config(conf *config.Database) *pgxpool.Config {
	URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
		conf.SSLMode,
	)

	dbConfig, err := pgxpool.ParseConfig(URL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConn
	dbConfig.MinConns = defaultMinConn
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	return dbConfig
}

func NewPostgres(conf *config.Config) (*postgresDB, error) {
	connPool, err := pgxpool.NewWithConfig(context.Background(), Config(&conf.Database))
	if err != nil {
		log.Fatalf("error while creating connection to the database, %v", err)
		return nil, err
	}

	err = connPool.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database")
		return nil, err
	}

	log.Println("Connected to the database!!")

	return &postgresDB{
		DB: connPool,
	}, nil
}
