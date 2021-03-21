package repository

// type QueueValuePostgres struct {
// 	db *sqlx.DB
// }

import (
	"fmt"

	broker "github.com/alipniczkij/web-broker"
)

type QueueValueRepo struct {
	queue *[]string
}

func NewQueueValueRepo(queue *[]string) *QueueValueRepo {
	return &QueueValueRepo{queue: queue}
}

func (q *QueueValueRepo) Get() (string, error) {
	if len(*q.queue) == 0 {
		return "", fmt.Errorf("Empty queue")
	}
	value := (*q.queue)[0]
	newQueue := (*q.queue)[1:]
	q.queue = &newQueue
	return value, nil
}

func (q *QueueValueRepo) Put(v broker.QueueValue) error {
	newQueue := append((*q.queue), v.Value)
	q.queue = &newQueue
	return nil
}
