package repository

import "github.com/alipniczkij/web-broker/internal/models"

type Queue interface {
	Get(*models.GetValue) (string, error)
	Put(*models.PutValue) error
}

type Repository struct {
	Queue
}

func NewRepository(fileName string) *Repository {
	return &Repository{
		Queue: NewQueueRepo(fileName),
	}
}
