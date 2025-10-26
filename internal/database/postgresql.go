package database

import (
	"NextShortLink/internal/config"
	"NextShortLink/internal/logger"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yxlimo/xormzap"
	"xorm.io/xorm"
)

var E *xorm.Engine

// InitDB inits global database connection
func InitDB() {
	cfg := DBConfig{
		Host:     config.C.GetString("database.host"),
		Port:     config.C.GetInt("database.port"),
		User:     config.C.GetString("database.user"),
		Name:     config.C.GetString("database.name"),
		Password: config.C.GetString("database.password"),
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	var err error
	E, err = xorm.NewEngine("postgres", dsn)
	if err != nil {
		logger.L.Fatal(err.Error())
	}

	E.SetLogger(xormzap.Logger(logger.L))

	if config.C.GetBool("debug") {
		E.ShowSQL(true)
	}

	err = E.Ping()
	if err != nil {
		logger.L.Fatal(err.Error())
	}

	logger.L.Debug("PostgresSQL initialized")
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// Close is a public method to close a database session
func Close(databaseSession *xorm.Session) {
	_ = databaseSession.Close()
}

// Rollback is a public method to roll back a transaction
func Rollback(session *xorm.Session) {
	err := session.Rollback()
	if err != nil {
		logger.L.Fatal(err.Error())
	}
}
