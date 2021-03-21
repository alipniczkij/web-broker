package service

import (
	broker "github.com/alipniczkij/web-broker"
	"github.com/alipniczkij/web-broker/pkg/repository"
)

type QueueValueService struct {
	queue repository.QueueValue
}

func NewQueueValueService(queue repository.QueueValue) *QueueValueService {
	return &QueueValueService{queue: queue}
}

func (s *QueueValueService) Get() (string, error) {
	return s.queue.Get()
}

func (s *QueueValueService) Put(v broker.QueueValue) error {
	return s.queue.Put(v)
}
