package postgres

import (
	"actualizer/internal/common/config"
	"actualizer/internal/common/connection"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type Connection struct {
	db *pgxpool.Pool
}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Connect(config config.DBConfig, ctx context.Context) (connection.Connection, error) {
	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName, config.DBSslMode))
	if err != nil {
		logrus.Fatalf("unable to create connection pool: %v\n", err)
		return connection.Connection{}, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		logrus.Fatalf("unable to ping DB: %v\n", err)
		return connection.Connection{}, err
	}

	return connection.Connection{Pool: pool}, nil
}

func (c *Connection) Disconnect(ctx context.Context) error {
	c.db.Close()
	return nil
}
