package interfaces

import (
	"actualizer/internal/common/dto"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type BookingRepo interface {
	GetAll(ctx context.Context, params dto.RequestParams) ([]dto.Booking, error)
	GetById(ctx context.Context, uuid uuid.UUID) (dto.Booking, error)
	Delete(ctx context.Context, uuid uuid.UUID) error
	Update(ctx context.Context, Booking *dto.Booking, uuid uuid.UUID) error
	Create(ctx context.Context, Booking *dto.Booking) error
	UpdateExpired(ctx context.Context) error
}
