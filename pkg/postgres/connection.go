package postgres

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/spf13/viper"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

type DatabaseManager interface {
	Connect() (*sqlx.DB, error)
	GetDB() (*sqlx.DB, error)
	Close() error
}

type Database struct {
	Master   string
	Database *sqlx.DB
}

type DatabaseHandlerFunc func(dest interface{}, isExec bool, query string, values ...interface{}) error

var (
	dbOnce     sync.Once
	dbInstance *Database
	DbDetails  string
)

func NewDatabase(master string) *Database {
	return &Database{
		Master: master,
	}
}

func InitConnection() DatabaseManager {
	dbOnce.Do(func() {
		details := viper.GetString("databasePostgres.details")
		DbDetails = details
		log.Info().Msgf("postgres: Initializing connection with details: %s", details)

		db := NewDatabase(DbDetails)
		_, err := db.Connect()
		if err != nil {
			log.Error().Err(err).Msg("postgres: Failed to connect to PostgreSQL")
			return
		}

		dbInstance = db

		log.Info().Msg("postgres: PostgresSQL initialized successfully")
	})

	return dbInstance
}

func (db *Database) CreateDatabaseHandler() DatabaseHandlerFunc {
	conn, err := db.GetDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get DB instance")
	}

	return func(dest interface{}, isExec bool, query string, values ...interface{}) error {
		if isExec {
			result, err := conn.Exec(query, values...)
			if err != nil {
				log.Error().Err(err).Msg("Failed to execute query")
				return err
			}
			rowsAffected, _ := result.RowsAffected()
			if rowsAffected == 0 {
				log.Warn().Msg("No rows affected by the query")
				return fmt.Errorf("postgres: no rows affected")
			}
			return nil
		}

		return conn.Select(dest, query, values...)
	}
}

func (db *Database) Connect() (*sqlx.DB, error) {
	if db.Master == "" {
		return nil, fmt.Errorf("connection string is empty")
	}

	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithServiceName(fmt.Sprintf("%s-postgres", viper.GetString("server.name"))))

	connStr := db.Master
	conn := sqlx.MustOpen("postgres", connStr)

	if err := conn.Ping(); err != nil {
		log.Error().Err(err).Msg("Failed to connect to PostgreSQL")
		return nil, err
	}

	conn.SetMaxOpenConns(viper.GetInt("databasePostgres.maxPool"))
	conn.SetMaxIdleConns(viper.GetInt("databasePostgres.minPool"))
	conn.SetConnMaxLifetime(time.Minute * 10)
	db.Database = conn

	return db.Database, nil
}

func (db *Database) GetDB() (*sqlx.DB, error) {
	if dbInstance == nil || dbInstance.Database == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	if dbInstance.Database.Stats().OpenConnections > 40 {
		fpcs := make([]uintptr, 1)
		n := runtime.Callers(2, fpcs)
		if n != 0 {
			fun := runtime.FuncForPC(fpcs[0] - 1)
			if fun != nil {
				log.Info().Msg("postgres: DB Conn more than 40, Caller from Func: " + fun.Name())
			}
		}
		log.Info().Msgf("postgres: DB Conn more than 40, currently: %d", dbInstance.Database.Stats().OpenConnections)
	}

	return dbInstance.Database, nil
}

func (db *Database) Close() error {
	if db.Database != nil {
		return db.Database.Close()
	}
	return nil
}
