package connection

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Connection struct {
	*mongo.Database
	*pgxpool.Pool
}
