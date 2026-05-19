package config

import (
	"context"
	"fmt"
	"os"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"sigap-sultan-be/src/common"
	"time"
)

//func StartDbConnection() *sqlx.DB {
//	db, err := sqlx.Connect("postgres", "user=postgres dbname=sigap_sultan sslmode=disable host=localhost")
//	if err != nil {
//		log.Error(err)
//	}
//
//	_, errExec := db.Exec(`SET search_path='dev'`)
//	if errExec != nil {
//		log.Error(errExec)
//	}
//
//	// Test the connection to the database
//	if err := db.Ping(); err != nil {
//		log.Fatal(err)
//	} else {
//		log.Info("Successfully Connected")
//	}
//
//	return db
//}

type DbConfig struct {
	EnvConfig *common.EnvConfig
}

func NewDbConfig(envConfig *common.EnvConfig) *DbConfig {
	return &DbConfig{EnvConfig: envConfig}
}

func (r *DbConfig) PgxPoolConfig() *pgxpool.Config {
	const defaultMaxConns = int32(100)
	const defaultMinConns = int32(10)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	// Your own Database URL
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?search_path=%s",
		r.EnvConfig.DbUsername,
		r.EnvConfig.DbPass,
		r.EnvConfig.DbHost,
		r.EnvConfig.DbPort,
		r.EnvConfig.DbName,
		r.EnvConfig.DbScheme,
	)

	dbConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	//dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
	//	log.Info("Before acquiring the connection pool to the database!!")
	//	return true
	//}
	//
	//dbConfig.AfterRelease = func(c *pgx.Conn) bool {
	//	log.Info("After releasing the connection pool to the database!!")
	//	return true
	//}
	//
	//dbConfig.BeforeClose = func(c *pgx.Conn) {
	//	log.Info("Closed the connection pool to the database!!")
	//}

	return dbConfig
}

func (r *DbConfig) StartDbConnection() *pgxpool.Pool {
	//conn, err := pgx.Connect(context.Background(), urlConnection)
	//if err != nil {
	//	log.Fatalf("Unable to connect to database: %v\n\n", err)
	//}

	// We use "pgxpool.Connect()" instead of "pgx.Connect()" because the vanilla driver is not safe
	// for concurrent connections (unlike the other Golang SQL drivers)
	// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx

	//conn, err := pgxpool.New(context.Background(), urlConnection)
	//dbConfig, errParseConfig := pgxpool.ParseConfig(urlConnection)
	conn, err := pgxpool.NewWithConfig(
		context.Background(),
		r.PgxPoolConfig(),
		//&pgxpool.Config{
		//	ConnConfig: &pgx.ConnConfig{
		//		Config: pgconn.Config{
		//			Host:     envConfig.DbHost,
		//			Port:     uint16(envConfig.DbPort),
		//			Database: envConfig.DbName,
		//			User:     envConfig.DbUsername,
		//		},
		//	},
		//	MaxConns:        100,
		//	MaxConnIdleTime: 5 * time.Minute,
		//	MaxConnLifetime: 60 * time.Minute,
		//},
	)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n\n", err)
	}

	return conn
}

func StartRedisConnection() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisClient.Ping(context.Background())

	return redisClient
}
