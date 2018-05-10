package service

type KeyValueStore interface {
	Get(int64) (string, error)
	Put(int64, string) error
}

type Service struct {
	db KeyValueStore
}

func (svc *Service) GetString(key int64) (string, error) {
	return svc.db.Get(key)
}

func (svc *Service) PutString(key int64, val string) error {
	return svc.db.Put(key, val)
}

func New(db KeyValueStore) *Service {
	return &Service{
		db: db,
	}
}
