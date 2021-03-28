package repository

import (
	"log"
	"sync"

	broker "github.com/alipniczkij/web-broker"
	"github.com/alipniczkij/web-broker/tools"
)

type QueueRepo struct {
	fileName string
}

func NewQueueRepo(fileName string) *QueueRepo {
	return &QueueRepo{fileName: fileName}
}

var m sync.Mutex

func (q *QueueRepo) Get(getReq *broker.GetValue) (string, error) {
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

func (q *QueueRepo) Put(putReq *broker.PutValue) error {
	m.Lock()
	defer m.Unlock()

	datas, err := tools.ReadJSON(q.fileName)
	if err != nil {
		return err
	}

	if _, found := datas[putReq.Key]; found {
		log.Printf("Try to append value %v", putReq.Value)
		datas[putReq.Key] = append(datas[putReq.Key], putReq.Value)
	} else {
		log.Printf("Try to set value %s", []string{putReq.Value})
		datas[putReq.Key] = []string{putReq.Value}
	}
	tools.WriteJSON(q.fileName, datas)
	return nil
}
