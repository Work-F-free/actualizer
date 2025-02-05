package repository

import (
	"actualizer/internal/common/config"
	"actualizer/internal/common/connection"
	"actualizer/internal/repository/interfaces"
	"actualizer/internal/repository/postgres"
	"golang.org/x/net/context"
)

type DBConnection interface {
	Connect(config config.DBConfig, ctx context.Context) (connection.Connection, error)
	Disconnect(ctx context.Context) error
}

type Repo struct {
	interfaces.BookingRepo
}

func New(cnt connection.Connection) *Repo {
	return &Repo{
		BookingRepo: postgres.NewBookingRepo(cnt.Pool),
	}
}
