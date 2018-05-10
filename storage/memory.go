package storage

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type KeyValueStore interface {
	Get(int64) (string, error)
	Put(int64, string) error
}

type DB struct {
	values map[int64]string
}

func (db *DB) Get(key int64) (string, error) {
	if val, ok := db.values[key]; ok {
		return val, nil
	}
	return "", ErrNotFound
}

func (db *DB) Put(key int64, val string) error {
	db.values[key] = val
	return nil
}

func New() *DB {
	return &DB{
		values: make(map[int64]string),
	}
}
