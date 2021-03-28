package repository

import (
	"sync"

	"github.com/alipniczkij/web-broker/internal/models"
	"github.com/alipniczkij/web-broker/tools"
)

type QueueRepo struct {
	fileName string
}

func NewQueueRepo(fileName string) *QueueRepo {
	return &QueueRepo{fileName: fileName}
}

var m sync.Mutex

func (q *QueueRepo) Get(getReq *models.GetValue) (string, error) {
	m.Lock()
	defer m.Unlock()

	datas, err := tools.ReadJSON(q.fileName)
	if err != nil {
		return "", err
	}
	if value, found := datas[getReq.Key]; found {
		if len(value) != 0 {
			datas[getReq.Key] = datas[getReq.Key][1:]
			err := tools.WriteJSON(q.fileName, datas)
			if err != nil {
				return "", err
			}
			return value[0], nil
		}
	}
	return "", nil
}

func (q *QueueRepo) Put(putReq *models.PutValue) error {
	m.Lock()
	defer m.Unlock()

	datas, err := tools.ReadJSON(q.fileName)
	if err != nil {
		return err
	}

	if _, found := datas[putReq.Key]; found {
		datas[putReq.Key] = append(datas[putReq.Key], putReq.Value)
	} else {
		datas[putReq.Key] = []string{putReq.Value}
	}
	tools.WriteJSON(q.fileName, datas)
	return nil
}
