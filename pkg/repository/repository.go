package repository

import broker "github.com/alipniczkij/web-broker"

type QueueValue interface {
	Get() (string, error)
	Put(broker.QueueValue) error
}

type Repository struct {
	QueueValue
}

func NewRepository(queue *[]string) *Repository {
	return &Repository{
		QueueValue: NewQueueValueRepo(queue),
	}
}
