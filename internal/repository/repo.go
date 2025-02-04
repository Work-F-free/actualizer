package repository

import (
	"golang.org/x/net/context"
	"seatPlanner/internal/common/config"
	"seatPlanner/internal/common/connection"
)

type DBConnection interface {
	Connect(config config.DBConfig, ctx context.Context) (connection.Connection, error)
	Disconnect(ctx context.Context) error
}

type Repo struct {
}

func New(cnt connection.Connection) *Repo {
	return &Repo{}
}
