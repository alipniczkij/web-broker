package repository

import broker "github.com/alipniczkij/web-broker"

type Queue interface {
	Get(*broker.GetValue) (string, error)
	Put(*broker.PutValue) error
}

type Repository struct {
	Queue
}

func NewRepository(fileName string) *Repository {
	return &Repository{
		Queue: NewQueueRepo(fileName),
	}
}
