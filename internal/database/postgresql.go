package database

import (
	"NextShortLink/internal/config"
	"NextShortLink/internal/logger"
	"fmt"

	"github.com/ghinknet/json"
	"github.com/ghinknet/xormzap"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

var E *xorm.Engine

// InitDB inits global database connection
func InitDB() {
	xorm.SetDefaultJSONHandler(JSON{})

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Get().Database.User,
		config.Get().Database.Password,
		config.Get().Database.Host,
		config.Get().Database.Port,
		config.Get().Database.Name,
	)

	var err error
	E, err = xorm.NewEngine("postgres", dsn)
	if err != nil {
		logger.L.Fatal("failed to init database", zap.Error(err), zap.String("dsn", dsn))
	}

	E.SetLogger(xormzap.Logger(logger.L))

	if config.Debug {
		E.ShowSQL(true)
	}

	if err = E.Ping(); err != nil {
		logger.L.Fatal("failed to ping database", zap.Error(err))
	}

	logger.L.Debug("PostgresSQL initialized")
}

type JSON struct{}

func (JSON) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (JSON) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// Close is a public method to close a database session
func Close(databaseSession *xorm.Session) {
	_ = databaseSession.Close()
}

// Rollback is a public method to roll back a transaction
func Rollback(session *xorm.Session) {
	if err := session.Rollback(); err != nil {
		logger.L.Fatal("failed to rollback transaction", zap.Error(err))
	}
}

// RollbackError is a public method to roll back a transaction with errors returned
func RollbackError(session *xorm.Session, err error) error {
	Rollback(session)
	return err
}
