package service

import (
	broker "github.com/alipniczkij/web-broker"
	"github.com/alipniczkij/web-broker/pkg/repository"
)

type QueueValue interface {
	Get() (string, error)
	Put(value broker.QueueValue) error
}

type Service struct {
	QueueValue
}

func NewService(queue *repository.Repository) *Service {
	return &Service{
		QueueValue: NewQueueValueService(queue.QueueValue),
	}
}
