package service

import (
	"actualizer/internal/repository/interfaces"
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type SchedulerService struct {
	repo    interfaces.BookingRepo
	started bool
}

func NewScheduler(repo interfaces.BookingRepo) *SchedulerService {
	return &SchedulerService{repo: repo}
}

func (s *SchedulerService) Start() error {

	if s.started {
		return errors.New("scheduler already started")
	}

	s.started = true
	go s.Process()

	return nil
}

func (s *SchedulerService) Stop() error {
	if !s.started {
		return errors.New("scheduler not started")
	}

	s.started = false
	logrus.Info("Scheduler stopped")
	return nil
}

func (s *SchedulerService) Process() error {
	ctx := context.Background()
	logrus.Info("Scheduler started")

	for s.started {
		time.Sleep(60 * time.Second)
		err := s.repo.UpdateExpired(ctx)
		if err != nil {
			logrus.Fatalf("error while updating booking data")
			return err
		}
	}

	return nil
}
