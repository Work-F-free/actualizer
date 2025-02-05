package postgres

import (
	"actualizer/internal/common/dto"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

const bookingTable = ""

type BookingRepo struct {
	db *pgxpool.Pool
	qb sq.StatementBuilderType
}

func NewBookingRepo(db *pgxpool.Pool) *BookingRepo {
	return &BookingRepo{
		db: db,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (s *BookingRepo) GetAll(ctx context.Context, params dto.RequestParams) ([]dto.Booking, error) {
	andExpr := s.AddAndExpr(params.Query)
	offset := params.Pagination["limit"] * (params.Pagination["page"] - 1)
	query, args, err := s.qb.Select("*").
		From(bookingTable).Where(andExpr).
		Limit(params.Pagination["limit"]).
		Offset(offset).ToSql()

	if err != nil {
		logrus.Fatalf("error while getting songs data from db: %s", err.Error())
		return []dto.Booking{}, err
	}

	row, err := s.db.Query(ctx, query, args...)

	booking, err := pgx.CollectRows(row, func(row pgx.CollectableRow) (dto.Booking, error) {
		var booking dto.Booking
		err = row.Scan(
			&booking.SeatId,
			&booking.Expired,
			&booking.FromDateTime,
			&booking.ToDateTime,
		)
		if err != nil {
			return dto.Booking{}, fmt.Errorf("error while getting entities from db: %w", err)
		}

		return booking, nil
	})

	if err != nil {
		logrus.Fatalf("error while getting song entities from db: %s", err.Error())
		return nil, err
	}

	return booking, nil
}

func (s *BookingRepo) GetById(ctx context.Context, uuid uuid.UUID) (dto.Booking, error) {
	query, args, err := s.qb.Select("*").From(bookingTable).Where(sq.Eq{"id": uuid}).ToSql()

	if err != nil {
		logrus.Fatalf("error while getting song data from db: %s", err.Error())
		return dto.Booking{}, err
	}

	row, err := s.db.Query(ctx, query, args...)

	booking, err := pgx.CollectOneRow(row, func(row pgx.CollectableRow) (dto.Booking, error) {
		var booking dto.Booking
		err = row.Scan(
			&booking.SeatId,
			&booking.Expired,
			&booking.FromDateTime,
			&booking.ToDateTime,
		)
		if err != nil {
			return dto.Booking{}, fmt.Errorf("error while getting entity from db: %w", err)
		}

		return booking, nil
	})

	if err != nil {
		logrus.Fatalf("error while getting song entity from db: %s", err.Error())
		return dto.Booking{}, err
	}

	return booking, nil
}

func (s *BookingRepo) Delete(ctx context.Context, uuid uuid.UUID) error {
	query, args, err := s.qb.Delete(bookingTable).Where(sq.Eq{"seat_id": uuid}).ToSql()

	if err != nil {
		logrus.Fatalf("error while creatinsg delete query: %s", err.Error())
		return fmt.Errorf("build delete query error: %w", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		logrus.Fatalf("error while deleting song entity: %s", err.Error())
		return fmt.Errorf("delete song error: %w", err)
	}

	return nil
}

func (s *BookingRepo) Update(ctx context.Context, Booking *dto.Booking, uuid uuid.UUID) error {
	updateQuery := s.qb.Update(bookingTable).
		SetMap(map[string]any{
			"seat_id":       Booking.SeatId,
			"expired":       Booking.Expired,
			"from_datetime": Booking.FromDateTime,
			"to_datetime":   Booking.ToDateTime,
		}).
		Where(sq.Eq{"seat_id": uuid})

	sqlQuery, arguments, err := updateQuery.ToSql()
	if err != nil {
		logrus.Fatalf("error while creating update entity: %s", err.Error())
		return fmt.Errorf("build query error: %w", err)
	}

	_, err = s.db.Exec(ctx, sqlQuery, arguments...)
	if err != nil {
		logrus.Fatalf("error while update song entity: %s", err.Error())
		return fmt.Errorf("update song error: %w", err)
	}
	return nil
}

func (s *BookingRepo) Create(ctx context.Context, Booking *dto.Booking) error {
	res := s.qb.Insert(bookingTable).
		Columns("seat_id", "expired", "from_datetime", "to_datetime").
		Values(Booking.SeatId, Booking.Expired, Booking.FromDateTime, Booking.ToDateTime)

	query, args, err := res.ToSql()
	if err != nil {
		logrus.Fatalf("error while creating insert query db: %s", err.Error())
		return fmt.Errorf("build insert query error: %w", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		logrus.Fatalf("error while inserting in db: %s", err.Error())
		return fmt.Errorf("insert song error: %w", err)
	}

	return nil
}

func (s *BookingRepo) UpdateExpired(ctx context.Context) error {
	res := s.qb.Update(bookingTable).Set("expired", true).Where(sq.LtOrEq{"to_datetime": time.Now()})
	query, args, err := res.ToSql()
	_, err = s.db.Exec(ctx, query, args...)

	if err != nil {
		logrus.Fatalf("error while inserting in db: %s", err.Error())
		return fmt.Errorf("insert song error: %w", err)
	}

	return nil
}

func (s *BookingRepo) AddAndExpr(params map[string]string) sq.And {
	andExpr := make(sq.And, 0, len(params))
	for k, v := range params {
		if v != "" {
			andExpr = append(andExpr, sq.Eq{k: v})
		}
	}

	return andExpr
}
