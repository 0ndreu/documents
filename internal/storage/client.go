package storage

import (
	"context"
	"errors"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

var (
	// errNilDriver is returned by call any method which use connection to database if connection is nil
	errNilDriver = errors.New("nil pgsql driver")
)

// Client interface
type Client interface {
	GetClient() (db *pgxpool.Pool, err error)
}

func (c Connection) GetClient() (db *pgxpool.Pool, err error) {
	if c.DB == nil {
		err = errNilDriver
		return
	}
	db = c.DB
	return
}

// newDatabase return pgsql connection
func newDatabase(conn string) (*pgxpool.Pool, error) {
	return pgxpool.Connect(context.Background(), conn)
}

// NewConnection pgsql connection
func NewConnection(conn string) (*Connection, error) {
	poll, err := newDatabase(conn)
	if err != nil {
		return nil, err
	}

	poll.Config().MaxConns = 15
	poll.Config().MaxConnIdleTime = 3 * time.Second
	poll.Config().HealthCheckPeriod = 3 * time.Second

	connection := Connection{
		DB: poll,
	}

	return &connection, nil
}

// Connection ...
type Connection struct {
	DB *pgxpool.Pool
}
