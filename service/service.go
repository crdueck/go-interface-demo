package service

import "github.com/crdueck/interface-demo/storage"

type KeyValueService interface {
	GetString(int64) (string, error)
	PutString(int64, string) error
}

type Service struct {
	db storage.KeyValueStore
}

func (svc *Service) GetString(key int64) (string, error) {
	return svc.db.Get(key)
}

func (svc *Service) PutString(key int64, val string) error {
	return svc.db.Put(key, val)
}

func New(db storage.KeyValueStore) *Service {
	return &Service{
		db: db,
	}
}
