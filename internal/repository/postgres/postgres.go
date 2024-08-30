package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qara-qurt/booking_service/config"
)

// Const values for configuring the database connection
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

// Config creates a new pgxpool.Config from the given config.Database
func Config(conf *config.Database) *pgxpool.Config {
	URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
		conf.SSLMode,
	)

	// Create a new pgxpool.Config
	dbConfig, err := pgxpool.ParseConfig(URL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	// Set the default values
	dbConfig.MaxConns = defaultMaxConn
	dbConfig.MinConns = defaultMinConn
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	return dbConfig
}

// NewPostgres creates a new postgresDB instance
func NewPostgres(conf *config.Config) (*postgresDB, error) {
	// Create a new connection pool with the given configuration
	connPool, err := pgxpool.NewWithConfig(context.Background(), Config(&conf.Database))
	if err != nil {
		log.Fatalf("error while creating connection to the database, %v", err)
		return nil, err
	}

	// Ping the database to check if the connection is successful
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
