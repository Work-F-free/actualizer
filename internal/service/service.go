package service

import (
	"actualizer/internal/repository"
)

type Service struct {
	Scheduler
}

type Scheduler interface {
	Start() error
	Stop() error
	Process() error
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		Scheduler: NewScheduler(repo),
	}
}
